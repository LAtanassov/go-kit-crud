// +build integration

package user_test

import (
	"context"
	"reflect"
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
		_, err := r.Read(context.TODO(), "read-non-existing-username")
		if err == nil {
			t.Errorf("Repository.Read(...) want error = %v", user.ErrUserNotFound)
		}

		if err != user.ErrUserNotFound {
			t.Errorf("Repository.Read(...) expect error = %v but got error = %v", user.ErrUserNotFound, err)
		}
	})

	t.Run("should read a user", func(t *testing.T) {
		want := user.New("read-username", "givenname", "familiyname")
		err = r.Create(context.TODO(), want)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		got, err := r.Read(context.TODO(), want.Username)
		if err != nil {
			t.Errorf("Repository.Read(...) error = %v", err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("should return ErrUserNotFound if trying delete non-existing user", func(t *testing.T) {
		err = r.Delete(context.TODO(), "delete-non-existing username")
		if err == nil {
			t.Errorf("Repository.Delete(...) want error = %v", user.ErrUserNotFound)
		}

		if err != user.ErrUserNotFound {
			t.Errorf("want err %v but got err %v", user.ErrUserNotFound, err)
		}
	})

	t.Run("should delete a user", func(t *testing.T) {
		u := user.New("delete-username", "givenname", "familiyname")
		err = r.Create(context.TODO(), u)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		err = r.Delete(context.TODO(), u.Username)
		if err != nil {
			t.Errorf("Repository.Delete(...) error = %v", err)
		}

		_, err := r.Read(context.TODO(), "delete-username")
		if err == nil {
			t.Errorf("Repository.Read(...) want error = %v", user.ErrUserNotFound)
		}

		if err != user.ErrUserNotFound {
			t.Errorf("want err %v but got err %v", user.ErrUserNotFound, err)
		}
	})
}
