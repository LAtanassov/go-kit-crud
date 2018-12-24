package user

import (
	"context"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"
)

// NewGRPCServer returns an gRPC server.
func NewGRPCServer(svc *Service) pb.UserServer {
	return &adapter{svc: svc}
}

type adapter struct {
	svc *Service
}

// Create forwards grpc calls to the service layer.
func (a *adapter) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	u := New(req.Username, req.Givenname, req.Familyname)
	err := a.svc.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReply{}, nil
}

// Read forwards grpc calls to the service layer.
func (a *adapter) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadReply, error) {
	_, err := a.svc.Read(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.ReadReply{}, nil
}

// Update forwards grpc calls to the service layer.
func (a *adapter) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	u := New(req.Username, req.Givenname, req.Familyname)
	err := a.svc.Update(ctx, u)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateReply{}, nil
}

// Delete forwards grpc calls to the service layer.
func (a *adapter) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	err := a.svc.Delete(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteReply{}, nil
}
