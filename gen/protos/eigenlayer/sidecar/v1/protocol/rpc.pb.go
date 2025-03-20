// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/protocol/rpc.proto

package protocol

import (
	_ "github.com/Layr-Labs/protobuf-libs/protos/eigenlayer/lib/annotations"
	_ "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/common"
	_ "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/eigenState"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_eigenlayer_sidecar_v1_protocol_rpc_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_protocol_rpc_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x31, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x69, 0x67,
	0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65,
	0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x8c, 0x18, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0xec, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x41, 0x76, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x42, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x76,
	0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x41, 0x76, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3a, 0x12, 0x38, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x2d, 0x61, 0x76, 0x73, 0x12, 0x84,
	0x02, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x48, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x80, 0xb5, 0x18, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x40, 0x12, 0x3e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0xa6, 0x02, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x4b,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4c, 0x2e, 0x65, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x80, 0xb5, 0x18, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x59, 0x12, 0x57, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d,
	0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0xf8,
	0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x45, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x47, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x12, 0x3b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x64, 0x2d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x12, 0xbc, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x36, 0x2e,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38,
	0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x7d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0xbf, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x45, 0x69, 0x67, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x3b, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x80, 0xb5,
	0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0xa4, 0x01, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x35, 0x2e,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x80, 0xb5,
	0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65,
	0x73, 0x12, 0xce, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12, 0x3b, 0x2e, 0x65, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x31, 0x12, 0x2f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x12, 0xe0, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x65, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x43, 0x12, 0x41, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0xeb, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0x42, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43,
	0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x37, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d,
	0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x73, 0x12, 0xf6, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0x44, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x65, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x48, 0x80, 0xb5, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x12, 0x3c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64,
	0x2d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0xf5, 0x01, 0x0a,
	0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0x44,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x80, 0xb5, 0x18,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x12, 0x3b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x7d, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x61, 0x6c, 0x73, 0x12, 0xaa, 0x02, 0x0a, 0x25, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x12, 0x4c,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d, 0x2e, 0x65,
	0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x80, 0xb5, 0x18,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5a, 0x12, 0x58, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x7d, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x64, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c,
	0x73, 0x42, 0x98, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x08, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x56, 0x50, 0xaa, 0x02, 0x1e, 0x45, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x56,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0xca, 0x02, 0x1e, 0x45, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0xe2, 0x02, 0x2a, 0x45, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x45, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_eigenlayer_sidecar_v1_protocol_rpc_proto_goTypes = []any{
	(*GetRegisteredAvsForOperatorRequest)(nil),            // 0: eigenlayer.sidecar.v1.protocol.GetRegisteredAvsForOperatorRequest
	(*GetDelegatedStrategiesForOperatorRequest)(nil),      // 1: eigenlayer.sidecar.v1.protocol.GetDelegatedStrategiesForOperatorRequest
	(*GetOperatorDelegatedStakeForStrategyRequest)(nil),   // 2: eigenlayer.sidecar.v1.protocol.GetOperatorDelegatedStakeForStrategyRequest
	(*GetDelegatedStakersForOperatorRequest)(nil),         // 3: eigenlayer.sidecar.v1.protocol.GetDelegatedStakersForOperatorRequest
	(*GetStakerSharesRequest)(nil),                        // 4: eigenlayer.sidecar.v1.protocol.GetStakerSharesRequest
	(*GetEigenStateChangesRequest)(nil),                   // 5: eigenlayer.sidecar.v1.protocol.GetEigenStateChangesRequest
	(*ListStrategiesRequest)(nil),                         // 6: eigenlayer.sidecar.v1.protocol.ListStrategiesRequest
	(*ListStakerStrategiesRequest)(nil),                   // 7: eigenlayer.sidecar.v1.protocol.ListStakerStrategiesRequest
	(*GetStrategyForStakerRequest)(nil),                   // 8: eigenlayer.sidecar.v1.protocol.GetStrategyForStakerRequest
	(*ListStakerQueuedWithdrawalsRequest)(nil),            // 9: eigenlayer.sidecar.v1.protocol.ListStakerQueuedWithdrawalsRequest
	(*ListStrategyQueuedWithdrawalsRequest)(nil),          // 10: eigenlayer.sidecar.v1.protocol.ListStrategyQueuedWithdrawalsRequest
	(*ListOperatorQueuedWithdrawalsRequest)(nil),          // 11: eigenlayer.sidecar.v1.protocol.ListOperatorQueuedWithdrawalsRequest
	(*ListOperatorStrategyQueuedWithdrawalsRequest)(nil),  // 12: eigenlayer.sidecar.v1.protocol.ListOperatorStrategyQueuedWithdrawalsRequest
	(*GetRegisteredAvsForOperatorResponse)(nil),           // 13: eigenlayer.sidecar.v1.protocol.GetRegisteredAvsForOperatorResponse
	(*GetDelegatedStrategiesForOperatorResponse)(nil),     // 14: eigenlayer.sidecar.v1.protocol.GetDelegatedStrategiesForOperatorResponse
	(*GetOperatorDelegatedStakeForStrategyResponse)(nil),  // 15: eigenlayer.sidecar.v1.protocol.GetOperatorDelegatedStakeForStrategyResponse
	(*GetDelegatedStakersForOperatorResponse)(nil),        // 16: eigenlayer.sidecar.v1.protocol.GetDelegatedStakersForOperatorResponse
	(*GetStakerSharesResponse)(nil),                       // 17: eigenlayer.sidecar.v1.protocol.GetStakerSharesResponse
	(*GetEigenStateChangesResponse)(nil),                  // 18: eigenlayer.sidecar.v1.protocol.GetEigenStateChangesResponse
	(*ListStrategiesResponse)(nil),                        // 19: eigenlayer.sidecar.v1.protocol.ListStrategiesResponse
	(*ListStakerStrategiesResponse)(nil),                  // 20: eigenlayer.sidecar.v1.protocol.ListStakerStrategiesResponse
	(*GetStrategyForStakerResponse)(nil),                  // 21: eigenlayer.sidecar.v1.protocol.GetStrategyForStakerResponse
	(*ListStakerQueuedWithdrawalsResponse)(nil),           // 22: eigenlayer.sidecar.v1.protocol.ListStakerQueuedWithdrawalsResponse
	(*ListStrategyQueuedWithdrawalsResponse)(nil),         // 23: eigenlayer.sidecar.v1.protocol.ListStrategyQueuedWithdrawalsResponse
	(*ListOperatorQueuedWithdrawalsResponse)(nil),         // 24: eigenlayer.sidecar.v1.protocol.ListOperatorQueuedWithdrawalsResponse
	(*ListOperatorStrategyQueuedWithdrawalsResponse)(nil), // 25: eigenlayer.sidecar.v1.protocol.ListOperatorStrategyQueuedWithdrawalsResponse
}
var file_eigenlayer_sidecar_v1_protocol_rpc_proto_depIdxs = []int32{
	0,  // 0: eigenlayer.sidecar.v1.protocol.Protocol.GetRegisteredAvsForOperator:input_type -> eigenlayer.sidecar.v1.protocol.GetRegisteredAvsForOperatorRequest
	1,  // 1: eigenlayer.sidecar.v1.protocol.Protocol.GetDelegatedStrategiesForOperator:input_type -> eigenlayer.sidecar.v1.protocol.GetDelegatedStrategiesForOperatorRequest
	2,  // 2: eigenlayer.sidecar.v1.protocol.Protocol.GetOperatorDelegatedStakeForStrategy:input_type -> eigenlayer.sidecar.v1.protocol.GetOperatorDelegatedStakeForStrategyRequest
	3,  // 3: eigenlayer.sidecar.v1.protocol.Protocol.GetDelegatedStakersForOperator:input_type -> eigenlayer.sidecar.v1.protocol.GetDelegatedStakersForOperatorRequest
	4,  // 4: eigenlayer.sidecar.v1.protocol.Protocol.GetStakerShares:input_type -> eigenlayer.sidecar.v1.protocol.GetStakerSharesRequest
	5,  // 5: eigenlayer.sidecar.v1.protocol.Protocol.GetEigenStateChanges:input_type -> eigenlayer.sidecar.v1.protocol.GetEigenStateChangesRequest
	6,  // 6: eigenlayer.sidecar.v1.protocol.Protocol.ListStrategies:input_type -> eigenlayer.sidecar.v1.protocol.ListStrategiesRequest
	7,  // 7: eigenlayer.sidecar.v1.protocol.Protocol.ListStakerStrategies:input_type -> eigenlayer.sidecar.v1.protocol.ListStakerStrategiesRequest
	8,  // 8: eigenlayer.sidecar.v1.protocol.Protocol.GetStrategyForStaker:input_type -> eigenlayer.sidecar.v1.protocol.GetStrategyForStakerRequest
	9,  // 9: eigenlayer.sidecar.v1.protocol.Protocol.ListStakerQueuedWithdrawals:input_type -> eigenlayer.sidecar.v1.protocol.ListStakerQueuedWithdrawalsRequest
	10, // 10: eigenlayer.sidecar.v1.protocol.Protocol.ListStrategyQueuedWithdrawals:input_type -> eigenlayer.sidecar.v1.protocol.ListStrategyQueuedWithdrawalsRequest
	11, // 11: eigenlayer.sidecar.v1.protocol.Protocol.ListOperatorQueuedWithdrawals:input_type -> eigenlayer.sidecar.v1.protocol.ListOperatorQueuedWithdrawalsRequest
	12, // 12: eigenlayer.sidecar.v1.protocol.Protocol.ListOperatorStrategyQueuedWithdrawals:input_type -> eigenlayer.sidecar.v1.protocol.ListOperatorStrategyQueuedWithdrawalsRequest
	13, // 13: eigenlayer.sidecar.v1.protocol.Protocol.GetRegisteredAvsForOperator:output_type -> eigenlayer.sidecar.v1.protocol.GetRegisteredAvsForOperatorResponse
	14, // 14: eigenlayer.sidecar.v1.protocol.Protocol.GetDelegatedStrategiesForOperator:output_type -> eigenlayer.sidecar.v1.protocol.GetDelegatedStrategiesForOperatorResponse
	15, // 15: eigenlayer.sidecar.v1.protocol.Protocol.GetOperatorDelegatedStakeForStrategy:output_type -> eigenlayer.sidecar.v1.protocol.GetOperatorDelegatedStakeForStrategyResponse
	16, // 16: eigenlayer.sidecar.v1.protocol.Protocol.GetDelegatedStakersForOperator:output_type -> eigenlayer.sidecar.v1.protocol.GetDelegatedStakersForOperatorResponse
	17, // 17: eigenlayer.sidecar.v1.protocol.Protocol.GetStakerShares:output_type -> eigenlayer.sidecar.v1.protocol.GetStakerSharesResponse
	18, // 18: eigenlayer.sidecar.v1.protocol.Protocol.GetEigenStateChanges:output_type -> eigenlayer.sidecar.v1.protocol.GetEigenStateChangesResponse
	19, // 19: eigenlayer.sidecar.v1.protocol.Protocol.ListStrategies:output_type -> eigenlayer.sidecar.v1.protocol.ListStrategiesResponse
	20, // 20: eigenlayer.sidecar.v1.protocol.Protocol.ListStakerStrategies:output_type -> eigenlayer.sidecar.v1.protocol.ListStakerStrategiesResponse
	21, // 21: eigenlayer.sidecar.v1.protocol.Protocol.GetStrategyForStaker:output_type -> eigenlayer.sidecar.v1.protocol.GetStrategyForStakerResponse
	22, // 22: eigenlayer.sidecar.v1.protocol.Protocol.ListStakerQueuedWithdrawals:output_type -> eigenlayer.sidecar.v1.protocol.ListStakerQueuedWithdrawalsResponse
	23, // 23: eigenlayer.sidecar.v1.protocol.Protocol.ListStrategyQueuedWithdrawals:output_type -> eigenlayer.sidecar.v1.protocol.ListStrategyQueuedWithdrawalsResponse
	24, // 24: eigenlayer.sidecar.v1.protocol.Protocol.ListOperatorQueuedWithdrawals:output_type -> eigenlayer.sidecar.v1.protocol.ListOperatorQueuedWithdrawalsResponse
	25, // 25: eigenlayer.sidecar.v1.protocol.Protocol.ListOperatorStrategyQueuedWithdrawals:output_type -> eigenlayer.sidecar.v1.protocol.ListOperatorStrategyQueuedWithdrawalsResponse
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_protocol_rpc_proto_init() }
func file_eigenlayer_sidecar_v1_protocol_rpc_proto_init() {
	if File_eigenlayer_sidecar_v1_protocol_rpc_proto != nil {
		return
	}
	file_eigenlayer_sidecar_v1_protocol_protocol_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eigenlayer_sidecar_v1_protocol_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_protocol_rpc_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_protocol_rpc_proto_depIdxs,
	}.Build()
	File_eigenlayer_sidecar_v1_protocol_rpc_proto = out.File
	file_eigenlayer_sidecar_v1_protocol_rpc_proto_rawDesc = nil
	file_eigenlayer_sidecar_v1_protocol_rpc_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_protocol_rpc_proto_depIdxs = nil
}
