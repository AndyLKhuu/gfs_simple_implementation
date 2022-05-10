// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: chunkserver.proto

package protos

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
	// Reads a range of bytes in a chunk.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error)
	// Receive potential write data from clients.
	ReceiveWriteData(ctx context.Context, in *WriteDataBundle, opts ...grpc.CallOption) (*Ack, error)
	// Commits a mutation operation on primary.
	PrimaryCommitMutate(ctx context.Context, in *ChunkHandler, opts ...grpc.CallOption) (*Ack, error)
	// Commits a mutation operation on secondary.
	SecondaryCommitMutate(ctx context.Context, in *ChunkHandler, opts ...grpc.CallOption) (*Ack, error)
	// Creates a new chunk.
	CreateNewChunk(ctx context.Context, in *ChunkHandler, opts ...grpc.CallOption) (*Ack, error)
}

type chunkServerClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServerClient(cc grpc.ClientConnInterface) ChunkServerClient {
	return &chunkServerClient{cc}
}

func (c *chunkServerClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error) {
	out := new(ReadReply)
	err := c.cc.Invoke(ctx, "/protos.ChunkServer/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerClient) ReceiveWriteData(ctx context.Context, in *WriteDataBundle, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/protos.ChunkServer/ReceiveWriteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerClient) PrimaryCommitMutate(ctx context.Context, in *ChunkHandler, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/protos.ChunkServer/PrimaryCommitMutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerClient) SecondaryCommitMutate(ctx context.Context, in *ChunkHandler, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/protos.ChunkServer/SecondaryCommitMutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerClient) CreateNewChunk(ctx context.Context, in *ChunkHandler, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/protos.ChunkServer/CreateNewChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServerServer is the server API for ChunkServer service.
// All implementations must embed UnimplementedChunkServerServer
// for forward compatibility
type ChunkServerServer interface {
	// Reads a range of bytes in a chunk.
	Read(context.Context, *ReadRequest) (*ReadReply, error)
	// Receive potential write data from clients.
	ReceiveWriteData(context.Context, *WriteDataBundle) (*Ack, error)
	// Commits a mutation operation on primary.
	PrimaryCommitMutate(context.Context, *ChunkHandler) (*Ack, error)
	// Commits a mutation operation on secondary.
	SecondaryCommitMutate(context.Context, *ChunkHandler) (*Ack, error)
	// Creates a new chunk.
	CreateNewChunk(context.Context, *ChunkHandler) (*Ack, error)
	mustEmbedUnimplementedChunkServerServer()
}

// UnimplementedChunkServerServer must be embedded to have forward compatible implementations.
type UnimplementedChunkServerServer struct {
}

func (UnimplementedChunkServerServer) Read(context.Context, *ReadRequest) (*ReadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedChunkServerServer) ReceiveWriteData(context.Context, *WriteDataBundle) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveWriteData not implemented")
}
func (UnimplementedChunkServerServer) PrimaryCommitMutate(context.Context, *ChunkHandler) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrimaryCommitMutate not implemented")
}
func (UnimplementedChunkServerServer) SecondaryCommitMutate(context.Context, *ChunkHandler) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecondaryCommitMutate not implemented")
}
func (UnimplementedChunkServerServer) CreateNewChunk(context.Context, *ChunkHandler) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewChunk not implemented")
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

func _ChunkServer_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChunkServer/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServer_ReceiveWriteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteDataBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).ReceiveWriteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChunkServer/ReceiveWriteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).ReceiveWriteData(ctx, req.(*WriteDataBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServer_PrimaryCommitMutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkHandler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).PrimaryCommitMutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChunkServer/PrimaryCommitMutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).PrimaryCommitMutate(ctx, req.(*ChunkHandler))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServer_SecondaryCommitMutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkHandler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).SecondaryCommitMutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChunkServer/SecondaryCommitMutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).SecondaryCommitMutate(ctx, req.(*ChunkHandler))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServer_CreateNewChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkHandler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).CreateNewChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChunkServer/CreateNewChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).CreateNewChunk(ctx, req.(*ChunkHandler))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkServer_ServiceDesc is the grpc.ServiceDesc for ChunkServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChunkServer",
	HandlerType: (*ChunkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _ChunkServer_Read_Handler,
		},
		{
			MethodName: "ReceiveWriteData",
			Handler:    _ChunkServer_ReceiveWriteData_Handler,
		},
		{
			MethodName: "PrimaryCommitMutate",
			Handler:    _ChunkServer_PrimaryCommitMutate_Handler,
		},
		{
			MethodName: "SecondaryCommitMutate",
			Handler:    _ChunkServer_SecondaryCommitMutate_Handler,
		},
		{
			MethodName: "CreateNewChunk",
			Handler:    _ChunkServer_CreateNewChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chunkserver.proto",
}
