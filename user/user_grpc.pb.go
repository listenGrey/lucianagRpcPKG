// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0--rc1
// source: user/user.proto

package user

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

// CheckExistClient is the client API for CheckExist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckExistClient interface {
	RegisterCheck(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Exist, error)
}

type checkExistClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckExistClient(cc grpc.ClientConnInterface) CheckExistClient {
	return &checkExistClient{cc}
}

func (c *checkExistClient) RegisterCheck(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Exist, error) {
	out := new(Exist)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.CheckExist/RegisterCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckExistServer is the server API for CheckExist service.
// All implementations must embed UnimplementedCheckExistServer
// for forward compatibility
type CheckExistServer interface {
	RegisterCheck(context.Context, *Email) (*Exist, error)
	mustEmbedUnimplementedCheckExistServer()
}

// UnimplementedCheckExistServer must be embedded to have forward compatible implementations.
type UnimplementedCheckExistServer struct {
}

func (UnimplementedCheckExistServer) RegisterCheck(context.Context, *Email) (*Exist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCheck not implemented")
}
func (UnimplementedCheckExistServer) mustEmbedUnimplementedCheckExistServer() {}

// UnsafeCheckExistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckExistServer will
// result in compilation errors.
type UnsafeCheckExistServer interface {
	mustEmbedUnimplementedCheckExistServer()
}

func RegisterCheckExistServer(s grpc.ServiceRegistrar, srv CheckExistServer) {
	s.RegisterService(&CheckExist_ServiceDesc, srv)
}

func _CheckExist_RegisterCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckExistServer).RegisterCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.CheckExist/RegisterCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckExistServer).RegisterCheck(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckExist_ServiceDesc is the grpc.ServiceDesc for CheckExist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckExist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.CheckExist",
	HandlerType: (*CheckExistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCheck",
			Handler:    _CheckExist_RegisterCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

// RegisterCheckClient is the client API for RegisterCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterCheckClient interface {
	RegisterCheck(ctx context.Context, in *RegisterFrom, opts ...grpc.CallOption) (*Success, error)
}

type registerCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterCheckClient(cc grpc.ClientConnInterface) RegisterCheckClient {
	return &registerCheckClient{cc}
}

func (c *registerCheckClient) RegisterCheck(ctx context.Context, in *RegisterFrom, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.RegisterCheck/RegisterCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterCheckServer is the server API for RegisterCheck service.
// All implementations must embed UnimplementedRegisterCheckServer
// for forward compatibility
type RegisterCheckServer interface {
	RegisterCheck(context.Context, *RegisterFrom) (*Success, error)
	mustEmbedUnimplementedRegisterCheckServer()
}

// UnimplementedRegisterCheckServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterCheckServer struct {
}

func (UnimplementedRegisterCheckServer) RegisterCheck(context.Context, *RegisterFrom) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCheck not implemented")
}
func (UnimplementedRegisterCheckServer) mustEmbedUnimplementedRegisterCheckServer() {}

// UnsafeRegisterCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterCheckServer will
// result in compilation errors.
type UnsafeRegisterCheckServer interface {
	mustEmbedUnimplementedRegisterCheckServer()
}

func RegisterRegisterCheckServer(s grpc.ServiceRegistrar, srv RegisterCheckServer) {
	s.RegisterService(&RegisterCheck_ServiceDesc, srv)
}

func _RegisterCheck_RegisterCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterFrom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterCheckServer).RegisterCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.RegisterCheck/RegisterCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterCheckServer).RegisterCheck(ctx, req.(*RegisterFrom))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterCheck_ServiceDesc is the grpc.ServiceDesc for RegisterCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.RegisterCheck",
	HandlerType: (*RegisterCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCheck",
			Handler:    _RegisterCheck_RegisterCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

// LoginCheckClient is the client API for LoginCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginCheckClient interface {
	LoginCheck(ctx context.Context, in *LoginForm, opts ...grpc.CallOption) (*LogInfo, error)
}

type loginCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginCheckClient(cc grpc.ClientConnInterface) LoginCheckClient {
	return &loginCheckClient{cc}
}

func (c *loginCheckClient) LoginCheck(ctx context.Context, in *LoginForm, opts ...grpc.CallOption) (*LogInfo, error) {
	out := new(LogInfo)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.LoginCheck/LoginCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginCheckServer is the server API for LoginCheck service.
// All implementations must embed UnimplementedLoginCheckServer
// for forward compatibility
type LoginCheckServer interface {
	LoginCheck(context.Context, *LoginForm) (*LogInfo, error)
	mustEmbedUnimplementedLoginCheckServer()
}

// UnimplementedLoginCheckServer must be embedded to have forward compatible implementations.
type UnimplementedLoginCheckServer struct {
}

func (UnimplementedLoginCheckServer) LoginCheck(context.Context, *LoginForm) (*LogInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginCheck not implemented")
}
func (UnimplementedLoginCheckServer) mustEmbedUnimplementedLoginCheckServer() {}

// UnsafeLoginCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginCheckServer will
// result in compilation errors.
type UnsafeLoginCheckServer interface {
	mustEmbedUnimplementedLoginCheckServer()
}

func RegisterLoginCheckServer(s grpc.ServiceRegistrar, srv LoginCheckServer) {
	s.RegisterService(&LoginCheck_ServiceDesc, srv)
}

func _LoginCheck_LoginCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginCheckServer).LoginCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.LoginCheck/LoginCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginCheckServer).LoginCheck(ctx, req.(*LoginForm))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginCheck_ServiceDesc is the grpc.ServiceDesc for LoginCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.LoginCheck",
	HandlerType: (*LoginCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginCheck",
			Handler:    _LoginCheck_LoginCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
