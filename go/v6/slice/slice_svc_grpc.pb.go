// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: slice/slice_svc.proto

package slice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PubSlice_GetSlice_FullMethodName      = "/v6.services.pub.PubSlice/GetSlice"
	PubSlice_CreateSlice_FullMethodName   = "/v6.services.pub.PubSlice/CreateSlice"
	PubSlice_GetMeta_FullMethodName       = "/v6.services.pub.PubSlice/GetMeta"
	PubSlice_CreateMeta_FullMethodName    = "/v6.services.pub.PubSlice/CreateMeta"
	PubSlice_AddFastLookup_FullMethodName = "/v6.services.pub.PubSlice/AddFastLookup"
	PubSlice_GetFastLookup_FullMethodName = "/v6.services.pub.PubSlice/GetFastLookup"
)

// PubSliceClient is the client API for PubSlice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSliceClient interface {
	GetSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error)
	CreateSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error)
	GetMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error)
	CreateMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error)
	AddFastLookup(ctx context.Context, in *FastLookup, opts ...grpc.CallOption) (*FastLookup, error)
	GetFastLookup(ctx context.Context, in *FastLookupRequest, opts ...grpc.CallOption) (*FastLookup, error)
}

type pubSliceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSliceClient(cc grpc.ClientConnInterface) PubSliceClient {
	return &pubSliceClient{cc}
}

func (c *pubSliceClient) GetSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Slice)
	err := c.cc.Invoke(ctx, PubSlice_GetSlice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSliceClient) CreateSlice(ctx context.Context, in *Slice, opts ...grpc.CallOption) (*Slice, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Slice)
	err := c.cc.Invoke(ctx, PubSlice_CreateSlice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSliceClient) GetMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Meta)
	err := c.cc.Invoke(ctx, PubSlice_GetMeta_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSliceClient) CreateMeta(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Meta)
	err := c.cc.Invoke(ctx, PubSlice_CreateMeta_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSliceClient) AddFastLookup(ctx context.Context, in *FastLookup, opts ...grpc.CallOption) (*FastLookup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FastLookup)
	err := c.cc.Invoke(ctx, PubSlice_AddFastLookup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSliceClient) GetFastLookup(ctx context.Context, in *FastLookupRequest, opts ...grpc.CallOption) (*FastLookup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FastLookup)
	err := c.cc.Invoke(ctx, PubSlice_GetFastLookup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSliceServer is the server API for PubSlice service.
// All implementations must embed UnimplementedPubSliceServer
// for forward compatibility.
type PubSliceServer interface {
	GetSlice(context.Context, *Slice) (*Slice, error)
	CreateSlice(context.Context, *Slice) (*Slice, error)
	GetMeta(context.Context, *Meta) (*Meta, error)
	CreateMeta(context.Context, *Meta) (*Meta, error)
	AddFastLookup(context.Context, *FastLookup) (*FastLookup, error)
	GetFastLookup(context.Context, *FastLookupRequest) (*FastLookup, error)
	mustEmbedUnimplementedPubSliceServer()
}

// UnimplementedPubSliceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPubSliceServer struct{}

func (UnimplementedPubSliceServer) GetSlice(context.Context, *Slice) (*Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlice not implemented")
}
func (UnimplementedPubSliceServer) CreateSlice(context.Context, *Slice) (*Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlice not implemented")
}
func (UnimplementedPubSliceServer) GetMeta(context.Context, *Meta) (*Meta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeta not implemented")
}
func (UnimplementedPubSliceServer) CreateMeta(context.Context, *Meta) (*Meta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeta not implemented")
}
func (UnimplementedPubSliceServer) AddFastLookup(context.Context, *FastLookup) (*FastLookup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFastLookup not implemented")
}
func (UnimplementedPubSliceServer) GetFastLookup(context.Context, *FastLookupRequest) (*FastLookup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFastLookup not implemented")
}
func (UnimplementedPubSliceServer) mustEmbedUnimplementedPubSliceServer() {}
func (UnimplementedPubSliceServer) testEmbeddedByValue()                  {}

// UnsafePubSliceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSliceServer will
// result in compilation errors.
type UnsafePubSliceServer interface {
	mustEmbedUnimplementedPubSliceServer()
}

func RegisterPubSliceServer(s grpc.ServiceRegistrar, srv PubSliceServer) {
	// If the following call pancis, it indicates UnimplementedPubSliceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PubSlice_ServiceDesc, srv)
}

func _PubSlice_GetSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSliceServer).GetSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSlice_GetSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSliceServer).GetSlice(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSlice_CreateSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSliceServer).CreateSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSlice_CreateSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSliceServer).CreateSlice(ctx, req.(*Slice))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSlice_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSliceServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSlice_GetMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSliceServer).GetMeta(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSlice_CreateMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSliceServer).CreateMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSlice_CreateMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSliceServer).CreateMeta(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSlice_AddFastLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastLookup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSliceServer).AddFastLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSlice_AddFastLookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSliceServer).AddFastLookup(ctx, req.(*FastLookup))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSlice_GetFastLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSliceServer).GetFastLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSlice_GetFastLookup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSliceServer).GetFastLookup(ctx, req.(*FastLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSlice_ServiceDesc is the grpc.ServiceDesc for PubSlice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSlice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubSlice",
	HandlerType: (*PubSliceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSlice",
			Handler:    _PubSlice_GetSlice_Handler,
		},
		{
			MethodName: "CreateSlice",
			Handler:    _PubSlice_CreateSlice_Handler,
		},
		{
			MethodName: "GetMeta",
			Handler:    _PubSlice_GetMeta_Handler,
		},
		{
			MethodName: "CreateMeta",
			Handler:    _PubSlice_CreateMeta_Handler,
		},
		{
			MethodName: "AddFastLookup",
			Handler:    _PubSlice_AddFastLookup_Handler,
		},
		{
			MethodName: "GetFastLookup",
			Handler:    _PubSlice_GetFastLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slice/slice_svc.proto",
}
