# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from eigenlayer.sidecar.v1.rewards import rewards_pb2 as eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2


class RewardsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetRewardsRoot = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetRewardsRoot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsRootRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsRootResponse.FromString,
                _registered_method=True)
        self.GenerateRewards = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateRewards',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsResponse.FromString,
                _registered_method=True)
        self.GenerateStakerOperators = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateStakerOperators',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateStakerOperatorsRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateStakerOperatorsResponse.FromString,
                _registered_method=True)
        self.BackfillStakerOperators = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/BackfillStakerOperators',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.BackfillStakerOperatorsRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.BackfillStakerOperatorsResponse.FromString,
                _registered_method=True)
        self.GenerateRewardsRoot = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateRewardsRoot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRootRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRootResponse.FromString,
                _registered_method=True)
        self.GetRewardsForSnapshot = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetRewardsForSnapshot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsForSnapshotRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsForSnapshotResponse.FromString,
                _registered_method=True)
        self.GetAttributableRewardsForSnapshot = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetAttributableRewardsForSnapshot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForSnapshotRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForSnapshotResponse.FromString,
                _registered_method=True)
        self.GetAttributableRewardsForDistributionRoot = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetAttributableRewardsForDistributionRoot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForDistributionRootRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForDistributionRootResponse.FromString,
                _registered_method=True)
        self.GetRewardsByAvsForDistributionRoot = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetRewardsByAvsForDistributionRoot',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsByAvsForDistributionRootRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsByAvsForDistributionRootResponse.FromString,
                _registered_method=True)
        self.GenerateClaimProof = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateClaimProof',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateClaimProofRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateClaimProofResponse.FromString,
                _registered_method=True)
        self.GetClaimableRewards = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetClaimableRewards',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimableRewardsRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimableRewardsResponse.FromString,
                _registered_method=True)
        self.GetTotalClaimedRewards = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetTotalClaimedRewards',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetTotalClaimedRewardsRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetTotalClaimedRewardsResponse.FromString,
                _registered_method=True)
        self.GetAvailableRewardsTokens = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetAvailableRewardsTokens',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAvailableRewardsTokensRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAvailableRewardsTokensResponse.FromString,
                _registered_method=True)
        self.GetSummarizedRewardsForEarner = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetSummarizedRewardsForEarner',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetSummarizedRewardsForEarnerRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetSummarizedRewardsForEarnerResponse.FromString,
                _registered_method=True)
        self.GetClaimedRewardsByBlock = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/GetClaimedRewardsByBlock',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimedRewardsByBlockRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimedRewardsByBlockResponse.FromString,
                _registered_method=True)
        self.ListClaimedRewardsByBlockRange = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/ListClaimedRewardsByBlockRange',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListClaimedRewardsByBlockRangeRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListClaimedRewardsByBlockRangeResponse.FromString,
                _registered_method=True)
        self.ListDistributionRoots = channel.unary_unary(
                '/eigenlayer.sidecar.v1.rewards.Rewards/ListDistributionRoots',
                request_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListDistributionRootsRequest.SerializeToString,
                response_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListDistributionRootsResponse.FromString,
                _registered_method=True)


class RewardsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetRewardsRoot(self, request, context):
        """GetRewardsRoot returns the posted on-chain root for the give block height
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateRewards(self, request, context):
        """GenerateRewards generates rewards for the given cutoff_date.
        If respondWithRewardsData is true, the response will include the rewards data, otherwise
        the sidecar will cache the data to be requested later.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateStakerOperators(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BackfillStakerOperators(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateRewardsRoot(self, request, context):
        """GenerateRewardsRoot generates a rewards root for the given snapshot.
        Returns an error if the rewards have not been generated for the snapshot.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRewardsForSnapshot(self, request, context):
        """GetRewardsForSnapshot returns the rewards for the provided snapshot.
        Useful if you generated the rewards and want to fetch them later.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAttributableRewardsForSnapshot(self, request, context):
        """GetAttributableRewardsForSnapshot returns the attributable rewards for the provided snapshot.
        This takes the cumulative rewards amounts and breaks them down across operators, avss, strategies, etc
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAttributableRewardsForDistributionRoot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRewardsByAvsForDistributionRoot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateClaimProof(self, request, context):
        """GenerateClaimProof generates a proof for the given earner address and tokens for claiming
        tokens against the RewardsCoordinator
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetClaimableRewards(self, request, context):
        """GetClaimableRewards returns the claimable rewards for the given earner address.
        This takes the current active tokens earned and subtracts total claimed.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTotalClaimedRewards(self, request, context):
        """GetTotalClaimedRewards returns the total claimed rewards for the given earner address, summed up to and including
        the provided blockHeight. If a blockHeight is omitted, the most recent indexed block is used.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAvailableRewardsTokens(self, request, context):
        """GetAvailableRewardsTokens returns the available rewards tokens for the given earner address
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetSummarizedRewardsForEarner(self, request, context):
        """GetSummarizedRewardsForEarner returns the summarized rewards for the given earner address, including:
        - total tokens earned
        - total tokens active (subset of earned that are in roots that have surpassed their activation delay)
        - total claimed
        - total claimable (total active - total claimed)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetClaimedRewardsByBlock(self, request, context):
        """GetClaimedRewardsByBlock returns the claimed rewards for the provided block height
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListClaimedRewardsByBlockRange(self, request, context):
        """ListClaimedRewardsByBlockRange returns the claimed rewards for the given earner address and block range,
        inclusive of the start and end block heights
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListDistributionRoots(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RewardsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetRewardsRoot': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRewardsRoot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsRootRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsRootResponse.SerializeToString,
            ),
            'GenerateRewards': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateRewards,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsResponse.SerializeToString,
            ),
            'GenerateStakerOperators': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateStakerOperators,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateStakerOperatorsRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateStakerOperatorsResponse.SerializeToString,
            ),
            'BackfillStakerOperators': grpc.unary_unary_rpc_method_handler(
                    servicer.BackfillStakerOperators,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.BackfillStakerOperatorsRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.BackfillStakerOperatorsResponse.SerializeToString,
            ),
            'GenerateRewardsRoot': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateRewardsRoot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRootRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRootResponse.SerializeToString,
            ),
            'GetRewardsForSnapshot': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRewardsForSnapshot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsForSnapshotRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsForSnapshotResponse.SerializeToString,
            ),
            'GetAttributableRewardsForSnapshot': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAttributableRewardsForSnapshot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForSnapshotRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForSnapshotResponse.SerializeToString,
            ),
            'GetAttributableRewardsForDistributionRoot': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAttributableRewardsForDistributionRoot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForDistributionRootRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForDistributionRootResponse.SerializeToString,
            ),
            'GetRewardsByAvsForDistributionRoot': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRewardsByAvsForDistributionRoot,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsByAvsForDistributionRootRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsByAvsForDistributionRootResponse.SerializeToString,
            ),
            'GenerateClaimProof': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateClaimProof,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateClaimProofRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateClaimProofResponse.SerializeToString,
            ),
            'GetClaimableRewards': grpc.unary_unary_rpc_method_handler(
                    servicer.GetClaimableRewards,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimableRewardsRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimableRewardsResponse.SerializeToString,
            ),
            'GetTotalClaimedRewards': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTotalClaimedRewards,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetTotalClaimedRewardsRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetTotalClaimedRewardsResponse.SerializeToString,
            ),
            'GetAvailableRewardsTokens': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAvailableRewardsTokens,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAvailableRewardsTokensRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAvailableRewardsTokensResponse.SerializeToString,
            ),
            'GetSummarizedRewardsForEarner': grpc.unary_unary_rpc_method_handler(
                    servicer.GetSummarizedRewardsForEarner,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetSummarizedRewardsForEarnerRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetSummarizedRewardsForEarnerResponse.SerializeToString,
            ),
            'GetClaimedRewardsByBlock': grpc.unary_unary_rpc_method_handler(
                    servicer.GetClaimedRewardsByBlock,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimedRewardsByBlockRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimedRewardsByBlockResponse.SerializeToString,
            ),
            'ListClaimedRewardsByBlockRange': grpc.unary_unary_rpc_method_handler(
                    servicer.ListClaimedRewardsByBlockRange,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListClaimedRewardsByBlockRangeRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListClaimedRewardsByBlockRangeResponse.SerializeToString,
            ),
            'ListDistributionRoots': grpc.unary_unary_rpc_method_handler(
                    servicer.ListDistributionRoots,
                    request_deserializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListDistributionRootsRequest.FromString,
                    response_serializer=eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListDistributionRootsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'eigenlayer.sidecar.v1.rewards.Rewards', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('eigenlayer.sidecar.v1.rewards.Rewards', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Rewards(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetRewardsRoot(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetRewardsRoot',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsRootRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsRootResponse.FromString,
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
    def GenerateRewards(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateRewards',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsResponse.FromString,
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
    def GenerateStakerOperators(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateStakerOperators',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateStakerOperatorsRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateStakerOperatorsResponse.FromString,
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
    def BackfillStakerOperators(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/BackfillStakerOperators',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.BackfillStakerOperatorsRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.BackfillStakerOperatorsResponse.FromString,
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
    def GenerateRewardsRoot(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateRewardsRoot',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRootRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateRewardsRootResponse.FromString,
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
    def GetRewardsForSnapshot(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetRewardsForSnapshot',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsForSnapshotRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsForSnapshotResponse.FromString,
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
    def GetAttributableRewardsForSnapshot(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetAttributableRewardsForSnapshot',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForSnapshotRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForSnapshotResponse.FromString,
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
    def GetAttributableRewardsForDistributionRoot(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetAttributableRewardsForDistributionRoot',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForDistributionRootRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAttributableRewardsForDistributionRootResponse.FromString,
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
    def GetRewardsByAvsForDistributionRoot(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetRewardsByAvsForDistributionRoot',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsByAvsForDistributionRootRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetRewardsByAvsForDistributionRootResponse.FromString,
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
    def GenerateClaimProof(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GenerateClaimProof',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateClaimProofRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GenerateClaimProofResponse.FromString,
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
    def GetClaimableRewards(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetClaimableRewards',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimableRewardsRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimableRewardsResponse.FromString,
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
    def GetTotalClaimedRewards(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetTotalClaimedRewards',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetTotalClaimedRewardsRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetTotalClaimedRewardsResponse.FromString,
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
    def GetAvailableRewardsTokens(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetAvailableRewardsTokens',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAvailableRewardsTokensRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetAvailableRewardsTokensResponse.FromString,
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
    def GetSummarizedRewardsForEarner(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetSummarizedRewardsForEarner',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetSummarizedRewardsForEarnerRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetSummarizedRewardsForEarnerResponse.FromString,
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
    def GetClaimedRewardsByBlock(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/GetClaimedRewardsByBlock',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimedRewardsByBlockRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.GetClaimedRewardsByBlockResponse.FromString,
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
    def ListClaimedRewardsByBlockRange(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/ListClaimedRewardsByBlockRange',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListClaimedRewardsByBlockRangeRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListClaimedRewardsByBlockRangeResponse.FromString,
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
    def ListDistributionRoots(request,
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
            '/eigenlayer.sidecar.v1.rewards.Rewards/ListDistributionRoots',
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListDistributionRootsRequest.SerializeToString,
            eigenlayer_dot_sidecar_dot_v1_dot_rewards_dot_rewards__pb2.ListDistributionRootsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
