// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: chunkserver.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The reply message containing the data from a file read.
type ReadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReadReply) Reset() {
	*x = ReadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReply) ProtoMessage() {}

func (x *ReadReply) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReply.ProtoReflect.Descriptor instead.
func (*ReadReply) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{0}
}

func (x *ReadReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// The request message containing the arguments to a read request.
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L  int32  `protobuf:"varint,1,opt,name=l,proto3" json:"l,omitempty"`
	R  int32  `protobuf:"varint,2,opt,name=r,proto3" json:"r,omitempty"`
	Ch uint64 `protobuf:"varint,3,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRequest) GetL() int32 {
	if x != nil {
		return x.L
	}
	return 0
}

func (x *ReadRequest) GetR() int32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *ReadRequest) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

// A message containing a chunk handle
type ChunkHandle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ch uint64 `protobuf:"varint,1,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *ChunkHandle) Reset() {
	*x = ChunkHandle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkHandle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkHandle) ProtoMessage() {}

func (x *ChunkHandle) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkHandle.ProtoReflect.Descriptor instead.
func (*ChunkHandle) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkHandle) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

type PrimaryCommitMutateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecondaryChunkServerAddresses []string `protobuf:"bytes,1,rep,name=secondaryChunkServerAddresses,proto3" json:"secondaryChunkServerAddresses,omitempty"`
	TxId                          string   `protobuf:"bytes,2,opt,name=txId,proto3" json:"txId,omitempty"`
	Ch                            uint64   `protobuf:"varint,3,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *PrimaryCommitMutateRequest) Reset() {
	*x = PrimaryCommitMutateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryCommitMutateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryCommitMutateRequest) ProtoMessage() {}

func (x *PrimaryCommitMutateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryCommitMutateRequest.ProtoReflect.Descriptor instead.
func (*PrimaryCommitMutateRequest) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{3}
}

func (x *PrimaryCommitMutateRequest) GetSecondaryChunkServerAddresses() []string {
	if x != nil {
		return x.SecondaryChunkServerAddresses
	}
	return nil
}

func (x *PrimaryCommitMutateRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *PrimaryCommitMutateRequest) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

type SecondaryCommitMutateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxIds []string `protobuf:"bytes,1,rep,name=txIds,proto3" json:"txIds,omitempty"`
	Ch    uint64   `protobuf:"varint,2,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *SecondaryCommitMutateRequest) Reset() {
	*x = SecondaryCommitMutateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondaryCommitMutateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondaryCommitMutateRequest) ProtoMessage() {}

func (x *SecondaryCommitMutateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondaryCommitMutateRequest.ProtoReflect.Descriptor instead.
func (*SecondaryCommitMutateRequest) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{4}
}

func (x *SecondaryCommitMutateRequest) GetTxIds() []string {
	if x != nil {
		return x.TxIds
	}
	return nil
}

func (x *SecondaryCommitMutateRequest) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

// A reply message containing an acknowledgement
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{5}
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A request message containing potential mutation data from clients.
type WriteDataBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	TxId   string `protobuf:"bytes,2,opt,name=txId,proto3" json:"txId,omitempty"`
	Size   uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Offset uint64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Ch     uint64 `protobuf:"varint,5,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *WriteDataBundle) Reset() {
	*x = WriteDataBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteDataBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteDataBundle) ProtoMessage() {}

func (x *WriteDataBundle) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteDataBundle.ProtoReflect.Descriptor instead.
func (*WriteDataBundle) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{6}
}

func (x *WriteDataBundle) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WriteDataBundle) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *WriteDataBundle) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WriteDataBundle) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *WriteDataBundle) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

// A request message containing a lease for a chunk
type LeaseBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeEnd int64  `protobuf:"varint,1,opt,name=timeEnd,proto3" json:"timeEnd,omitempty"`
	Ch      uint64 `protobuf:"varint,2,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *LeaseBundle) Reset() {
	*x = LeaseBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaseBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaseBundle) ProtoMessage() {}

func (x *LeaseBundle) ProtoReflect() protoreflect.Message {
	mi := &file_chunkserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaseBundle.ProtoReflect.Descriptor instead.
func (*LeaseBundle) Descriptor() ([]byte, []int) {
	return file_chunkserver_proto_rawDescGZIP(), []int{7}
}

func (x *LeaseBundle) GetTimeEnd() int64 {
	if x != nil {
		return x.TimeEnd
	}
	return 0
}

func (x *LeaseBundle) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

var File_chunkserver_proto protoreflect.FileDescriptor

var file_chunkserver_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x1f, 0x0a, 0x09, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x63, 0x68, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x63, 0x68, 0x22, 0x86, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x1d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x78, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x63, 0x68, 0x22,
	0x44, 0x0a, 0x1c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x78, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x78, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x63, 0x68, 0x22, 0x1f, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x75, 0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x78, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x63, 0x68, 0x22, 0x37, 0x0a,
	0x0b, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x63, 0x68, 0x32, 0xb0, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x63, 0x6b, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x15, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c,
	0x65, 0x61, 0x73, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x66, 0x73,
	0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chunkserver_proto_rawDescOnce sync.Once
	file_chunkserver_proto_rawDescData = file_chunkserver_proto_rawDesc
)

func file_chunkserver_proto_rawDescGZIP() []byte {
	file_chunkserver_proto_rawDescOnce.Do(func() {
		file_chunkserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_chunkserver_proto_rawDescData)
	})
	return file_chunkserver_proto_rawDescData
}

var file_chunkserver_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chunkserver_proto_goTypes = []interface{}{
	(*ReadReply)(nil),                    // 0: protos.ReadReply
	(*ReadRequest)(nil),                  // 1: protos.ReadRequest
	(*ChunkHandle)(nil),                  // 2: protos.ChunkHandle
	(*PrimaryCommitMutateRequest)(nil),   // 3: protos.PrimaryCommitMutateRequest
	(*SecondaryCommitMutateRequest)(nil), // 4: protos.SecondaryCommitMutateRequest
	(*Ack)(nil),                          // 5: protos.Ack
	(*WriteDataBundle)(nil),              // 6: protos.WriteDataBundle
	(*LeaseBundle)(nil),                  // 7: protos.LeaseBundle
}
var file_chunkserver_proto_depIdxs = []int32{
	1, // 0: protos.ChunkServer.Read:input_type -> protos.ReadRequest
	6, // 1: protos.ChunkServer.ReceiveWriteData:input_type -> protos.WriteDataBundle
	3, // 2: protos.ChunkServer.PrimaryCommitMutate:input_type -> protos.PrimaryCommitMutateRequest
	4, // 3: protos.ChunkServer.SecondaryCommitMutate:input_type -> protos.SecondaryCommitMutateRequest
	2, // 4: protos.ChunkServer.CreateNewChunk:input_type -> protos.ChunkHandle
	2, // 5: protos.ChunkServer.RemoveChunk:input_type -> protos.ChunkHandle
	7, // 6: protos.ChunkServer.ReceiveLease:input_type -> protos.LeaseBundle
	0, // 7: protos.ChunkServer.Read:output_type -> protos.ReadReply
	5, // 8: protos.ChunkServer.ReceiveWriteData:output_type -> protos.Ack
	5, // 9: protos.ChunkServer.PrimaryCommitMutate:output_type -> protos.Ack
	5, // 10: protos.ChunkServer.SecondaryCommitMutate:output_type -> protos.Ack
	5, // 11: protos.ChunkServer.CreateNewChunk:output_type -> protos.Ack
	5, // 12: protos.ChunkServer.RemoveChunk:output_type -> protos.Ack
	5, // 13: protos.ChunkServer.ReceiveLease:output_type -> protos.Ack
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chunkserver_proto_init() }
func file_chunkserver_proto_init() {
	if File_chunkserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chunkserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkHandle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryCommitMutateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondaryCommitMutateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteDataBundle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chunkserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaseBundle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chunkserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chunkserver_proto_goTypes,
		DependencyIndexes: file_chunkserver_proto_depIdxs,
		MessageInfos:      file_chunkserver_proto_msgTypes,
	}.Build()
	File_chunkserver_proto = out.File
	file_chunkserver_proto_rawDesc = nil
	file_chunkserver_proto_goTypes = nil
	file_chunkserver_proto_depIdxs = nil
}
