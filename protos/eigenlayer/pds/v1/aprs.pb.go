// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/pds/v1/aprs.proto

package v1

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

type GetDailyOperatorStrategyAprsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	Date            string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetDailyOperatorStrategyAprsRequest) Reset() {
	*x = GetDailyOperatorStrategyAprsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyOperatorStrategyAprsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyOperatorStrategyAprsRequest) ProtoMessage() {}

func (x *GetDailyOperatorStrategyAprsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyOperatorStrategyAprsRequest.ProtoReflect.Descriptor instead.
func (*GetDailyOperatorStrategyAprsRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_pds_v1_aprs_proto_rawDescGZIP(), []int{0}
}

func (x *GetDailyOperatorStrategyAprsRequest) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

func (x *GetDailyOperatorStrategyAprsRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type OperatorStrategyApr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strategy string `protobuf:"bytes,1,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Apr      string `protobuf:"bytes,2,opt,name=apr,proto3" json:"apr,omitempty"`
}

func (x *OperatorStrategyApr) Reset() {
	*x = OperatorStrategyApr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorStrategyApr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorStrategyApr) ProtoMessage() {}

func (x *OperatorStrategyApr) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorStrategyApr.ProtoReflect.Descriptor instead.
func (*OperatorStrategyApr) Descriptor() ([]byte, []int) {
	return file_eigenlayer_pds_v1_aprs_proto_rawDescGZIP(), []int{1}
}

func (x *OperatorStrategyApr) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *OperatorStrategyApr) GetApr() string {
	if x != nil {
		return x.Apr
	}
	return ""
}

type GetDailyOperatorStrategyAprsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aprs []*OperatorStrategyApr `protobuf:"bytes,1,rep,name=aprs,proto3" json:"aprs,omitempty"`
}

func (x *GetDailyOperatorStrategyAprsResponse) Reset() {
	*x = GetDailyOperatorStrategyAprsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyOperatorStrategyAprsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyOperatorStrategyAprsResponse) ProtoMessage() {}

func (x *GetDailyOperatorStrategyAprsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyOperatorStrategyAprsResponse.ProtoReflect.Descriptor instead.
func (*GetDailyOperatorStrategyAprsResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_pds_v1_aprs_proto_rawDescGZIP(), []int{2}
}

func (x *GetDailyOperatorStrategyAprsResponse) GetAprs() []*OperatorStrategyApr {
	if x != nil {
		return x.Aprs
	}
	return nil
}

type GetDailyAprForEarnerStrategyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EarnerAddress string `protobuf:"bytes,1,opt,name=earner_address,json=earnerAddress,proto3" json:"earner_address,omitempty"`
	Strategy      string `protobuf:"bytes,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Date          string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetDailyAprForEarnerStrategyRequest) Reset() {
	*x = GetDailyAprForEarnerStrategyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyAprForEarnerStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyAprForEarnerStrategyRequest) ProtoMessage() {}

func (x *GetDailyAprForEarnerStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyAprForEarnerStrategyRequest.ProtoReflect.Descriptor instead.
func (*GetDailyAprForEarnerStrategyRequest) Descriptor() ([]byte, []int) {
	return file_eigenlayer_pds_v1_aprs_proto_rawDescGZIP(), []int{3}
}

func (x *GetDailyAprForEarnerStrategyRequest) GetEarnerAddress() string {
	if x != nil {
		return x.EarnerAddress
	}
	return ""
}

func (x *GetDailyAprForEarnerStrategyRequest) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *GetDailyAprForEarnerStrategyRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetDailyAprForEarnerStrategyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apr string `protobuf:"bytes,1,opt,name=apr,proto3" json:"apr,omitempty"`
}

func (x *GetDailyAprForEarnerStrategyResponse) Reset() {
	*x = GetDailyAprForEarnerStrategyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyAprForEarnerStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyAprForEarnerStrategyResponse) ProtoMessage() {}

func (x *GetDailyAprForEarnerStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eigenlayer_pds_v1_aprs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyAprForEarnerStrategyResponse.ProtoReflect.Descriptor instead.
func (*GetDailyAprForEarnerStrategyResponse) Descriptor() ([]byte, []int) {
	return file_eigenlayer_pds_v1_aprs_proto_rawDescGZIP(), []int{4}
}

func (x *GetDailyAprForEarnerStrategyResponse) GetApr() string {
	if x != nil {
		return x.Apr
	}
	return ""
}

var File_eigenlayer_pds_v1_aprs_proto protoreflect.FileDescriptor

var file_eigenlayer_pds_v1_aprs_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61,
	0x70, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x41, 0x70, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x70,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x70, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x72, 0x22,
	0x67, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x70, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x61, 0x70, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41,
	0x70, 0x72, 0x52, 0x04, 0x61, 0x70, 0x72, 0x73, 0x22, 0x7c, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x41, 0x70, 0x72, 0x46, 0x6f, 0x72, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x38, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x41, 0x70, 0x72, 0x46, 0x6f, 0x72, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x72,
	0x32, 0xd1, 0x03, 0x0a, 0x04, 0x41, 0x70, 0x72, 0x73, 0x12, 0xda, 0x01, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x70, 0x72, 0x73, 0x12, 0x3b, 0x2e, 0x65, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x70, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x70, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x37, 0x2f,
	0x61, 0x70, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x7b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x7d, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x2d, 0x61, 0x70, 0x72, 0x73, 0x2f,
	0x7b, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x12, 0xeb, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x41, 0x70, 0x72, 0x46, 0x6f, 0x72, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x3b, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x70, 0x72, 0x46, 0x6f, 0x72, 0x45,
	0x61, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x70, 0x72, 0x46, 0x6f, 0x72, 0x45, 0x61, 0x72, 0x6e,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x12, 0x48, 0x2f, 0x61, 0x70, 0x72,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x65, 0x61,
	0x72, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x7d, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x2d, 0x61, 0x70, 0x72, 0x2f, 0x7b, 0x64,
	0x61, 0x74, 0x65, 0x7d, 0x42, 0xe4, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x41, 0x70, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79,
	0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x64, 0x73, 0x2f, 0x61, 0x70, 0x72, 0x73, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x50, 0x41, 0xaa, 0x02, 0x16, 0x45, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x64, 0x73, 0x2e, 0x41, 0x70, 0x72, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x16, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x50,
	0x64, 0x73, 0x5c, 0x41, 0x70, 0x72, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x45, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x50, 0x64, 0x73, 0x5c, 0x41, 0x70, 0x72, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x19, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x50, 0x64,
	0x73, 0x3a, 0x3a, 0x41, 0x70, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eigenlayer_pds_v1_aprs_proto_rawDescOnce sync.Once
	file_eigenlayer_pds_v1_aprs_proto_rawDescData = file_eigenlayer_pds_v1_aprs_proto_rawDesc
)

func file_eigenlayer_pds_v1_aprs_proto_rawDescGZIP() []byte {
	file_eigenlayer_pds_v1_aprs_proto_rawDescOnce.Do(func() {
		file_eigenlayer_pds_v1_aprs_proto_rawDescData = protoimpl.X.CompressGZIP(file_eigenlayer_pds_v1_aprs_proto_rawDescData)
	})
	return file_eigenlayer_pds_v1_aprs_proto_rawDescData
}

var file_eigenlayer_pds_v1_aprs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eigenlayer_pds_v1_aprs_proto_goTypes = []any{
	(*GetDailyOperatorStrategyAprsRequest)(nil),  // 0: eigenlayer.pds.aprs.v1.GetDailyOperatorStrategyAprsRequest
	(*OperatorStrategyApr)(nil),                  // 1: eigenlayer.pds.aprs.v1.OperatorStrategyApr
	(*GetDailyOperatorStrategyAprsResponse)(nil), // 2: eigenlayer.pds.aprs.v1.GetDailyOperatorStrategyAprsResponse
	(*GetDailyAprForEarnerStrategyRequest)(nil),  // 3: eigenlayer.pds.aprs.v1.GetDailyAprForEarnerStrategyRequest
	(*GetDailyAprForEarnerStrategyResponse)(nil), // 4: eigenlayer.pds.aprs.v1.GetDailyAprForEarnerStrategyResponse
}
var file_eigenlayer_pds_v1_aprs_proto_depIdxs = []int32{
	1, // 0: eigenlayer.pds.aprs.v1.GetDailyOperatorStrategyAprsResponse.aprs:type_name -> eigenlayer.pds.aprs.v1.OperatorStrategyApr
	0, // 1: eigenlayer.pds.aprs.v1.Aprs.GetDailyOperatorStrategyAprs:input_type -> eigenlayer.pds.aprs.v1.GetDailyOperatorStrategyAprsRequest
	3, // 2: eigenlayer.pds.aprs.v1.Aprs.GetDailyAprForEarnerStrategy:input_type -> eigenlayer.pds.aprs.v1.GetDailyAprForEarnerStrategyRequest
	2, // 3: eigenlayer.pds.aprs.v1.Aprs.GetDailyOperatorStrategyAprs:output_type -> eigenlayer.pds.aprs.v1.GetDailyOperatorStrategyAprsResponse
	4, // 4: eigenlayer.pds.aprs.v1.Aprs.GetDailyAprForEarnerStrategy:output_type -> eigenlayer.pds.aprs.v1.GetDailyAprForEarnerStrategyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eigenlayer_pds_v1_aprs_proto_init() }
func file_eigenlayer_pds_v1_aprs_proto_init() {
	if File_eigenlayer_pds_v1_aprs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eigenlayer_pds_v1_aprs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetDailyOperatorStrategyAprsRequest); i {
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
		file_eigenlayer_pds_v1_aprs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OperatorStrategyApr); i {
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
		file_eigenlayer_pds_v1_aprs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetDailyOperatorStrategyAprsResponse); i {
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
		file_eigenlayer_pds_v1_aprs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetDailyAprForEarnerStrategyRequest); i {
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
		file_eigenlayer_pds_v1_aprs_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetDailyAprForEarnerStrategyResponse); i {
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
			RawDescriptor: file_eigenlayer_pds_v1_aprs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eigenlayer_pds_v1_aprs_proto_goTypes,
		DependencyIndexes: file_eigenlayer_pds_v1_aprs_proto_depIdxs,
		MessageInfos:      file_eigenlayer_pds_v1_aprs_proto_msgTypes,
	}.Build()
	File_eigenlayer_pds_v1_aprs_proto = out.File
	file_eigenlayer_pds_v1_aprs_proto_rawDesc = nil
	file_eigenlayer_pds_v1_aprs_proto_goTypes = nil
	file_eigenlayer_pds_v1_aprs_proto_depIdxs = nil
}
