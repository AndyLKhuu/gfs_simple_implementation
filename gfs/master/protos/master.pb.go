// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: master.proto

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

type FileCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	RepFactor int64  `protobuf:"varint,2,opt,name=repFactor,proto3" json:"repFactor,omitempty"`
}

func (x *FileCreateRequest) Reset() {
	*x = FileCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileCreateRequest) ProtoMessage() {}

func (x *FileCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileCreateRequest.ProtoReflect.Descriptor instead.
func (*FileCreateRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

func (x *FileCreateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileCreateRequest) GetRepFactor() int64 {
	if x != nil {
		return x.RepFactor
	}
	return 0
}

type FileRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileRemoveRequest) Reset() {
	*x = FileRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRemoveRequest) ProtoMessage() {}

func (x *FileRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRemoveRequest.ProtoReflect.Descriptor instead.
func (*FileRemoveRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{1}
}

func (x *FileRemoveRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ChunkHandle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ch uint64 `protobuf:"varint,1,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *ChunkHandle) Reset() {
	*x = ChunkHandle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkHandle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkHandle) ProtoMessage() {}

func (x *ChunkHandle) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[2]
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
	return file_master_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkHandle) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

type RenewChunkLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Ch   uint64 `protobuf:"varint,2,opt,name=ch,proto3" json:"ch,omitempty"`
}

func (x *RenewChunkLeaseRequest) Reset() {
	*x = RenewChunkLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewChunkLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewChunkLeaseRequest) ProtoMessage() {}

func (x *RenewChunkLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewChunkLeaseRequest.ProtoReflect.Descriptor instead.
func (*RenewChunkLeaseRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{3}
}

func (x *RenewChunkLeaseRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *RenewChunkLeaseRequest) GetCh() uint64 {
	if x != nil {
		return x.Ch
	}
	return 0
}

// The request message containing the user's name.
type ChunkServerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChunkServerID) Reset() {
	*x = ChunkServerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkServerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkServerID) ProtoMessage() {}

func (x *ChunkServerID) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkServerID.ProtoReflect.Descriptor instead.
func (*ChunkServerID) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{4}
}

func (x *ChunkServerID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The reply message containing the system's fixed chunk size.
type ChunkSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ChunkSize) Reset() {
	*x = ChunkSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkSize) ProtoMessage() {}

func (x *ChunkSize) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkSize.ProtoReflect.Descriptor instead.
func (*ChunkSize) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{5}
}

func (x *ChunkSize) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// The place holder message for a client write request.
type ClientWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientWriteRequest) Reset() {
	*x = ClientWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientWriteRequest) ProtoMessage() {}

func (x *ClientWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientWriteRequest.ProtoReflect.Descriptor instead.
func (*ClientWriteRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{6}
}

// The empty message for a client's inquiry on the system chunk size.
type SystemChunkSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SystemChunkSizeRequest) Reset() {
	*x = SystemChunkSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemChunkSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemChunkSizeRequest) ProtoMessage() {}

func (x *SystemChunkSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemChunkSizeRequest.ProtoReflect.Descriptor instead.
func (*SystemChunkSizeRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{7}
}

// The request message for a file's chunk size
type ChunkLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	ChunkIdx uint32 `protobuf:"varint,2,opt,name=chunkIdx,proto3" json:"chunkIdx,omitempty"`
}

func (x *ChunkLocationRequest) Reset() {
	*x = ChunkLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkLocationRequest) ProtoMessage() {}

func (x *ChunkLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkLocationRequest.ProtoReflect.Descriptor instead.
func (*ChunkLocationRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{8}
}

func (x *ChunkLocationRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ChunkLocationRequest) GetChunkIdx() uint32 {
	if x != nil {
		return x.ChunkIdx
	}
	return 0
}

// The reply message for a file's chunk handler and its replica locations.
type ChunkLocationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkServerIds []string `protobuf:"bytes,1,rep,name=chunkServerIds,proto3" json:"chunkServerIds,omitempty"`
	Primary        string   `protobuf:"bytes,2,opt,name=primary,proto3" json:"primary,omitempty"`
	ChunkHandle    uint64   `protobuf:"varint,3,opt,name=chunkHandle,proto3" json:"chunkHandle,omitempty"`
}

func (x *ChunkLocationReply) Reset() {
	*x = ChunkLocationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkLocationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkLocationReply) ProtoMessage() {}

func (x *ChunkLocationReply) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkLocationReply.ProtoReflect.Descriptor instead.
func (*ChunkLocationReply) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{9}
}

func (x *ChunkLocationReply) GetChunkServerIds() []string {
	if x != nil {
		return x.ChunkServerIds
	}
	return nil
}

func (x *ChunkLocationReply) GetPrimary() string {
	if x != nil {
		return x.Primary
	}
	return ""
}

func (x *ChunkLocationReply) GetChunkHandle() uint64 {
	if x != nil {
		return x.ChunkHandle
	}
	return 0
}

// The response message acknowledging a transaction.
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[10]
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
	return file_master_proto_rawDescGZIP(), []int{10}
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_master_proto protoreflect.FileDescriptor

var file_master_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x70, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x27, 0x0a,
	0x11, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x63, 0x68, 0x22, 0x3c, 0x0a, 0x16, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x63, 0x68, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x78, 0x22, 0x78, 0x0a,
	0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x1f, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xdb, 0x03, 0x0a, 0x06, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x6b, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x19,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41,
	0x63, 0x6b, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41,
	0x63, 0x6b, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x65, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x66, 0x73, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_master_proto_rawDescOnce sync.Once
	file_master_proto_rawDescData = file_master_proto_rawDesc
)

func file_master_proto_rawDescGZIP() []byte {
	file_master_proto_rawDescOnce.Do(func() {
		file_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_proto_rawDescData)
	})
	return file_master_proto_rawDescData
}

var file_master_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_master_proto_goTypes = []interface{}{
	(*FileCreateRequest)(nil),      // 0: master.FileCreateRequest
	(*FileRemoveRequest)(nil),      // 1: master.FileRemoveRequest
	(*ChunkHandle)(nil),            // 2: master.ChunkHandle
	(*RenewChunkLeaseRequest)(nil), // 3: master.RenewChunkLeaseRequest
	(*ChunkServerID)(nil),          // 4: master.ChunkServerID
	(*ChunkSize)(nil),              // 5: master.ChunkSize
	(*ClientWriteRequest)(nil),     // 6: master.ClientWriteRequest
	(*SystemChunkSizeRequest)(nil), // 7: master.SystemChunkSizeRequest
	(*ChunkLocationRequest)(nil),   // 8: master.ChunkLocationRequest
	(*ChunkLocationReply)(nil),     // 9: master.ChunkLocationReply
	(*Ack)(nil),                    // 10: master.Ack
}
var file_master_proto_depIdxs = []int32{
	4,  // 0: master.Master.SendHeartBeatMessage:input_type -> master.ChunkServerID
	8,  // 1: master.Master.GetChunkLocation:input_type -> master.ChunkLocationRequest
	7,  // 2: master.Master.GetSystemChunkSize:input_type -> master.SystemChunkSizeRequest
	6,  // 3: master.Master.ReceiveClientWriteRequest:input_type -> master.ClientWriteRequest
	0,  // 4: master.Master.CreateFile:input_type -> master.FileCreateRequest
	1,  // 5: master.Master.RemoveFile:input_type -> master.FileRemoveRequest
	3,  // 6: master.Master.RenewChunkLease:input_type -> master.RenewChunkLeaseRequest
	10, // 7: master.Master.SendHeartBeatMessage:output_type -> master.Ack
	9,  // 8: master.Master.GetChunkLocation:output_type -> master.ChunkLocationReply
	5,  // 9: master.Master.GetSystemChunkSize:output_type -> master.ChunkSize
	10, // 10: master.Master.ReceiveClientWriteRequest:output_type -> master.Ack
	10, // 11: master.Master.CreateFile:output_type -> master.Ack
	10, // 12: master.Master.RemoveFile:output_type -> master.Ack
	10, // 13: master.Master.RenewChunkLease:output_type -> master.Ack
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_master_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileCreateRequest); i {
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
		file_master_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRemoveRequest); i {
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
		file_master_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_master_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewChunkLeaseRequest); i {
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
		file_master_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkServerID); i {
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
		file_master_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkSize); i {
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
		file_master_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientWriteRequest); i {
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
		file_master_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemChunkSizeRequest); i {
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
		file_master_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkLocationRequest); i {
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
		file_master_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkLocationReply); i {
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
		file_master_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_master_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
		MessageInfos:      file_master_proto_msgTypes,
	}.Build()
	File_master_proto = out.File
	file_master_proto_rawDesc = nil
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}
