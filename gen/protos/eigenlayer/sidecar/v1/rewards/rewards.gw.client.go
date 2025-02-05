// Code generated by protoc-gen-grpc-gateway-client. DO NOT EDIT.
// source: eigenlayer/sidecar/v1/rewards/rewards.proto

package rewards

import (
	context "context"
	fmt "fmt"
	gateway "github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
	url "net/url"
)

// RewardsGatewayClient is the interface for Rewards service client.
type RewardsGatewayClient interface {
	// GetRewardsRoot returns the posted on-chain root for the give block height
	GetRewardsRoot(context.Context, *GetRewardsRootRequest) (*GetRewardsRootResponse, error)
	// GenerateRewards generates rewards for the given cutoff_date.
	// If respondWithRewardsData is true, the response will include the rewards data, otherwise
	// the sidecar will cache the data to be requested later.
	GenerateRewards(context.Context, *GenerateRewardsRequest) (*GenerateRewardsResponse, error)
	GenerateStakerOperators(context.Context, *GenerateStakerOperatorsRequest) (*GenerateStakerOperatorsResponse, error)
	BackfillStakerOperators(context.Context, *BackfillStakerOperatorsRequest) (*BackfillStakerOperatorsResponse, error)
	// GenerateRewardsRoot generates a rewards root for the given snapshot.
	// Returns an error if the rewards have not been generated for the snapshot.
	GenerateRewardsRoot(context.Context, *GenerateRewardsRootRequest) (*GenerateRewardsRootResponse, error)
	// GetRewardsForSnapshot returns the rewards for the provided snapshot.
	// Useful if you generated the rewards and want to fetch them later.
	GetRewardsForSnapshot(context.Context, *GetRewardsForSnapshotRequest) (*GetRewardsForSnapshotResponse, error)
	// GetAttributableRewardsForSnapshot returns the attributable rewards for the provided snapshot.
	// This takes the cumulative rewards amounts and breaks them down across operators, avss, strategies, etc
	GetAttributableRewardsForSnapshot(context.Context, *GetAttributableRewardsForSnapshotRequest) (*GetAttributableRewardsForSnapshotResponse, error)
	GetAttributableRewardsForDistributionRoot(context.Context, *GetAttributableRewardsForDistributionRootRequest) (*GetAttributableRewardsForDistributionRootResponse, error)
	GetRewardsByAvsForDistributionRoot(context.Context, *GetRewardsByAvsForDistributionRootRequest) (*GetRewardsByAvsForDistributionRootResponse, error)
	// GenerateClaimProof generates a proof for the given earner address and tokens for claiming
	// tokens against the RewardsCoordinator
	GenerateClaimProof(context.Context, *GenerateClaimProofRequest) (*GenerateClaimProofResponse, error)
	// GetClaimableRewards returns the claimable rewards for the given earner address.
	// This takes the current active tokens earned and subtracts total claimed.
	GetClaimableRewards(context.Context, *GetClaimableRewardsRequest) (*GetClaimableRewardsResponse, error)
	// GetTotalClaimedRewards returns the total claimed rewards for the given earner address, summed up to and including
	// the provided blockHeight. If a blockHeight is omitted, the most recent indexed block is used.
	GetTotalClaimedRewards(context.Context, *GetTotalClaimedRewardsRequest) (*GetTotalClaimedRewardsResponse, error)
	// GetAvailableRewardsTokens returns the available rewards tokens for the given earner address
	GetAvailableRewardsTokens(context.Context, *GetAvailableRewardsTokensRequest) (*GetAvailableRewardsTokensResponse, error)
	// GetSummarizedRewardsForEarner returns the summarized rewards for the given earner address, including:
	// - total tokens earned
	// - total tokens active (subset of earned that are in roots that have surpassed their activation delay)
	// - total claimed
	// - total claimable (total active - total claimed)
	GetSummarizedRewardsForEarner(context.Context, *GetSummarizedRewardsForEarnerRequest) (*GetSummarizedRewardsForEarnerResponse, error)
	// GetClaimedRewardsByBlock returns the claimed rewards for the provided block height
	GetClaimedRewardsByBlock(context.Context, *GetClaimedRewardsByBlockRequest) (*GetClaimedRewardsByBlockResponse, error)
	// ListClaimedRewardsByBlockRange returns the claimed rewards for the given earner address and block range,
	// inclusive of the start and end block heights
	ListClaimedRewardsByBlockRange(context.Context, *ListClaimedRewardsByBlockRangeRequest) (*ListClaimedRewardsByBlockRangeResponse, error)
	ListDistributionRoots(context.Context, *ListDistributionRootsRequest) (*ListDistributionRootsResponse, error)
}

func NewRewardsGatewayClient(c gateway.Client) RewardsGatewayClient {
	return &rewardsGatewayClient{
		gwc: c,
	}
}

type rewardsGatewayClient struct {
	gwc gateway.Client
}

func (c *rewardsGatewayClient) GetRewardsRoot(ctx context.Context, req *GetRewardsRootRequest) (*GetRewardsRootResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/blocks/{block_height}/rewards-root")
	gwReq.SetPathParam("block_height", fmt.Sprintf("%v", req.BlockHeight))
	return gateway.DoRequest[GetRewardsRootResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GenerateRewards(ctx context.Context, req *GenerateRewardsRequest) (*GenerateRewardsResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/rewards/v1/generate-rewards")
	gwReq.SetBody(req)
	return gateway.DoRequest[GenerateRewardsResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GenerateStakerOperators(ctx context.Context, req *GenerateStakerOperatorsRequest) (*GenerateStakerOperatorsResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/rewards/v1/generate-staker-operators")
	gwReq.SetBody(req)
	return gateway.DoRequest[GenerateStakerOperatorsResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) BackfillStakerOperators(ctx context.Context, req *BackfillStakerOperatorsRequest) (*BackfillStakerOperatorsResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/rewards/v1/backfill-staker-operators")
	gwReq.SetBody(req)
	return gateway.DoRequest[BackfillStakerOperatorsResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GenerateRewardsRoot(ctx context.Context, req *GenerateRewardsRootRequest) (*GenerateRewardsRootResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/rewards/v1/generate-rewards-root")
	gwReq.SetBody(req)
	return gateway.DoRequest[GenerateRewardsRootResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetRewardsForSnapshot(ctx context.Context, req *GetRewardsForSnapshotRequest) (*GetRewardsForSnapshotResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/rewards/{snapshot}")
	gwReq.SetPathParam("snapshot", fmt.Sprintf("%v", req.Snapshot))
	return gateway.DoRequest[GetRewardsForSnapshotResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetAttributableRewardsForSnapshot(ctx context.Context, req *GetAttributableRewardsForSnapshotRequest) (*GetAttributableRewardsForSnapshotResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/attributable-rewards/{snapshot}")
	gwReq.SetPathParam("snapshot", fmt.Sprintf("%v", req.Snapshot))
	return gateway.DoRequest[GetAttributableRewardsForSnapshotResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetAttributableRewardsForDistributionRoot(ctx context.Context, req *GetAttributableRewardsForDistributionRootRequest) (*GetAttributableRewardsForDistributionRootResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/attributable-rewards-by-root/{distribution_root}")
	gwReq.SetPathParam("distribution_root", fmt.Sprintf("%v", req.DistributionRoot))
	return gateway.DoRequest[GetAttributableRewardsForDistributionRootResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetRewardsByAvsForDistributionRoot(ctx context.Context, req *GetRewardsByAvsForDistributionRootRequest) (*GetRewardsByAvsForDistributionRootResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/avs-rewards-by-root/{distribution_root}")
	gwReq.SetPathParam("distribution_root", fmt.Sprintf("%v", req.DistributionRoot))
	return gateway.DoRequest[GetRewardsByAvsForDistributionRootResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GenerateClaimProof(ctx context.Context, req *GenerateClaimProofRequest) (*GenerateClaimProofResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/rewards/v1/claim-proof")
	gwReq.SetBody(req)
	return gateway.DoRequest[GenerateClaimProofResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetClaimableRewards(ctx context.Context, req *GetClaimableRewardsRequest) (*GetClaimableRewardsResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/earners/{earner_address}/claimable-rewards")
	gwReq.SetPathParam("earner_address", fmt.Sprintf("%v", req.EarnerAddress))
	q := url.Values{}
	q.Add("blockHeight", fmt.Sprintf("%v", req.BlockHeight))
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetClaimableRewardsResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetTotalClaimedRewards(ctx context.Context, req *GetTotalClaimedRewardsRequest) (*GetTotalClaimedRewardsResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/earners/{earner_address}/total-claimed-rewards")
	gwReq.SetPathParam("earner_address", fmt.Sprintf("%v", req.EarnerAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetTotalClaimedRewardsResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetAvailableRewardsTokens(ctx context.Context, req *GetAvailableRewardsTokensRequest) (*GetAvailableRewardsTokensResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/earners/{earner_address}/available-rewards-tokens")
	gwReq.SetPathParam("earner_address", fmt.Sprintf("%v", req.EarnerAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetAvailableRewardsTokensResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetSummarizedRewardsForEarner(ctx context.Context, req *GetSummarizedRewardsForEarnerRequest) (*GetSummarizedRewardsForEarnerResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/earners/{earner_address}/summarized-rewards")
	gwReq.SetPathParam("earner_address", fmt.Sprintf("%v", req.EarnerAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetSummarizedRewardsForEarnerResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) GetClaimedRewardsByBlock(ctx context.Context, req *GetClaimedRewardsByBlockRequest) (*GetClaimedRewardsByBlockResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/blocks/{block_height}/claimed-rewards")
	gwReq.SetPathParam("block_height", fmt.Sprintf("%v", req.BlockHeight))
	return gateway.DoRequest[GetClaimedRewardsByBlockResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) ListClaimedRewardsByBlockRange(ctx context.Context, req *ListClaimedRewardsByBlockRangeRequest) (*ListClaimedRewardsByBlockRangeResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/earners/{earner_address}/claimed-rewards")
	gwReq.SetPathParam("earner_address", fmt.Sprintf("%v", req.EarnerAddress))
	q := url.Values{}
	q.Add("startBlockHeight", fmt.Sprintf("%v", req.StartBlockHeight))
	q.Add("endBlockHeight", fmt.Sprintf("%v", req.EndBlockHeight))
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[ListClaimedRewardsByBlockRangeResponse](ctx, gwReq)
}

func (c *rewardsGatewayClient) ListDistributionRoots(ctx context.Context, req *ListDistributionRootsRequest) (*ListDistributionRootsResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rewards/v1/distribution-roots")
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[ListDistributionRootsResponse](ctx, gwReq)
}
