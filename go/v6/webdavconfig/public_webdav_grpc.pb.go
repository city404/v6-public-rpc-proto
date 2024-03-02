// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: webdav/public_webdav.proto

package webdavconfig

import (
	context "context"
	common "github.com/city404/v6-public-rpc-proto/go/v6/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PubDavConfig_Get_FullMethodName              = "/v6.services.pub.PubDavConfig/Get"
	PubDavConfig_Update_FullMethodName           = "/v6.services.pub.PubDavConfig/Update"
	PubDavConfig_Enable_FullMethodName           = "/v6.services.pub.PubDavConfig/Enable"
	PubDavConfig_Disable_FullMethodName          = "/v6.services.pub.PubDavConfig/Disable"
	PubDavConfig_ValidateUserName_FullMethodName = "/v6.services.pub.PubDavConfig/ValidateUserName"
)

// PubDavConfigClient is the client API for PubDavConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubDavConfigClient interface {
	// rpc Create (DavConfig) returns (DavConfig) {}
	Get(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error)
	Update(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error)
	Enable(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error)
	Disable(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error)
	ValidateUserName(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*common.UserNameValidateResponse, error)
}

type pubDavConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewPubDavConfigClient(cc grpc.ClientConnInterface) PubDavConfigClient {
	return &pubDavConfigClient{cc}
}

func (c *pubDavConfigClient) Get(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error) {
	out := new(DavConfig)
	err := c.cc.Invoke(ctx, PubDavConfig_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubDavConfigClient) Update(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error) {
	out := new(DavConfig)
	err := c.cc.Invoke(ctx, PubDavConfig_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubDavConfigClient) Enable(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error) {
	out := new(DavConfig)
	err := c.cc.Invoke(ctx, PubDavConfig_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubDavConfigClient) Disable(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*DavConfig, error) {
	out := new(DavConfig)
	err := c.cc.Invoke(ctx, PubDavConfig_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubDavConfigClient) ValidateUserName(ctx context.Context, in *DavConfig, opts ...grpc.CallOption) (*common.UserNameValidateResponse, error) {
	out := new(common.UserNameValidateResponse)
	err := c.cc.Invoke(ctx, PubDavConfig_ValidateUserName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubDavConfigServer is the server API for PubDavConfig service.
// All implementations must embed UnimplementedPubDavConfigServer
// for forward compatibility
type PubDavConfigServer interface {
	// rpc Create (DavConfig) returns (DavConfig) {}
	Get(context.Context, *DavConfig) (*DavConfig, error)
	Update(context.Context, *DavConfig) (*DavConfig, error)
	Enable(context.Context, *DavConfig) (*DavConfig, error)
	Disable(context.Context, *DavConfig) (*DavConfig, error)
	ValidateUserName(context.Context, *DavConfig) (*common.UserNameValidateResponse, error)
	mustEmbedUnimplementedPubDavConfigServer()
}

// UnimplementedPubDavConfigServer must be embedded to have forward compatible implementations.
type UnimplementedPubDavConfigServer struct {
}

func (UnimplementedPubDavConfigServer) Get(context.Context, *DavConfig) (*DavConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubDavConfigServer) Update(context.Context, *DavConfig) (*DavConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPubDavConfigServer) Enable(context.Context, *DavConfig) (*DavConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedPubDavConfigServer) Disable(context.Context, *DavConfig) (*DavConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedPubDavConfigServer) ValidateUserName(context.Context, *DavConfig) (*common.UserNameValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserName not implemented")
}
func (UnimplementedPubDavConfigServer) mustEmbedUnimplementedPubDavConfigServer() {}

// UnsafePubDavConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubDavConfigServer will
// result in compilation errors.
type UnsafePubDavConfigServer interface {
	mustEmbedUnimplementedPubDavConfigServer()
}

func RegisterPubDavConfigServer(s grpc.ServiceRegistrar, srv PubDavConfigServer) {
	s.RegisterService(&PubDavConfig_ServiceDesc, srv)
}

func _PubDavConfig_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DavConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubDavConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubDavConfig_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubDavConfigServer).Get(ctx, req.(*DavConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubDavConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DavConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubDavConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubDavConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubDavConfigServer).Update(ctx, req.(*DavConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubDavConfig_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DavConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubDavConfigServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubDavConfig_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubDavConfigServer).Enable(ctx, req.(*DavConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubDavConfig_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DavConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubDavConfigServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubDavConfig_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubDavConfigServer).Disable(ctx, req.(*DavConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubDavConfig_ValidateUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DavConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubDavConfigServer).ValidateUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubDavConfig_ValidateUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubDavConfigServer).ValidateUserName(ctx, req.(*DavConfig))
	}
	return interceptor(ctx, in, info, handler)
}

// PubDavConfig_ServiceDesc is the grpc.ServiceDesc for PubDavConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubDavConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubDavConfig",
	HandlerType: (*PubDavConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PubDavConfig_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PubDavConfig_Update_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _PubDavConfig_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _PubDavConfig_Disable_Handler,
		},
		{
			MethodName: "ValidateUserName",
			Handler:    _PubDavConfig_ValidateUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webdav/public_webdav.proto",
}
