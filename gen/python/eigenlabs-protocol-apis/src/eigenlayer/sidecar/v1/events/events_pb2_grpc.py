# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from eigenlayer.sidecar.v1.events import events_pb2 as eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2


class EventsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.StreamEigenStateChanges = channel.unary_stream(
                '/eigenlayer.sidecar.v1.events.Events/StreamEigenStateChanges',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamEigenStateChangesRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamEigenStateChangesResponse.FromString,
                _registered_method=True)
        self.StreamIndexedBlocks = channel.unary_stream(
                '/eigenlayer.sidecar.v1.events.Events/StreamIndexedBlocks',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamIndexedBlocksRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamIndexedBlocksResponse.FromString,
                _registered_method=True)


class EventsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def StreamEigenStateChanges(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StreamIndexedBlocks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EventsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'StreamEigenStateChanges': grpc.unary_stream_rpc_method_handler(
                    servicer.StreamEigenStateChanges,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamEigenStateChangesRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamEigenStateChangesResponse.SerializeToString,
            ),
            'StreamIndexedBlocks': grpc.unary_stream_rpc_method_handler(
                    servicer.StreamIndexedBlocks,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamIndexedBlocksRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamIndexedBlocksResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'eigenlayer.sidecar.v1.events.Events', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('eigenlayer.sidecar.v1.events.Events', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Events(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def StreamEigenStateChanges(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/eigenlayer.sidecar.v1.events.Events/StreamEigenStateChanges',
            eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamEigenStateChangesRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamEigenStateChangesResponse.FromString,
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
    def StreamIndexedBlocks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/eigenlayer.sidecar.v1.events.Events/StreamIndexedBlocks',
            eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamIndexedBlocksRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_events_dot_events__pb2.StreamIndexedBlocksResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
