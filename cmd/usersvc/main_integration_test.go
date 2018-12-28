// +build integration

package main_test

import (
	"context"
	"testing"

	grpc "google.golang.org/grpc"

	"google.golang.org/grpc/status"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"
	"google.golang.org/grpc/codes"
)

func TestService_Create(t *testing.T) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	t.Run("should create a new user", func(t *testing.T) {
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: "unique-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("client.Create() error = %v", err)
		}
	})

	t.Run("should return AlreadyExists if the username already exists", func(t *testing.T) {

		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: "twice-username", Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("client.Create() error = %v", err)
		}

		_, err = client.Create(context.TODO(), &pb.CreateRequest{Username: "twice-username", Givenname: "givenname", Familyname: "familyname"})
		if err == nil {
			t.Errorf("client.Create() expects error with code = %v", codes.AlreadyExists)
		}

		c := status.Code(err)
		if c != codes.AlreadyExists {
			t.Errorf("client.Create() expects error with code = %v, but got err with code %v", codes.AlreadyExists, c)
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
			t.Errorf("client.Create() error = %v", err)
		}

		_, err = client.Read(context.TODO(), &pb.ReadRequest{Username: username})
		if err != nil {
			t.Errorf("client.Read() error = %v", err)
		}
	})

	t.Run("should return NotFound if the username does not exist", func(t *testing.T) {
		_, err := client.Read(context.TODO(), &pb.ReadRequest{Username: "non-existing-username"})
		if err == nil {
			t.Errorf("client.Read() expects error with code = %v", codes.NotFound)
		}

		c := status.Code(err)
		if c != codes.NotFound {
			t.Errorf("client.Read() expects error with code = %v, but got err with code %v", codes.NotFound, c)
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
			t.Errorf("client.Create() error = %v", err)
		}

		_, err = client.Update(context.TODO(), &pb.UpdateRequest{Username: "update-username", Givenname: "another-givenname", Familyname: "familyname"})
		if err != nil {
			t.Errorf("client.Update() error = %v", err)
		}
	})

	t.Run("should return ErrUserNotFound if the username does not exist", func(t *testing.T) {
		_, err = client.Update(context.TODO(), &pb.UpdateRequest{Username: "non-existing-username", Givenname: "another-givenname", Familyname: "familyname"})
		if err == nil {
			t.Errorf("client.Update() expects error with code = %v", codes.NotFound)
		}

		c := status.Code(err)
		if c != codes.NotFound {
			t.Errorf("client.Update() expects error with code = %v, but got err with code %v", codes.NotFound, c)
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
			t.Errorf("client.Delete() expects error with code = %v", codes.NotFound)
		}

		c := status.Code(err)
		if c != codes.NotFound {
			t.Errorf("client.Delete() expects error with code = %v, but got err with code %v", codes.NotFound, c)
		}
	})
}
