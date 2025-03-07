// Code generated by protoc-gen-grpc-gateway-client. DO NOT EDIT.
// source: eigenlayer/sidecar/v1/sidecar/rpc.proto

package sidecar

import (
	context "context"
	fmt "fmt"
	gateway "github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
	url "net/url"
)

// RpcGatewayClient is the interface for Rpc service client.
type RpcGatewayClient interface {
	GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error)
	GetStateRoot(context.Context, *GetStateRootRequest) (*GetStateRootResponse, error)
	// About returns information about the running sidecar process
	About(context.Context, *AboutRequest) (*AboutResponse, error)
}

func NewRpcGatewayClient(c gateway.Client) RpcGatewayClient {
	return &rpcGatewayClient{
		gwc: c,
	}
}

type rpcGatewayClient struct {
	gwc gateway.Client
}

func (c *rpcGatewayClient) GetBlockHeight(ctx context.Context, req *GetBlockHeightRequest) (*GetBlockHeightResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rpc/v1/latest-block")
	q := url.Values{}
	q.Add("verified", fmt.Sprintf("%v", req.Verified))
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[GetBlockHeightResponse](ctx, gwReq)
}

func (c *rpcGatewayClient) GetStateRoot(ctx context.Context, req *GetStateRootRequest) (*GetStateRootResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rpc/v1/state-roots/{blockNumber}")
	gwReq.SetPathParam("blockNumber", fmt.Sprintf("%v", req.BlockNumber))
	return gateway.DoRequest[GetStateRootResponse](ctx, gwReq)
}

func (c *rpcGatewayClient) About(ctx context.Context, req *AboutRequest) (*AboutResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/rpc/v1/about")
	return gateway.DoRequest[AboutResponse](ctx, gwReq)
}
