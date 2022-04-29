// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// JuegoClient is the client API for Juego service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JuegoClient interface {
	Jugar(ctx context.Context, in *JuegoRequest, opts ...grpc.CallOption) (*JuegoReply, error)
}

type juegoClient struct {
	cc grpc.ClientConnInterface
}

func NewJuegoClient(cc grpc.ClientConnInterface) JuegoClient {
	return &juegoClient{cc}
}

func (c *juegoClient) Jugar(ctx context.Context, in *JuegoRequest, opts ...grpc.CallOption) (*JuegoReply, error) {
	out := new(JuegoReply)
	err := c.cc.Invoke(ctx, "/proto.Juego/Jugar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JuegoServer is the server API for Juego service.
// All implementations must embed UnimplementedJuegoServer
// for forward compatibility
type JuegoServer interface {
	Jugar(context.Context, *JuegoRequest) (*JuegoReply, error)
	mustEmbedUnimplementedJuegoServer()
}

// UnimplementedJuegoServer must be embedded to have forward compatible implementations.
type UnimplementedJuegoServer struct {
}

func (UnimplementedJuegoServer) Jugar(context.Context, *JuegoRequest) (*JuegoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Jugar not implemented")
}
func (UnimplementedJuegoServer) mustEmbedUnimplementedJuegoServer() {}

// UnsafeJuegoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JuegoServer will
// result in compilation errors.
type UnsafeJuegoServer interface {
	mustEmbedUnimplementedJuegoServer()
}

func RegisterJuegoServer(s grpc.ServiceRegistrar, srv JuegoServer) {
	s.RegisterService(&Juego_ServiceDesc, srv)
}

func _Juego_Jugar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JuegoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JuegoServer).Jugar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Juego/Jugar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JuegoServer).Jugar(ctx, req.(*JuegoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Juego_ServiceDesc is the grpc.ServiceDesc for Juego service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Juego_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Juego",
	HandlerType: (*JuegoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Jugar",
			Handler:    _Juego_Jugar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/juego.proto",
}
