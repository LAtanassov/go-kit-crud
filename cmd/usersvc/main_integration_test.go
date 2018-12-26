// +build integration

package main_test

import (
	"context"
	"testing"

	grpc "google.golang.org/grpc"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"
	"github.com/LAtanassov/go-kit-crud/pkg/user"
)

func TestService_create(t *testing.T) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	t.Run("should create a new user", func(t *testing.T) {
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: "unique-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}
	})

	t.Run("should return ErrUserAlreadyExists if the username already exists", func(t *testing.T) {

		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: "twice-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		_, err = client.Create(context.TODO(), &pb.CreateRequest{Username: "twice-username", Givenname: "givenname", Familyname: "familyname"})
		if err == nil {
			t.Errorf("Service.Create() expects error = %v", err)

		}

		if err != user.ErrUserAlreadyExists {
			t.Errorf("Service.Create() expects error = %v, but got err %v", user.ErrUserAlreadyExists, err)

		}
	})

}

func TestService_Read(t *testing.T) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	t.Run("should read an existing new user", func(t *testing.T) {
		username := "read-username"
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: username, Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		_, err = client.Read(context.TODO(), &pb.ReadRequest{Username: username})
		if err != nil {
			t.Errorf("Service.Read() error = %v", err)

		}
	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		_, err := client.Read(context.TODO(), &pb.ReadRequest{Username: "non-existing-username"})
		if err == nil {
			t.Errorf("Service.Read() want error = %v", user.ErrUserAlreadyExists)
		}
	})
}

func TestService_Update(t *testing.T) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	t.Run("should update an existing new user", func(t *testing.T) {

		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: "update-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		_, err = client.Update(context.TODO(), &pb.UpdateRequest{Username: "update-username", Givenname: "another-givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("Service.Update() error = %v", err)
		}
	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		_, err = client.Update(context.TODO(), &pb.UpdateRequest{Username: "non-existing-username", Givenname: "another-givenname", Familyname: "familyname"})
		if err == nil {
			t.Errorf("Service.Update() want error = %v", user.ErrUserNotFound)
		}
		if err != user.ErrUserNotFound {
			t.Errorf("Service.Update() expects error = %v, but got err %v", user.ErrUserNotFound, err)
		}
	})
}

func TestService_Delete(t *testing.T) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	t.Run("should delete an existing user", func(t *testing.T) {
		username := "delete-username"
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: username, Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("Service.Create() error = %v", err)

		}

		_, err = client.Delete(context.TODO(), &pb.DeleteRequest{Username: username})
		if err != nil {
			t.Errorf("Service.Delete() error = %v", err)
		}
	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		_, err = client.Delete(context.TODO(), &pb.DeleteRequest{Username: "non-existing-username"})
		if err == nil {
			t.Errorf("Service.Delete() want error = %v", user.ErrUserNotFound)
		}
		if err != user.ErrUserNotFound {
			t.Errorf("Service.Delete() expects error = %v, but got err %v", user.ErrUserNotFound, err)
		}
	})
}
