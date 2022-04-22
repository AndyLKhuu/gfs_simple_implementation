// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: chunkserver.proto

package chunkserver

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

// ChunkServerClient is the client API for ChunkServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkServerClient interface {
	// Sends a heart beat message
	SendHeartBeatMessage(ctx context.Context, in *HeartBeatMessage, opts ...grpc.CallOption) (*HeartBeatMessageReply, error)
	// Sends another greeting
	SayHelloAgain(ctx context.Context, in *HeartBeatMessage, opts ...grpc.CallOption) (*HeartBeatMessageReply, error)
}

type chunkServerClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServerClient(cc grpc.ClientConnInterface) ChunkServerClient {
	return &chunkServerClient{cc}
}

func (c *chunkServerClient) SendHeartBeatMessage(ctx context.Context, in *HeartBeatMessage, opts ...grpc.CallOption) (*HeartBeatMessageReply, error) {
	out := new(HeartBeatMessageReply)
	err := c.cc.Invoke(ctx, "/master.ChunkServer/SendHeartBeatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerClient) SayHelloAgain(ctx context.Context, in *HeartBeatMessage, opts ...grpc.CallOption) (*HeartBeatMessageReply, error) {
	out := new(HeartBeatMessageReply)
	err := c.cc.Invoke(ctx, "/master.ChunkServer/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServerServer is the server API for ChunkServer service.
// All implementations must embed UnimplementedChunkServerServer
// for forward compatibility
type ChunkServerServer interface {
	// Sends a heart beat message
	SendHeartBeatMessage(context.Context, *HeartBeatMessage) (*HeartBeatMessageReply, error)
	// Sends another greeting
	SayHelloAgain(context.Context, *HeartBeatMessage) (*HeartBeatMessageReply, error)
	mustEmbedUnimplementedChunkServerServer()
}

// UnimplementedChunkServerServer must be embedded to have forward compatible implementations.
type UnimplementedChunkServerServer struct {
}

func (UnimplementedChunkServerServer) SendHeartBeatMessage(context.Context, *HeartBeatMessage) (*HeartBeatMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartBeatMessage not implemented")
}
func (UnimplementedChunkServerServer) SayHelloAgain(context.Context, *HeartBeatMessage) (*HeartBeatMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}
func (UnimplementedChunkServerServer) mustEmbedUnimplementedChunkServerServer() {}

// UnsafeChunkServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServerServer will
// result in compilation errors.
type UnsafeChunkServerServer interface {
	mustEmbedUnimplementedChunkServerServer()
}

func RegisterChunkServerServer(s grpc.ServiceRegistrar, srv ChunkServerServer) {
	s.RegisterService(&ChunkServer_ServiceDesc, srv)
}

func _ChunkServer_SendHeartBeatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).SendHeartBeatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.ChunkServer/SendHeartBeatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).SendHeartBeatMessage(ctx, req.(*HeartBeatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServer_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.ChunkServer/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).SayHelloAgain(ctx, req.(*HeartBeatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkServer_ServiceDesc is the grpc.ServiceDesc for ChunkServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.ChunkServer",
	HandlerType: (*ChunkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHeartBeatMessage",
			Handler:    _ChunkServer_SendHeartBeatMessage_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _ChunkServer_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chunkserver.proto",
}
