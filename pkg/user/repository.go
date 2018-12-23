package user

import (
	"errors"
	"sync"
)

var (
	// ErrUserNotFound is returned when the caller expects an user but it does not exist in the repository.
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists is returned when the caller tries to create a new user with an existing username.
	ErrUserAlreadyExists = errors.New("user already exists")
)

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

// InMemoryRepository stores user within an inmemory map.
type InMemoryRepository struct {
	sync.RWMutex
	users map[string]User
}

// NewInMemoryRepository returns an NewInMemoryRepository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[string]User),
	}
}

// Create assignns an user to the inmemory map.
func (r *InMemoryRepository) Create(u User) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.users[u.username]
	if ok {
		return ErrUserAlreadyExists
	}

	r.users[u.username] = u
	return nil
}

// Read returns an existing user
func (r *InMemoryRepository) Read(username string) (User, error) {
	r.RLock()
	defer r.RUnlock()

	u, ok := r.users[username]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return u, nil
}

// Update an existing user
func (r *InMemoryRepository) Update(u User) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.users[u.username]
	if !ok {
		return ErrUserNotFound
	}

	r.users[u.username] = u

	return nil
}

// Delete an existing user
func (r *InMemoryRepository) Delete(username string) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.users[username]
	if !ok {
		return ErrUserNotFound
	}

	delete(r.users, username)

	return nil
}
