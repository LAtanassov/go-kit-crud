package user

import "context"

// Service is used for a layered architecture.
type Service interface {
	// Create a new user and persit it.
	Create(ctx context.Context, user User) error
	// Read an existing user.
	Read(ctx context.Context, username string) (User, error)
	// Update an existing user.
	Update(ctx context.Context, u User) error
	// Delete an existing user.
	Delete(ctx context.Context, username string) error
}

// Middleware is a chainable behavior modifier for Service.
type Middleware func(Service) Service

// ServiceImpl repesents the service layer that provides CRUD operations using a Repository.
// The service layer should handle:
// - all Repository-specific errors and expose only Serivce-specific errors to its consumers
type ServiceImpl struct {
	repo Repository
}

// Repository defines an upstream dependency used by Service.
type Repository interface {
	Create(ctx context.Context, user User) error
	Read(ctx context.Context, username string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, username string) error
}

// NewService allocates and returns (user) service.
func NewService(r Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: r,
	}
}

// Create a new user and persit it.
func (s *ServiceImpl) Create(ctx context.Context, user User) error {
	return s.repo.Create(ctx, user)
}

// Read an existing user.
func (s *ServiceImpl) Read(ctx context.Context, username string) (User, error) {
	return s.repo.Read(ctx, username)
}

// Update an existing user.
func (s *ServiceImpl) Update(ctx context.Context, u User) error {
	return s.repo.Update(ctx, u)
}

// Delete an existing user.
func (s *ServiceImpl) Delete(ctx context.Context, username string) error {
	return s.repo.Delete(ctx, username)
}
