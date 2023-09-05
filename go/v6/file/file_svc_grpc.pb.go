// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: file/file_svc.proto

package file

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
	File_GetSlice_FullMethodName      = "/v6.services.pub.File/GetSlice"
	File_CreateSlice_FullMethodName   = "/v6.services.pub.File/CreateSlice"
	File_GetMeta_FullMethodName       = "/v6.services.pub.File/GetMeta"
	File_CreateMeta_FullMethodName    = "/v6.services.pub.File/CreateMeta"
	File_AddFastLookup_FullMethodName = "/v6.services.pub.File/AddFastLookup"
	File_GetFastLookup_FullMethodName = "/v6.services.pub.File/GetFastLookup"
)

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	GetSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error)
	CreateSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error)
	GetMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error)
	CreateMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error)
	AddFastLookup(ctx context.Context, in *FastLookup, opts ...grpc.CallOption) (*FastLookup, error)
	GetFastLookup(ctx context.Context, in *FastLookupRequest, opts ...grpc.CallOption) (*FastLookup, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) GetSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error) {
	out := new(Slice)
	err := c.cc.Invoke(ctx, File_GetSlice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) CreateSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error) {
	out := new(Slice)
	err := c.cc.Invoke(ctx, File_CreateSlice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) GetMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error) {
	out := new(Meta)
	err := c.cc.Invoke(ctx, File_GetMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) CreateMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error) {
	out := new(Meta)
	err := c.cc.Invoke(ctx, File_CreateMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) AddFastLookup(ctx context.Context, in *FastLookup, opts ...grpc.CallOption) (*FastLookup, error) {
	out := new(FastLookup)
	err := c.cc.Invoke(ctx, File_AddFastLookup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) GetFastLookup(ctx context.Context, in *FastLookupRequest, opts ...grpc.CallOption) (*FastLookup, error) {
	out := new(FastLookup)
	err := c.cc.Invoke(ctx, File_GetFastLookup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility
type FileServer interface {
	GetSlice(context.Context, *Slice) (*Slice, error)
	CreateSlice(context.Context, *Slice) (*Slice, error)
	GetMeta(context.Context, *Meta) (*Meta, error)
	CreateMeta(context.Context, *Meta) (*Meta, error)
	AddFastLookup(context.Context, *FastLookup) (*FastLookup, error)
	GetFastLookup(context.Context, *FastLookupRequest) (*FastLookup, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (UnimplementedFileServer) GetSlice(context.Context, *Slice) (*Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlice not implemented")
}
func (UnimplementedFileServer) CreateSlice(context.Context, *Slice) (*Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlice not implemented")
}
func (UnimplementedFileServer) GetMeta(context.Context, *Meta) (*Meta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeta not implemented")
}
func (UnimplementedFileServer) CreateMeta(context.Context, *Meta) (*Meta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeta not implemented")
}
func (UnimplementedFileServer) AddFastLookup(context.Context, *FastLookup) (*FastLookup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFastLookup not implemented")
}
func (UnimplementedFileServer) GetFastLookup(context.Context, *FastLookupRequest) (*FastLookup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFastLookup not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_GetSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_GetSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetSlice(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_CreateSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).CreateSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_CreateSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).CreateSlice(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_GetMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetMeta(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_CreateMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).CreateMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_CreateMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).CreateMeta(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_AddFastLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastLookup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AddFastLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_AddFastLookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AddFastLookup(ctx, req.(*FastLookup))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_GetFastLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetFastLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_GetFastLookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetFastLookup(ctx, req.(*FastLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSlice",
			Handler:    _File_GetSlice_Handler,
		},
		{
			MethodName: "CreateSlice",
			Handler:    _File_CreateSlice_Handler,
		},
		{
			MethodName: "GetMeta",
			Handler:    _File_GetMeta_Handler,
		},
		{
			MethodName: "CreateMeta",
			Handler:    _File_CreateMeta_Handler,
		},
		{
			MethodName: "AddFastLookup",
			Handler:    _File_AddFastLookup_Handler,
		},
		{
			MethodName: "GetFastLookup",
			Handler:    _File_GetFastLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/file_svc.proto",
}
