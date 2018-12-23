package user_test

import (
	"reflect"
	"testing"

	"github.com/LAtanassov/go-kit-crud/pkg/user"
)

func TestService_create(t *testing.T) {

	t.Run("should create a new user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)

		err := s.Create("username", "givenname", "familyname")
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}
	})

	t.Run("should return ErrUserAlreadyExists if the username already exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)

		err := s.Create("username", "givenname", "familyname")
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		err = s.Create("username", "givenname", "familyname")
		if err == nil {
			t.Errorf("Service.Create() expects error = %v", err)

		}

		if err != user.ErrUserAlreadyExists {
			t.Errorf("Service.Create() expects error = %v, but got err %v", user.ErrUserAlreadyExists, err)

		}
	})

}

func TestService_Read(t *testing.T) {
	t.Run("should read an existing new user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		got := user.New("username", "givenname", "familyname")

		err := s.Create("username", got.GivenName, got.FamilyName)
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		want, err := s.Read("username")
		if err != nil {
			t.Errorf("Service.Read() error = %v", err)

		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Service.Read() = %v, want %v", got, want)
		}

	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)

		_, err := s.Read("username")
		if err == nil {
			t.Errorf("Service.Read() expect error = %v", user.ErrUserNotFound)

		}
		if err != user.ErrUserNotFound {
			t.Errorf("Service.Read() expects error = %v, but got err %v", user.ErrUserNotFound, err)

		}
	})
}

func TestService_Update(t *testing.T) {
	t.Run("should update an existing new user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		u := user.New("username", "givenname", "familyname")

		err := s.Create("username", u.GivenName, u.FamilyName)
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		u.FamilyName = "Another-Family"
		err = s.Update(u)
		if err != nil {
			t.Errorf("Service.Update() error = %v", err)

		}
	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		u := user.New("username", "givenname", "familyname")

		err := s.Update(u)
		if err == nil {
			t.Errorf("Service.Update() expect error = %v", user.ErrUserNotFound)

		}
		if err != user.ErrUserNotFound {
			t.Errorf("Service.Update() expects error = %v, but got err %v", user.ErrUserNotFound, err)

		}
	})
}

func TestService_Delete(t *testing.T) {
	t.Run("should delete an existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)

		err := s.Create("username", "givenname", "familyname")
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		err = s.Delete("username")
		if err != nil {
			t.Errorf("Service.Delete() error = %v", err)
		}
	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)

		err := s.Delete("username")
		if err == nil {
			t.Errorf("Service.Update() expect error = %v", user.ErrUserNotFound)

		}
		if err != user.ErrUserNotFound {
			t.Errorf("Service.Update() expects error = %v, but got err %v", user.ErrUserNotFound, err)
		}
	})
}
