// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: sftp/public_sftp.proto

package sftpconfig

import (
	context "context"
	common "github.com/city404/v6-public-rpc-proto/go/v6/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PubSftpConfig_Get_FullMethodName              = "/v6.services.pub.PubSftpConfig/Get"
	PubSftpConfig_Update_FullMethodName           = "/v6.services.pub.PubSftpConfig/Update"
	PubSftpConfig_Enable_FullMethodName           = "/v6.services.pub.PubSftpConfig/Enable"
	PubSftpConfig_Disable_FullMethodName          = "/v6.services.pub.PubSftpConfig/Disable"
	PubSftpConfig_UpdateKeys_FullMethodName       = "/v6.services.pub.PubSftpConfig/UpdateKeys"
	PubSftpConfig_ValidateUserName_FullMethodName = "/v6.services.pub.PubSftpConfig/ValidateUserName"
)

// PubSftpConfigClient is the client API for PubSftpConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSftpConfigClient interface {
	// rpc Create (DavConfig) returns (DavConfig) {}
	Get(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error)
	Update(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error)
	Enable(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error)
	Disable(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error)
	UpdateKeys(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error)
	ValidateUserName(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*common.UserNameValidateResponse, error)
}

type pubSftpConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSftpConfigClient(cc grpc.ClientConnInterface) PubSftpConfigClient {
	return &pubSftpConfigClient{cc}
}

func (c *pubSftpConfigClient) Get(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) Update(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) Enable(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) Disable(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) UpdateKeys(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_UpdateKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) ValidateUserName(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*common.UserNameValidateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.UserNameValidateResponse)
	err := c.cc.Invoke(ctx, PubSftpConfig_ValidateUserName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSftpConfigServer is the server API for PubSftpConfig service.
// All implementations must embed UnimplementedPubSftpConfigServer
// for forward compatibility.
type PubSftpConfigServer interface {
	// rpc Create (DavConfig) returns (DavConfig) {}
	Get(context.Context, *SftpConfig) (*SftpConfig, error)
	Update(context.Context, *SftpConfig) (*SftpConfig, error)
	Enable(context.Context, *SftpConfig) (*SftpConfig, error)
	Disable(context.Context, *SftpConfig) (*SftpConfig, error)
	UpdateKeys(context.Context, *SftpConfig) (*SftpConfig, error)
	ValidateUserName(context.Context, *SftpConfig) (*common.UserNameValidateResponse, error)
	mustEmbedUnimplementedPubSftpConfigServer()
}

// UnimplementedPubSftpConfigServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPubSftpConfigServer struct{}

func (UnimplementedPubSftpConfigServer) Get(context.Context, *SftpConfig) (*SftpConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubSftpConfigServer) Update(context.Context, *SftpConfig) (*SftpConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPubSftpConfigServer) Enable(context.Context, *SftpConfig) (*SftpConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedPubSftpConfigServer) Disable(context.Context, *SftpConfig) (*SftpConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedPubSftpConfigServer) UpdateKeys(context.Context, *SftpConfig) (*SftpConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeys not implemented")
}
func (UnimplementedPubSftpConfigServer) ValidateUserName(context.Context, *SftpConfig) (*common.UserNameValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserName not implemented")
}
func (UnimplementedPubSftpConfigServer) mustEmbedUnimplementedPubSftpConfigServer() {}
func (UnimplementedPubSftpConfigServer) testEmbeddedByValue()                       {}

// UnsafePubSftpConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSftpConfigServer will
// result in compilation errors.
type UnsafePubSftpConfigServer interface {
	mustEmbedUnimplementedPubSftpConfigServer()
}

func RegisterPubSftpConfigServer(s grpc.ServiceRegistrar, srv PubSftpConfigServer) {
	// If the following call pancis, it indicates UnimplementedPubSftpConfigServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PubSftpConfig_ServiceDesc, srv)
}

func _PubSftpConfig_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSftpConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSftpConfig_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSftpConfigServer).Get(ctx, req.(*SftpConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSftpConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSftpConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSftpConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSftpConfigServer).Update(ctx, req.(*SftpConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSftpConfig_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSftpConfigServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSftpConfig_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSftpConfigServer).Enable(ctx, req.(*SftpConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSftpConfig_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSftpConfigServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSftpConfig_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSftpConfigServer).Disable(ctx, req.(*SftpConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSftpConfig_UpdateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSftpConfigServer).UpdateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSftpConfig_UpdateKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSftpConfigServer).UpdateKeys(ctx, req.(*SftpConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSftpConfig_ValidateUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SftpConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSftpConfigServer).ValidateUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSftpConfig_ValidateUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSftpConfigServer).ValidateUserName(ctx, req.(*SftpConfig))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSftpConfig_ServiceDesc is the grpc.ServiceDesc for PubSftpConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSftpConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubSftpConfig",
	HandlerType: (*PubSftpConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PubSftpConfig_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PubSftpConfig_Update_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _PubSftpConfig_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _PubSftpConfig_Disable_Handler,
		},
		{
			MethodName: "UpdateKeys",
			Handler:    _PubSftpConfig_UpdateKeys_Handler,
		},
		{
			MethodName: "ValidateUserName",
			Handler:    _PubSftpConfig_ValidateUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sftp/public_sftp.proto",
}
