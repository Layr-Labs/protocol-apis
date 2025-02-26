// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/slashing/common.proto

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

type SlashingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description     string                     `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Operator        string                     `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	TransactionHash string                     `protobuf:"bytes,3,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	LogIndex        uint64                     `protobuf:"varint,4,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	BlockNumber     uint64                     `protobuf:"varint,5,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Strategies      []*SlashingEvent_Strategy  `protobuf:"bytes,6,rep,name=strategies,proto3" json:"strategies,omitempty"`
	OperatorSet     *SlashingEvent_OperatorSet `protobuf:"bytes,7,opt,name=operatorSet,proto3" json:"operatorSet,omitempty"`
}

func (x *SlashingEvent) Reset() {
	*x = SlashingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlashingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlashingEvent) ProtoMessage() {}

func (x *SlashingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlashingEvent.ProtoReflect.Descriptor instead.
func (*SlashingEvent) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescGZIP(), []int{0}
}

func (x *SlashingEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SlashingEvent) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *SlashingEvent) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *SlashingEvent) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *SlashingEvent) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *SlashingEvent) GetStrategies() []*SlashingEvent_Strategy {
	if x != nil {
		return x.Strategies
	}
	return nil
}

func (x *SlashingEvent) GetOperatorSet() *SlashingEvent_OperatorSet {
	if x != nil {
		return x.OperatorSet
	}
	return nil
}

type SlashedStaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staker          string `protobuf:"bytes,1,opt,name=staker,proto3" json:"staker,omitempty"`
	TransactionHash string `protobuf:"bytes,2,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	LogIndex        uint64 `protobuf:"varint,3,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	BlockNumber     uint64 `protobuf:"varint,4,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Strategy        string `protobuf:"bytes,5,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Shares          string `protobuf:"bytes,6,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (x *SlashedStaker) Reset() {
	*x = SlashedStaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlashedStaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlashedStaker) ProtoMessage() {}

func (x *SlashedStaker) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlashedStaker.ProtoReflect.Descriptor instead.
func (*SlashedStaker) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescGZIP(), []int{1}
}

func (x *SlashedStaker) GetStaker() string {
	if x != nil {
		return x.Staker
	}
	return ""
}

func (x *SlashedStaker) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *SlashedStaker) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *SlashedStaker) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *SlashedStaker) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *SlashedStaker) GetShares() string {
	if x != nil {
		return x.Shares
	}
	return ""
}

type StakerSlashingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staker        string         `protobuf:"bytes,1,opt,name=staker,proto3" json:"staker,omitempty"`
	SlashingEvent *SlashingEvent `protobuf:"bytes,2,opt,name=slashingEvent,proto3" json:"slashingEvent,omitempty"`
}

func (x *StakerSlashingEvent) Reset() {
	*x = StakerSlashingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakerSlashingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakerSlashingEvent) ProtoMessage() {}

func (x *StakerSlashingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakerSlashingEvent.ProtoReflect.Descriptor instead.
func (*StakerSlashingEvent) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescGZIP(), []int{2}
}

func (x *StakerSlashingEvent) GetStaker() string {
	if x != nil {
		return x.Staker
	}
	return ""
}

func (x *StakerSlashingEvent) GetSlashingEvent() *SlashingEvent {
	if x != nil {
		return x.SlashingEvent
	}
	return nil
}

type SlashingEvent_Strategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strategy           string `protobuf:"bytes,1,opt,name=strategy,proto3" json:"strategy,omitempty"`
	WadSlashed         string `protobuf:"bytes,2,opt,name=wadSlashed,proto3" json:"wadSlashed,omitempty"`
	TotalSharesSlashed string `protobuf:"bytes,3,opt,name=totalSharesSlashed,proto3" json:"totalSharesSlashed,omitempty"`
}

func (x *SlashingEvent_Strategy) Reset() {
	*x = SlashingEvent_Strategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlashingEvent_Strategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlashingEvent_Strategy) ProtoMessage() {}

func (x *SlashingEvent_Strategy) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlashingEvent_Strategy.ProtoReflect.Descriptor instead.
func (*SlashingEvent_Strategy) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SlashingEvent_Strategy) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *SlashingEvent_Strategy) GetWadSlashed() string {
	if x != nil {
		return x.WadSlashed
	}
	return ""
}

func (x *SlashingEvent_Strategy) GetTotalSharesSlashed() string {
	if x != nil {
		return x.TotalSharesSlashed
	}
	return ""
}

type SlashingEvent_OperatorSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Avs string `protobuf:"bytes,2,opt,name=avs,proto3" json:"avs,omitempty"`
}

func (x *SlashingEvent_OperatorSet) Reset() {
	*x = SlashingEvent_OperatorSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlashingEvent_OperatorSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlashingEvent_OperatorSet) ProtoMessage() {}

func (x *SlashingEvent_OperatorSet) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlashingEvent_OperatorSet.ProtoReflect.Descriptor instead.
func (*SlashingEvent_OperatorSet) Descriptor() ([]byte, []int) {
	return file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SlashingEvent_OperatorSet) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SlashingEvent_OperatorSet) GetAvs() string {
	if x != nil {
		return x.Avs
	}
	return ""
}

var File_eigenlayer_sidecar_v1_slashing_common_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_slashing_common_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65,
	0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x04, 0x0a,
	0x0d, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x52, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x0b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x1a, 0x76, 0x0a, 0x08, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x64, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x64, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65,
	0x64, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73,
	0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65,
	0x64, 0x1a, 0x2f, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x76, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x0d, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0d, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0d,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x9b, 0x02,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x56, 0x53, 0xaa, 0x02, 0x1e, 0x45, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x1e, 0x45, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x2a, 0x45, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescOnce sync.Once
	file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescData = file_eigenlayer_sidecar_v1_slashing_common_proto_rawDesc
)

func file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescGZIP() []byte {
	file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescOnce.Do(func() {
		file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescData)
	})
	return file_eigenlayer_sidecar_v1_slashing_common_proto_rawDescData
}

var file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eigenlayer_sidecar_v1_slashing_common_proto_goTypes = []any{
	(*SlashingEvent)(nil),             // 0: eigenlayer.sidecar.v1.slashing.SlashingEvent
	(*SlashedStaker)(nil),             // 1: eigenlayer.sidecar.v1.slashing.SlashedStaker
	(*StakerSlashingEvent)(nil),       // 2: eigenlayer.sidecar.v1.slashing.StakerSlashingEvent
	(*SlashingEvent_Strategy)(nil),    // 3: eigenlayer.sidecar.v1.slashing.SlashingEvent.Strategy
	(*SlashingEvent_OperatorSet)(nil), // 4: eigenlayer.sidecar.v1.slashing.SlashingEvent.OperatorSet
}
var file_eigenlayer_sidecar_v1_slashing_common_proto_depIdxs = []int32{
	3, // 0: eigenlayer.sidecar.v1.slashing.SlashingEvent.strategies:type_name -> eigenlayer.sidecar.v1.slashing.SlashingEvent.Strategy
	4, // 1: eigenlayer.sidecar.v1.slashing.SlashingEvent.operatorSet:type_name -> eigenlayer.sidecar.v1.slashing.SlashingEvent.OperatorSet
	0, // 2: eigenlayer.sidecar.v1.slashing.StakerSlashingEvent.slashingEvent:type_name -> eigenlayer.sidecar.v1.slashing.SlashingEvent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_slashing_common_proto_init() }
func file_eigenlayer_sidecar_v1_slashing_common_proto_init() {
	if File_eigenlayer_sidecar_v1_slashing_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SlashingEvent); i {
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
		file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SlashedStaker); i {
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
		file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StakerSlashingEvent); i {
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
		file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SlashingEvent_Strategy); i {
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
		file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SlashingEvent_OperatorSet); i {
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
			RawDescriptor: file_eigenlayer_sidecar_v1_slashing_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_slashing_common_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_slashing_common_proto_depIdxs,
		MessageInfos:      file_eigenlayer_sidecar_v1_slashing_common_proto_msgTypes,
	}.Build()
	File_eigenlayer_sidecar_v1_slashing_common_proto = out.File
	file_eigenlayer_sidecar_v1_slashing_common_proto_rawDesc = nil
	file_eigenlayer_sidecar_v1_slashing_common_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_slashing_common_proto_depIdxs = nil
}
