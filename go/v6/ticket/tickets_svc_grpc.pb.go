// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: ticket/tickets_svc.proto

package ticket

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
	PubTicket_Get_FullMethodName               = "/v6.services.pub.PubTicket/Get"
	PubTicket_List_FullMethodName              = "/v6.services.pub.PubTicket/List"
	PubTicket_Create_FullMethodName            = "/v6.services.pub.PubTicket/Create"
	PubTicket_GetContactInfo_FullMethodName    = "/v6.services.pub.PubTicket/GetContactInfo"
	PubTicket_UpdateContactInfo_FullMethodName = "/v6.services.pub.PubTicket/UpdateContactInfo"
	PubTicket_ListConversation_FullMethodName  = "/v6.services.pub.PubTicket/ListConversation"
	PubTicket_Comment_FullMethodName           = "/v6.services.pub.PubTicket/Comment"
)

// PubTicketClient is the client API for PubTicket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubTicketClient interface {
	Get(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error)
	List(ctx context.Context, in *TicketListRequest, opts ...grpc.CallOption) (*TicketListResponse, error)
	Create(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error)
	GetContactInfo(ctx context.Context, in *ContactInfo, opts ...grpc.CallOption) (*ContactInfo, error)
	UpdateContactInfo(ctx context.Context, in *ContactInfo, opts ...grpc.CallOption) (*ContactInfo, error)
	ListConversation(ctx context.Context, in *ConversationListRequest, opts ...grpc.CallOption) (*ConversationListResponse, error)
	Comment(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error)
}

type pubTicketClient struct {
	cc grpc.ClientConnInterface
}

func NewPubTicketClient(cc grpc.ClientConnInterface) PubTicketClient {
	return &pubTicketClient{cc}
}

func (c *pubTicketClient) Get(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ticket)
	err := c.cc.Invoke(ctx, PubTicket_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubTicketClient) List(ctx context.Context, in *TicketListRequest, opts ...grpc.CallOption) (*TicketListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketListResponse)
	err := c.cc.Invoke(ctx, PubTicket_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubTicketClient) Create(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ticket)
	err := c.cc.Invoke(ctx, PubTicket_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubTicketClient) GetContactInfo(ctx context.Context, in *ContactInfo, opts ...grpc.CallOption) (*ContactInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContactInfo)
	err := c.cc.Invoke(ctx, PubTicket_GetContactInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubTicketClient) UpdateContactInfo(ctx context.Context, in *ContactInfo, opts ...grpc.CallOption) (*ContactInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContactInfo)
	err := c.cc.Invoke(ctx, PubTicket_UpdateContactInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubTicketClient) ListConversation(ctx context.Context, in *ConversationListRequest, opts ...grpc.CallOption) (*ConversationListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConversationListResponse)
	err := c.cc.Invoke(ctx, PubTicket_ListConversation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubTicketClient) Comment(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ticket)
	err := c.cc.Invoke(ctx, PubTicket_Comment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubTicketServer is the server API for PubTicket service.
// All implementations must embed UnimplementedPubTicketServer
// for forward compatibility.
type PubTicketServer interface {
	Get(context.Context, *Ticket) (*Ticket, error)
	List(context.Context, *TicketListRequest) (*TicketListResponse, error)
	Create(context.Context, *Ticket) (*Ticket, error)
	GetContactInfo(context.Context, *ContactInfo) (*ContactInfo, error)
	UpdateContactInfo(context.Context, *ContactInfo) (*ContactInfo, error)
	ListConversation(context.Context, *ConversationListRequest) (*ConversationListResponse, error)
	Comment(context.Context, *Ticket) (*Ticket, error)
	mustEmbedUnimplementedPubTicketServer()
}

// UnimplementedPubTicketServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPubTicketServer struct{}

func (UnimplementedPubTicketServer) Get(context.Context, *Ticket) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubTicketServer) List(context.Context, *TicketListRequest) (*TicketListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPubTicketServer) Create(context.Context, *Ticket) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPubTicketServer) GetContactInfo(context.Context, *ContactInfo) (*ContactInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactInfo not implemented")
}
func (UnimplementedPubTicketServer) UpdateContactInfo(context.Context, *ContactInfo) (*ContactInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContactInfo not implemented")
}
func (UnimplementedPubTicketServer) ListConversation(context.Context, *ConversationListRequest) (*ConversationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConversation not implemented")
}
func (UnimplementedPubTicketServer) Comment(context.Context, *Ticket) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedPubTicketServer) mustEmbedUnimplementedPubTicketServer() {}
func (UnimplementedPubTicketServer) testEmbeddedByValue()                   {}

// UnsafePubTicketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubTicketServer will
// result in compilation errors.
type UnsafePubTicketServer interface {
	mustEmbedUnimplementedPubTicketServer()
}

func RegisterPubTicketServer(s grpc.ServiceRegistrar, srv PubTicketServer) {
	// If the following call pancis, it indicates UnimplementedPubTicketServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PubTicket_ServiceDesc, srv)
}

func _PubTicket_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).Get(ctx, req.(*Ticket))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubTicket_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).List(ctx, req.(*TicketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubTicket_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).Create(ctx, req.(*Ticket))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubTicket_GetContactInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).GetContactInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_GetContactInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).GetContactInfo(ctx, req.(*ContactInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubTicket_UpdateContactInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).UpdateContactInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_UpdateContactInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).UpdateContactInfo(ctx, req.(*ContactInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubTicket_ListConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).ListConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_ListConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).ListConversation(ctx, req.(*ConversationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubTicket_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubTicketServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubTicket_Comment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubTicketServer).Comment(ctx, req.(*Ticket))
	}
	return interceptor(ctx, in, info, handler)
}

// PubTicket_ServiceDesc is the grpc.ServiceDesc for PubTicket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubTicket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubTicket",
	HandlerType: (*PubTicketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PubTicket_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PubTicket_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PubTicket_Create_Handler,
		},
		{
			MethodName: "GetContactInfo",
			Handler:    _PubTicket_GetContactInfo_Handler,
		},
		{
			MethodName: "UpdateContactInfo",
			Handler:    _PubTicket_UpdateContactInfo_Handler,
		},
		{
			MethodName: "ListConversation",
			Handler:    _PubTicket_ListConversation_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _PubTicket_Comment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket/tickets_svc.proto",
}
