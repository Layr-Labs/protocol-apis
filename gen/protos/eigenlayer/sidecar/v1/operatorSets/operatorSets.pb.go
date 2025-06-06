// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/operatorSets/operatorSets.proto

package operatorSets

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListOperatorsForStakerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StakerAddress string                 `protobuf:"bytes,1,opt,name=stakerAddress,proto3" json:"stakerAddress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperatorsForStakerRequest) Reset() {
	*x = ListOperatorsForStakerRequest{}
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperatorsForStakerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorsForStakerRequest) ProtoMessage() {}

func (x *ListOperatorsForStakerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorsForStakerRequest.ProtoReflect.Descriptor instead.
func (*ListOperatorsForStakerRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP(), []int{0}
}

func (x *ListOperatorsForStakerRequest) GetStakerAddress() string {
	if x != nil {
		return x.StakerAddress
	}
	return ""
}

type ListOperatorsForStakerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operators     []*Operator            `protobuf:"bytes,1,rep,name=operators,proto3" json:"operators,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperatorsForStakerResponse) Reset() {
	*x = ListOperatorsForStakerResponse{}
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperatorsForStakerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorsForStakerResponse) ProtoMessage() {}

func (x *ListOperatorsForStakerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorsForStakerResponse.ProtoReflect.Descriptor instead.
func (*ListOperatorsForStakerResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP(), []int{1}
}

func (x *ListOperatorsForStakerResponse) GetOperators() []*Operator {
	if x != nil {
		return x.Operators
	}
	return nil
}

type ListOperatorsForStrategyRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	StrategyAddress string                 `protobuf:"bytes,1,opt,name=strategyAddress,proto3" json:"strategyAddress,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListOperatorsForStrategyRequest) Reset() {
	*x = ListOperatorsForStrategyRequest{}
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperatorsForStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorsForStrategyRequest) ProtoMessage() {}

func (x *ListOperatorsForStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorsForStrategyRequest.ProtoReflect.Descriptor instead.
func (*ListOperatorsForStrategyRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP(), []int{2}
}

func (x *ListOperatorsForStrategyRequest) GetStrategyAddress() string {
	if x != nil {
		return x.StrategyAddress
	}
	return ""
}

type ListOperatorsForStrategyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operators     []*Operator            `protobuf:"bytes,1,rep,name=operators,proto3" json:"operators,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperatorsForStrategyResponse) Reset() {
	*x = ListOperatorsForStrategyResponse{}
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperatorsForStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorsForStrategyResponse) ProtoMessage() {}

func (x *ListOperatorsForStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorsForStrategyResponse.ProtoReflect.Descriptor instead.
func (*ListOperatorsForStrategyResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP(), []int{3}
}

func (x *ListOperatorsForStrategyResponse) GetOperators() []*Operator {
	if x != nil {
		return x.Operators
	}
	return nil
}

type ListOperatorsForAvsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AvsAddress    string                 `protobuf:"bytes,1,opt,name=avsAddress,proto3" json:"avsAddress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperatorsForAvsRequest) Reset() {
	*x = ListOperatorsForAvsRequest{}
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperatorsForAvsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorsForAvsRequest) ProtoMessage() {}

func (x *ListOperatorsForAvsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorsForAvsRequest.ProtoReflect.Descriptor instead.
func (*ListOperatorsForAvsRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP(), []int{4}
}

func (x *ListOperatorsForAvsRequest) GetAvsAddress() string {
	if x != nil {
		return x.AvsAddress
	}
	return ""
}

type ListOperatorsForAvsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operators     []*Operator            `protobuf:"bytes,1,rep,name=operators,proto3" json:"operators,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperatorsForAvsResponse) Reset() {
	*x = ListOperatorsForAvsResponse{}
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperatorsForAvsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperatorsForAvsResponse) ProtoMessage() {}

func (x *ListOperatorsForAvsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperatorsForAvsResponse.ProtoReflect.Descriptor instead.
func (*ListOperatorsForAvsResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP(), []int{5}
}

func (x *ListOperatorsForAvsResponse) GetOperators() []*Operator {
	if x != nil {
		return x.Operators
	}
	return nil
}

var File_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDesc = string([]byte{
	0x0a, 0x35, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x6c, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x22, 0x4b, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a,
	0x20, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f,
	0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x3c, 0x0a,
	0x1a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f,
	0x72, 0x41, 0x76, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x76, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x76, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x69, 0x0a, 0x1b, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x41,
	0x76, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x42, 0xb9, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x73, 0x42, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x56, 0x4f, 0xaa,
	0x02, 0x22, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2e, 0x56, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x74, 0x73, 0xca, 0x02, 0x22, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0xe2, 0x02, 0x2e, 0x45, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x45, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescOnce sync.Once
	file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescData []byte
)

func file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescGZIP() []byte {
	file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescOnce.Do(func() {
		file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDesc), len(file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDesc)))
	})
	return file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDescData
}

var file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_goTypes = []any{
	(*ListOperatorsForStakerRequest)(nil),    // 0: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerRequest
	(*ListOperatorsForStakerResponse)(nil),   // 1: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerResponse
	(*ListOperatorsForStrategyRequest)(nil),  // 2: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyRequest
	(*ListOperatorsForStrategyResponse)(nil), // 3: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyResponse
	(*ListOperatorsForAvsRequest)(nil),       // 4: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsRequest
	(*ListOperatorsForAvsResponse)(nil),      // 5: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsResponse
	(*Operator)(nil),                         // 6: eigenlayer.sidecar.v1.operatorSets.Operator
}
var file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_depIdxs = []int32{
	6, // 0: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerResponse.operators:type_name -> eigenlayer.sidecar.v1.operatorSets.Operator
	6, // 1: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyResponse.operators:type_name -> eigenlayer.sidecar.v1.operatorSets.Operator
	6, // 2: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsResponse.operators:type_name -> eigenlayer.sidecar.v1.operatorSets.Operator
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_init() }
func file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_init() {
	if File_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto != nil {
		return
	}
	file_eigenlayer_sidecar_v1_operatorSets_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDesc), len(file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_depIdxs,
		MessageInfos:      file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_msgTypes,
	}.Build()
	File_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto = out.File
	file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_depIdxs = nil
}
