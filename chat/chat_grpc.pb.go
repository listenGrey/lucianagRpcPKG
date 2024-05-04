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

// GetChatServiceClient is the client API for GetChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetChatServiceClient interface {
	GetChat(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chat, error)
}

type getChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetChatServiceClient(cc grpc.ClientConnInterface) GetChatServiceClient {
	return &getChatServiceClient{cc}
}

func (c *getChatServiceClient) GetChat(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.GetChatService/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChatServiceServer is the server API for GetChatService service.
// All implementations must embed UnimplementedGetChatServiceServer
// for forward compatibility
type GetChatServiceServer interface {
	GetChat(context.Context, *ID) (*Chat, error)
	mustEmbedUnimplementedGetChatServiceServer()
}

// UnimplementedGetChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGetChatServiceServer struct {
}

func (UnimplementedGetChatServiceServer) GetChat(context.Context, *ID) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedGetChatServiceServer) mustEmbedUnimplementedGetChatServiceServer() {}

// UnsafeGetChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetChatServiceServer will
// result in compilation errors.
type UnsafeGetChatServiceServer interface {
	mustEmbedUnimplementedGetChatServiceServer()
}

func RegisterGetChatServiceServer(s grpc.ServiceRegistrar, srv GetChatServiceServer) {
	s.RegisterService(&GetChatService_ServiceDesc, srv)
}

func _GetChatService_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChatServiceServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.GetChatService/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChatServiceServer).GetChat(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// GetChatService_ServiceDesc is the grpc.ServiceDesc for GetChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.GetChatService",
	HandlerType: (*GetChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChat",
			Handler:    _GetChatService_GetChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

// GetChatsServiceClient is the client API for GetChatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetChatsServiceClient interface {
	GetChats(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chats, error)
}

type getChatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetChatsServiceClient(cc grpc.ClientConnInterface) GetChatsServiceClient {
	return &getChatsServiceClient{cc}
}

func (c *getChatsServiceClient) GetChats(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chats, error) {
	out := new(Chats)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.GetChatsService/GetChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChatsServiceServer is the server API for GetChatsService service.
// All implementations must embed UnimplementedGetChatsServiceServer
// for forward compatibility
type GetChatsServiceServer interface {
	GetChats(context.Context, *ID) (*Chats, error)
	mustEmbedUnimplementedGetChatsServiceServer()
}

// UnimplementedGetChatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGetChatsServiceServer struct {
}

func (UnimplementedGetChatsServiceServer) GetChats(context.Context, *ID) (*Chats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedGetChatsServiceServer) mustEmbedUnimplementedGetChatsServiceServer() {}

// UnsafeGetChatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetChatsServiceServer will
// result in compilation errors.
type UnsafeGetChatsServiceServer interface {
	mustEmbedUnimplementedGetChatsServiceServer()
}

func RegisterGetChatsServiceServer(s grpc.ServiceRegistrar, srv GetChatsServiceServer) {
	s.RegisterService(&GetChatsService_ServiceDesc, srv)
}

func _GetChatsService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChatsServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.GetChatsService/GetChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChatsServiceServer).GetChats(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// GetChatsService_ServiceDesc is the grpc.ServiceDesc for GetChatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetChatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.GetChatsService",
	HandlerType: (*GetChatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChats",
			Handler:    _GetChatsService_GetChats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}
