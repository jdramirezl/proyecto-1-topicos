// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/message.proto

package message

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	AddMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ConsumeMessage(ctx context.Context, opts ...grpc.CallOption) (MessageService_ConsumeMessageClient, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) AddMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/MessageService/AddMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) RemoveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/MessageService/RemoveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ConsumeMessage(ctx context.Context, opts ...grpc.CallOption) (MessageService_ConsumeMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageService_ServiceDesc.Streams[0], "/MessageService/ConsumeMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceConsumeMessageClient{stream}
	return x, nil
}

type MessageService_ConsumeMessageClient interface {
	Send(*ConsumeMessageRequest) error
	Recv() (*ConsumeMessageResponse, error)
	grpc.ClientStream
}

type messageServiceConsumeMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceConsumeMessageClient) Send(m *ConsumeMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceConsumeMessageClient) Recv() (*ConsumeMessageResponse, error) {
	m := new(ConsumeMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	AddMessage(context.Context, *MessageRequest) (*empty.Empty, error)
	RemoveMessage(context.Context, *MessageRequest) (*empty.Empty, error)
	ConsumeMessage(MessageService_ConsumeMessageServer) error
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) AddMessage(context.Context, *MessageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessage not implemented")
}
func (UnimplementedMessageServiceServer) RemoveMessage(context.Context, *MessageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMessage not implemented")
}
func (UnimplementedMessageServiceServer) ConsumeMessage(MessageService_ConsumeMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeMessage not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_AddMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).AddMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/AddMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).AddMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_RemoveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).RemoveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/RemoveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).RemoveMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ConsumeMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).ConsumeMessage(&messageServiceConsumeMessageServer{stream})
}

type MessageService_ConsumeMessageServer interface {
	Send(*ConsumeMessageResponse) error
	Recv() (*ConsumeMessageRequest, error)
	grpc.ServerStream
}

type messageServiceConsumeMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceConsumeMessageServer) Send(m *ConsumeMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceConsumeMessageServer) Recv() (*ConsumeMessageRequest, error) {
	m := new(ConsumeMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMessage",
			Handler:    _MessageService_AddMessage_Handler,
		},
		{
			MethodName: "RemoveMessage",
			Handler:    _MessageService_RemoveMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeMessage",
			Handler:       _MessageService_ConsumeMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/message.proto",
}
