package user_test

import (
	"reflect"
	"testing"

	"github.com/LAtanassov/go-kit-crud/pkg/user"
)

func TestInMemoryRepository_create(t *testing.T) {

	t.Run("should create a user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		u := user.New("username", "firstname", "familyname")

		err := r.Create(u)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}
	})

	t.Run("should return ErrUserAlreadyExists if username is already used", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		u := user.New("username", "firstname", "familyname")

		err := r.Create(u)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		err = r.Create(u)
		if err == nil {
			t.Errorf("Repository.Create(...) expect error = %v", user.ErrUserAlreadyExists)
		}

		if err != user.ErrUserAlreadyExists {
			t.Errorf("Repository.Create(...) expect error = %v but got error = %v", user.ErrUserAlreadyExists, err)
		}
	})

	t.Run("should read an existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		want := user.New("username", "firstname", "familyname")

		err := r.Create(want)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		got, err := r.Read("username")
		if err != nil {
			t.Errorf("Repository.Read(...) error = %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Service.Read() = %v, want %v", got, want)
		}
	})

	t.Run("should read an non-existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		want := user.New("username", "firstname", "familyname")

		err := r.Create(want)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		_, err = r.Read("another-username")
		if err == nil {
			t.Errorf("Repository.Read(...) want error = %v", user.ErrUserNotFound)
		}

		if err != user.ErrUserNotFound {
			t.Errorf("Repository.Read(...) expect error = %v but got error = %v", user.ErrUserNotFound, err)
		}
	})

	t.Run("should update a existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		u := user.New("username", "firstname", "familyname")

		err := r.Create(u)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		u.FamilyName = "newFamilyName"
		err = r.Update(u)
		if err != nil {
			t.Errorf("Repository.Update(...) error = %v", err)
		}
	})

	t.Run("should update a non-existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		oldOne := user.New("old-username", "firstname", "familyname")

		err := r.Create(oldOne)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		newOne := user.New("new-username", "firstname", "familyname")
		err = r.Update(newOne)
		if err == nil {
			t.Errorf("Repository.Update(...) error = %v", user.ErrUserNotFound)
		}
		if err != user.ErrUserNotFound {
			t.Errorf("Repository.Update(...) expect error = %v but got error = %v", user.ErrUserNotFound, err)
		}
	})

	t.Run("should delete a existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		u := user.New("username", "firstname", "familyname")

		err := r.Create(u)
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}

		err = r.Delete("username")
		if err != nil {
			t.Errorf("Repository.Delete(...) error = %v", err)
		}
	})

	t.Run("should delete a non-existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()

		err := r.Delete("username")
		if err == nil {
			t.Errorf("Repository.Delete(...) expect error = %v", user.ErrUserNotFound)
		}

		if err != user.ErrUserNotFound {
			t.Errorf("Repository.Delete(...) expect error = %v but got error %v", user.ErrUserNotFound, err)
		}
	})

}
