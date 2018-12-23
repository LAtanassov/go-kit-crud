package user

import "errors"

// Service repesents the service layer that provides CRUD operations using a Repository.
// The service layer should handle:
// - all Repository-specific errors and expose only Serivce-specific errors to its consumers
type Service struct {
	repo Repository
}

// Repository defines an upstream dependency used by Service.
type Repository interface {
	Create(user User) error
	Read(username string) error
	Update(user User) error
	Delete(username string) error
}

// User is an identity object persisted in the Repository.
type User struct {
	username   string
	GivenName  string
	FamilyName string
}

// New returns a user
func New(username, givenname, familyname string) User {
	return User{
		username:   username,
		GivenName:  givenname,
		FamilyName: familyname,
	}
}

// NewService allocates and returns (user) service.
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) create(username, givenname, familyname string) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *Service) read(username string) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *Service) update(username, givenname, familyname string) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *Service) delete(username string) (User, error) {
	return User{}, errors.New("not implemented")
}
