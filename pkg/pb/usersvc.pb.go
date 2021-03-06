// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usersvc.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CreateRequest contains all fields necessary to create a new user.
type CreateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Givenname            string   `protobuf:"bytes,2,opt,name=givenname,proto3" json:"givenname,omitempty"`
	Familyname           string   `protobuf:"bytes,3,opt,name=familyname,proto3" json:"familyname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateRequest) GetGivenname() string {
	if m != nil {
		return m.Givenname
	}
	return ""
}

func (m *CreateRequest) GetFamilyname() string {
	if m != nil {
		return m.Familyname
	}
	return ""
}

type CreateReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReply) Reset()         { *m = CreateReply{} }
func (m *CreateReply) String() string { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()    {}
func (*CreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{1}
}
func (m *CreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReply.Unmarshal(m, b)
}
func (m *CreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReply.Marshal(b, m, deterministic)
}
func (dst *CreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReply.Merge(dst, src)
}
func (m *CreateReply) XXX_Size() int {
	return xxx_messageInfo_CreateReply.Size(m)
}
func (m *CreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReply proto.InternalMessageInfo

func (m *CreateReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ReadRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{2}
}
func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (dst *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(dst, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ReadReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadReply) Reset()         { *m = ReadReply{} }
func (m *ReadReply) String() string { return proto.CompactTextString(m) }
func (*ReadReply) ProtoMessage()    {}
func (*ReadReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{3}
}
func (m *ReadReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadReply.Unmarshal(m, b)
}
func (m *ReadReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadReply.Marshal(b, m, deterministic)
}
func (dst *ReadReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadReply.Merge(dst, src)
}
func (m *ReadReply) XXX_Size() int {
	return xxx_messageInfo_ReadReply.Size(m)
}
func (m *ReadReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReadReply proto.InternalMessageInfo

func (m *ReadReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// UpdateRequest contains all fields necessary to update an existing user.
type UpdateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Givenname            string   `protobuf:"bytes,2,opt,name=givenname,proto3" json:"givenname,omitempty"`
	Familyname           string   `protobuf:"bytes,3,opt,name=familyname,proto3" json:"familyname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{4}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateRequest) GetGivenname() string {
	if m != nil {
		return m.Givenname
	}
	return ""
}

func (m *UpdateRequest) GetFamilyname() string {
	if m != nil {
		return m.Familyname
	}
	return ""
}

type UpdateReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReply) Reset()         { *m = UpdateReply{} }
func (m *UpdateReply) String() string { return proto.CompactTextString(m) }
func (*UpdateReply) ProtoMessage()    {}
func (*UpdateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{5}
}
func (m *UpdateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReply.Unmarshal(m, b)
}
func (m *UpdateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReply.Marshal(b, m, deterministic)
}
func (dst *UpdateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReply.Merge(dst, src)
}
func (m *UpdateReply) XXX_Size() int {
	return xxx_messageInfo_UpdateReply.Size(m)
}
func (m *UpdateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReply proto.InternalMessageInfo

func (m *UpdateReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type DeleteRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{6}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type DeleteReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReply) Reset()         { *m = DeleteReply{} }
func (m *DeleteReply) String() string { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()    {}
func (*DeleteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_usersvc_0e3f5da68e3f87e1, []int{7}
}
func (m *DeleteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReply.Unmarshal(m, b)
}
func (m *DeleteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReply.Marshal(b, m, deterministic)
}
func (dst *DeleteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReply.Merge(dst, src)
}
func (m *DeleteReply) XXX_Size() int {
	return xxx_messageInfo_DeleteReply.Size(m)
}
func (m *DeleteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReply proto.InternalMessageInfo

func (m *DeleteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "pb.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "pb.CreateReply")
	proto.RegisterType((*ReadRequest)(nil), "pb.ReadRequest")
	proto.RegisterType((*ReadReply)(nil), "pb.ReadReply")
	proto.RegisterType((*UpdateRequest)(nil), "pb.UpdateRequest")
	proto.RegisterType((*UpdateReply)(nil), "pb.UpdateReply")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteReply)(nil), "pb.DeleteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/pb.User/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error) {
	out := new(ReadReply)
	err := c.cc.Invoke(ctx, "/pb.User/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := c.cc.Invoke(ctx, "/pb.User/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := c.cc.Invoke(ctx, "/pb.User/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	Read(context.Context, *ReadRequest) (*ReadReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _User_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usersvc.proto",
}

func init() { proto.RegisterFile("usersvc.proto", fileDescriptor_usersvc_0e3f5da68e3f87e1) }

var fileDescriptor_usersvc_0e3f5da68e3f87e1 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x18, 0x6c, 0xda, 0x12, 0xcc, 0x94, 0x45, 0xfd, 0x4e, 0x12, 0xfc, 0x63, 0x4f, 0x15, 0x61, 0x0f,
	0xfa, 0x08, 0xfa, 0x04, 0x81, 0x3e, 0x40, 0x62, 0x3f, 0x25, 0x90, 0xb6, 0xeb, 0x26, 0x2d, 0xe4,
	0xed, 0x7c, 0x34, 0xc9, 0x7e, 0xc6, 0x6e, 0x20, 0x87, 0x9e, 0x7a, 0x4b, 0x66, 0x86, 0xf9, 0x66,
	0x26, 0x81, 0xda, 0xd7, 0xec, 0xea, 0xc3, 0x87, 0xb1, 0x6e, 0xd7, 0xec, 0x68, 0x6a, 0x0b, 0x5d,
	0x42, 0xbd, 0x39, 0xce, 0x1b, 0xce, 0xf8, 0x7b, 0xcf, 0x75, 0x43, 0x29, 0x2e, 0x3a, 0xd5, 0x36,
	0xdf, 0xf0, 0x4d, 0xf4, 0x18, 0x2d, 0x93, 0xec, 0xff, 0x9d, 0x6e, 0x91, 0x7c, 0x95, 0x07, 0xde,
	0x7a, 0x72, 0xea, 0xc9, 0x23, 0x40, 0xf7, 0xc0, 0x67, 0xbe, 0x29, 0xab, 0xd6, 0xd3, 0x33, 0x4f,
	0x07, 0x88, 0x7e, 0xc0, 0xa2, 0x3f, 0x65, 0xab, 0x96, 0xae, 0x30, 0x63, 0xe7, 0xfe, 0x6e, 0x74,
	0x8f, 0xfa, 0x09, 0x8b, 0x8c, 0xf3, 0xf5, 0x09, 0x49, 0xf4, 0x1d, 0x12, 0x91, 0x8e, 0x3b, 0x95,
	0x50, 0x2b, 0xbb, 0x3e, 0x57, 0xab, 0xfe, 0xd4, 0x78, 0x96, 0x67, 0xa8, 0x77, 0xae, 0xf8, 0xa4,
	0x2c, 0x9d, 0x5b, 0x2f, 0x1e, 0x75, 0x7b, 0xf9, 0x89, 0x30, 0x5f, 0xd5, 0xec, 0xc8, 0x20, 0x96,
	0x35, 0xe9, 0xda, 0xd8, 0xc2, 0x0c, 0x3e, 0x62, 0x7a, 0x19, 0x42, 0xb6, 0x6a, 0xf5, 0x84, 0x96,
	0x98, 0x77, 0x8b, 0x91, 0xa7, 0x82, 0x99, 0x53, 0x75, 0x04, 0x44, 0x69, 0x10, 0x4b, 0x23, 0x71,
	0x1e, 0x0c, 0x29, 0xce, 0x41, 0x61, 0xd1, 0x4b, 0x66, 0xd1, 0x0f, 0xca, 0x8a, 0x3e, 0xa8, 0xa4,
	0x27, 0x45, 0xec, 0xff, 0xbe, 0xd7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x5e, 0x82, 0x42,
	0x8e, 0x02, 0x00, 0x00,
}
