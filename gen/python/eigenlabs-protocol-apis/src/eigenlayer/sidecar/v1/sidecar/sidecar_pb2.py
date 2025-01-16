# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eigenlayer/sidecar/v1/sidecar/sidecar.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'eigenlayer/sidecar/v1/sidecar/sidecar.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+eigenlayer/sidecar/v1/sidecar/sidecar.proto\x12\x1d\x65igenlayer.sidecar.v1.sidecar\x1a\x1cgoogle/api/annotations.proto\"3\n\x15GetBlockHeightRequest\x12\x1a\n\x08verified\x18\x01 \x01(\x08R\x08verified\"X\n\x16GetBlockHeightResponse\x12 \n\x0b\x62lockNumber\x18\x01 \x01(\x04R\x0b\x62lockNumber\x12\x1c\n\tblockHash\x18\x02 \x01(\tR\tblockHash\"7\n\x13GetStateRootRequest\x12 \n\x0b\x62lockNumber\x18\x01 \x01(\x04R\x0b\x62lockNumber\"\x80\x01\n\x14GetStateRootResponse\x12&\n\x0e\x65thBlockNumber\x18\x01 \x01(\x04R\x0e\x65thBlockNumber\x12\"\n\x0c\x65thBlockHash\x18\x02 \x01(\tR\x0c\x65thBlockHash\x12\x1c\n\tstateRoot\x18\x03 \x01(\tR\tstateRoot2\xce\x02\n\x03Rpc\x12\x9e\x01\n\x0eGetBlockHeight\x12\x34.eigenlayer.sidecar.v1.sidecar.GetBlockHeightRequest\x1a\x35.eigenlayer.sidecar.v1.sidecar.GetBlockHeightResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/rpc/v1/latest-block:\x01*\x12\xa5\x01\n\x0cGetStateRoot\x12\x32.eigenlayer.sidecar.v1.sidecar.GetStateRootRequest\x1a\x33.eigenlayer.sidecar.v1.sidecar.GetStateRootResponse\",\x82\xd3\xe4\x93\x02&\"!/rpc/v1/state-roots/{blockNumber}:\x01*B\x96\x02\n!com.eigenlayer.sidecar.v1.sidecarB\x0cSidecarProtoP\x01ZKgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/sidecar\xa2\x02\x04\x45SVS\xaa\x02\x1d\x45igenlayer.Sidecar.V1.Sidecar\xca\x02\x1d\x45igenlayer\\Sidecar\\V1\\Sidecar\xe2\x02)Eigenlayer\\Sidecar\\V1\\Sidecar\\GPBMetadata\xea\x02 Eigenlayer::Sidecar::V1::Sidecarb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'eigenlayer.sidecar.v1.sidecar.sidecar_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!com.eigenlayer.sidecar.v1.sidecarB\014SidecarProtoP\001ZKgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/sidecar\242\002\004ESVS\252\002\035Eigenlayer.Sidecar.V1.Sidecar\312\002\035Eigenlayer\\Sidecar\\V1\\Sidecar\342\002)Eigenlayer\\Sidecar\\V1\\Sidecar\\GPBMetadata\352\002 Eigenlayer::Sidecar::V1::Sidecar'
  _globals['_RPC'].methods_by_name['GetBlockHeight']._loaded_options = None
  _globals['_RPC'].methods_by_name['GetBlockHeight']._serialized_options = b'\202\323\344\223\002\031\"\024/rpc/v1/latest-block:\001*'
  _globals['_RPC'].methods_by_name['GetStateRoot']._loaded_options = None
  _globals['_RPC'].methods_by_name['GetStateRoot']._serialized_options = b'\202\323\344\223\002&\"!/rpc/v1/state-roots/{blockNumber}:\001*'
  _globals['_GETBLOCKHEIGHTREQUEST']._serialized_start=108
  _globals['_GETBLOCKHEIGHTREQUEST']._serialized_end=159
  _globals['_GETBLOCKHEIGHTRESPONSE']._serialized_start=161
  _globals['_GETBLOCKHEIGHTRESPONSE']._serialized_end=249
  _globals['_GETSTATEROOTREQUEST']._serialized_start=251
  _globals['_GETSTATEROOTREQUEST']._serialized_end=306
  _globals['_GETSTATEROOTRESPONSE']._serialized_start=309
  _globals['_GETSTATEROOTRESPONSE']._serialized_end=437
  _globals['_RPC']._serialized_start=440
  _globals['_RPC']._serialized_end=774
# @@protoc_insertion_point(module_scope)