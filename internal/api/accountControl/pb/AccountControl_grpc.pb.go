// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: AccountControl.proto

package pb

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
	AccountControl_RegistrationAccount_FullMethodName     = "/AccountControl/RegistrationAccount"
	AccountControl_AuthorizationAccount_FullMethodName    = "/AccountControl/AuthorizationAccount"
	AccountControl_ChangePasswordAccount_FullMethodName   = "/AccountControl/ChangePasswordAccount"
	AccountControl_DeleteAccount_FullMethodName           = "/AccountControl/DeleteAccount"
	AccountControl_IsAuthorizedSessionData_FullMethodName = "/AccountControl/IsAuthorizedSessionData"
	AccountControl_DeleteSessionData_FullMethodName       = "/AccountControl/DeleteSessionData"
)

// AccountControlClient is the client API for AccountControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountControlClient interface {
	RegistrationAccount(ctx context.Context, in *User, opts ...grpc.CallOption) (*SessionData, error)
	AuthorizationAccount(ctx context.Context, in *User, opts ...grpc.CallOption) (*SessionData, error)
	ChangePasswordAccount(ctx context.Context, in *ChangePasswordData, opts ...grpc.CallOption) (*Null, error)
	DeleteAccount(ctx context.Context, in *FullInfoUser, opts ...grpc.CallOption) (*Null, error)
	IsAuthorizedSessionData(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*AccountID, error)
	DeleteSessionData(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*Null, error)
}

type accountControlClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountControlClient(cc grpc.ClientConnInterface) AccountControlClient {
	return &accountControlClient{cc}
}

func (c *accountControlClient) RegistrationAccount(ctx context.Context, in *User, opts ...grpc.CallOption) (*SessionData, error) {
	out := new(SessionData)
	err := c.cc.Invoke(ctx, AccountControl_RegistrationAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountControlClient) AuthorizationAccount(ctx context.Context, in *User, opts ...grpc.CallOption) (*SessionData, error) {
	out := new(SessionData)
	err := c.cc.Invoke(ctx, AccountControl_AuthorizationAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountControlClient) ChangePasswordAccount(ctx context.Context, in *ChangePasswordData, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, AccountControl_ChangePasswordAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountControlClient) DeleteAccount(ctx context.Context, in *FullInfoUser, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, AccountControl_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountControlClient) IsAuthorizedSessionData(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*AccountID, error) {
	out := new(AccountID)
	err := c.cc.Invoke(ctx, AccountControl_IsAuthorizedSessionData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountControlClient) DeleteSessionData(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, AccountControl_DeleteSessionData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountControlServer is the server API for AccountControl service.
// All implementations must embed UnimplementedAccountControlServer
// for forward compatibility
type AccountControlServer interface {
	RegistrationAccount(context.Context, *User) (*SessionData, error)
	AuthorizationAccount(context.Context, *User) (*SessionData, error)
	ChangePasswordAccount(context.Context, *ChangePasswordData) (*Null, error)
	DeleteAccount(context.Context, *FullInfoUser) (*Null, error)
	IsAuthorizedSessionData(context.Context, *SessionData) (*AccountID, error)
	DeleteSessionData(context.Context, *SessionData) (*Null, error)
}

// UnimplementedAccountControlServer must be embedded to have forward compatible implementations.
type UnimplementedAccountControlServer struct {
}

func (UnimplementedAccountControlServer) RegistrationAccount(context.Context, *User) (*SessionData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationAccount not implemented")
}
func (UnimplementedAccountControlServer) AuthorizationAccount(context.Context, *User) (*SessionData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizationAccount not implemented")
}
func (UnimplementedAccountControlServer) ChangePasswordAccount(context.Context, *ChangePasswordData) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePasswordAccount not implemented")
}
func (UnimplementedAccountControlServer) DeleteAccount(context.Context, *FullInfoUser) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountControlServer) IsAuthorizedSessionData(context.Context, *SessionData) (*AccountID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorizedSessionData not implemented")
}
func (UnimplementedAccountControlServer) DeleteSessionData(context.Context, *SessionData) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSessionData not implemented")
}

// UnsafeAccountControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountControlServer will
// result in compilation errors.
type UnsafeAccountControlServer interface {
	mustEmbedUnimplementedAccountControlServer()
}

func RegisterAccountControlServer(s grpc.ServiceRegistrar, srv AccountControlServer) {
	s.RegisterService(&AccountControl_ServiceDesc, srv)
}

func _AccountControl_RegistrationAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountControlServer).RegistrationAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountControl_RegistrationAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountControlServer).RegistrationAccount(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountControl_AuthorizationAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountControlServer).AuthorizationAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountControl_AuthorizationAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountControlServer).AuthorizationAccount(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountControl_ChangePasswordAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountControlServer).ChangePasswordAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountControl_ChangePasswordAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountControlServer).ChangePasswordAccount(ctx, req.(*ChangePasswordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountControl_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullInfoUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountControlServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountControl_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountControlServer).DeleteAccount(ctx, req.(*FullInfoUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountControl_IsAuthorizedSessionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountControlServer).IsAuthorizedSessionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountControl_IsAuthorizedSessionData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountControlServer).IsAuthorizedSessionData(ctx, req.(*SessionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountControl_DeleteSessionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountControlServer).DeleteSessionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountControl_DeleteSessionData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountControlServer).DeleteSessionData(ctx, req.(*SessionData))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountControl_ServiceDesc is the grpc.ServiceDesc for AccountControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccountControl",
	HandlerType: (*AccountControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistrationAccount",
			Handler:    _AccountControl_RegistrationAccount_Handler,
		},
		{
			MethodName: "AuthorizationAccount",
			Handler:    _AccountControl_AuthorizationAccount_Handler,
		},
		{
			MethodName: "ChangePasswordAccount",
			Handler:    _AccountControl_ChangePasswordAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountControl_DeleteAccount_Handler,
		},
		{
			MethodName: "IsAuthorizedSessionData",
			Handler:    _AccountControl_IsAuthorizedSessionData_Handler,
		},
		{
			MethodName: "DeleteSessionData",
			Handler:    _AccountControl_DeleteSessionData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AccountControl.proto",
}
