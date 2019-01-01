// +build integration

package user_test

import (
	"context"
	"testing"

	"github.com/LAtanassov/go-kit-crud/pkg/user"
)

func TestMySQLRepository(t *testing.T) {

	t.Run("should create a user", func(t *testing.T) {
		r, err := user.NewSQLRepository("mysql", "UserService:password@tcp(localhost:3306)/UserDB")
		if err != nil {
			t.Errorf("NewSQLRepository(...) error = %v", err)
		}

		err = r.Create(context.TODO(), user.New("username", "firstname", "familyname"))
		if err != nil {
			t.Errorf("Repository.Create(...) error = %v", err)
		}
	})
}
