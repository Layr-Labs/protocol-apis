// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/sidecar/sidecar.proto

package sidecar

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetBlockHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verified bool `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *GetBlockHeightRequest) Reset() {
	*x = GetBlockHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeightRequest) ProtoMessage() {}

func (x *GetBlockHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeightRequest.ProtoReflect.Descriptor instead.
func (*GetBlockHeightRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockHeightRequest) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type GetBlockHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockHash   string `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
}

func (x *GetBlockHeightResponse) Reset() {
	*x = GetBlockHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeightResponse) ProtoMessage() {}

func (x *GetBlockHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeightResponse.ProtoReflect.Descriptor instead.
func (*GetBlockHeightResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockHeightResponse) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *GetBlockHeightResponse) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

type GetStateRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
}

func (x *GetStateRootRequest) Reset() {
	*x = GetStateRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateRootRequest) ProtoMessage() {}

func (x *GetStateRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateRootRequest.ProtoReflect.Descriptor instead.
func (*GetStateRootRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP(), []int{2}
}

func (x *GetStateRootRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

type GetStateRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthBlockNumber uint64 `protobuf:"varint,1,opt,name=ethBlockNumber,proto3" json:"ethBlockNumber,omitempty"`
	EthBlockHash   string `protobuf:"bytes,2,opt,name=ethBlockHash,proto3" json:"ethBlockHash,omitempty"`
	StateRoot      string `protobuf:"bytes,3,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
}

func (x *GetStateRootResponse) Reset() {
	*x = GetStateRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateRootResponse) ProtoMessage() {}

func (x *GetStateRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateRootResponse.ProtoReflect.Descriptor instead.
func (*GetStateRootResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP(), []int{3}
}

func (x *GetStateRootResponse) GetEthBlockNumber() uint64 {
	if x != nil {
		return x.EthBlockNumber
	}
	return 0
}

func (x *GetStateRootResponse) GetEthBlockHash() string {
	if x != nil {
		return x.EthBlockHash
	}
	return ""
}

func (x *GetStateRootResponse) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

type AboutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AboutRequest) Reset() {
	*x = AboutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AboutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AboutRequest) ProtoMessage() {}

func (x *AboutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AboutRequest.ProtoReflect.Descriptor instead.
func (*AboutRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP(), []int{4}
}

type AboutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Commit  string `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	Chain   string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *AboutResponse) Reset() {
	*x = AboutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AboutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AboutResponse) ProtoMessage() {}

func (x *AboutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AboutResponse.ProtoReflect.Descriptor instead.
func (*AboutResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP(), []int{5}
}

func (x *AboutResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AboutResponse) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *AboutResponse) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

var File_eigenlayer_sidecar_v1_sidecar_sidecar_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f,
	0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x65,
	0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22,
	0x58, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x65,
	0x74, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x74, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x74, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x74, 0x68, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x0d, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x42, 0x96,
	0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x42, 0x0c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x56, 0x53, 0xaa, 0x02, 0x1d, 0x45, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0xca, 0x02, 0x1d, 0x45, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0xe2, 0x02, 0x29, 0x45, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x3a, 0x3a, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescOnce sync.Once
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescData = file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDesc
)

func file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescGZIP() []byte {
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescOnce.Do(func() {
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescData = protoimpl.X.CompressGZIP(file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescData)
	})
	return file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDescData
}

var file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_goTypes = []any{
	(*GetBlockHeightRequest)(nil),  // 0: eigenlayer.sidecar.v1.sidecar.GetBlockHeightRequest
	(*GetBlockHeightResponse)(nil), // 1: eigenlayer.sidecar.v1.sidecar.GetBlockHeightResponse
	(*GetStateRootRequest)(nil),    // 2: eigenlayer.sidecar.v1.sidecar.GetStateRootRequest
	(*GetStateRootResponse)(nil),   // 3: eigenlayer.sidecar.v1.sidecar.GetStateRootResponse
	(*AboutRequest)(nil),           // 4: eigenlayer.sidecar.v1.sidecar.AboutRequest
	(*AboutResponse)(nil),          // 5: eigenlayer.sidecar.v1.sidecar.AboutResponse
}
var file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_init() }
func file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_init() {
	if File_eigenlayer_sidecar_v1_sidecar_sidecar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetBlockHeightRequest); i {
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
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetBlockHeightResponse); i {
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
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetStateRootRequest); i {
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
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetStateRootResponse); i {
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
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AboutRequest); i {
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
		file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AboutResponse); i {
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
			RawDescriptor: file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_depIdxs,
		MessageInfos:      file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_msgTypes,
	}.Build()
	File_eigenlayer_sidecar_v1_sidecar_sidecar_proto = out.File
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_rawDesc = nil
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_depIdxs = nil
}
