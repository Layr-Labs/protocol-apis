// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: eigenlayer/sidecar/v1/sidecar/rpc.proto

package sidecar

import (
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

var File_eigenlayer_sidecar_v1_sidecar_rpc_proto protoreflect.FileDescriptor

var file_eigenlayer_sidecar_v1_sidecar_rpc_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x65, 0x69, 0x67, 0x65, 0x6e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xff, 0x05, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x9b, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x34,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0xa2, 0x01, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x32, 0x2e, 0x65, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2d, 0x72, 0x6f, 0x6f, 0x74, 0x73,
	0x2f, 0x7b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x7d, 0x12, 0x79,
	0x0a, 0x05, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x2b, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x99, 0x01, 0x0a, 0x0c, 0x4c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x32, 0x2e, 0x65, 0x69, 0x67,
	0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x4c,
	0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x65,
	0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x42, 0x92, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x42, 0x08, 0x52, 0x70, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0xa2, 0x02, 0x04, 0x45, 0x53, 0x56, 0x53, 0xaa, 0x02, 0x1d, 0x45, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0xca, 0x02, 0x1d, 0x45, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0xe2, 0x02, 0x29, 0x45, 0x69,
	0x67, 0x65, 0x6e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x45, 0x69, 0x67, 0x65, 0x6e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_eigenlayer_sidecar_v1_sidecar_rpc_proto_goTypes = []any{
	(*GetBlockHeightRequest)(nil),  // 0: eigenlayer.sidecar.v1.sidecar.GetBlockHeightRequest
	(*GetStateRootRequest)(nil),    // 1: eigenlayer.sidecar.v1.sidecar.GetStateRootRequest
	(*AboutRequest)(nil),           // 2: eigenlayer.sidecar.v1.sidecar.AboutRequest
	(*LoadContractRequest)(nil),    // 3: eigenlayer.sidecar.v1.sidecar.LoadContractRequest
	(*LoadContractsRequest)(nil),   // 4: eigenlayer.sidecar.v1.sidecar.LoadContractsRequest
	(*GetBlockHeightResponse)(nil), // 5: eigenlayer.sidecar.v1.sidecar.GetBlockHeightResponse
	(*GetStateRootResponse)(nil),   // 6: eigenlayer.sidecar.v1.sidecar.GetStateRootResponse
	(*AboutResponse)(nil),          // 7: eigenlayer.sidecar.v1.sidecar.AboutResponse
	(*LoadContractResponse)(nil),   // 8: eigenlayer.sidecar.v1.sidecar.LoadContractResponse
	(*LoadContractsResponse)(nil),  // 9: eigenlayer.sidecar.v1.sidecar.LoadContractsResponse
}
var file_eigenlayer_sidecar_v1_sidecar_rpc_proto_depIdxs = []int32{
	0, // 0: eigenlayer.sidecar.v1.sidecar.Rpc.GetBlockHeight:input_type -> eigenlayer.sidecar.v1.sidecar.GetBlockHeightRequest
	1, // 1: eigenlayer.sidecar.v1.sidecar.Rpc.GetStateRoot:input_type -> eigenlayer.sidecar.v1.sidecar.GetStateRootRequest
	2, // 2: eigenlayer.sidecar.v1.sidecar.Rpc.About:input_type -> eigenlayer.sidecar.v1.sidecar.AboutRequest
	3, // 3: eigenlayer.sidecar.v1.sidecar.Rpc.LoadContract:input_type -> eigenlayer.sidecar.v1.sidecar.LoadContractRequest
	4, // 4: eigenlayer.sidecar.v1.sidecar.Rpc.LoadContracts:input_type -> eigenlayer.sidecar.v1.sidecar.LoadContractsRequest
	5, // 5: eigenlayer.sidecar.v1.sidecar.Rpc.GetBlockHeight:output_type -> eigenlayer.sidecar.v1.sidecar.GetBlockHeightResponse
	6, // 6: eigenlayer.sidecar.v1.sidecar.Rpc.GetStateRoot:output_type -> eigenlayer.sidecar.v1.sidecar.GetStateRootResponse
	7, // 7: eigenlayer.sidecar.v1.sidecar.Rpc.About:output_type -> eigenlayer.sidecar.v1.sidecar.AboutResponse
	8, // 8: eigenlayer.sidecar.v1.sidecar.Rpc.LoadContract:output_type -> eigenlayer.sidecar.v1.sidecar.LoadContractResponse
	9, // 9: eigenlayer.sidecar.v1.sidecar.Rpc.LoadContracts:output_type -> eigenlayer.sidecar.v1.sidecar.LoadContractsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eigenlayer_sidecar_v1_sidecar_rpc_proto_init() }
func file_eigenlayer_sidecar_v1_sidecar_rpc_proto_init() {
	if File_eigenlayer_sidecar_v1_sidecar_rpc_proto != nil {
		return
	}
	file_eigenlayer_sidecar_v1_sidecar_sidecar_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eigenlayer_sidecar_v1_sidecar_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eigenlayer_sidecar_v1_sidecar_rpc_proto_goTypes,
		DependencyIndexes: file_eigenlayer_sidecar_v1_sidecar_rpc_proto_depIdxs,
	}.Build()
	File_eigenlayer_sidecar_v1_sidecar_rpc_proto = out.File
	file_eigenlayer_sidecar_v1_sidecar_rpc_proto_rawDesc = nil
	file_eigenlayer_sidecar_v1_sidecar_rpc_proto_goTypes = nil
	file_eigenlayer_sidecar_v1_sidecar_rpc_proto_depIdxs = nil
}
