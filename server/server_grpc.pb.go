// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

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

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	// Ack method
	Ack(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	// Server side streaming
	Subscribe(ctx context.Context, in *Event, opts ...grpc.CallOption) (EventService_SubscribeClient, error)
	// Client side streaming
	Publish(ctx context.Context, opts ...grpc.CallOption) (EventService_PublishClient, error)
	// Bidirectional streaming
	Route(ctx context.Context, opts ...grpc.CallOption) (EventService_RouteClient, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Ack(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/server.EventService/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Subscribe(ctx context.Context, in *Event, opts ...grpc.CallOption) (EventService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[0], "/server.EventService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (EventService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[1], "/server.EventService/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServicePublishClient{stream}
	return x, nil
}

type EventService_PublishClient interface {
	Send(*Event) error
	CloseAndRecv() (*Event, error)
	grpc.ClientStream
}

type eventServicePublishClient struct {
	grpc.ClientStream
}

func (x *eventServicePublishClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServicePublishClient) CloseAndRecv() (*Event, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) Route(ctx context.Context, opts ...grpc.CallOption) (EventService_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[2], "/server.EventService/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceRouteClient{stream}
	return x, nil
}

type EventService_RouteClient interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventServiceRouteClient struct {
	grpc.ClientStream
}

func (x *eventServiceRouteClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServiceRouteClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	// Ack method
	Ack(context.Context, *Event) (*Event, error)
	// Server side streaming
	Subscribe(*Event, EventService_SubscribeServer) error
	// Client side streaming
	Publish(EventService_PublishServer) error
	// Bidirectional streaming
	Route(EventService_RouteServer) error
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) Ack(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedEventServiceServer) Subscribe(*Event, EventService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventServiceServer) Publish(EventService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedEventServiceServer) Route(EventService_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.EventService/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Ack(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Event)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServiceServer).Subscribe(m, &eventServiceSubscribeServer{stream})
}

type EventService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).Publish(&eventServicePublishServer{stream})
}

type EventService_PublishServer interface {
	SendAndClose(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type eventServicePublishServer struct {
	grpc.ServerStream
}

func (x *eventServicePublishServer) SendAndClose(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServicePublishServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventService_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).Route(&eventServiceRouteServer{stream})
}

type EventService_RouteServer interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type eventServiceRouteServer struct {
	grpc.ServerStream
}

func (x *eventServiceRouteServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServiceRouteServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ack",
			Handler:    _EventService_Ack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _EventService_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _EventService_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}
