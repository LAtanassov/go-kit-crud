package main_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"
	grpc "google.golang.org/grpc"
)

func BenchmarkService_Create(b *testing.B) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	users := generate(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: users[i], Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			b.Errorf("Service.Create() error = %v", err)
		}
	}
}

func BenchmarkService_Read(b *testing.B) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	users := generate(b.N)

	for i := 0; i < b.N; i++ {
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: users[i], Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			b.Errorf("Service.Create() error = %v", err)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = client.Read(context.TODO(), &pb.ReadRequest{Username: users[i]})
		if err != nil {
			b.Errorf("Service.Read() error = %v", err)

		}
	}
}

func BenchmarkService_Update(b *testing.B) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	users := generate(b.N)

	for i := 0; i < b.N; i++ {
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: users[i], Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			b.Errorf("Service.Create() error = %v", err)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = client.Update(context.TODO(), &pb.UpdateRequest{Username: users[i], Givenname: "another-givenname", Familyname: "familyname"})
		if err != nil {
			b.Errorf("Service.Read() error = %v", err)

		}
	}
}

func BenchmarkService_Delete(b *testing.B) {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	users := generate(b.N)

	for i := 0; i < b.N; i++ {
		_, err := client.Create(context.TODO(), &pb.CreateRequest{Username: users[i], Givenname: "givenname", Familyname: "familyname"})
		if err != nil {
			b.Errorf("Service.Create() error = %v", err)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = client.Delete(context.TODO(), &pb.DeleteRequest{Username: users[i]})
		if err != nil {
			b.Errorf("Service.Read() error = %v", err)

		}
	}

}

func generate(n int) []string {
	users := make([]string, n)
	for i := 0; i < n; i++ {
		users[i] = big.NewInt(rand.Int63()).Text(16)
	}
	return users
}
