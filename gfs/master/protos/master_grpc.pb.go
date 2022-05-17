// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: master.proto

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

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterClient interface {
	// Sends a heart beat message to the client
	SendHeartBeatMessage(ctx context.Context, in *ChunkServerID, opts ...grpc.CallOption) (*Ack, error)
	// Retrieves the chunk handler of a chunk and the locations of its replicas
	GetChunkLocation(ctx context.Context, in *ChunkLocationRequest, opts ...grpc.CallOption) (*ChunkLocationReply, error)
	// Retrieve the fixed system chunksize
	GetSystemChunkSize(ctx context.Context, in *SystemChunkSizeRequest, opts ...grpc.CallOption) (*ChunkSize, error)
	// ClientWriteRequest
	ReceiveClientWriteRequest(ctx context.Context, in *ClientWriteRequest, opts ...grpc.CallOption) (*Ack, error)
	// Create a file.
	CreateFile(ctx context.Context, in *FileCreateRequest, opts ...grpc.CallOption) (*Ack, error)
	// Remove a file.
	RemoveFile(ctx context.Context, in *FileRemoveRequest, opts ...grpc.CallOption) (*Ack, error)
	// Renew lease on chunk.
	RenewChunkLease(ctx context.Context, in *ChunkHandle, opts ...grpc.CallOption) (*Ack, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) SendHeartBeatMessage(ctx context.Context, in *ChunkServerID, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/master.Master/SendHeartBeatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetChunkLocation(ctx context.Context, in *ChunkLocationRequest, opts ...grpc.CallOption) (*ChunkLocationReply, error) {
	out := new(ChunkLocationReply)
	err := c.cc.Invoke(ctx, "/master.Master/GetChunkLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetSystemChunkSize(ctx context.Context, in *SystemChunkSizeRequest, opts ...grpc.CallOption) (*ChunkSize, error) {
	out := new(ChunkSize)
	err := c.cc.Invoke(ctx, "/master.Master/GetSystemChunkSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ReceiveClientWriteRequest(ctx context.Context, in *ClientWriteRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/master.Master/ReceiveClientWriteRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) CreateFile(ctx context.Context, in *FileCreateRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/master.Master/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) RemoveFile(ctx context.Context, in *FileRemoveRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/master.Master/RemoveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) RenewChunkLease(ctx context.Context, in *ChunkHandle, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/master.Master/RenewChunkLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
// All implementations must embed UnimplementedMasterServer
// for forward compatibility
type MasterServer interface {
	// Sends a heart beat message to the client
	SendHeartBeatMessage(context.Context, *ChunkServerID) (*Ack, error)
	// Retrieves the chunk handler of a chunk and the locations of its replicas
	GetChunkLocation(context.Context, *ChunkLocationRequest) (*ChunkLocationReply, error)
	// Retrieve the fixed system chunksize
	GetSystemChunkSize(context.Context, *SystemChunkSizeRequest) (*ChunkSize, error)
	// ClientWriteRequest
	ReceiveClientWriteRequest(context.Context, *ClientWriteRequest) (*Ack, error)
	// Create a file.
	CreateFile(context.Context, *FileCreateRequest) (*Ack, error)
	// Remove a file.
	RemoveFile(context.Context, *FileRemoveRequest) (*Ack, error)
	// Renew lease on chunk.
	RenewChunkLease(context.Context, *ChunkHandle) (*Ack, error)
	mustEmbedUnimplementedMasterServer()
}

// UnimplementedMasterServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (UnimplementedMasterServer) SendHeartBeatMessage(context.Context, *ChunkServerID) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartBeatMessage not implemented")
}
func (UnimplementedMasterServer) GetChunkLocation(context.Context, *ChunkLocationRequest) (*ChunkLocationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunkLocation not implemented")
}
func (UnimplementedMasterServer) GetSystemChunkSize(context.Context, *SystemChunkSizeRequest) (*ChunkSize, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemChunkSize not implemented")
}
func (UnimplementedMasterServer) ReceiveClientWriteRequest(context.Context, *ClientWriteRequest) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveClientWriteRequest not implemented")
}
func (UnimplementedMasterServer) CreateFile(context.Context, *FileCreateRequest) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedMasterServer) RemoveFile(context.Context, *FileRemoveRequest) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFile not implemented")
}
func (UnimplementedMasterServer) RenewChunkLease(context.Context, *ChunkHandle) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewChunkLease not implemented")
}
func (UnimplementedMasterServer) mustEmbedUnimplementedMasterServer() {}

// UnsafeMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServer will
// result in compilation errors.
type UnsafeMasterServer interface {
	mustEmbedUnimplementedMasterServer()
}

func RegisterMasterServer(s grpc.ServiceRegistrar, srv MasterServer) {
	s.RegisterService(&Master_ServiceDesc, srv)
}

func _Master_SendHeartBeatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkServerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).SendHeartBeatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/SendHeartBeatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).SendHeartBeatMessage(ctx, req.(*ChunkServerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetChunkLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetChunkLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/GetChunkLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetChunkLocation(ctx, req.(*ChunkLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetSystemChunkSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemChunkSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetSystemChunkSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/GetSystemChunkSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetSystemChunkSize(ctx, req.(*SystemChunkSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ReceiveClientWriteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ReceiveClientWriteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/ReceiveClientWriteRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ReceiveClientWriteRequest(ctx, req.(*ClientWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreateFile(ctx, req.(*FileCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_RemoveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RemoveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/RemoveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RemoveFile(ctx, req.(*FileRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_RenewChunkLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkHandle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RenewChunkLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/RenewChunkLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RenewChunkLease(ctx, req.(*ChunkHandle))
	}
	return interceptor(ctx, in, info, handler)
}

// Master_ServiceDesc is the grpc.ServiceDesc for Master service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Master_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHeartBeatMessage",
			Handler:    _Master_SendHeartBeatMessage_Handler,
		},
		{
			MethodName: "GetChunkLocation",
			Handler:    _Master_GetChunkLocation_Handler,
		},
		{
			MethodName: "GetSystemChunkSize",
			Handler:    _Master_GetSystemChunkSize_Handler,
		},
		{
			MethodName: "ReceiveClientWriteRequest",
			Handler:    _Master_ReceiveClientWriteRequest_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _Master_CreateFile_Handler,
		},
		{
			MethodName: "RemoveFile",
			Handler:    _Master_RemoveFile_Handler,
		},
		{
			MethodName: "RenewChunkLease",
			Handler:    _Master_RenewChunkLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
