package user

// Service repesents the service layer that provides CRUD operations using a Repository.
// The service layer should handle:
// - all Repository-specific errors and expose only Serivce-specific errors to its consumers
type Service struct {
	repo Repository
}

// Repository defines an upstream dependency used by Service.
type Repository interface {
	Create(user User) error
	Read(username string) (User, error)
	Update(user User) error
	Delete(username string) error
}

// NewService allocates and returns (user) service.
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// Create a new user and persit it.
func (s *Service) Create(username, givenname, familyname string) error {
	u := New(username, givenname, familyname)
	return s.repo.Create(u)
}

// Read an existing user.
func (s *Service) Read(username string) (User, error) {
	return s.repo.Read(username)
}

// Update an existing user.
func (s *Service) Update(u User) error {
	return s.repo.Update(u)
}

// Delete an existing user.
func (s *Service) Delete(username string) error {
	return s.repo.Delete(username)
}
