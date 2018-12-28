package user_test

import (
	"context"
	"testing"

	"google.golang.org/grpc/status"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"
	"github.com/LAtanassov/go-kit-crud/pkg/user"
	"google.golang.org/grpc/codes"
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

	t.Run("should return AlreadyExists if username already exists", func(t *testing.T) {
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

		_, err = g.Create(ctx, req)
		if err == nil {
			t.Errorf("GRPCServerpter.Create() want error with code = %v", codes.AlreadyExists)
			return
		}

		c := status.Code(err)
		if c != codes.AlreadyExists {
			t.Errorf("GRPCServerpter.Create() want code = %v but got %v", codes.AlreadyExists, c)
			return
		}
	})
}

func Test_adapter_Read(t *testing.T) {
	t.Run("should create an existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.CreateRequest{Username: "read-username", Givenname: "givenname", Familyname: "familyname"}

		_, err := g.Create(context.TODO(), req)
		if err != nil {
			t.Errorf("GRPCServerpter.Create() error = %v", err)
			return
		}

		_, err = g.Read(context.TODO(), &pb.ReadRequest{Username: "read-username"})
		if err != nil {
			t.Errorf("GRPCServerpter.Read() error = %v", err)
			return
		}
	})

	t.Run("should return NotFound if username does not exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.ReadRequest{Username: "username"}
		ctx := context.TODO()

		_, err := g.Read(ctx, req)

		if err == nil {
			t.Errorf("GRPCServerpter.Read() expect error with code = %v", codes.NotFound)
			return
		}

		c := status.Code(err)
		if c != codes.NotFound {
			t.Errorf("GRPCServerpter.Read() want code = %v but got %v", codes.NotFound, c)
			return
		}
	})
}

func Test_adapter_Update(t *testing.T) {
	t.Run("should update an existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		_, err := g.Create(context.TODO(), &pb.CreateRequest{Username: "update-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("GRPCServerpter.Create() error = %v", err)
			return
		}

		_, err = g.Update(context.TODO(), &pb.UpdateRequest{Username: "update-username", Givenname: "another-givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("GRPCServerpter.Update() error = %v", err)
			return
		}
	})

	t.Run("should return NotFound if username does not exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.UpdateRequest{Username: "username", Givenname: "givenname", Familyname: "familyname"}
		ctx := context.TODO()

		_, err := g.Update(ctx, req)

		if err == nil {
			t.Errorf("GRPCServerpter.Update() expect error with code = %v", codes.NotFound)
			return
		}

		c := status.Code(err)
		if c != codes.NotFound {
			t.Errorf("GRPCServerpter.Read() want code = %v but got %v", codes.NotFound, c)
			return
		}
	})
}

func Test_adapter_Delete(t *testing.T) {
	t.Run("should delete an existing user", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		_, err := g.Create(context.TODO(), &pb.CreateRequest{Username: "delete-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("GRPCServerpter.Create() error = %v", err)
			return
		}

		_, err = g.Delete(context.TODO(), &pb.DeleteRequest{Username: "delete-username"})
		if err != nil {
			t.Errorf("GRPCServerpter.Delete() error = %v", err)
			return
		}
	})

	t.Run("should return NotFound if username does not exists", func(t *testing.T) {
		r := user.NewInMemoryRepository()
		s := user.NewService(r)
		g := user.NewGRPCServer(s)

		req := &pb.DeleteRequest{Username: "username"}
		ctx := context.TODO()

		_, err := g.Delete(ctx, req)

		if err == nil {
			t.Errorf("GRPCServerpter.Delete() expect error with code = %v", codes.NotFound)
			return
		}

		c := status.Code(err)
		if c != codes.NotFound {
			t.Errorf("GRPCServerpter.Read() want code = %v but got %v", codes.NotFound, c)
			return
		}
	})
}
