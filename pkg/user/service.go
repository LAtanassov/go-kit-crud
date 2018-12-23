package user

import "context"

// Service repesents the service layer that provides CRUD operations using a Repository.
// The service layer should handle:
// - all Repository-specific errors and expose only Serivce-specific errors to its consumers
type Service struct {
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
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// Create a new user and persit it.
func (s *Service) Create(ctx context.Context, user User) error {
	return s.repo.Create(ctx, user)
}

// Read an existing user.
func (s *Service) Read(ctx context.Context, username string) (User, error) {
	return s.repo.Read(ctx, username)
}

// Update an existing user.
func (s *Service) Update(ctx context.Context, u User) error {
	return s.repo.Update(ctx, u)
}

// Delete an existing user.
func (s *Service) Delete(ctx context.Context, username string) error {
	return s.repo.Delete(ctx, username)
}
