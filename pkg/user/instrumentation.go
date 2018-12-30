package user

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
)

// NewInstrumentingMiddleware returns an instance of the instrumented middleware.
func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next Service) Service {
		return &instrumentingService{
			requestCount:   counter,
			requestLatency: latency,
			Service:        next,
		}
	}
}

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// Create a new user and persit it.
func (s *instrumentingService) Create(ctx context.Context, user User) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Create").Add(1)
		s.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Create(ctx, user)
}

// Read an existing user.
func (s *instrumentingService) Read(ctx context.Context, username string) (user User, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Read").Add(1)
		s.requestLatency.With("method", "Read").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Read(ctx, username)
}

// Update an existing user.
func (s *instrumentingService) Update(ctx context.Context, u User) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Update").Add(1)
		s.requestLatency.With("method", "Update").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Update(ctx, u)
}

// Delete an existing user.
func (s *instrumentingService) Delete(ctx context.Context, username string) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Delete").Add(1)
		s.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Delete(ctx, username)
}
