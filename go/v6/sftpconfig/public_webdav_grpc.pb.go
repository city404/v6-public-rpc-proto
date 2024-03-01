// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: sftp/public_webdav.proto

package sftpconfig

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
	PubSftpConfig_Get_FullMethodName     = "/v6.services.pub.PubSftpConfig/Get"
	PubSftpConfig_Update_FullMethodName  = "/v6.services.pub.PubSftpConfig/Update"
	PubSftpConfig_Enable_FullMethodName  = "/v6.services.pub.PubSftpConfig/Enable"
	PubSftpConfig_Disable_FullMethodName = "/v6.services.pub.PubSftpConfig/Disable"
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
}

type pubSftpConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSftpConfigClient(cc grpc.ClientConnInterface) PubSftpConfigClient {
	return &pubSftpConfigClient{cc}
}

func (c *pubSftpConfigClient) Get(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) Update(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) Enable(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSftpConfigClient) Disable(ctx context.Context, in *SftpConfig, opts ...grpc.CallOption) (*SftpConfig, error) {
	out := new(SftpConfig)
	err := c.cc.Invoke(ctx, PubSftpConfig_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSftpConfigServer is the server API for PubSftpConfig service.
// All implementations must embed UnimplementedPubSftpConfigServer
// for forward compatibility
type PubSftpConfigServer interface {
	// rpc Create (DavConfig) returns (DavConfig) {}
	Get(context.Context, *SftpConfig) (*SftpConfig, error)
	Update(context.Context, *SftpConfig) (*SftpConfig, error)
	Enable(context.Context, *SftpConfig) (*SftpConfig, error)
	Disable(context.Context, *SftpConfig) (*SftpConfig, error)
	mustEmbedUnimplementedPubSftpConfigServer()
}

// UnimplementedPubSftpConfigServer must be embedded to have forward compatible implementations.
type UnimplementedPubSftpConfigServer struct {
}

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
func (UnimplementedPubSftpConfigServer) mustEmbedUnimplementedPubSftpConfigServer() {}

// UnsafePubSftpConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSftpConfigServer will
// result in compilation errors.
type UnsafePubSftpConfigServer interface {
	mustEmbedUnimplementedPubSftpConfigServer()
}

func RegisterPubSftpConfigServer(s grpc.ServiceRegistrar, srv PubSftpConfigServer) {
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sftp/public_webdav.proto",
}
