// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: share/file_share.proto

package fileshare

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
	PubFileShare_Create_FullMethodName          = "/v6.services.pub.PubFileShare/Create"
	PubFileShare_Get_FullMethodName             = "/v6.services.pub.PubFileShare/Get"
	PubFileShare_Like_FullMethodName            = "/v6.services.pub.PubFileShare/Like"
	PubFileShare_Dislike_FullMethodName         = "/v6.services.pub.PubFileShare/Dislike"
	PubFileShare_ComplaintShare_FullMethodName  = "/v6.services.pub.PubFileShare/ComplaintShare"
	PubFileShare_List_FullMethodName            = "/v6.services.pub.PubFileShare/List"
	PubFileShare_ListComplaint_FullMethodName   = "/v6.services.pub.PubFileShare/ListComplaint"
	PubFileShare_Delete_FullMethodName          = "/v6.services.pub.PubFileShare/Delete"
	PubFileShare_DeleteComplaint_FullMethodName = "/v6.services.pub.PubFileShare/DeleteComplaint"
	PubFileShare_Save_FullMethodName            = "/v6.services.pub.PubFileShare/Save"
)

// PubFileShareClient is the client API for PubFileShare service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubFileShareClient interface {
	Create(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error)
	Get(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error)
	Like(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error)
	Dislike(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error)
	ComplaintShare(ctx context.Context, in *Complaint, opts ...grpc.CallOption) (*Complaint, error)
	List(ctx context.Context, in *FileShareListRequest, opts ...grpc.CallOption) (*FileShareListResponse, error)
	ListComplaint(ctx context.Context, in *ComplaintListRequest, opts ...grpc.CallOption) (*ComplaintListResponse, error)
	Delete(ctx context.Context, in *FileShareDeleteRequest, opts ...grpc.CallOption) (*FileShareDeleteResponse, error)
	DeleteComplaint(ctx context.Context, in *ComplaintDeleteRequest, opts ...grpc.CallOption) (*ComplaintDeleteResponse, error)
	Save(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error)
}

type pubFileShareClient struct {
	cc grpc.ClientConnInterface
}

func NewPubFileShareClient(cc grpc.ClientConnInterface) PubFileShareClient {
	return &pubFileShareClient{cc}
}

func (c *pubFileShareClient) Create(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShare)
	err := c.cc.Invoke(ctx, PubFileShare_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) Get(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShare)
	err := c.cc.Invoke(ctx, PubFileShare_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) Like(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShare)
	err := c.cc.Invoke(ctx, PubFileShare_Like_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) Dislike(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShare)
	err := c.cc.Invoke(ctx, PubFileShare_Dislike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) ComplaintShare(ctx context.Context, in *Complaint, opts ...grpc.CallOption) (*Complaint, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Complaint)
	err := c.cc.Invoke(ctx, PubFileShare_ComplaintShare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) List(ctx context.Context, in *FileShareListRequest, opts ...grpc.CallOption) (*FileShareListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShareListResponse)
	err := c.cc.Invoke(ctx, PubFileShare_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) ListComplaint(ctx context.Context, in *ComplaintListRequest, opts ...grpc.CallOption) (*ComplaintListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ComplaintListResponse)
	err := c.cc.Invoke(ctx, PubFileShare_ListComplaint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) Delete(ctx context.Context, in *FileShareDeleteRequest, opts ...grpc.CallOption) (*FileShareDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShareDeleteResponse)
	err := c.cc.Invoke(ctx, PubFileShare_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) DeleteComplaint(ctx context.Context, in *ComplaintDeleteRequest, opts ...grpc.CallOption) (*ComplaintDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ComplaintDeleteResponse)
	err := c.cc.Invoke(ctx, PubFileShare_DeleteComplaint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubFileShareClient) Save(ctx context.Context, in *FileShare, opts ...grpc.CallOption) (*FileShare, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileShare)
	err := c.cc.Invoke(ctx, PubFileShare_Save_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubFileShareServer is the server API for PubFileShare service.
// All implementations must embed UnimplementedPubFileShareServer
// for forward compatibility.
type PubFileShareServer interface {
	Create(context.Context, *FileShare) (*FileShare, error)
	Get(context.Context, *FileShare) (*FileShare, error)
	Like(context.Context, *FileShare) (*FileShare, error)
	Dislike(context.Context, *FileShare) (*FileShare, error)
	ComplaintShare(context.Context, *Complaint) (*Complaint, error)
	List(context.Context, *FileShareListRequest) (*FileShareListResponse, error)
	ListComplaint(context.Context, *ComplaintListRequest) (*ComplaintListResponse, error)
	Delete(context.Context, *FileShareDeleteRequest) (*FileShareDeleteResponse, error)
	DeleteComplaint(context.Context, *ComplaintDeleteRequest) (*ComplaintDeleteResponse, error)
	Save(context.Context, *FileShare) (*FileShare, error)
	mustEmbedUnimplementedPubFileShareServer()
}

// UnimplementedPubFileShareServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPubFileShareServer struct{}

func (UnimplementedPubFileShareServer) Create(context.Context, *FileShare) (*FileShare, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPubFileShareServer) Get(context.Context, *FileShare) (*FileShare, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubFileShareServer) Like(context.Context, *FileShare) (*FileShare, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedPubFileShareServer) Dislike(context.Context, *FileShare) (*FileShare, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dislike not implemented")
}
func (UnimplementedPubFileShareServer) ComplaintShare(context.Context, *Complaint) (*Complaint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComplaintShare not implemented")
}
func (UnimplementedPubFileShareServer) List(context.Context, *FileShareListRequest) (*FileShareListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPubFileShareServer) ListComplaint(context.Context, *ComplaintListRequest) (*ComplaintListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComplaint not implemented")
}
func (UnimplementedPubFileShareServer) Delete(context.Context, *FileShareDeleteRequest) (*FileShareDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPubFileShareServer) DeleteComplaint(context.Context, *ComplaintDeleteRequest) (*ComplaintDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComplaint not implemented")
}
func (UnimplementedPubFileShareServer) Save(context.Context, *FileShare) (*FileShare, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedPubFileShareServer) mustEmbedUnimplementedPubFileShareServer() {}
func (UnimplementedPubFileShareServer) testEmbeddedByValue()                      {}

// UnsafePubFileShareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubFileShareServer will
// result in compilation errors.
type UnsafePubFileShareServer interface {
	mustEmbedUnimplementedPubFileShareServer()
}

func RegisterPubFileShareServer(s grpc.ServiceRegistrar, srv PubFileShareServer) {
	// If the following call pancis, it indicates UnimplementedPubFileShareServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PubFileShare_ServiceDesc, srv)
}

func _PubFileShare_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).Create(ctx, req.(*FileShare))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).Get(ctx, req.(*FileShare))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).Like(ctx, req.(*FileShare))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_Dislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).Dislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_Dislike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).Dislike(ctx, req.(*FileShare))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_ComplaintShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Complaint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).ComplaintShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_ComplaintShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).ComplaintShare(ctx, req.(*Complaint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShareListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).List(ctx, req.(*FileShareListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_ListComplaint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplaintListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).ListComplaint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_ListComplaint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).ListComplaint(ctx, req.(*ComplaintListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShareDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).Delete(ctx, req.(*FileShareDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_DeleteComplaint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplaintDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).DeleteComplaint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_DeleteComplaint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).DeleteComplaint(ctx, req.(*ComplaintDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubFileShare_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubFileShareServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubFileShare_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubFileShareServer).Save(ctx, req.(*FileShare))
	}
	return interceptor(ctx, in, info, handler)
}

// PubFileShare_ServiceDesc is the grpc.ServiceDesc for PubFileShare service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubFileShare_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubFileShare",
	HandlerType: (*PubFileShareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PubFileShare_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PubFileShare_Get_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _PubFileShare_Like_Handler,
		},
		{
			MethodName: "Dislike",
			Handler:    _PubFileShare_Dislike_Handler,
		},
		{
			MethodName: "ComplaintShare",
			Handler:    _PubFileShare_ComplaintShare_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PubFileShare_List_Handler,
		},
		{
			MethodName: "ListComplaint",
			Handler:    _PubFileShare_ListComplaint_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PubFileShare_Delete_Handler,
		},
		{
			MethodName: "DeleteComplaint",
			Handler:    _PubFileShare_DeleteComplaint_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _PubFileShare_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "share/file_share.proto",
}
