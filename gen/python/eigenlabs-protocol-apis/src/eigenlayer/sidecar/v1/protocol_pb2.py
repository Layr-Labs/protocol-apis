# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eigenlayer/sidecar/v1/protocol.proto
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
    'eigenlayer/sidecar/v1/protocol.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$eigenlayer/sidecar/v1/protocol.proto\x12\x1e\x65igenlayer.sidecar.protocol.v1\x1a\x1cgoogle/api/annotations.proto\"\x88\x01\n\"GetRegisteredAvsForOperatorRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"J\n#GetRegisteredAvsForOperatorResponse\x12#\n\ravs_addresses\x18\x01 \x03(\tR\x0c\x61vsAddresses\"\x8e\x01\n(GetDelegatedStrategiesForOperatorRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"Z\n)GetDelegatedStrategiesForOperatorResponse\x12-\n\x12strategy_addresses\x18\x01 \x03(\tR\x11strategyAddresses\"\xbc\x01\n+GetOperatorDelegatedStakeForStrategyRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12)\n\x10strategy_address\x18\x02 \x01(\tR\x0fstrategyAddress\x12&\n\x0c\x62lock_height\x18\x03 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"k\n,GetOperatorDelegatedStakeForStrategyResponse\x12\x16\n\x06shares\x18\x01 \x01(\tR\x06shares\x12#\n\ravs_addresses\x18\x02 \x03(\tR\x0c\x61vsAddresses\"\xc9\x01\n%GetDelegatedStakersForOperatorRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x12\x1f\n\x0bpage_number\x18\x03 \x01(\rR\npageNumber\x12\x1b\n\tpage_size\x18\x04 \x01(\rR\x08pageSizeB\x0f\n\r_block_height\"O\n&GetDelegatedStakersForOperatorResponse\x12%\n\x0estaker_address\x18\x01 \x03(\tR\rstakerAddress\"x\n\x16GetStakerSharesRequest\x12%\n\x0estaker_address\x18\x01 \x01(\tR\rstakerAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"\x91\x01\n\x0bStakerShare\x12\x1a\n\x08strategy\x18\x01 \x01(\tR\x08strategy\x12\x16\n\x06shares\x18\x02 \x01(\tR\x06shares\x12)\n\x10operator_address\x18\x03 \x01(\tR\x0foperatorAddress\x12#\n\ravs_addresses\x18\x04 \x03(\tR\x0c\x61vsAddresses\"^\n\x17GetStakerSharesResponse\x12\x43\n\x06shares\x18\x01 \x03(\x0b\x32+.eigenlayer.sidecar.protocol.v1.StakerShareR\x06shares2\xcf\t\n\x08Protocol\x12\xe8\x01\n\x1bGetRegisteredAvsForOperator\x12\x42.eigenlayer.sidecar.protocol.v1.GetRegisteredAvsForOperatorRequest\x1a\x43.eigenlayer.sidecar.protocol.v1.GetRegisteredAvsForOperatorResponse\"@\x82\xd3\xe4\x93\x02:\x12\x38/protocol/v1/operators/{operator_address}/registered-avs\x12\x80\x02\n!GetDelegatedStrategiesForOperator\x12H.eigenlayer.sidecar.protocol.v1.GetDelegatedStrategiesForOperatorRequest\x1aI.eigenlayer.sidecar.protocol.v1.GetDelegatedStrategiesForOperatorResponse\"F\x82\xd3\xe4\x93\x02@\x12>/protocol/v1/operators/{operator_address}/delegated-strategies\x12\xa2\x02\n$GetOperatorDelegatedStakeForStrategy\x12K.eigenlayer.sidecar.protocol.v1.GetOperatorDelegatedStakeForStrategyRequest\x1aL.eigenlayer.sidecar.protocol.v1.GetOperatorDelegatedStakeForStrategyResponse\"_\x82\xd3\xe4\x93\x02Y\x12W/protocol/v1/operators/{operator_address}/strategies/{strategy_address}/delegated-stake\x12\xf4\x01\n\x1eGetDelegatedStakersForOperator\x12\x45.eigenlayer.sidecar.protocol.v1.GetDelegatedStakersForOperatorRequest\x1a\x46.eigenlayer.sidecar.protocol.v1.GetDelegatedStakersForOperatorResponse\"C\x82\xd3\xe4\x93\x02=\x12;/protocol/v1/operators/{operator_address}/delegated-stakers\x12\xb8\x01\n\x0fGetStakerShares\x12\x36.eigenlayer.sidecar.protocol.v1.GetStakerSharesRequest\x1a\x37.eigenlayer.sidecar.protocol.v1.GetStakerSharesResponse\"4\x82\xd3\xe4\x93\x02.\x12,/protocol/v1/stakers/{staker_address}/sharesB\x9c\x02\n\"com.eigenlayer.sidecar.protocol.v1B\rProtocolProtoP\x01ZLgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/protocol/v1\xa2\x02\x03\x45SP\xaa\x02\x1e\x45igenlayer.Sidecar.Protocol.V1\xca\x02\x1e\x45igenlayer\\Sidecar\\Protocol\\V1\xe2\x02*Eigenlayer\\Sidecar\\Protocol\\V1\\GPBMetadata\xea\x02!Eigenlayer::Sidecar::Protocol::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'eigenlayer.sidecar.v1.protocol_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\"com.eigenlayer.sidecar.protocol.v1B\rProtocolProtoP\001ZLgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/protocol/v1\242\002\003ESP\252\002\036Eigenlayer.Sidecar.Protocol.V1\312\002\036Eigenlayer\\Sidecar\\Protocol\\V1\342\002*Eigenlayer\\Sidecar\\Protocol\\V1\\GPBMetadata\352\002!Eigenlayer::Sidecar::Protocol::V1'
  _globals['_PROTOCOL'].methods_by_name['GetRegisteredAvsForOperator']._loaded_options = None
  _globals['_PROTOCOL'].methods_by_name['GetRegisteredAvsForOperator']._serialized_options = b'\202\323\344\223\002:\0228/protocol/v1/operators/{operator_address}/registered-avs'
  _globals['_PROTOCOL'].methods_by_name['GetDelegatedStrategiesForOperator']._loaded_options = None
  _globals['_PROTOCOL'].methods_by_name['GetDelegatedStrategiesForOperator']._serialized_options = b'\202\323\344\223\002@\022>/protocol/v1/operators/{operator_address}/delegated-strategies'
  _globals['_PROTOCOL'].methods_by_name['GetOperatorDelegatedStakeForStrategy']._loaded_options = None
  _globals['_PROTOCOL'].methods_by_name['GetOperatorDelegatedStakeForStrategy']._serialized_options = b'\202\323\344\223\002Y\022W/protocol/v1/operators/{operator_address}/strategies/{strategy_address}/delegated-stake'
  _globals['_PROTOCOL'].methods_by_name['GetDelegatedStakersForOperator']._loaded_options = None
  _globals['_PROTOCOL'].methods_by_name['GetDelegatedStakersForOperator']._serialized_options = b'\202\323\344\223\002=\022;/protocol/v1/operators/{operator_address}/delegated-stakers'
  _globals['_PROTOCOL'].methods_by_name['GetStakerShares']._loaded_options = None
  _globals['_PROTOCOL'].methods_by_name['GetStakerShares']._serialized_options = b'\202\323\344\223\002.\022,/protocol/v1/stakers/{staker_address}/shares'
  _globals['_GETREGISTEREDAVSFOROPERATORREQUEST']._serialized_start=103
  _globals['_GETREGISTEREDAVSFOROPERATORREQUEST']._serialized_end=239
  _globals['_GETREGISTEREDAVSFOROPERATORRESPONSE']._serialized_start=241
  _globals['_GETREGISTEREDAVSFOROPERATORRESPONSE']._serialized_end=315
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORREQUEST']._serialized_start=318
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORREQUEST']._serialized_end=460
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORRESPONSE']._serialized_start=462
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORRESPONSE']._serialized_end=552
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYREQUEST']._serialized_start=555
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYREQUEST']._serialized_end=743
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYRESPONSE']._serialized_start=745
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYRESPONSE']._serialized_end=852
  _globals['_GETDELEGATEDSTAKERSFOROPERATORREQUEST']._serialized_start=855
  _globals['_GETDELEGATEDSTAKERSFOROPERATORREQUEST']._serialized_end=1056
  _globals['_GETDELEGATEDSTAKERSFOROPERATORRESPONSE']._serialized_start=1058
  _globals['_GETDELEGATEDSTAKERSFOROPERATORRESPONSE']._serialized_end=1137
  _globals['_GETSTAKERSHARESREQUEST']._serialized_start=1139
  _globals['_GETSTAKERSHARESREQUEST']._serialized_end=1259
  _globals['_STAKERSHARE']._serialized_start=1262
  _globals['_STAKERSHARE']._serialized_end=1407
  _globals['_GETSTAKERSHARESRESPONSE']._serialized_start=1409
  _globals['_GETSTAKERSHARESRESPONSE']._serialized_end=1503
  _globals['_PROTOCOL']._serialized_start=1506
  _globals['_PROTOCOL']._serialized_end=2737
# @@protoc_insertion_point(module_scope)
