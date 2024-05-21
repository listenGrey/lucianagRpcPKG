// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0--rc1
// source: ask/ask.proto

package ask

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

// ResponseServiceClient is the client API for ResponseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResponseServiceClient interface {
	Response(ctx context.Context, in *ChatContext, opts ...grpc.CallOption) (*Result, error)
}

type responseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResponseServiceClient(cc grpc.ClientConnInterface) ResponseServiceClient {
	return &responseServiceClient{cc}
}

func (c *responseServiceClient) Response(ctx context.Context, in *ChatContext, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/lucianagRpcPKG.ResponseService/Response", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResponseServiceServer is the server API for ResponseService service.
// All implementations must embed UnimplementedResponseServiceServer
// for forward compatibility
type ResponseServiceServer interface {
	Response(context.Context, *ChatContext) (*Result, error)
	mustEmbedUnimplementedResponseServiceServer()
}

// UnimplementedResponseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResponseServiceServer struct {
}

func (UnimplementedResponseServiceServer) Response(context.Context, *ChatContext) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Response not implemented")
}
func (UnimplementedResponseServiceServer) mustEmbedUnimplementedResponseServiceServer() {}

// UnsafeResponseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResponseServiceServer will
// result in compilation errors.
type UnsafeResponseServiceServer interface {
	mustEmbedUnimplementedResponseServiceServer()
}

func RegisterResponseServiceServer(s grpc.ServiceRegistrar, srv ResponseServiceServer) {
	s.RegisterService(&ResponseService_ServiceDesc, srv)
}

func _ResponseService_Response_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponseServiceServer).Response(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lucianagRpcPKG.ResponseService/Response",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponseServiceServer).Response(ctx, req.(*ChatContext))
	}
	return interceptor(ctx, in, info, handler)
}

// ResponseService_ServiceDesc is the grpc.ServiceDesc for ResponseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResponseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lucianagRpcPKG.ResponseService",
	HandlerType: (*ResponseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Response",
			Handler:    _ResponseService_Response_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ask/ask.proto",
}
