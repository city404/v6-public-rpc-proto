// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: userfile/user_file.proto

package userfile

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
	PubUserFile_Create_FullMethodName = "/v6.services.pub.PubUserFile/Create"
	PubUserFile_Get_FullMethodName    = "/v6.services.pub.PubUserFile/Get"
)

// PubUserFileClient is the client API for PubUserFile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubUserFileClient interface {
	Create(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
	Get(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
}

type pubUserFileClient struct {
	cc grpc.ClientConnInterface
}

func NewPubUserFileClient(cc grpc.ClientConnInterface) PubUserFileClient {
	return &pubUserFileClient{cc}
}

func (c *pubUserFileClient) Create(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, PubUserFile_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Get(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, PubUserFile_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubUserFileServer is the server API for PubUserFile service.
// All implementations must embed UnimplementedPubUserFileServer
// for forward compatibility
type PubUserFileServer interface {
	Create(context.Context, *File) (*File, error)
	Get(context.Context, *File) (*File, error)
	mustEmbedUnimplementedPubUserFileServer()
}

// UnimplementedPubUserFileServer must be embedded to have forward compatible implementations.
type UnimplementedPubUserFileServer struct {
}

func (UnimplementedPubUserFileServer) Create(context.Context, *File) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPubUserFileServer) Get(context.Context, *File) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubUserFileServer) mustEmbedUnimplementedPubUserFileServer() {}

// UnsafePubUserFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubUserFileServer will
// result in compilation errors.
type UnsafePubUserFileServer interface {
	mustEmbedUnimplementedPubUserFileServer()
}

func RegisterPubUserFileServer(s grpc.ServiceRegistrar, srv PubUserFileServer) {
	s.RegisterService(&PubUserFile_ServiceDesc, srv)
}

func _PubUserFile_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Create(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Get(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

// PubUserFile_ServiceDesc is the grpc.ServiceDesc for PubUserFile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubUserFile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubUserFile",
	HandlerType: (*PubUserFileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PubUserFile_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PubUserFile_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userfile/user_file.proto",
}
