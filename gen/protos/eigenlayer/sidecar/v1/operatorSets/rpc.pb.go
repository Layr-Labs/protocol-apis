// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/operatorSets/rpc.proto

package operatorSets

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_eigenlayer_sidecar_v1_operatorSets_rpc_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x74, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x05, 0x0a, 0x0c, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x12, 0xdb, 0x01, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x6b, 0x65, 0x72, 0x12, 0x41, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0xe6, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x12, 0x43, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x65, 0x69, 0x67, 0x65,
	0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x37, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0xcc, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x76, 0x73, 0x12, 0x3e, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x76,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x76,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2e, 0x12, 0x2c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x74, 0x73, 0x2f, 0x61, 0x76, 0x73, 0x73, 0x2f, 0x7b, 0x61, 0x76, 0x73, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x42,
	0xb0, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x73, 0x42, 0x08, 0x52, 0x70, 0x63, 0x50,
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
}

var file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_goTypes = []any{
	(*ListOperatorsForStakerRequest)(nil),    // 0: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerRequest
	(*ListOperatorsForStrategyRequest)(nil),  // 1: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyRequest
	(*ListOperatorsForAvsRequest)(nil),       // 2: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsRequest
	(*ListOperatorsForStakerResponse)(nil),   // 3: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerResponse
	(*ListOperatorsForStrategyResponse)(nil), // 4: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyResponse
	(*ListOperatorsForAvsResponse)(nil),      // 5: eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsResponse
}
var file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_depIdxs = []int32{
	0, // 0: eigenlayer.sidecar.v1.operatorSets.OperatorSets.ListOperatorsForStaker:input_type -> eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerRequest
	1, // 1: eigenlayer.sidecar.v1.operatorSets.OperatorSets.ListOperatorsForStrategy:input_type -> eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyRequest
	2, // 2: eigenlayer.sidecar.v1.operatorSets.OperatorSets.ListOperatorsForAvs:input_type -> eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsRequest
	3, // 3: eigenlayer.sidecar.v1.operatorSets.OperatorSets.ListOperatorsForStaker:output_type -> eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStakerResponse
	4, // 4: eigenlayer.sidecar.v1.operatorSets.OperatorSets.ListOperatorsForStrategy:output_type -> eigenlayer.sidecar.v1.operatorSets.ListOperatorsForStrategyResponse
	5, // 5: eigenlayer.sidecar.v1.operatorSets.OperatorSets.ListOperatorsForAvs:output_type -> eigenlayer.sidecar.v1.operatorSets.ListOperatorsForAvsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_init() }
func file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_init() {
	if File_eigenlayer_sidecar_v1_operatorSets_rpc_proto != nil {
		return
	}
	file_eigenlayer_sidecar_v1_operatorSets_operatorSets_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_depIdxs,
	}.Build()
	File_eigenlayer_sidecar_v1_operatorSets_rpc_proto = out.File
	file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_rawDesc = nil
	file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_operatorSets_rpc_proto_depIdxs = nil
}
