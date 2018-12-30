package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"
	"time"
	"net/http"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"

	"google.golang.org/grpc"

	"github.com/LAtanassov/go-kit-crud/pkg/user"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Configuration represents a set of environment variables loaded at startup e.g.:
// #!/bin/sh
// export USERSVC_ADDR=8080
// export USERSVC_TIMEOUTSEC=15
// export USERSVC_OPSADDR=9090
type Configuration struct {
	Addr    string `default:":8080"`
	TimeoutSec int    `default:"15"`

	OpsAddr string `default:":9090"`
}

func main() {

	logger := newStdErrLogger()

	config, err := loadConfiguration()
	if err != nil {
		logger.Log("error", errors.Wrap(err, "could not load configuration"))
		os.Exit(1)
	}

	// === service layer ===

	fieldKeys := []string{"method"}
	repo := user.NewInMemoryRepository()

	var svc user.Service
	svc = user.NewService(repo)
	svc = user.NewLoggingMiddleware(kitlog.With(logger, "component", "UserService"))(svc)
	svc = user.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "UserService",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "UserService",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys))(svc)

	// === transport layer ===

	server := grpc.NewServer(
		grpc.ConnectionTimeout(time.Duration(config.TimeoutSec) * time.Second),
	)

	pb.RegisterUserServer(server, user.NewGRPCServer(svc))

	lis, err := net.Listen("tcp", config.Addr)
	if err != nil {
		logger.Log("error", errors.Wrapf(err, "failed listen to %s", config.Addr))
		os.Exit(1)
	}
	defer lis.Close()

	go func(server *grpc.Server, lis net.Listener, logger kitlog.Logger) {

		err = server.Serve(lis) // blocking
		if err != nil {
			logger.Log("err", errors.Wrap(err, "failed to start grpc service"))
			// TODO: this migth not work here
			os.Exit(1)
		}

	}(server, lis, logger)

	logger.Log("service", "grpc user service", "listen", config.Addr)

	// === ops layer ===

	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())
	//handler.Handle("/_status/liveness", t.MakeLivenessHandler())
	//handler.Handle("/_status/readiness", t.MakeReadinessHandler())

	opsServer := &http.Server{
		Addr:         config.OpsAddr,
		Handler:      handler,
		ReadTimeout:  time.Duration(config.TimeoutSec) * time.Second,
		WriteTimeout: time.Duration(config.TimeoutSec) * time.Second,
	}

	go func(server *http.Server, logger kitlog.Logger) {
		err := server.ListenAndServe()
		if err != nil {
			logger.Log("error", errors.Wrap(err, "failed to start operations service"))
			// TODO: this migth not work here
			os.Exit(1)
		}
	}(opsServer, logger)

	logger.Log("service", "operations service", "listen", config.OpsAddr)

	// === shutdown ===

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	logger.Log("service", "stopping")
	server.GracefulStop()
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	opsServer.Shutdown(ctx)
	
	logger.Log("service", "stopped")
}

func newStdErrLogger() kitlog.Logger {
	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	return logger
}

func loadConfiguration() (Configuration, error) {
	// from env.
	var env Configuration
	err := envconfig.Process("usersvc", &env)
	if err != nil {
		return Configuration{}, err
	}

	// from cli
	addr := flag.String("usersvc.addr", env.Addr, "service address")
	timeoutSec := flag.Int("usersvc.timeout.sec", env.TimeoutSec, "tcp connection timeout in seconds")
	opsAddr := flag.String("usersvc.ops.addr", env.OpsAddr, "operation service address")

	return Configuration{
		Addr: *addr,
		TimeoutSec: *timeoutSec,
		OpsAddr : *opsAddr,
	}, nil
}
