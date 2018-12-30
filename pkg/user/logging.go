package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

// NewLoggingMiddleware returns a new instance of a logging middleware.
func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingService{logger, next}
	}
}

type loggingService struct {
	logger log.Logger
	Service
}

// Create a new user and persit it.
func (s *loggingService) Create(ctx context.Context, user User) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Create",
			"username", user.username,

			"err", fmt.Sprintf("%+v", err),

			"took", time.Since(begin),
		)
	}(time.Now())

	return s.Service.Create(ctx, user)
}

// Read an existing user.
func (s *loggingService) Read(ctx context.Context, username string) (user User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Read",
			"username", username,

			"err", fmt.Sprintf("%+v", err),

			"took", time.Since(begin),
		)
	}(time.Now())

	return s.Service.Read(ctx, username)
}

// Update an existing user.
func (s *loggingService) Update(ctx context.Context, u User) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Update",
			"username", u.username,

			"err", fmt.Sprintf("%+v", err),

			"took", time.Since(begin),
		)
	}(time.Now())

	return s.Service.Update(ctx, u)
}

// Delete an existing user.
func (s *loggingService) Delete(ctx context.Context, username string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Delete",
			"username", username,

			"err", fmt.Sprintf("%+v", err),

			"took", time.Since(begin),
		)
	}(time.Now())

	return s.Service.Delete(ctx, username)
}
