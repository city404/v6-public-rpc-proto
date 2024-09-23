// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: user/user_svc.proto

package user

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
	PubUser_Get_FullMethodName                      = "/v6.services.pub.PubUser/Get"
	PubUser_Login_FullMethodName                    = "/v6.services.pub.PubUser/Login"
	PubUser_Refresh_FullMethodName                  = "/v6.services.pub.PubUser/Refresh"
	PubUser_Logoff_FullMethodName                   = "/v6.services.pub.PubUser/Logoff"
	PubUser_ResetPassword_FullMethodName            = "/v6.services.pub.PubUser/ResetPassword"
	PubUser_ChangePassword_FullMethodName           = "/v6.services.pub.PubUser/ChangePassword"
	PubUser_Register_FullMethodName                 = "/v6.services.pub.PubUser/Register"
	PubUser_Delete_FullMethodName                   = "/v6.services.pub.PubUser/Delete"
	PubUser_Update_FullMethodName                   = "/v6.services.pub.PubUser/Update"
	PubUser_SendSmsVerifyCode_FullMethodName        = "/v6.services.pub.PubUser/SendSmsVerifyCode"
	PubUser_SendSmsVerifyCodeNotUser_FullMethodName = "/v6.services.pub.PubUser/SendSmsVerifyCodeNotUser"
	PubUser_VerifyAuthToken_FullMethodName          = "/v6.services.pub.PubUser/VerifyAuthToken"
	PubUser_CreateAuthToken_FullMethodName          = "/v6.services.pub.PubUser/CreateAuthToken"
	PubUser_ValidateUserInfo_FullMethodName         = "/v6.services.pub.PubUser/ValidateUserInfo"
)

// PubUserClient is the client API for PubUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubUserClient interface {
	Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	Logoff(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error)
	ResetPassword(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*User, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*User, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	SendSmsVerifyCode(ctx context.Context, in *SmsVeifyCodeSendRequest, opts ...grpc.CallOption) (*SmsVeifyCodeSendResponse, error)
	SendSmsVerifyCodeNotUser(ctx context.Context, in *SmsVeifyCodeSendRequestNotUser, opts ...grpc.CallOption) (*SmsVeifyCodeSendResponse, error)
	VerifyAuthToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*OauthTokenCheckResponse, error)
	CreateAuthToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*OauthTokenResponse, error)
	ValidateUserInfo(ctx context.Context, in *UserValidateInfo, opts ...grpc.CallOption) (*common.UserNameValidateResponse, error)
}

type pubUserClient struct {
	cc grpc.ClientConnInterface
}

func NewPubUserClient(cc grpc.ClientConnInterface) PubUserClient {
	return &pubUserClient{cc}
}

func (c *pubUserClient) Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PubUser_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Refresh(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Token)
	err := c.cc.Invoke(ctx, PubUser_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Logoff(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Logoff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) ResetPassword(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Delete(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, PubUser_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) SendSmsVerifyCode(ctx context.Context, in *SmsVeifyCodeSendRequest, opts ...grpc.CallOption) (*SmsVeifyCodeSendResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmsVeifyCodeSendResponse)
	err := c.cc.Invoke(ctx, PubUser_SendSmsVerifyCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) SendSmsVerifyCodeNotUser(ctx context.Context, in *SmsVeifyCodeSendRequestNotUser, opts ...grpc.CallOption) (*SmsVeifyCodeSendResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SmsVeifyCodeSendResponse)
	err := c.cc.Invoke(ctx, PubUser_SendSmsVerifyCodeNotUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) VerifyAuthToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*OauthTokenCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OauthTokenCheckResponse)
	err := c.cc.Invoke(ctx, PubUser_VerifyAuthToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) CreateAuthToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*OauthTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OauthTokenResponse)
	err := c.cc.Invoke(ctx, PubUser_CreateAuthToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubUserClient) ValidateUserInfo(ctx context.Context, in *UserValidateInfo, opts ...grpc.CallOption) (*common.UserNameValidateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.UserNameValidateResponse)
	err := c.cc.Invoke(ctx, PubUser_ValidateUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubUserServer is the server API for PubUser service.
// All implementations must embed UnimplementedPubUserServer
// for forward compatibility.
type PubUserServer interface {
	Get(context.Context, *User) (*User, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *Token) (*Token, error)
	Logoff(context.Context, *Token) (*User, error)
	ResetPassword(context.Context, *LoginRequest) (*User, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*User, error)
	Register(context.Context, *RegisterRequest) (*User, error)
	Delete(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	SendSmsVerifyCode(context.Context, *SmsVeifyCodeSendRequest) (*SmsVeifyCodeSendResponse, error)
	SendSmsVerifyCodeNotUser(context.Context, *SmsVeifyCodeSendRequestNotUser) (*SmsVeifyCodeSendResponse, error)
	VerifyAuthToken(context.Context, *LoginRequest) (*OauthTokenCheckResponse, error)
	CreateAuthToken(context.Context, *LoginRequest) (*OauthTokenResponse, error)
	ValidateUserInfo(context.Context, *UserValidateInfo) (*common.UserNameValidateResponse, error)
	mustEmbedUnimplementedPubUserServer()
}

// UnimplementedPubUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPubUserServer struct{}

func (UnimplementedPubUserServer) Get(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPubUserServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPubUserServer) Refresh(context.Context, *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedPubUserServer) Logoff(context.Context, *Token) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logoff not implemented")
}
func (UnimplementedPubUserServer) ResetPassword(context.Context, *LoginRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedPubUserServer) ChangePassword(context.Context, *ChangePasswordRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedPubUserServer) Register(context.Context, *RegisterRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedPubUserServer) Delete(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPubUserServer) Update(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPubUserServer) SendSmsVerifyCode(context.Context, *SmsVeifyCodeSendRequest) (*SmsVeifyCodeSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsVerifyCode not implemented")
}
func (UnimplementedPubUserServer) SendSmsVerifyCodeNotUser(context.Context, *SmsVeifyCodeSendRequestNotUser) (*SmsVeifyCodeSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsVerifyCodeNotUser not implemented")
}
func (UnimplementedPubUserServer) VerifyAuthToken(context.Context, *LoginRequest) (*OauthTokenCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAuthToken not implemented")
}
func (UnimplementedPubUserServer) CreateAuthToken(context.Context, *LoginRequest) (*OauthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthToken not implemented")
}
func (UnimplementedPubUserServer) ValidateUserInfo(context.Context, *UserValidateInfo) (*common.UserNameValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserInfo not implemented")
}
func (UnimplementedPubUserServer) mustEmbedUnimplementedPubUserServer() {}
func (UnimplementedPubUserServer) testEmbeddedByValue()                 {}

// UnsafePubUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubUserServer will
// result in compilation errors.
type UnsafePubUserServer interface {
	mustEmbedUnimplementedPubUserServer()
}

func RegisterPubUserServer(s grpc.ServiceRegistrar, srv PubUserServer) {
	// If the following call pancis, it indicates UnimplementedPubUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PubUser_ServiceDesc, srv)
}

func _PubUser_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Get(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Refresh(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Logoff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Logoff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Logoff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Logoff(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).ResetPassword(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Delete(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_SendSmsVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsVeifyCodeSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).SendSmsVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_SendSmsVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).SendSmsVerifyCode(ctx, req.(*SmsVeifyCodeSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_SendSmsVerifyCodeNotUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsVeifyCodeSendRequestNotUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).SendSmsVerifyCodeNotUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_SendSmsVerifyCodeNotUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).SendSmsVerifyCodeNotUser(ctx, req.(*SmsVeifyCodeSendRequestNotUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_VerifyAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).VerifyAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_VerifyAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).VerifyAuthToken(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_CreateAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).CreateAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_CreateAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).CreateAuthToken(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubUser_ValidateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserValidateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubUserServer).ValidateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubUser_ValidateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubUserServer).ValidateUserInfo(ctx, req.(*UserValidateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// PubUser_ServiceDesc is the grpc.ServiceDesc for PubUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v6.services.pub.PubUser",
	HandlerType: (*PubUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PubUser_Get_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _PubUser_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _PubUser_Refresh_Handler,
		},
		{
			MethodName: "Logoff",
			Handler:    _PubUser_Logoff_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _PubUser_ResetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _PubUser_ChangePassword_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PubUser_Register_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PubUser_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PubUser_Update_Handler,
		},
		{
			MethodName: "SendSmsVerifyCode",
			Handler:    _PubUser_SendSmsVerifyCode_Handler,
		},
		{
			MethodName: "SendSmsVerifyCodeNotUser",
			Handler:    _PubUser_SendSmsVerifyCodeNotUser_Handler,
		},
		{
			MethodName: "VerifyAuthToken",
			Handler:    _PubUser_VerifyAuthToken_Handler,
		},
		{
			MethodName: "CreateAuthToken",
			Handler:    _PubUser_CreateAuthToken_Handler,
		},
		{
			MethodName: "ValidateUserInfo",
			Handler:    _PubUser_ValidateUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user_svc.proto",
}
