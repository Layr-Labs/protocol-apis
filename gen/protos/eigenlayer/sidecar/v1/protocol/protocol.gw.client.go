// Code generated by protoc-gen-grpc-gateway-client. DO NOT EDIT.
// source: eigenlayer/sidecar/v1/protocol/protocol.proto

package protocol

import (
	context "context"
	fmt "fmt"
	gateway "github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
	url "net/url"
)

// ProtocolGatewayClient is the interface for Protocol service client.
type ProtocolGatewayClient interface {
	// GetRegisteredAvsForOperator returns the list of registered AVs for a given operator
	// BlockHeight is optional, otherwise latest is used.
	GetRegisteredAvsForOperator(context.Context, *GetRegisteredAvsForOperatorRequest) (*GetRegisteredAvsForOperatorResponse, error)
	// GetDelegatedStrategiesForOperator returns strategies an Operator has delegated
	GetDelegatedStrategiesForOperator(context.Context, *GetDelegatedStrategiesForOperatorRequest) (*GetDelegatedStrategiesForOperatorResponse, error)
	// GetOperatorDelegatedStakeForStrategy returns the amount of delegated stake for a given strategy for an operator
	GetOperatorDelegatedStakeForStrategy(context.Context, *GetOperatorDelegatedStakeForStrategyRequest) (*GetOperatorDelegatedStakeForStrategyResponse, error)
	// GetDelegatedStakersForOperator returns the list of stakers that have delegated to an operator.
	// BlockHeight is optional, otherwise latest is used.
	GetDelegatedStakersForOperator(context.Context, *GetDelegatedStakersForOperatorRequest) (*GetDelegatedStakersForOperatorResponse, error)
	// GetStakerShares returns the shares of a staker, and optionally, the operator operator they've delegated to and
	// the AVS the operator has registered with.
	GetStakerShares(context.Context, *GetStakerSharesRequest) (*GetStakerSharesResponse, error)
	GetEigenStateChanges(context.Context, *GetEigenStateChangesRequest) (*GetEigenStateChangesResponse, error)
}

func NewProtocolGatewayClient(c gateway.Client) ProtocolGatewayClient {
	return &protocolGatewayClient{
		gwc: c,
	}
}

type protocolGatewayClient struct {
	gwc gateway.Client
}

func (c *protocolGatewayClient) GetRegisteredAvsForOperator(ctx context.Context, req *GetRegisteredAvsForOperatorRequest) (*GetRegisteredAvsForOperatorResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/protocol/v1/operators/{operator_address}/registered-avs")
	gwReq.SetPathParam("operator_address", fmt.Sprintf("%v", req.OperatorAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetRegisteredAvsForOperatorResponse](ctx, gwReq)
}

func (c *protocolGatewayClient) GetDelegatedStrategiesForOperator(ctx context.Context, req *GetDelegatedStrategiesForOperatorRequest) (*GetDelegatedStrategiesForOperatorResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/protocol/v1/operators/{operator_address}/delegated-strategies")
	gwReq.SetPathParam("operator_address", fmt.Sprintf("%v", req.OperatorAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetDelegatedStrategiesForOperatorResponse](ctx, gwReq)
}

func (c *protocolGatewayClient) GetOperatorDelegatedStakeForStrategy(ctx context.Context, req *GetOperatorDelegatedStakeForStrategyRequest) (*GetOperatorDelegatedStakeForStrategyResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/protocol/v1/operators/{operator_address}/strategies/{strategy_address}/delegated-stake")
	gwReq.SetPathParam("operator_address", fmt.Sprintf("%v", req.OperatorAddress))
	gwReq.SetPathParam("strategy_address", fmt.Sprintf("%v", req.StrategyAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetOperatorDelegatedStakeForStrategyResponse](ctx, gwReq)
}

func (c *protocolGatewayClient) GetDelegatedStakersForOperator(ctx context.Context, req *GetDelegatedStakersForOperatorRequest) (*GetDelegatedStakersForOperatorResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/protocol/v1/operators/{operator_address}/delegated-stakers")
	gwReq.SetPathParam("operator_address", fmt.Sprintf("%v", req.OperatorAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	if req.Pagination != nil {
		q.Add("pagination.pageNumber", fmt.Sprintf("%v", req.Pagination.PageNumber))
		q.Add("pagination.pageSize", fmt.Sprintf("%v", req.Pagination.PageSize))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetDelegatedStakersForOperatorResponse](ctx, gwReq)
}

func (c *protocolGatewayClient) GetStakerShares(ctx context.Context, req *GetStakerSharesRequest) (*GetStakerSharesResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/protocol/v1/stakers/{staker_address}/shares")
	gwReq.SetPathParam("staker_address", fmt.Sprintf("%v", req.StakerAddress))
	q := url.Values{}
	if req.BlockHeight != nil {
		q.Add("blockHeight", fmt.Sprintf("%v", *req.BlockHeight))
	}
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetStakerSharesResponse](ctx, gwReq)
}

func (c *protocolGatewayClient) GetEigenStateChanges(ctx context.Context, req *GetEigenStateChangesRequest) (*GetEigenStateChangesResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/protocol/v1/eigen-state-changes")
	q := url.Values{}
	q.Add("blockHeight", fmt.Sprintf("%v", req.BlockHeight))
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetEigenStateChangesResponse](ctx, gwReq)
}
