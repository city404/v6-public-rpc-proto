// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: userfile/public_user_file.proto

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
	PubUserFile_Create_FullMethodName              = "/v6.services.pub.PubUserFile/Create"
	PubUserFile_Get_FullMethodName                 = "/v6.services.pub.PubUserFile/Get"
	PubUserFile_Rename_FullMethodName              = "/v6.services.pub.PubUserFile/Rename"
	PubUserFile_Trash_FullMethodName               = "/v6.services.pub.PubUserFile/Trash"
	PubUserFile_Move_FullMethodName                = "/v6.services.pub.PubUserFile/Move"
	PubUserFile_Copy_FullMethodName                = "/v6.services.pub.PubUserFile/Copy"
	PubUserFile_Delete_FullMethodName              = "/v6.services.pub.PubUserFile/Delete"
	PubUserFile_DeleteTrash_FullMethodName         = "/v6.services.pub.PubUserFile/DeleteTrash"
	PubUserFile_Recover_FullMethodName             = "/v6.services.pub.PubUserFile/Recover"
	PubUserFile_BatchOperation_FullMethodName      = "/v6.services.pub.PubUserFile/BatchOperation"
	PubUserFile_List_FullMethodName                = "/v6.services.pub.PubUserFile/List"
	PubUserFile_ListTrash_FullMethodName           = "/v6.services.pub.PubUserFile/ListTrash"
	PubUserFile_Search_FullMethodName              = "/v6.services.pub.PubUserFile/Search"
	PubUserFile_CreateDownloadOffer_FullMethodName = "/v6.services.pub.PubUserFile/CreateDownloadOffer"
)

// PubUserFileClient is the client API for PubUserFile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubUserFileClient interface {
	Create(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
	Get(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
	// rpc Update (File) returns (File) {}
	Rename(ctx context.Context, in *File, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	Trash(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	Move(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	Copy(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	Delete(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	DeleteTrash(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	Recover(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	BatchOperation(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error)
	List(ctx context.Context, in *FileListRequest, opts ...grpc.CallOption) (*FileListResponse, error)
	ListTrash(ctx context.Context, in *FileListRequest, opts ...grpc.CallOption) (*FileListResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*FileListResponse, error)
	CreateDownloadOffer(ctx context.Context, in *RTCDownloadInfo, opts ...grpc.CallOption) (*RTCDownloadResponse, error)
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

func (c *pubUserFileClient) Rename(ctx context.Context, in *File, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Rename_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Trash(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Trash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Move(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Move_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Copy(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Copy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Delete(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) DeleteTrash(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_DeleteTrash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Recover(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Recover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) BatchOperation(ctx context.Context, in *BatchOperationRequest, opts ...grpc.CallOption) (*BatchOperationResponse, error) {
	out := new(BatchOperationResponse)
	err := c.cc.Invoke(ctx, PubUserFile_BatchOperation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) List(ctx context.Context, in *FileListRequest, opts ...grpc.CallOption) (*FileListResponse, error) {
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, PubUserFile_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) ListTrash(ctx context.Context, in *FileListRequest, opts ...grpc.CallOption) (*FileListResponse, error) {
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, PubUserFile_ListTrash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*FileListResponse, error) {
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, PubUserFile_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) CreateDownloadOffer(ctx context.Context, in *RTCDownloadInfo, opts ...grpc.CallOption) (*RTCDownloadResponse, error) {
	out := new(RTCDownloadResponse)
	err := c.cc.Invoke(ctx, PubUserFile_CreateDownloadOffer_FullMethodName, in, out, opts...)
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
	// rpc Update (File) returns (File) {}
	Rename(context.Context, *File) (*BatchOperationResponse, error)
	Trash(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	Move(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	Copy(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	Delete(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	DeleteTrash(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	Recover(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	BatchOperation(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error)
	List(context.Context, *FileListRequest) (*FileListResponse, error)
	ListTrash(context.Context, *FileListRequest) (*FileListResponse, error)
	Search(context.Context, *SearchRequest) (*FileListResponse, error)
	CreateDownloadOffer(context.Context, *RTCDownloadInfo) (*RTCDownloadResponse, error)
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
func (UnimplementedPubUserFileServer) Rename(context.Context, *File) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedPubUserFileServer) Trash(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trash not implemented")
}
func (UnimplementedPubUserFileServer) Move(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedPubUserFileServer) Copy(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedPubUserFileServer) Delete(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPubUserFileServer) DeleteTrash(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrash not implemented")
}
func (UnimplementedPubUserFileServer) Recover(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (UnimplementedPubUserFileServer) BatchOperation(context.Context, *BatchOperationRequest) (*BatchOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOperation not implemented")
}
func (UnimplementedPubUserFileServer) List(context.Context, *FileListRequest) (*FileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPubUserFileServer) ListTrash(context.Context, *FileListRequest) (*FileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrash not implemented")
}
func (UnimplementedPubUserFileServer) Search(context.Context, *SearchRequest) (*FileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedPubUserFileServer) CreateDownloadOffer(context.Context, *RTCDownloadInfo) (*RTCDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownloadOffer not implemented")
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

func _PubUserFile_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Rename_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Rename(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Trash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Trash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Trash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Trash(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Move(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Copy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Copy(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Delete(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_DeleteTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).DeleteTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_DeleteTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).DeleteTrash(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Recover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Recover(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_BatchOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).BatchOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_BatchOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).BatchOperation(ctx, req.(*BatchOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).List(ctx, req.(*FileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_ListTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).ListTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_ListTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).ListTrash(ctx, req.(*FileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_CreateDownloadOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RTCDownloadInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).CreateDownloadOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_CreateDownloadOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).CreateDownloadOffer(ctx, req.(*RTCDownloadInfo))
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
		{
			MethodName: "Rename",
			Handler:    _PubUserFile_Rename_Handler,
		},
		{
			MethodName: "Trash",
			Handler:    _PubUserFile_Trash_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _PubUserFile_Move_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _PubUserFile_Copy_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PubUserFile_Delete_Handler,
		},
		{
			MethodName: "DeleteTrash",
			Handler:    _PubUserFile_DeleteTrash_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _PubUserFile_Recover_Handler,
		},
		{
			MethodName: "BatchOperation",
			Handler:    _PubUserFile_BatchOperation_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PubUserFile_List_Handler,
		},
		{
			MethodName: "ListTrash",
			Handler:    _PubUserFile_ListTrash_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _PubUserFile_Search_Handler,
		},
		{
			MethodName: "CreateDownloadOffer",
			Handler:    _PubUserFile_CreateDownloadOffer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userfile/public_user_file.proto",
}
