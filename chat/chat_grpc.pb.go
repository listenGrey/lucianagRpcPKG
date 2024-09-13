// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0--rc1
// source: chat/chat.proto

package chat

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

// GetChatListClient is the client API for GetChatList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetChatListClient interface {
	GetChatList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ChatList, error)
}

type getChatListClient struct {
	cc grpc.ClientConnInterface
}

func NewGetChatListClient(cc grpc.ClientConnInterface) GetChatListClient {
	return &getChatListClient{cc}
}

func (c *getChatListClient) GetChatList(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ChatList, error) {
	out := new(ChatList)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.GetChatList/GetChatList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChatListServer is the server API for GetChatList service.
// All implementations must embed UnimplementedGetChatListServer
// for forward compatibility
type GetChatListServer interface {
	GetChatList(context.Context, *ID) (*ChatList, error)
	mustEmbedUnimplementedGetChatListServer()
}

// UnimplementedGetChatListServer must be embedded to have forward compatible implementations.
type UnimplementedGetChatListServer struct {
}

func (UnimplementedGetChatListServer) GetChatList(context.Context, *ID) (*ChatList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatList not implemented")
}
func (UnimplementedGetChatListServer) mustEmbedUnimplementedGetChatListServer() {}

// UnsafeGetChatListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetChatListServer will
// result in compilation errors.
type UnsafeGetChatListServer interface {
	mustEmbedUnimplementedGetChatListServer()
}

func RegisterGetChatListServer(s grpc.ServiceRegistrar, srv GetChatListServer) {
	s.RegisterService(&GetChatList_ServiceDesc, srv)
}

func _GetChatList_GetChatList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChatListServer).GetChatList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.GetChatList/GetChatList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChatListServer).GetChatList(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// GetChatList_ServiceDesc is the grpc.ServiceDesc for GetChatList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetChatList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.GetChatList",
	HandlerType: (*GetChatListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChatList",
			Handler:    _GetChatList_GetChatList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

// GetChatClient is the client API for GetChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetChatClient interface {
	GetChat(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chat, error)
}

type getChatClient struct {
	cc grpc.ClientConnInterface
}

func NewGetChatClient(cc grpc.ClientConnInterface) GetChatClient {
	return &getChatClient{cc}
}

func (c *getChatClient) GetChat(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.GetChat/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChatServer is the server API for GetChat service.
// All implementations must embed UnimplementedGetChatServer
// for forward compatibility
type GetChatServer interface {
	GetChat(context.Context, *ID) (*Chat, error)
	mustEmbedUnimplementedGetChatServer()
}

// UnimplementedGetChatServer must be embedded to have forward compatible implementations.
type UnimplementedGetChatServer struct {
}

func (UnimplementedGetChatServer) GetChat(context.Context, *ID) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedGetChatServer) mustEmbedUnimplementedGetChatServer() {}

// UnsafeGetChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetChatServer will
// result in compilation errors.
type UnsafeGetChatServer interface {
	mustEmbedUnimplementedGetChatServer()
}

func RegisterGetChatServer(s grpc.ServiceRegistrar, srv GetChatServer) {
	s.RegisterService(&GetChat_ServiceDesc, srv)
}

func _GetChat_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChatServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.GetChat/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChatServer).GetChat(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// GetChat_ServiceDesc is the grpc.ServiceDesc for GetChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.GetChat",
	HandlerType: (*GetChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChat",
			Handler:    _GetChat_GetChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

// NewChatClient is the client API for NewChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewChatClient interface {
	NewChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Success, error)
}

type newChatClient struct {
	cc grpc.ClientConnInterface
}

func NewNewChatClient(cc grpc.ClientConnInterface) NewChatClient {
	return &newChatClient{cc}
}

func (c *newChatClient) NewChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.NewChat/NewChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewChatServer is the server API for NewChat service.
// All implementations must embed UnimplementedNewChatServer
// for forward compatibility
type NewChatServer interface {
	NewChat(context.Context, *Chat) (*Success, error)
	mustEmbedUnimplementedNewChatServer()
}

// UnimplementedNewChatServer must be embedded to have forward compatible implementations.
type UnimplementedNewChatServer struct {
}

func (UnimplementedNewChatServer) NewChat(context.Context, *Chat) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewChat not implemented")
}
func (UnimplementedNewChatServer) mustEmbedUnimplementedNewChatServer() {}

// UnsafeNewChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewChatServer will
// result in compilation errors.
type UnsafeNewChatServer interface {
	mustEmbedUnimplementedNewChatServer()
}

func RegisterNewChatServer(s grpc.ServiceRegistrar, srv NewChatServer) {
	s.RegisterService(&NewChat_ServiceDesc, srv)
}

func _NewChat_NewChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewChatServer).NewChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.NewChat/NewChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewChatServer).NewChat(ctx, req.(*Chat))
	}
	return interceptor(ctx, in, info, handler)
}

// NewChat_ServiceDesc is the grpc.ServiceDesc for NewChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.NewChat",
	HandlerType: (*NewChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewChat",
			Handler:    _NewChat_NewChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

// RenameChatClient is the client API for RenameChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RenameChatClient interface {
	RenameChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Success, error)
}

type renameChatClient struct {
	cc grpc.ClientConnInterface
}

func NewRenameChatClient(cc grpc.ClientConnInterface) RenameChatClient {
	return &renameChatClient{cc}
}

func (c *renameChatClient) RenameChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.RenameChat/RenameChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RenameChatServer is the server API for RenameChat service.
// All implementations must embed UnimplementedRenameChatServer
// for forward compatibility
type RenameChatServer interface {
	RenameChat(context.Context, *Chat) (*Success, error)
	mustEmbedUnimplementedRenameChatServer()
}

// UnimplementedRenameChatServer must be embedded to have forward compatible implementations.
type UnimplementedRenameChatServer struct {
}

func (UnimplementedRenameChatServer) RenameChat(context.Context, *Chat) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameChat not implemented")
}
func (UnimplementedRenameChatServer) mustEmbedUnimplementedRenameChatServer() {}

// UnsafeRenameChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RenameChatServer will
// result in compilation errors.
type UnsafeRenameChatServer interface {
	mustEmbedUnimplementedRenameChatServer()
}

func RegisterRenameChatServer(s grpc.ServiceRegistrar, srv RenameChatServer) {
	s.RegisterService(&RenameChat_ServiceDesc, srv)
}

func _RenameChat_RenameChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenameChatServer).RenameChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.RenameChat/RenameChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenameChatServer).RenameChat(ctx, req.(*Chat))
	}
	return interceptor(ctx, in, info, handler)
}

// RenameChat_ServiceDesc is the grpc.ServiceDesc for RenameChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RenameChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.RenameChat",
	HandlerType: (*RenameChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RenameChat",
			Handler:    _RenameChat_RenameChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

// DeleteChatClient is the client API for DeleteChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteChatClient interface {
	DeleteChat(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Success, error)
}

type deleteChatClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteChatClient(cc grpc.ClientConnInterface) DeleteChatClient {
	return &deleteChatClient{cc}
}

func (c *deleteChatClient) DeleteChat(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.DeleteChat/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteChatServer is the server API for DeleteChat service.
// All implementations must embed UnimplementedDeleteChatServer
// for forward compatibility
type DeleteChatServer interface {
	DeleteChat(context.Context, *ID) (*Success, error)
	mustEmbedUnimplementedDeleteChatServer()
}

// UnimplementedDeleteChatServer must be embedded to have forward compatible implementations.
type UnimplementedDeleteChatServer struct {
}

func (UnimplementedDeleteChatServer) DeleteChat(context.Context, *ID) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedDeleteChatServer) mustEmbedUnimplementedDeleteChatServer() {}

// UnsafeDeleteChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteChatServer will
// result in compilation errors.
type UnsafeDeleteChatServer interface {
	mustEmbedUnimplementedDeleteChatServer()
}

func RegisterDeleteChatServer(s grpc.ServiceRegistrar, srv DeleteChatServer) {
	s.RegisterService(&DeleteChat_ServiceDesc, srv)
}

func _DeleteChat_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteChatServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.DeleteChat/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteChatServer).DeleteChat(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteChat_ServiceDesc is the grpc.ServiceDesc for DeleteChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.DeleteChat",
	HandlerType: (*DeleteChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteChat",
			Handler:    _DeleteChat_DeleteChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}
