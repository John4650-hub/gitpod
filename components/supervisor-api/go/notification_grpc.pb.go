// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	// Prompts the user and asks for a decision. Typically called by some external
	// process. If the list of actions is empty this service returns immediately,
	// otherwise it blocks until the user has made their choice.
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	// Subscribe to notifications. Typically called by the IDE.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (NotificationService_SubscribeClient, error)
	// Report a user's choice as a response to a notification. Typically called by
	// the IDE.
	Respond(ctx context.Context, in *RespondRequest, opts ...grpc.CallOption) (*RespondResponse, error)
	// Called by the notifications clients to inform supervisor about
	// which is the latest client actively used by the user.
	// Useful for gp open / gp preview to trigger actions in the right client
	SetActiveClient(ctx context.Context, in *SetActiveClientRequest, opts ...grpc.CallOption) (*SetActiveClientResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/supervisor.NotificationService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (NotificationService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[0], "/supervisor.NotificationService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type notificationServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *notificationServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) Respond(ctx context.Context, in *RespondRequest, opts ...grpc.CallOption) (*RespondResponse, error) {
	out := new(RespondResponse)
	err := c.cc.Invoke(ctx, "/supervisor.NotificationService/Respond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SetActiveClient(ctx context.Context, in *SetActiveClientRequest, opts ...grpc.CallOption) (*SetActiveClientResponse, error) {
	out := new(SetActiveClientResponse)
	err := c.cc.Invoke(ctx, "/supervisor.NotificationService/SetActiveClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	// Prompts the user and asks for a decision. Typically called by some external
	// process. If the list of actions is empty this service returns immediately,
	// otherwise it blocks until the user has made their choice.
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	// Subscribe to notifications. Typically called by the IDE.
	Subscribe(*SubscribeRequest, NotificationService_SubscribeServer) error
	// Report a user's choice as a response to a notification. Typically called by
	// the IDE.
	Respond(context.Context, *RespondRequest) (*RespondResponse, error)
	// Called by the notifications clients to inform supervisor about
	// which is the latest client actively used by the user.
	// Useful for gp open / gp preview to trigger actions in the right client
	SetActiveClient(context.Context, *SetActiveClientRequest) (*SetActiveClientResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Notify(context.Context, *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedNotificationServiceServer) Subscribe(*SubscribeRequest, NotificationService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedNotificationServiceServer) Respond(context.Context, *RespondRequest) (*RespondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Respond not implemented")
}
func (UnimplementedNotificationServiceServer) SetActiveClient(context.Context, *SetActiveClientRequest) (*SetActiveClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActiveClient not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.NotificationService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).Subscribe(m, &notificationServiceSubscribeServer{stream})
}

type NotificationService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type notificationServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *notificationServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_Respond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Respond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.NotificationService/Respond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Respond(ctx, req.(*RespondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SetActiveClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetActiveClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SetActiveClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.NotificationService/SetActiveClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SetActiveClient(ctx, req.(*SetActiveClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _NotificationService_Notify_Handler,
		},
		{
			MethodName: "Respond",
			Handler:    _NotificationService_Respond_Handler,
		},
		{
			MethodName: "SetActiveClient",
			Handler:    _NotificationService_SetActiveClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NotificationService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}
