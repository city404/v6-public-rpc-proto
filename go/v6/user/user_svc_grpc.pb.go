// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user/user_svc.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PubUser_Get_FullMethodName     = "/v6.services.pub.PubUser/Get"
	PubUser_Login_FullMethodName   = "/v6.services.pub.PubUser/Login"
	PubUser_Refresh_FullMethodName = "/v6.services.pub.PubUser/Refresh"
	PubUser_Logoff_FullMethodName  = "/v6.services.pub.PubUser/Logoff"
	PubUser_Create_FullMethodName  = "/v6.services.pub.PubUser/Create"
	PubUser_Delete_FullMethodName  = "/v6.services.pub.PubUser/Delete"
)

// PubUserClient is the client API for PubUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubUserClient interface {
	Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *Token, opts ...grpc.CallOption) (*LoginResponse, error)
	Logoff(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type pubUserClient struct {
	cc grpc.ClientConnInterface
}

func NewPubUserClient(cc grpc.ClientConnInterface) PubUserClient {
	return &pubUserClient{cc}
}

func (c *pubUserClient) Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PubUser_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Refresh(ctx context.Context, in *Token, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PubUser_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Logoff(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Logoff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Delete(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubUserServer is the server API for PubUser service.
// All implementations must embed UnimplementedPubUserServer
// for forward compatibility
type PubUserServer interface {
	Get(context.Context, *User) (*User, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *Token) (*LoginResponse, error)
	Logoff(context.Context, *Token) (*User, error)
	Create(context.Context, *User) (*User, error)
	Delete(context.Context, *User) (*User, error)
	mustEmbedUnimplementedPubUserServer()
}

// UnimplementedPubUserServer must be embedded to have forward compatible implementations.
type UnimplementedPubUserServer struct {
}

func (UnimplementedPubUserServer) Get(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubUserServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPubUserServer) Refresh(context.Context, *Token) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedPubUserServer) Logoff(context.Context, *Token) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logoff not implemented")
}
func (UnimplementedPubUserServer) Create(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPubUserServer) Delete(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPubUserServer) mustEmbedUnimplementedPubUserServer() {}

// UnsafePubUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubUserServer will
// result in compilation errors.
type UnsafePubUserServer interface {
	mustEmbedUnimplementedPubUserServer()
}

func RegisterPubUserServer(s grpc.ServiceRegistrar, srv PubUserServer) {
	s.RegisterService(&PubUser_ServiceDesc, srv)
}

func _PubUser_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Get(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Refresh(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Logoff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Logoff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Logoff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Logoff(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Delete(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// PubUser_ServiceDesc is the grpc.ServiceDesc for PubUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubUser",
	HandlerType: (*PubUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PubUser_Get_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _PubUser_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _PubUser_Refresh_Handler,
		},
		{
			MethodName: "Logoff",
			Handler:    _PubUser_Logoff_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PubUser_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PubUser_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user_svc.proto",
}
