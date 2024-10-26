// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: messages/messages.proto

package messages

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
	MessagesService_GetMessage_FullMethodName         = "/messages.MessagesService/getMessage"
	MessagesService_GetAndResetMetrics_FullMethodName = "/messages.MessagesService/getAndResetMetrics"
)

// MessagesServiceClient is the client API for MessagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesServiceClient interface {
	GetMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
	GetAndResetMetrics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MetricsReply, error)
}

type messagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesServiceClient(cc grpc.ClientConnInterface) MessagesServiceClient {
	return &messagesServiceClient{cc}
}

func (c *messagesServiceClient) GetMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, MessagesService_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesServiceClient) GetAndResetMetrics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MetricsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricsReply)
	err := c.cc.Invoke(ctx, MessagesService_GetAndResetMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServiceServer is the server API for MessagesService service.
// All implementations must embed UnimplementedMessagesServiceServer
// for forward compatibility.
type MessagesServiceServer interface {
	GetMessage(context.Context, *MessageRequest) (*MessageReply, error)
	GetAndResetMetrics(context.Context, *Empty) (*MetricsReply, error)
	mustEmbedUnimplementedMessagesServiceServer()
}

// UnimplementedMessagesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagesServiceServer struct{}

func (UnimplementedMessagesServiceServer) GetMessage(context.Context, *MessageRequest) (*MessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessagesServiceServer) GetAndResetMetrics(context.Context, *Empty) (*MetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAndResetMetrics not implemented")
}
func (UnimplementedMessagesServiceServer) mustEmbedUnimplementedMessagesServiceServer() {}
func (UnimplementedMessagesServiceServer) testEmbeddedByValue()                         {}

// UnsafeMessagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServiceServer will
// result in compilation errors.
type UnsafeMessagesServiceServer interface {
	mustEmbedUnimplementedMessagesServiceServer()
}

func RegisterMessagesServiceServer(s grpc.ServiceRegistrar, srv MessagesServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessagesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessagesService_ServiceDesc, srv)
}

func _MessagesService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagesService_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).GetMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagesService_GetAndResetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).GetAndResetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagesService_GetAndResetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).GetAndResetMetrics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagesService_ServiceDesc is the grpc.ServiceDesc for MessagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messages.MessagesService",
	HandlerType: (*MessagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMessage",
			Handler:    _MessagesService_GetMessage_Handler,
		},
		{
			MethodName: "getAndResetMetrics",
			Handler:    _MessagesService_GetAndResetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages/messages.proto",
}
