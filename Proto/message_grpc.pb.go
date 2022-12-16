// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: Proto/message.proto

package Proto

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	EnviarEscuadron(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error)
	EnviarEscuadron2(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error)
	EnviarEscuadron3(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error)
	EnviarEscuadron4(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error)
	EstadoLab(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error)
	EstadoLab2(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error)
	EstadoLab3(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error)
	EstadoLab4(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) EnviarEscuadron(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error) {
	out := new(ResEnviarEscuadron)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EnviarEscuadron", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EnviarEscuadron2(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error) {
	out := new(ResEnviarEscuadron)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EnviarEscuadron2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EnviarEscuadron3(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error) {
	out := new(ResEnviarEscuadron)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EnviarEscuadron3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EnviarEscuadron4(ctx context.Context, in *ReqEnviarEscuadron, opts ...grpc.CallOption) (*ResEnviarEscuadron, error) {
	out := new(ResEnviarEscuadron)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EnviarEscuadron4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EstadoLab(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error) {
	out := new(ResEstadoLab)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EstadoLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EstadoLab2(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error) {
	out := new(ResEstadoLab)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EstadoLab2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EstadoLab3(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error) {
	out := new(ResEstadoLab)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EstadoLab3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) EstadoLab4(ctx context.Context, in *ReqEstadoLab, opts ...grpc.CallOption) (*ResEstadoLab, error) {
	out := new(ResEstadoLab)
	err := c.cc.Invoke(ctx, "/grpc.MessageService/EstadoLab4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	EnviarEscuadron(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error)
	EnviarEscuadron2(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error)
	EnviarEscuadron3(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error)
	EnviarEscuadron4(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error)
	EstadoLab(context.Context, *ReqEstadoLab) (*ResEstadoLab, error)
	EstadoLab2(context.Context, *ReqEstadoLab) (*ResEstadoLab, error)
	EstadoLab3(context.Context, *ReqEstadoLab) (*ResEstadoLab, error)
	EstadoLab4(context.Context, *ReqEstadoLab) (*ResEstadoLab, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) EnviarEscuadron(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEscuadron not implemented")
}
func (UnimplementedMessageServiceServer) EnviarEscuadron2(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEscuadron2 not implemented")
}
func (UnimplementedMessageServiceServer) EnviarEscuadron3(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEscuadron3 not implemented")
}
func (UnimplementedMessageServiceServer) EnviarEscuadron4(context.Context, *ReqEnviarEscuadron) (*ResEnviarEscuadron, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEscuadron4 not implemented")
}
func (UnimplementedMessageServiceServer) EstadoLab(context.Context, *ReqEstadoLab) (*ResEstadoLab, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoLab not implemented")
}
func (UnimplementedMessageServiceServer) EstadoLab2(context.Context, *ReqEstadoLab) (*ResEstadoLab, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoLab2 not implemented")
}
func (UnimplementedMessageServiceServer) EstadoLab3(context.Context, *ReqEstadoLab) (*ResEstadoLab, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoLab3 not implemented")
}
func (UnimplementedMessageServiceServer) EstadoLab4(context.Context, *ReqEstadoLab) (*ResEstadoLab, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoLab4 not implemented")
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

func _MessageService_EnviarEscuadron_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEnviarEscuadron)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EnviarEscuadron(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EnviarEscuadron",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EnviarEscuadron(ctx, req.(*ReqEnviarEscuadron))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EnviarEscuadron2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEnviarEscuadron)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EnviarEscuadron2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EnviarEscuadron2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EnviarEscuadron2(ctx, req.(*ReqEnviarEscuadron))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EnviarEscuadron3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEnviarEscuadron)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EnviarEscuadron3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EnviarEscuadron3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EnviarEscuadron3(ctx, req.(*ReqEnviarEscuadron))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EnviarEscuadron4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEnviarEscuadron)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EnviarEscuadron4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EnviarEscuadron4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EnviarEscuadron4(ctx, req.(*ReqEnviarEscuadron))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EstadoLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEstadoLab)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EstadoLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EstadoLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EstadoLab(ctx, req.(*ReqEstadoLab))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EstadoLab2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEstadoLab)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EstadoLab2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EstadoLab2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EstadoLab2(ctx, req.(*ReqEstadoLab))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EstadoLab3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEstadoLab)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EstadoLab3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EstadoLab3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EstadoLab3(ctx, req.(*ReqEstadoLab))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_EstadoLab4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqEstadoLab)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).EstadoLab4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MessageService/EstadoLab4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).EstadoLab4(ctx, req.(*ReqEstadoLab))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarEscuadron",
			Handler:    _MessageService_EnviarEscuadron_Handler,
		},
		{
			MethodName: "EnviarEscuadron2",
			Handler:    _MessageService_EnviarEscuadron2_Handler,
		},
		{
			MethodName: "EnviarEscuadron3",
			Handler:    _MessageService_EnviarEscuadron3_Handler,
		},
		{
			MethodName: "EnviarEscuadron4",
			Handler:    _MessageService_EnviarEscuadron4_Handler,
		},
		{
			MethodName: "EstadoLab",
			Handler:    _MessageService_EstadoLab_Handler,
		},
		{
			MethodName: "EstadoLab2",
			Handler:    _MessageService_EstadoLab2_Handler,
		},
		{
			MethodName: "EstadoLab3",
			Handler:    _MessageService_EstadoLab3_Handler,
		},
		{
			MethodName: "EstadoLab4",
			Handler:    _MessageService_EstadoLab4_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/message.proto",
}
