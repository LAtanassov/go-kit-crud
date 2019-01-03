package user

import (
	"context"
	"errors"
	"sync"

	"database/sql"

	"github.com/go-sql-driver/mysql"

	"github.com/VividCortex/mysqlerr"
)

var (
	// ErrUserNotFound is returned when the caller expects an user but it does not exist in the repository.
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists is returned when the caller tries to create a new user with an existing username.
	ErrUserAlreadyExists = errors.New("user already exists")
)

// User is an identity object persisted in the Repository.
type User struct {
	Username   string
	GivenName  string
	FamilyName string
}

// New returns a user
func New(username, givenname, familyname string) User {
	return User{
		Username:   username,
		GivenName:  givenname,
		FamilyName: familyname,
	}
}

// === InMemory ===

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
func (r *InMemoryRepository) Create(_ context.Context, u User) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.users[u.Username]
	if ok {
		return ErrUserAlreadyExists
	}

	r.users[u.Username] = u
	return nil
}

// Read returns an existing user
func (r *InMemoryRepository) Read(_ context.Context, username string) (User, error) {
	r.RLock()
	defer r.RUnlock()

	u, ok := r.users[username]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return u, nil
}

// Update an existing user
func (r *InMemoryRepository) Update(_ context.Context, u User) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.users[u.Username]
	if !ok {
		return ErrUserNotFound
	}

	r.users[u.Username] = u

	return nil
}

// Delete an existing user
func (r *InMemoryRepository) Delete(_ context.Context, username string) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.users[username]
	if !ok {
		return ErrUserNotFound
	}

	delete(r.users, username)

	return nil
}

// === MySQL ===

// SQLRepository stores users within a sql database
type SQLRepository struct {
	db *sql.DB
}

// NewSQLRepository returns an new SQLRepository
func NewSQLRepository(driver, dsn string) (*SQLRepository, error) {

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return &SQLRepository{
		db: db,
	}, nil
}

// Create and writes a user into the mysql database
func (r *SQLRepository) Create(ctx context.Context, u User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO Users(Username, Givenname, Familyname) VALUES (?, ?, ?)", u.Username, u.GivenName, u.FamilyName)
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
			return ErrUserAlreadyExists
		}
	}
	return err
}

// Create and writes a user into the mysql database
func (r *SQLRepository) Read(ctx context.Context, username string) (User, error) {
	u := User{}
	err := r.db.QueryRowContext(ctx, "SELECT u.Username, u.Givenname, u.Familyname FROM Users as u WHERE u.Username = ?", username).Scan(&u.Username, &u.GivenName, &u.FamilyName)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, ErrUserNotFound
		}
		return u, err
	}
	return u, nil
}

// Delete a user from mysql database
func (r *SQLRepository) Delete(ctx context.Context, username string) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM Users WHERE Username = ?", username)
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowCnt == 0 {
		return ErrUserNotFound
	}

	return err
}
