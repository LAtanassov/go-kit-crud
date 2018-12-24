package user_test

import (
	"context"
	"testing"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"
	"github.com/LAtanassov/go-kit-crud/pkg/user"
)

func Test_adapter_Create(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.CreateRequest{Username: "username", Givenname: "givenname", Familyname: "familyname"}
		ctx := context.TODO()

		_, err := g.Create(ctx, req)

		if err != nil {
			t.Errorf("GRPCServerpter.Create() error = %v", err)
			return
		}
	})
}

func Test_adapter_Read(t *testing.T) {
	t.Run("should return user not found if username does not exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.ReadRequest{Username: "username"}
		ctx := context.TODO()

		_, err := g.Read(ctx, req)

		if err == nil {
			t.Errorf("GRPCServerpter.Read() expect error = %v", user.ErrUserNotFound)
			return
		}
	})
}

func Test_adapter_Update(t *testing.T) {
	t.Run("should return user not found if username does not exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.UpdateRequest{Username: "username", Givenname: "givenname", Familyname: "familyname"}
		ctx := context.TODO()

		_, err := g.Update(ctx, req)

		if err == nil {
			t.Errorf("GRPCServerpter.Update() expect error = %v", user.ErrUserNotFound)
			return
		}
	})
}

func Test_adapter_Delete(t *testing.T) {
	t.Run("should return user not found if username does not exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.DeleteRequest{Username: "username"}
		ctx := context.TODO()

		_, err := g.Delete(ctx, req)

		if err == nil {
			t.Errorf("GRPCServerpter.Delete() expect error = %v", user.ErrUserNotFound)
			return
		}
	})
}
