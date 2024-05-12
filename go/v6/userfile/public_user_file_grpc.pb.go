// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
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
	PubUserFile_Create_FullMethodName                    = "/v6.services.pub.PubUserFile/Create"
	PubUserFile_Get_FullMethodName                       = "/v6.services.pub.PubUserFile/Get"
	PubUserFile_Rename_FullMethodName                    = "/v6.services.pub.PubUserFile/Rename"
	PubUserFile_Trash_FullMethodName                     = "/v6.services.pub.PubUserFile/Trash"
	PubUserFile_Move_FullMethodName                      = "/v6.services.pub.PubUserFile/Move"
	PubUserFile_Copy_FullMethodName                      = "/v6.services.pub.PubUserFile/Copy"
	PubUserFile_Delete_FullMethodName                    = "/v6.services.pub.PubUserFile/Delete"
	PubUserFile_DeleteTrash_FullMethodName               = "/v6.services.pub.PubUserFile/DeleteTrash"
	PubUserFile_Recover_FullMethodName                   = "/v6.services.pub.PubUserFile/Recover"
	PubUserFile_BatchOperation_FullMethodName            = "/v6.services.pub.PubUserFile/BatchOperation"
	PubUserFile_List_FullMethodName                      = "/v6.services.pub.PubUserFile/List"
	PubUserFile_ListTrash_FullMethodName                 = "/v6.services.pub.PubUserFile/ListTrash"
	PubUserFile_Search_FullMethodName                    = "/v6.services.pub.PubUserFile/Search"
	PubUserFile_CreateManageRTCOffer_FullMethodName      = "/v6.services.pub.PubUserFile/CreateManageRTCOffer"
	PubUserFile_SendClientIceCandidate_FullMethodName    = "/v6.services.pub.PubUserFile/SendClientIceCandidate"
	PubUserFile_GetServerIceCandidate_FullMethodName     = "/v6.services.pub.PubUserFile/GetServerIceCandidate"
	PubUserFile_ParseFileSlice_FullMethodName            = "/v6.services.pub.PubUserFile/ParseFileSlice"
	PubUserFile_GetSliceDownloadAddress_FullMethodName   = "/v6.services.pub.PubUserFile/GetSliceDownloadAddress"
	PubUserFile_GetDownloadAndPreviewInfo_FullMethodName = "/v6.services.pub.PubUserFile/GetDownloadAndPreviewInfo"
	PubUserFile_PreviewDoc_FullMethodName                = "/v6.services.pub.PubUserFile/PreviewDoc"
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
	CreateManageRTCOffer(ctx context.Context, in *ManageRTCRequest, opts ...grpc.CallOption) (*ManageRTCResponse, error)
	SendClientIceCandidate(ctx context.Context, in *SendIceCandidateRequest, opts ...grpc.CallOption) (*SendIceCandidateResponse, error)
	GetServerIceCandidate(ctx context.Context, in *GetIceCandidateRequest, opts ...grpc.CallOption) (*GetIceCandidateResponse, error)
	ParseFileSlice(ctx context.Context, in *File, opts ...grpc.CallOption) (*ParseFileSliceResponse, error)
	GetSliceDownloadAddress(ctx context.Context, in *SliceDownloadAddressRequest, opts ...grpc.CallOption) (*SliceDownloadAddressResponse, error)
	GetDownloadAndPreviewInfo(ctx context.Context, in *File, opts ...grpc.CallOption) (*DownloadAndPreviewInfo, error)
	PreviewDoc(ctx context.Context, in *File, opts ...grpc.CallOption) (*DocFilePreview, error)
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

func (c *pubUserFileClient) CreateManageRTCOffer(ctx context.Context, in *ManageRTCRequest, opts ...grpc.CallOption) (*ManageRTCResponse, error) {
	out := new(ManageRTCResponse)
	err := c.cc.Invoke(ctx, PubUserFile_CreateManageRTCOffer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) SendClientIceCandidate(ctx context.Context, in *SendIceCandidateRequest, opts ...grpc.CallOption) (*SendIceCandidateResponse, error) {
	out := new(SendIceCandidateResponse)
	err := c.cc.Invoke(ctx, PubUserFile_SendClientIceCandidate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) GetServerIceCandidate(ctx context.Context, in *GetIceCandidateRequest, opts ...grpc.CallOption) (*GetIceCandidateResponse, error) {
	out := new(GetIceCandidateResponse)
	err := c.cc.Invoke(ctx, PubUserFile_GetServerIceCandidate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) ParseFileSlice(ctx context.Context, in *File, opts ...grpc.CallOption) (*ParseFileSliceResponse, error) {
	out := new(ParseFileSliceResponse)
	err := c.cc.Invoke(ctx, PubUserFile_ParseFileSlice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) GetSliceDownloadAddress(ctx context.Context, in *SliceDownloadAddressRequest, opts ...grpc.CallOption) (*SliceDownloadAddressResponse, error) {
	out := new(SliceDownloadAddressResponse)
	err := c.cc.Invoke(ctx, PubUserFile_GetSliceDownloadAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) GetDownloadAndPreviewInfo(ctx context.Context, in *File, opts ...grpc.CallOption) (*DownloadAndPreviewInfo, error) {
	out := new(DownloadAndPreviewInfo)
	err := c.cc.Invoke(ctx, PubUserFile_GetDownloadAndPreviewInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserFileClient) PreviewDoc(ctx context.Context, in *File, opts ...grpc.CallOption) (*DocFilePreview, error) {
	out := new(DocFilePreview)
	err := c.cc.Invoke(ctx, PubUserFile_PreviewDoc_FullMethodName, in, out, opts...)
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
	CreateManageRTCOffer(context.Context, *ManageRTCRequest) (*ManageRTCResponse, error)
	SendClientIceCandidate(context.Context, *SendIceCandidateRequest) (*SendIceCandidateResponse, error)
	GetServerIceCandidate(context.Context, *GetIceCandidateRequest) (*GetIceCandidateResponse, error)
	ParseFileSlice(context.Context, *File) (*ParseFileSliceResponse, error)
	GetSliceDownloadAddress(context.Context, *SliceDownloadAddressRequest) (*SliceDownloadAddressResponse, error)
	GetDownloadAndPreviewInfo(context.Context, *File) (*DownloadAndPreviewInfo, error)
	PreviewDoc(context.Context, *File) (*DocFilePreview, error)
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
func (UnimplementedPubUserFileServer) CreateManageRTCOffer(context.Context, *ManageRTCRequest) (*ManageRTCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManageRTCOffer not implemented")
}
func (UnimplementedPubUserFileServer) SendClientIceCandidate(context.Context, *SendIceCandidateRequest) (*SendIceCandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClientIceCandidate not implemented")
}
func (UnimplementedPubUserFileServer) GetServerIceCandidate(context.Context, *GetIceCandidateRequest) (*GetIceCandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerIceCandidate not implemented")
}
func (UnimplementedPubUserFileServer) ParseFileSlice(context.Context, *File) (*ParseFileSliceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseFileSlice not implemented")
}
func (UnimplementedPubUserFileServer) GetSliceDownloadAddress(context.Context, *SliceDownloadAddressRequest) (*SliceDownloadAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliceDownloadAddress not implemented")
}
func (UnimplementedPubUserFileServer) GetDownloadAndPreviewInfo(context.Context, *File) (*DownloadAndPreviewInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadAndPreviewInfo not implemented")
}
func (UnimplementedPubUserFileServer) PreviewDoc(context.Context, *File) (*DocFilePreview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDoc not implemented")
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

func _PubUserFile_CreateManageRTCOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageRTCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).CreateManageRTCOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_CreateManageRTCOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).CreateManageRTCOffer(ctx, req.(*ManageRTCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_SendClientIceCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendIceCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).SendClientIceCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_SendClientIceCandidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).SendClientIceCandidate(ctx, req.(*SendIceCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_GetServerIceCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIceCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).GetServerIceCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_GetServerIceCandidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).GetServerIceCandidate(ctx, req.(*GetIceCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_ParseFileSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).ParseFileSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_ParseFileSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).ParseFileSlice(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_GetSliceDownloadAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SliceDownloadAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).GetSliceDownloadAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_GetSliceDownloadAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).GetSliceDownloadAddress(ctx, req.(*SliceDownloadAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_GetDownloadAndPreviewInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).GetDownloadAndPreviewInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_GetDownloadAndPreviewInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).GetDownloadAndPreviewInfo(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUserFile_PreviewDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserFileServer).PreviewDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUserFile_PreviewDoc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserFileServer).PreviewDoc(ctx, req.(*File))
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
			MethodName: "CreateManageRTCOffer",
			Handler:    _PubUserFile_CreateManageRTCOffer_Handler,
		},
		{
			MethodName: "SendClientIceCandidate",
			Handler:    _PubUserFile_SendClientIceCandidate_Handler,
		},
		{
			MethodName: "GetServerIceCandidate",
			Handler:    _PubUserFile_GetServerIceCandidate_Handler,
		},
		{
			MethodName: "ParseFileSlice",
			Handler:    _PubUserFile_ParseFileSlice_Handler,
		},
		{
			MethodName: "GetSliceDownloadAddress",
			Handler:    _PubUserFile_GetSliceDownloadAddress_Handler,
		},
		{
			MethodName: "GetDownloadAndPreviewInfo",
			Handler:    _PubUserFile_GetDownloadAndPreviewInfo_Handler,
		},
		{
			MethodName: "PreviewDoc",
			Handler:    _PubUserFile_PreviewDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userfile/public_user_file.proto",
}
