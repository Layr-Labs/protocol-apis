// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: eigenlayer/sidecar/v1/rewards.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Rewards_GetRewardsRoot_FullMethodName                    = "/eigenlayer.sidecar.rewards.v1.Rewards/GetRewardsRoot"
	Rewards_GenerateRewards_FullMethodName                   = "/eigenlayer.sidecar.rewards.v1.Rewards/GenerateRewards"
	Rewards_GetRewardsForSnapshot_FullMethodName             = "/eigenlayer.sidecar.rewards.v1.Rewards/GetRewardsForSnapshot"
	Rewards_GetAttributableRewardsForSnapshot_FullMethodName = "/eigenlayer.sidecar.rewards.v1.Rewards/GetAttributableRewardsForSnapshot"
	Rewards_GenerateClaimProof_FullMethodName                = "/eigenlayer.sidecar.rewards.v1.Rewards/GenerateClaimProof"
	Rewards_GetAvailableRewards_FullMethodName               = "/eigenlayer.sidecar.rewards.v1.Rewards/GetAvailableRewards"
	Rewards_GetTotalClaimedRewards_FullMethodName            = "/eigenlayer.sidecar.rewards.v1.Rewards/GetTotalClaimedRewards"
	Rewards_GetAvailableRewardsTokens_FullMethodName         = "/eigenlayer.sidecar.rewards.v1.Rewards/GetAvailableRewardsTokens"
	Rewards_GetSummarizedRewardsForEarner_FullMethodName     = "/eigenlayer.sidecar.rewards.v1.Rewards/GetSummarizedRewardsForEarner"
	Rewards_GetClaimedRewardsByBlock_FullMethodName          = "/eigenlayer.sidecar.rewards.v1.Rewards/GetClaimedRewardsByBlock"
)

// RewardsClient is the client API for Rewards service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RewardsClient interface {
	// GetRewardsRoot returns the rewards root for the given block number
	GetRewardsRoot(ctx context.Context, in *GetRewardsRootRequest, opts ...grpc.CallOption) (*GetRewardsRootResponse, error)
	// GenerateRewards generates rewards for the given snapshot.
	// If respondWithRewardsData is true, the response will include the rewards data, otherwise
	// the sidecar will cache the data to be requested later.
	GenerateRewards(ctx context.Context, in *GenerateRewardsRequest, opts ...grpc.CallOption) (*GenerateRewardsResponse, error)
	// GetRewardsForSnapshot returns the rewards for the provided snapshot.
	// Useful if you generated the rewards and want to fetch them later.
	GetRewardsForSnapshot(ctx context.Context, in *GetRewardsForSnapshotRequest, opts ...grpc.CallOption) (*GetRewardsForSnapshotResponse, error)
	// GetAttributableRewardsForSnapshot returns the attributable rewards for the provided snapshot.
	// This takes the cumulative rewards amounts and breaks them down across operators, avss, strategies, etc
	GetAttributableRewardsForSnapshot(ctx context.Context, in *GetAttributableRewardsForSnapshotRequest, opts ...grpc.CallOption) (*GetAttributableRewardsForSnapshotResponse, error)
	// GenerateClaimProof generates a proof for the given earner address and tokens for claiming
	// tokens against the RewardsCoordinator
	GenerateClaimProof(ctx context.Context, in *GenerateClaimProofRequest, opts ...grpc.CallOption) (*GenerateClaimProofResponse, error)
	// GetAvailableRewards returns the available rewards for the given earner address
	// This takes the amount earned from the current active root and subtracts total claimed.
	GetAvailableRewards(ctx context.Context, in *GetAvailableRewardsRequest, opts ...grpc.CallOption) (*GetAvailableRewardsResponse, error)
	// GetTotalClaimedRewards returns the total claimed rewards for the given earner address
	// BlockHeight is optional. If omitted, the latest block height is used.
	GetTotalClaimedRewards(ctx context.Context, in *GetTotalClaimedRewardsRequest, opts ...grpc.CallOption) (*GetTotalClaimedRewardsResponse, error)
	// GetAvailableRewardsTokens returns the available rewards tokens for the given earner address
	GetAvailableRewardsTokens(ctx context.Context, in *GetAvailableRewardsTokensRequest, opts ...grpc.CallOption) (*GetAvailableRewardsTokensResponse, error)
	// GetSummarizedRewardsForEarner returns the summarized rewards for the given earner address, including how much
	// theyve earned, how much is currently claimable, and how much has been claimed.
	GetSummarizedRewardsForEarner(ctx context.Context, in *GetSummarizedRewardsForEarnerRequest, opts ...grpc.CallOption) (*GetSummarizedRewardsForEarnerResponse, error)
	// GetClaimedRewardsByBlock returns the claimed rewards for the given block height
	GetClaimedRewardsByBlock(ctx context.Context, in *GetClaimedRewardsByBlockRequest, opts ...grpc.CallOption) (*GetClaimedRewardsByBlockResponse, error)
}

type rewardsClient struct {
	cc grpc.ClientConnInterface
}

func NewRewardsClient(cc grpc.ClientConnInterface) RewardsClient {
	return &rewardsClient{cc}
}

func (c *rewardsClient) GetRewardsRoot(ctx context.Context, in *GetRewardsRootRequest, opts ...grpc.CallOption) (*GetRewardsRootResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRewardsRootResponse)
	err := c.cc.Invoke(ctx, Rewards_GetRewardsRoot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GenerateRewards(ctx context.Context, in *GenerateRewardsRequest, opts ...grpc.CallOption) (*GenerateRewardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateRewardsResponse)
	err := c.cc.Invoke(ctx, Rewards_GenerateRewards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetRewardsForSnapshot(ctx context.Context, in *GetRewardsForSnapshotRequest, opts ...grpc.CallOption) (*GetRewardsForSnapshotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRewardsForSnapshotResponse)
	err := c.cc.Invoke(ctx, Rewards_GetRewardsForSnapshot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetAttributableRewardsForSnapshot(ctx context.Context, in *GetAttributableRewardsForSnapshotRequest, opts ...grpc.CallOption) (*GetAttributableRewardsForSnapshotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttributableRewardsForSnapshotResponse)
	err := c.cc.Invoke(ctx, Rewards_GetAttributableRewardsForSnapshot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GenerateClaimProof(ctx context.Context, in *GenerateClaimProofRequest, opts ...grpc.CallOption) (*GenerateClaimProofResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateClaimProofResponse)
	err := c.cc.Invoke(ctx, Rewards_GenerateClaimProof_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetAvailableRewards(ctx context.Context, in *GetAvailableRewardsRequest, opts ...grpc.CallOption) (*GetAvailableRewardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAvailableRewardsResponse)
	err := c.cc.Invoke(ctx, Rewards_GetAvailableRewards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetTotalClaimedRewards(ctx context.Context, in *GetTotalClaimedRewardsRequest, opts ...grpc.CallOption) (*GetTotalClaimedRewardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTotalClaimedRewardsResponse)
	err := c.cc.Invoke(ctx, Rewards_GetTotalClaimedRewards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetAvailableRewardsTokens(ctx context.Context, in *GetAvailableRewardsTokensRequest, opts ...grpc.CallOption) (*GetAvailableRewardsTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAvailableRewardsTokensResponse)
	err := c.cc.Invoke(ctx, Rewards_GetAvailableRewardsTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetSummarizedRewardsForEarner(ctx context.Context, in *GetSummarizedRewardsForEarnerRequest, opts ...grpc.CallOption) (*GetSummarizedRewardsForEarnerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSummarizedRewardsForEarnerResponse)
	err := c.cc.Invoke(ctx, Rewards_GetSummarizedRewardsForEarner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsClient) GetClaimedRewardsByBlock(ctx context.Context, in *GetClaimedRewardsByBlockRequest, opts ...grpc.CallOption) (*GetClaimedRewardsByBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClaimedRewardsByBlockResponse)
	err := c.cc.Invoke(ctx, Rewards_GetClaimedRewardsByBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RewardsServer is the server API for Rewards service.
// All implementations should embed UnimplementedRewardsServer
// for forward compatibility.
type RewardsServer interface {
	// GetRewardsRoot returns the rewards root for the given block number
	GetRewardsRoot(context.Context, *GetRewardsRootRequest) (*GetRewardsRootResponse, error)
	// GenerateRewards generates rewards for the given snapshot.
	// If respondWithRewardsData is true, the response will include the rewards data, otherwise
	// the sidecar will cache the data to be requested later.
	GenerateRewards(context.Context, *GenerateRewardsRequest) (*GenerateRewardsResponse, error)
	// GetRewardsForSnapshot returns the rewards for the provided snapshot.
	// Useful if you generated the rewards and want to fetch them later.
	GetRewardsForSnapshot(context.Context, *GetRewardsForSnapshotRequest) (*GetRewardsForSnapshotResponse, error)
	// GetAttributableRewardsForSnapshot returns the attributable rewards for the provided snapshot.
	// This takes the cumulative rewards amounts and breaks them down across operators, avss, strategies, etc
	GetAttributableRewardsForSnapshot(context.Context, *GetAttributableRewardsForSnapshotRequest) (*GetAttributableRewardsForSnapshotResponse, error)
	// GenerateClaimProof generates a proof for the given earner address and tokens for claiming
	// tokens against the RewardsCoordinator
	GenerateClaimProof(context.Context, *GenerateClaimProofRequest) (*GenerateClaimProofResponse, error)
	// GetAvailableRewards returns the available rewards for the given earner address
	// This takes the amount earned from the current active root and subtracts total claimed.
	GetAvailableRewards(context.Context, *GetAvailableRewardsRequest) (*GetAvailableRewardsResponse, error)
	// GetTotalClaimedRewards returns the total claimed rewards for the given earner address
	// BlockHeight is optional. If omitted, the latest block height is used.
	GetTotalClaimedRewards(context.Context, *GetTotalClaimedRewardsRequest) (*GetTotalClaimedRewardsResponse, error)
	// GetAvailableRewardsTokens returns the available rewards tokens for the given earner address
	GetAvailableRewardsTokens(context.Context, *GetAvailableRewardsTokensRequest) (*GetAvailableRewardsTokensResponse, error)
	// GetSummarizedRewardsForEarner returns the summarized rewards for the given earner address, including how much
	// theyve earned, how much is currently claimable, and how much has been claimed.
	GetSummarizedRewardsForEarner(context.Context, *GetSummarizedRewardsForEarnerRequest) (*GetSummarizedRewardsForEarnerResponse, error)
	// GetClaimedRewardsByBlock returns the claimed rewards for the given block height
	GetClaimedRewardsByBlock(context.Context, *GetClaimedRewardsByBlockRequest) (*GetClaimedRewardsByBlockResponse, error)
}

// UnimplementedRewardsServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRewardsServer struct{}

func (UnimplementedRewardsServer) GetRewardsRoot(context.Context, *GetRewardsRootRequest) (*GetRewardsRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewardsRoot not implemented")
}
func (UnimplementedRewardsServer) GenerateRewards(context.Context, *GenerateRewardsRequest) (*GenerateRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRewards not implemented")
}
func (UnimplementedRewardsServer) GetRewardsForSnapshot(context.Context, *GetRewardsForSnapshotRequest) (*GetRewardsForSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewardsForSnapshot not implemented")
}
func (UnimplementedRewardsServer) GetAttributableRewardsForSnapshot(context.Context, *GetAttributableRewardsForSnapshotRequest) (*GetAttributableRewardsForSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributableRewardsForSnapshot not implemented")
}
func (UnimplementedRewardsServer) GenerateClaimProof(context.Context, *GenerateClaimProofRequest) (*GenerateClaimProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateClaimProof not implemented")
}
func (UnimplementedRewardsServer) GetAvailableRewards(context.Context, *GetAvailableRewardsRequest) (*GetAvailableRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableRewards not implemented")
}
func (UnimplementedRewardsServer) GetTotalClaimedRewards(context.Context, *GetTotalClaimedRewardsRequest) (*GetTotalClaimedRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalClaimedRewards not implemented")
}
func (UnimplementedRewardsServer) GetAvailableRewardsTokens(context.Context, *GetAvailableRewardsTokensRequest) (*GetAvailableRewardsTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableRewardsTokens not implemented")
}
func (UnimplementedRewardsServer) GetSummarizedRewardsForEarner(context.Context, *GetSummarizedRewardsForEarnerRequest) (*GetSummarizedRewardsForEarnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummarizedRewardsForEarner not implemented")
}
func (UnimplementedRewardsServer) GetClaimedRewardsByBlock(context.Context, *GetClaimedRewardsByBlockRequest) (*GetClaimedRewardsByBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimedRewardsByBlock not implemented")
}
func (UnimplementedRewardsServer) testEmbeddedByValue() {}

// UnsafeRewardsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RewardsServer will
// result in compilation errors.
type UnsafeRewardsServer interface {
	mustEmbedUnimplementedRewardsServer()
}

func RegisterRewardsServer(s grpc.ServiceRegistrar, srv RewardsServer) {
	// If the following call pancis, it indicates UnimplementedRewardsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rewards_ServiceDesc, srv)
}

func _Rewards_GetRewardsRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRewardsRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetRewardsRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetRewardsRoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetRewardsRoot(ctx, req.(*GetRewardsRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GenerateRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GenerateRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GenerateRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GenerateRewards(ctx, req.(*GenerateRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetRewardsForSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRewardsForSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetRewardsForSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetRewardsForSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetRewardsForSnapshot(ctx, req.(*GetRewardsForSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetAttributableRewardsForSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributableRewardsForSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetAttributableRewardsForSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetAttributableRewardsForSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetAttributableRewardsForSnapshot(ctx, req.(*GetAttributableRewardsForSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GenerateClaimProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateClaimProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GenerateClaimProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GenerateClaimProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GenerateClaimProof(ctx, req.(*GenerateClaimProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetAvailableRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetAvailableRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetAvailableRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetAvailableRewards(ctx, req.(*GetAvailableRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetTotalClaimedRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalClaimedRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetTotalClaimedRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetTotalClaimedRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetTotalClaimedRewards(ctx, req.(*GetTotalClaimedRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetAvailableRewardsTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableRewardsTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetAvailableRewardsTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetAvailableRewardsTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetAvailableRewardsTokens(ctx, req.(*GetAvailableRewardsTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetSummarizedRewardsForEarner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummarizedRewardsForEarnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetSummarizedRewardsForEarner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetSummarizedRewardsForEarner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetSummarizedRewardsForEarner(ctx, req.(*GetSummarizedRewardsForEarnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rewards_GetClaimedRewardsByBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClaimedRewardsByBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServer).GetClaimedRewardsByBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rewards_GetClaimedRewardsByBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServer).GetClaimedRewardsByBlock(ctx, req.(*GetClaimedRewardsByBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rewards_ServiceDesc is the grpc.ServiceDesc for Rewards service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rewards_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eigenlayer.sidecar.rewards.v1.Rewards",
	HandlerType: (*RewardsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRewardsRoot",
			Handler:    _Rewards_GetRewardsRoot_Handler,
		},
		{
			MethodName: "GenerateRewards",
			Handler:    _Rewards_GenerateRewards_Handler,
		},
		{
			MethodName: "GetRewardsForSnapshot",
			Handler:    _Rewards_GetRewardsForSnapshot_Handler,
		},
		{
			MethodName: "GetAttributableRewardsForSnapshot",
			Handler:    _Rewards_GetAttributableRewardsForSnapshot_Handler,
		},
		{
			MethodName: "GenerateClaimProof",
			Handler:    _Rewards_GenerateClaimProof_Handler,
		},
		{
			MethodName: "GetAvailableRewards",
			Handler:    _Rewards_GetAvailableRewards_Handler,
		},
		{
			MethodName: "GetTotalClaimedRewards",
			Handler:    _Rewards_GetTotalClaimedRewards_Handler,
		},
		{
			MethodName: "GetAvailableRewardsTokens",
			Handler:    _Rewards_GetAvailableRewardsTokens_Handler,
		},
		{
			MethodName: "GetSummarizedRewardsForEarner",
			Handler:    _Rewards_GetSummarizedRewardsForEarner_Handler,
		},
		{
			MethodName: "GetClaimedRewardsByBlock",
			Handler:    _Rewards_GetClaimedRewardsByBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eigenlayer/sidecar/v1/rewards.proto",
}