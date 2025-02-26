// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/slashing/slashing.proto

package slashing

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListStakerSlashingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StakerAddress string `protobuf:"bytes,1,opt,name=stakerAddress,proto3" json:"stakerAddress,omitempty"`
}

func (x *ListStakerSlashingHistoryRequest) Reset() {
	*x = ListStakerSlashingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStakerSlashingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStakerSlashingHistoryRequest) ProtoMessage() {}

func (x *ListStakerSlashingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStakerSlashingHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListStakerSlashingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{0}
}

func (x *ListStakerSlashingHistoryRequest) GetStakerAddress() string {
	if x != nil {
		return x.StakerAddress
	}
	return ""
}

type ListStakerSlashingHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StakerSlashingEvents []*StakerSlashingEvent `protobuf:"bytes,1,rep,name=stakerSlashingEvents,proto3" json:"stakerSlashingEvents,omitempty"`
}

func (x *ListStakerSlashingHistoryResponse) Reset() {
	*x = ListStakerSlashingHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStakerSlashingHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStakerSlashingHistoryResponse) ProtoMessage() {}

func (x *ListStakerSlashingHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStakerSlashingHistoryResponse.ProtoReflect.Descriptor instead.
func (*ListStakerSlashingHistoryResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{1}
}

func (x *ListStakerSlashingHistoryResponse) GetStakerSlashingEvents() []*StakerSlashingEvent {
	if x != nil {
		return x.StakerSlashingEvents
	}
	return nil
}

type ListOperatorSlashingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorAddress string `protobuf:"bytes,1,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
}

func (x *ListOperatorSlashingHistoryRequest) Reset() {
	*x = ListOperatorSlashingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOperatorSlashingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorSlashingHistoryRequest) ProtoMessage() {}

func (x *ListOperatorSlashingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorSlashingHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListOperatorSlashingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{2}
}

func (x *ListOperatorSlashingHistoryRequest) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

type ListOperatorSlashingHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlashingEvents []*SlashingEvent `protobuf:"bytes,1,rep,name=slashingEvents,proto3" json:"slashingEvents,omitempty"`
}

func (x *ListOperatorSlashingHistoryResponse) Reset() {
	*x = ListOperatorSlashingHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOperatorSlashingHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorSlashingHistoryResponse) ProtoMessage() {}

func (x *ListOperatorSlashingHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorSlashingHistoryResponse.ProtoReflect.Descriptor instead.
func (*ListOperatorSlashingHistoryResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{3}
}

func (x *ListOperatorSlashingHistoryResponse) GetSlashingEvents() []*SlashingEvent {
	if x != nil {
		return x.SlashingEvents
	}
	return nil
}

type ListAvsSlashingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvsAddress string `protobuf:"bytes,1,opt,name=avsAddress,proto3" json:"avsAddress,omitempty"`
}

func (x *ListAvsSlashingHistoryRequest) Reset() {
	*x = ListAvsSlashingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAvsSlashingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvsSlashingHistoryRequest) ProtoMessage() {}

func (x *ListAvsSlashingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvsSlashingHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListAvsSlashingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{4}
}

func (x *ListAvsSlashingHistoryRequest) GetAvsAddress() string {
	if x != nil {
		return x.AvsAddress
	}
	return ""
}

type ListAvsSlashingHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlashingEvents []*SlashingEvent `protobuf:"bytes,1,rep,name=slashingEvents,proto3" json:"slashingEvents,omitempty"`
}

func (x *ListAvsSlashingHistoryResponse) Reset() {
	*x = ListAvsSlashingHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAvsSlashingHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvsSlashingHistoryResponse) ProtoMessage() {}

func (x *ListAvsSlashingHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvsSlashingHistoryResponse.ProtoReflect.Descriptor instead.
func (*ListAvsSlashingHistoryResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{5}
}

func (x *ListAvsSlashingHistoryResponse) GetSlashingEvents() []*SlashingEvent {
	if x != nil {
		return x.SlashingEvents
	}
	return nil
}

type ListAvsOperatorSetSlashingHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvsAddress    string `protobuf:"bytes,1,opt,name=avsAddress,proto3" json:"avsAddress,omitempty"`
	OperatorSetId uint64 `protobuf:"varint,2,opt,name=operatorSetId,proto3" json:"operatorSetId,omitempty"`
}

func (x *ListAvsOperatorSetSlashingHistoryRequest) Reset() {
	*x = ListAvsOperatorSetSlashingHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAvsOperatorSetSlashingHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvsOperatorSetSlashingHistoryRequest) ProtoMessage() {}

func (x *ListAvsOperatorSetSlashingHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvsOperatorSetSlashingHistoryRequest.ProtoReflect.Descriptor instead.
func (*ListAvsOperatorSetSlashingHistoryRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{6}
}

func (x *ListAvsOperatorSetSlashingHistoryRequest) GetAvsAddress() string {
	if x != nil {
		return x.AvsAddress
	}
	return ""
}

func (x *ListAvsOperatorSetSlashingHistoryRequest) GetOperatorSetId() uint64 {
	if x != nil {
		return x.OperatorSetId
	}
	return 0
}

type ListAvsOperatorSetSlashingHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlashingEvents []*SlashingEvent `protobuf:"bytes,1,rep,name=slashingEvents,proto3" json:"slashingEvents,omitempty"`
}

func (x *ListAvsOperatorSetSlashingHistoryResponse) Reset() {
	*x = ListAvsOperatorSetSlashingHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAvsOperatorSetSlashingHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvsOperatorSetSlashingHistoryResponse) ProtoMessage() {}

func (x *ListAvsOperatorSetSlashingHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvsOperatorSetSlashingHistoryResponse.ProtoReflect.Descriptor instead.
func (*ListAvsOperatorSetSlashingHistoryResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP(), []int{7}
}

func (x *ListAvsOperatorSetSlashingHistoryResponse) GetSlashingEvents() []*SlashingEvent {
	if x != nil {
		return x.SlashingEvents
	}
	return nil
}

var File_eigenlayer_sidecar_v1_slashing_slashing_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x20, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x14, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x14,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x4e, 0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x7c, 0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x0e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x3f, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x73, 0x53, 0x6c, 0x61,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x76, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x73, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x77, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x73, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x28,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x74, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x76, 0x73, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76,
	0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x82,
	0x01, 0x0a, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x0e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x9d, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x0d, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x56, 0x53,
	0xaa, 0x02, 0x1e, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0xca, 0x02, 0x1e, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0xe2, 0x02, 0x2a, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c,
	0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x21, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescOnce sync.Once
	file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescData = file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDesc
)

func file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescGZIP() []byte {
	file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescOnce.Do(func() {
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescData = protoimpl.X.CompressGZIP(file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescData)
	})
	return file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDescData
}

var file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eigenlayer_sidecar_v1_slashing_slashing_proto_goTypes = []any{
	(*ListStakerSlashingHistoryRequest)(nil),          // 0: eigenlayer.sidecar.v1.slashing.ListStakerSlashingHistoryRequest
	(*ListStakerSlashingHistoryResponse)(nil),         // 1: eigenlayer.sidecar.v1.slashing.ListStakerSlashingHistoryResponse
	(*ListOperatorSlashingHistoryRequest)(nil),        // 2: eigenlayer.sidecar.v1.slashing.ListOperatorSlashingHistoryRequest
	(*ListOperatorSlashingHistoryResponse)(nil),       // 3: eigenlayer.sidecar.v1.slashing.ListOperatorSlashingHistoryResponse
	(*ListAvsSlashingHistoryRequest)(nil),             // 4: eigenlayer.sidecar.v1.slashing.ListAvsSlashingHistoryRequest
	(*ListAvsSlashingHistoryResponse)(nil),            // 5: eigenlayer.sidecar.v1.slashing.ListAvsSlashingHistoryResponse
	(*ListAvsOperatorSetSlashingHistoryRequest)(nil),  // 6: eigenlayer.sidecar.v1.slashing.ListAvsOperatorSetSlashingHistoryRequest
	(*ListAvsOperatorSetSlashingHistoryResponse)(nil), // 7: eigenlayer.sidecar.v1.slashing.ListAvsOperatorSetSlashingHistoryResponse
	(*StakerSlashingEvent)(nil),                       // 8: eigenlayer.sidecar.v1.slashing.StakerSlashingEvent
	(*SlashingEvent)(nil),                             // 9: eigenlayer.sidecar.v1.slashing.SlashingEvent
}
var file_eigenlayer_sidecar_v1_slashing_slashing_proto_depIdxs = []int32{
	8, // 0: eigenlayer.sidecar.v1.slashing.ListStakerSlashingHistoryResponse.stakerSlashingEvents:type_name -> eigenlayer.sidecar.v1.slashing.StakerSlashingEvent
	9, // 1: eigenlayer.sidecar.v1.slashing.ListOperatorSlashingHistoryResponse.slashingEvents:type_name -> eigenlayer.sidecar.v1.slashing.SlashingEvent
	9, // 2: eigenlayer.sidecar.v1.slashing.ListAvsSlashingHistoryResponse.slashingEvents:type_name -> eigenlayer.sidecar.v1.slashing.SlashingEvent
	9, // 3: eigenlayer.sidecar.v1.slashing.ListAvsOperatorSetSlashingHistoryResponse.slashingEvents:type_name -> eigenlayer.sidecar.v1.slashing.SlashingEvent
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_slashing_slashing_proto_init() }
func file_eigenlayer_sidecar_v1_slashing_slashing_proto_init() {
	if File_eigenlayer_sidecar_v1_slashing_slashing_proto != nil {
		return
	}
	file_eigenlayer_sidecar_v1_slashing_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListStakerSlashingHistoryRequest); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListStakerSlashingHistoryResponse); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListOperatorSlashingHistoryRequest); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListOperatorSlashingHistoryResponse); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListAvsSlashingHistoryRequest); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListAvsSlashingHistoryResponse); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListAvsOperatorSetSlashingHistoryRequest); i {
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
		file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListAvsOperatorSetSlashingHistoryResponse); i {
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
			RawDescriptor: file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_slashing_slashing_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_slashing_slashing_proto_depIdxs,
		MessageInfos:      file_eigenlayer_sidecar_v1_slashing_slashing_proto_msgTypes,
	}.Build()
	File_eigenlayer_sidecar_v1_slashing_slashing_proto = out.File
	file_eigenlayer_sidecar_v1_slashing_slashing_proto_rawDesc = nil
	file_eigenlayer_sidecar_v1_slashing_slashing_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_slashing_slashing_proto_depIdxs = nil
}
