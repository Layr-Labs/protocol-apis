# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from eigenlayer.sidecar.v1.protocol import protocol_pb2 as eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2


class ProtocolStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetRegisteredAvsForOperator = channel.unary_unary(
                '/eigenlayer.sidecar.v1.protocol.Protocol/GetRegisteredAvsForOperator',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetRegisteredAvsForOperatorRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetRegisteredAvsForOperatorResponse.FromString,
                _registered_method=True)
        self.GetDelegatedStrategiesForOperator = channel.unary_unary(
                '/eigenlayer.sidecar.v1.protocol.Protocol/GetDelegatedStrategiesForOperator',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStrategiesForOperatorRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStrategiesForOperatorResponse.FromString,
                _registered_method=True)
        self.GetOperatorDelegatedStakeForStrategy = channel.unary_unary(
                '/eigenlayer.sidecar.v1.protocol.Protocol/GetOperatorDelegatedStakeForStrategy',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetOperatorDelegatedStakeForStrategyRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetOperatorDelegatedStakeForStrategyResponse.FromString,
                _registered_method=True)
        self.GetDelegatedStakersForOperator = channel.unary_unary(
                '/eigenlayer.sidecar.v1.protocol.Protocol/GetDelegatedStakersForOperator',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStakersForOperatorRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStakersForOperatorResponse.FromString,
                _registered_method=True)
        self.GetStakerShares = channel.unary_unary(
                '/eigenlayer.sidecar.v1.protocol.Protocol/GetStakerShares',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetStakerSharesRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetStakerSharesResponse.FromString,
                _registered_method=True)


class ProtocolServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetRegisteredAvsForOperator(self, request, context):
        """GetRegisteredAvsForOperator returns the list of registered AVs for a given operator
        BlockHeight is optional, otherwise latest is used.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDelegatedStrategiesForOperator(self, request, context):
        """GetDelegatedStrategiesForOperator returns strategies an Operator has delegated
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOperatorDelegatedStakeForStrategy(self, request, context):
        """GetOperatorDelegatedStakeForStrategy returns the amount of delegated stake for a given strategy for an operator
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDelegatedStakersForOperator(self, request, context):
        """GetDelegatedStakersForOperator returns the list of stakers that have delegated to an operator.
        BlockHeight is optional, otherwise latest is used.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetStakerShares(self, request, context):
        """GetStakerShares returns the shares of a staker, and optionally, the operator operator they've delegated to and
        the AVS the operator has registered with.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProtocolServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetRegisteredAvsForOperator': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRegisteredAvsForOperator,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetRegisteredAvsForOperatorRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetRegisteredAvsForOperatorResponse.SerializeToString,
            ),
            'GetDelegatedStrategiesForOperator': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDelegatedStrategiesForOperator,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStrategiesForOperatorRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStrategiesForOperatorResponse.SerializeToString,
            ),
            'GetOperatorDelegatedStakeForStrategy': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOperatorDelegatedStakeForStrategy,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetOperatorDelegatedStakeForStrategyRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetOperatorDelegatedStakeForStrategyResponse.SerializeToString,
            ),
            'GetDelegatedStakersForOperator': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDelegatedStakersForOperator,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStakersForOperatorRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStakersForOperatorResponse.SerializeToString,
            ),
            'GetStakerShares': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStakerShares,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetStakerSharesRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetStakerSharesResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'eigenlayer.sidecar.v1.protocol.Protocol', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('eigenlayer.sidecar.v1.protocol.Protocol', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Protocol(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetRegisteredAvsForOperator(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/eigenlayer.sidecar.v1.protocol.Protocol/GetRegisteredAvsForOperator',
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetRegisteredAvsForOperatorRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetRegisteredAvsForOperatorResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetDelegatedStrategiesForOperator(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/eigenlayer.sidecar.v1.protocol.Protocol/GetDelegatedStrategiesForOperator',
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStrategiesForOperatorRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStrategiesForOperatorResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetOperatorDelegatedStakeForStrategy(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/eigenlayer.sidecar.v1.protocol.Protocol/GetOperatorDelegatedStakeForStrategy',
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetOperatorDelegatedStakeForStrategyRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetOperatorDelegatedStakeForStrategyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetDelegatedStakersForOperator(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/eigenlayer.sidecar.v1.protocol.Protocol/GetDelegatedStakersForOperator',
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStakersForOperatorRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetDelegatedStakersForOperatorResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetStakerShares(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/eigenlayer.sidecar.v1.protocol.Protocol/GetStakerShares',
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetStakerSharesRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_protocol_dot_protocol__pb2.GetStakerSharesResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
