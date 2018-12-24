package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"

	"google.golang.org/grpc"

	"github.com/LAtanassov/go-kit-crud/pkg/user"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	kitlog "github.com/go-kit/kit/log"
)

// Configuration represents a set of environment variables loaded at startup e.g.:
// #!/bin/sh
// export USERSVC_ADDR=8080
// export USERSVC_TIMEOUTSEC=15
type Configuration struct {
	TCPAddr    string `default:":8080"`
	TimeoutSec int    `default:"15"`
}

func main() {

	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	// === configuration ===

	var c Configuration

	err := envconfig.Process("usersvc", &c)
	if err != nil {
		logger.Log("err", errors.Wrap(err, "failed to process environment variables"))
		os.Exit(1)
	}

	var tcpAddr string
	tcpAddr = *flag.String("usersvc.addr", c.TCPAddr, "tcp address")
	var timeoutSec int
	timeoutSec = *flag.Int("usersvc.timeout.sec", c.TimeoutSec, "tcp connection timeout in seconds")

	// === service layer ===

	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)

	// === transport layer ===

	server := grpc.NewServer(
		grpc.ConnectionTimeout(time.Duration(timeoutSec) * time.Second),
	)

	pb.RegisterUserServer(server, user.NewGRPCServer(service))

	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		logger.Log("err", errors.Wrapf(err, "failed tcp listen to %s", tcpAddr))
		os.Exit(1)
	}
	defer lis.Close()

	go func(server *grpc.Server, lis net.Listener, logger kitlog.Logger) {

		err = server.Serve(lis) // blocking
		if err != nil {
			logger.Log("err", errors.Wrapf(err, "failed tcp listen to %s", tcpAddr))
			os.Exit(1)
		}

	}(server, lis, logger)

	logger.Log("service", "started", "listen", tcpAddr)

	// === shutdown ===

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	logger.Log("service", "stopping")
	server.GracefulStop()
	logger.Log("service", "stopped")
}
