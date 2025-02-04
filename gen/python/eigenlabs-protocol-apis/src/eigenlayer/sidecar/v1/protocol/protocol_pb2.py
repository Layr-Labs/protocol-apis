# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eigenlayer/sidecar/v1/protocol/protocol.proto
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
    'eigenlayer/sidecar/v1/protocol/protocol.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from eigenlayer.sidecar.v1.eigenState import eigenState_pb2 as eigenlayer_dot_sidecar_dot_v1_dot_eigenState_dot_eigenState__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from eigenlayer.sidecar.v1.common import types_pb2 as eigenlayer_dot_sidecar_dot_v1_dot_common_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-eigenlayer/sidecar/v1/protocol/protocol.proto\x12\x1e\x65igenlayer.sidecar.v1.protocol\x1a\x31\x65igenlayer/sidecar/v1/eigenState/eigenState.proto\x1a\x1cgoogle/api/annotations.proto\x1a(eigenlayer/sidecar/v1/common/types.proto\"\x88\x01\n\"GetRegisteredAvsForOperatorRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"J\n#GetRegisteredAvsForOperatorResponse\x12#\n\ravs_addresses\x18\x01 \x03(\tR\x0c\x61vsAddresses\"\x8e\x01\n(GetDelegatedStrategiesForOperatorRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"Z\n)GetDelegatedStrategiesForOperatorResponse\x12-\n\x12strategy_addresses\x18\x01 \x03(\tR\x11strategyAddresses\"\xbc\x01\n+GetOperatorDelegatedStakeForStrategyRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12)\n\x10strategy_address\x18\x02 \x01(\tR\x0fstrategyAddress\x12&\n\x0c\x62lock_height\x18\x03 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"k\n,GetOperatorDelegatedStakeForStrategyResponse\x12\x16\n\x06shares\x18\x01 \x01(\tR\x06shares\x12#\n\ravs_addresses\x18\x02 \x03(\tR\x0c\x61vsAddresses\"\xe9\x01\n%GetDelegatedStakersForOperatorRequest\x12)\n\x10operator_address\x18\x01 \x01(\tR\x0foperatorAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x12M\n\npagination\x18\x03 \x01(\x0b\x32(.eigenlayer.sidecar.v1.common.PaginationH\x01R\npagination\x88\x01\x01\x42\x0f\n\r_block_heightB\r\n\x0b_pagination\"\xad\x01\n&GetDelegatedStakersForOperatorResponse\x12)\n\x10staker_addresses\x18\x01 \x03(\tR\x0fstakerAddresses\x12J\n\tnext_page\x18\x02 \x01(\x0b\x32(.eigenlayer.sidecar.v1.common.PaginationH\x00R\x08nextPage\x88\x01\x01\x42\x0c\n\n_next_page\"x\n\x16GetStakerSharesRequest\x12%\n\x0estaker_address\x18\x01 \x01(\tR\rstakerAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"\xab\x01\n\x0bStakerShare\x12\x1a\n\x08strategy\x18\x01 \x01(\tR\x08strategy\x12\x16\n\x06shares\x18\x02 \x01(\tR\x06shares\x12.\n\x10operator_address\x18\x03 \x01(\tH\x00R\x0foperatorAddress\x88\x01\x01\x12#\n\ravs_addresses\x18\x04 \x03(\tR\x0c\x61vsAddressesB\x13\n\x11_operator_address\"^\n\x17GetStakerSharesResponse\x12\x43\n\x06shares\x18\x01 \x03(\x0b\x32+.eigenlayer.sidecar.v1.protocol.StakerShareR\x06shares\"@\n\x1bGetEigenStateChangesRequest\x12!\n\x0c\x62lock_height\x18\x01 \x01(\x04R\x0b\x62lockHeight\"l\n\x1cGetEigenStateChangesResponse\x12L\n\x07\x63hanges\x18\x01 \x03(\x0b\x32\x32.eigenlayer.sidecar.v1.eigenState.EigenStateChangeR\x07\x63hanges2\x90\x0b\n\x08Protocol\x12\xe8\x01\n\x1bGetRegisteredAvsForOperator\x12\x42.eigenlayer.sidecar.v1.protocol.GetRegisteredAvsForOperatorRequest\x1a\x43.eigenlayer.sidecar.v1.protocol.GetRegisteredAvsForOperatorResponse\"@\x82\xd3\xe4\x93\x02:\x12\x38/protocol/v1/operators/{operator_address}/registered-avs\x12\x80\x02\n!GetDelegatedStrategiesForOperator\x12H.eigenlayer.sidecar.v1.protocol.GetDelegatedStrategiesForOperatorRequest\x1aI.eigenlayer.sidecar.v1.protocol.GetDelegatedStrategiesForOperatorResponse\"F\x82\xd3\xe4\x93\x02@\x12>/protocol/v1/operators/{operator_address}/delegated-strategies\x12\xa2\x02\n$GetOperatorDelegatedStakeForStrategy\x12K.eigenlayer.sidecar.v1.protocol.GetOperatorDelegatedStakeForStrategyRequest\x1aL.eigenlayer.sidecar.v1.protocol.GetOperatorDelegatedStakeForStrategyResponse\"_\x82\xd3\xe4\x93\x02Y\x12W/protocol/v1/operators/{operator_address}/strategies/{strategy_address}/delegated-stake\x12\xf4\x01\n\x1eGetDelegatedStakersForOperator\x12\x45.eigenlayer.sidecar.v1.protocol.GetDelegatedStakersForOperatorRequest\x1a\x46.eigenlayer.sidecar.v1.protocol.GetDelegatedStakersForOperatorResponse\"C\x82\xd3\xe4\x93\x02=\x12;/protocol/v1/operators/{operator_address}/delegated-stakers\x12\xb8\x01\n\x0fGetStakerShares\x12\x36.eigenlayer.sidecar.v1.protocol.GetStakerSharesRequest\x1a\x37.eigenlayer.sidecar.v1.protocol.GetStakerSharesResponse\"4\x82\xd3\xe4\x93\x02.\x12,/protocol/v1/stakers/{staker_address}/shares\x12\xbe\x01\n\x14GetEigenStateChanges\x12;.eigenlayer.sidecar.v1.protocol.GetEigenStateChangesRequest\x1a<.eigenlayer.sidecar.v1.protocol.GetEigenStateChangesResponse\"+\x82\xd3\xe4\x93\x02%\" /protocol/v1/eigen-state-changes:\x01*B\x9d\x02\n\"com.eigenlayer.sidecar.v1.protocolB\rProtocolProtoP\x01ZLgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/protocol\xa2\x02\x04\x45SVP\xaa\x02\x1e\x45igenlayer.Sidecar.V1.Protocol\xca\x02\x1e\x45igenlayer\\Sidecar\\V1\\Protocol\xe2\x02*Eigenlayer\\Sidecar\\V1\\Protocol\\GPBMetadata\xea\x02!Eigenlayer::Sidecar::V1::Protocolb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'eigenlayer.sidecar.v1.protocol.protocol_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\"com.eigenlayer.sidecar.v1.protocolB\rProtocolProtoP\001ZLgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/protocol\242\002\004ESVP\252\002\036Eigenlayer.Sidecar.V1.Protocol\312\002\036Eigenlayer\\Sidecar\\V1\\Protocol\342\002*Eigenlayer\\Sidecar\\V1\\Protocol\\GPBMetadata\352\002!Eigenlayer::Sidecar::V1::Protocol'
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
  _globals['_PROTOCOL'].methods_by_name['GetEigenStateChanges']._loaded_options = None
  _globals['_PROTOCOL'].methods_by_name['GetEigenStateChanges']._serialized_options = b'\202\323\344\223\002%\" /protocol/v1/eigen-state-changes:\001*'
  _globals['_GETREGISTEREDAVSFOROPERATORREQUEST']._serialized_start=205
  _globals['_GETREGISTEREDAVSFOROPERATORREQUEST']._serialized_end=341
  _globals['_GETREGISTEREDAVSFOROPERATORRESPONSE']._serialized_start=343
  _globals['_GETREGISTEREDAVSFOROPERATORRESPONSE']._serialized_end=417
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORREQUEST']._serialized_start=420
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORREQUEST']._serialized_end=562
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORRESPONSE']._serialized_start=564
  _globals['_GETDELEGATEDSTRATEGIESFOROPERATORRESPONSE']._serialized_end=654
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYREQUEST']._serialized_start=657
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYREQUEST']._serialized_end=845
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYRESPONSE']._serialized_start=847
  _globals['_GETOPERATORDELEGATEDSTAKEFORSTRATEGYRESPONSE']._serialized_end=954
  _globals['_GETDELEGATEDSTAKERSFOROPERATORREQUEST']._serialized_start=957
  _globals['_GETDELEGATEDSTAKERSFOROPERATORREQUEST']._serialized_end=1190
  _globals['_GETDELEGATEDSTAKERSFOROPERATORRESPONSE']._serialized_start=1193
  _globals['_GETDELEGATEDSTAKERSFOROPERATORRESPONSE']._serialized_end=1366
  _globals['_GETSTAKERSHARESREQUEST']._serialized_start=1368
  _globals['_GETSTAKERSHARESREQUEST']._serialized_end=1488
  _globals['_STAKERSHARE']._serialized_start=1491
  _globals['_STAKERSHARE']._serialized_end=1662
  _globals['_GETSTAKERSHARESRESPONSE']._serialized_start=1664
  _globals['_GETSTAKERSHARESRESPONSE']._serialized_end=1758
  _globals['_GETEIGENSTATECHANGESREQUEST']._serialized_start=1760
  _globals['_GETEIGENSTATECHANGESREQUEST']._serialized_end=1824
  _globals['_GETEIGENSTATECHANGESRESPONSE']._serialized_start=1826
  _globals['_GETEIGENSTATECHANGESRESPONSE']._serialized_end=1934
  _globals['_PROTOCOL']._serialized_start=1937
  _globals['_PROTOCOL']._serialized_end=3361
# @@protoc_insertion_point(module_scope)
