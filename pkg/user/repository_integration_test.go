// +build integration

package user_test

import (
	"context"
	"testing"

	"github.com/LAtanassov/go-kit-crud/pkg/user"
)

func TestMySQLRepository(t *testing.T) {

	r, err := user.NewSQLRepository("mysql", "UserService:password@tcp(localhost:3306)/UserDB")
	if err != nil {
		t.Errorf("NewSQLRepository(...) error = %v", err)
	}

	t.Run("should insert a new user", func(t *testing.T) {
		err := r.Create(context.TODO(), user.New("create-username", "givenname", "familiyname"))
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}
	})

	t.Run("should return a new user", func(t *testing.T) {
		err := r.Create(context.TODO(), user.New("create-twice-username", "givenname", "familiyname"))
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		err = r.Create(context.TODO(), user.New("create-twice-username", "givenname", "familiyname"))
		if err == nil {
			t.Errorf("Repository.Create(...) want error = %v", user.ErrUserAlreadyExists)
		}
		if err != user.ErrUserAlreadyExists {
			t.Errorf("Repository.Create(...) expect error = %v but got error = %v", user.ErrUserAlreadyExists, err)
		}
	})

	t.Run("should return ErrUserNotFound if read non-existing username", func(t *testing.T) {
		_, err := r.Read(context.TODO(), "read-username")
		if err == nil {
			t.Errorf("Repository.Read(...) want error = %v", user.ErrUserNotFound)
		}

		if err != user.ErrUserNotFound {
			t.Errorf("Repository.Read(...) expect error = %v but got error = %v", user.ErrUserNotFound, err)
		}
	})
}
