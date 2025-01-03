# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eigenlayer/sidecar/v1/api.proto
# Protobuf Python Version: 5.29.2
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
    2,
    '',
    'eigenlayer/sidecar/v1/api.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1f\x65igenlayer/sidecar/v1/api.proto\x12\x19\x65igenlayer.sidecar.api.v1\x1a\x1cgoogle/api/annotations.proto\"3\n\x15GetBlockHeightRequest\x12\x1a\n\x08verified\x18\x01 \x01(\x08R\x08verified\"X\n\x16GetBlockHeightResponse\x12 \n\x0b\x62lockNumber\x18\x01 \x01(\x04R\x0b\x62lockNumber\x12\x1c\n\tblockHash\x18\x02 \x01(\tR\tblockHash\"7\n\x13GetStateRootRequest\x12 \n\x0b\x62lockNumber\x18\x01 \x01(\x04R\x0b\x62lockNumber\"\x80\x01\n\x14GetStateRootResponse\x12&\n\x0e\x65thBlockNumber\x18\x01 \x01(\x04R\x0e\x65thBlockNumber\x12\"\n\x0c\x65thBlockHash\x18\x02 \x01(\tR\x0c\x65thBlockHash\x12\x1c\n\tstateRoot\x18\x03 \x01(\tR\tstateRoot2\xbe\x02\n\x03Rpc\x12\x96\x01\n\x0eGetBlockHeight\x12\x30.eigenlayer.sidecar.api.v1.GetBlockHeightRequest\x1a\x31.eigenlayer.sidecar.api.v1.GetBlockHeightResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/rpc/v1/latest-block:\x01*\x12\x9d\x01\n\x0cGetStateRoot\x12..eigenlayer.sidecar.api.v1.GetStateRootRequest\x1a/.eigenlayer.sidecar.api.v1.GetStateRootResponse\",\x82\xd3\xe4\x93\x02&\"!/rpc/v1/state-roots/{blockNumber}:\x01*B\xf9\x01\n\x1d\x63om.eigenlayer.sidecar.api.v1B\x08\x41piProtoP\x01ZGgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/api/v1\xa2\x02\x03\x45SA\xaa\x02\x19\x45igenlayer.Sidecar.Api.V1\xca\x02\x19\x45igenlayer\\Sidecar\\Api\\V1\xe2\x02%Eigenlayer\\Sidecar\\Api\\V1\\GPBMetadata\xea\x02\x1c\x45igenlayer::Sidecar::Api::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'eigenlayer.sidecar.v1.api_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\035com.eigenlayer.sidecar.api.v1B\010ApiProtoP\001ZGgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/api/v1\242\002\003ESA\252\002\031Eigenlayer.Sidecar.Api.V1\312\002\031Eigenlayer\\Sidecar\\Api\\V1\342\002%Eigenlayer\\Sidecar\\Api\\V1\\GPBMetadata\352\002\034Eigenlayer::Sidecar::Api::V1'
  _globals['_RPC'].methods_by_name['GetBlockHeight']._loaded_options = None
  _globals['_RPC'].methods_by_name['GetBlockHeight']._serialized_options = b'\202\323\344\223\002\031\"\024/rpc/v1/latest-block:\001*'
  _globals['_RPC'].methods_by_name['GetStateRoot']._loaded_options = None
  _globals['_RPC'].methods_by_name['GetStateRoot']._serialized_options = b'\202\323\344\223\002&\"!/rpc/v1/state-roots/{blockNumber}:\001*'
  _globals['_GETBLOCKHEIGHTREQUEST']._serialized_start=92
  _globals['_GETBLOCKHEIGHTREQUEST']._serialized_end=143
  _globals['_GETBLOCKHEIGHTRESPONSE']._serialized_start=145
  _globals['_GETBLOCKHEIGHTRESPONSE']._serialized_end=233
  _globals['_GETSTATEROOTREQUEST']._serialized_start=235
  _globals['_GETSTATEROOTREQUEST']._serialized_end=290
  _globals['_GETSTATEROOTRESPONSE']._serialized_start=293
  _globals['_GETSTATEROOTRESPONSE']._serialized_end=421
  _globals['_RPC']._serialized_start=424
  _globals['_RPC']._serialized_end=742
# @@protoc_insertion_point(module_scope)
