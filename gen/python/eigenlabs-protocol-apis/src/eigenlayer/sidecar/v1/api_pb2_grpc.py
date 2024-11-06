# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from eigenlayer.sidecar.v1 import api_pb2 as eigenlayer_dot_sidecar_dot_v1_dot_api__pb2


class RpcStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetBlockHeight = channel.unary_unary(
                '/eigenlayer.sidecar.api.v1.Rpc/GetBlockHeight',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetBlockHeightRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetBlockHeightResponse.FromString,
                _registered_method=True)
        self.GetStateRoot = channel.unary_unary(
                '/eigenlayer.sidecar.api.v1.Rpc/GetStateRoot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetStateRootRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetStateRootResponse.FromString,
                _registered_method=True)


class RpcServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetBlockHeight(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetStateRoot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RpcServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetBlockHeight': grpc.unary_unary_rpc_method_handler(
                    servicer.GetBlockHeight,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetBlockHeightRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetBlockHeightResponse.SerializeToString,
            ),
            'GetStateRoot': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStateRoot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetStateRootRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetStateRootResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'eigenlayer.sidecar.api.v1.Rpc', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('eigenlayer.sidecar.api.v1.Rpc', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Rpc(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetBlockHeight(request,
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
            '/eigenlayer.sidecar.api.v1.Rpc/GetBlockHeight',
            eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetBlockHeightRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetBlockHeightResponse.FromString,
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
    def GetStateRoot(request,
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
            '/eigenlayer.sidecar.api.v1.Rpc/GetStateRoot',
            eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetStateRootRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_api__pb2.GetStateRootResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
