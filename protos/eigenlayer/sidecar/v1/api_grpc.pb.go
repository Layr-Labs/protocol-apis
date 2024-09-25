// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: eigenlayer/sidecar/v1/api.proto

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
	Rpc_GetBlockHeight_FullMethodName            = "/eigenlayer.sidecar.api.v1.Rpc/GetBlockHeight"
	Rpc_GetStateRoot_FullMethodName              = "/eigenlayer.sidecar.api.v1.Rpc/GetStateRoot"
	Rpc_GetRewardsRoot_FullMethodName            = "/eigenlayer.sidecar.api.v1.Rpc/GetRewardsRoot"
	Rpc_GenerateClaimProof_FullMethodName        = "/eigenlayer.sidecar.api.v1.Rpc/GenerateClaimProof"
	Rpc_GetAvailableRewards_FullMethodName       = "/eigenlayer.sidecar.api.v1.Rpc/GetAvailableRewards"
	Rpc_GetAvailableRewardsTokens_FullMethodName = "/eigenlayer.sidecar.api.v1.Rpc/GetAvailableRewardsTokens"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error)
	GetStateRoot(ctx context.Context, in *GetStateRootRequest, opts ...grpc.CallOption) (*GetStateRootResponse, error)
	GetRewardsRoot(ctx context.Context, in *GetRewardsRootRequest, opts ...grpc.CallOption) (*GetRewardsRootResponse, error)
	GenerateClaimProof(ctx context.Context, in *GenerateClaimProofRequest, opts ...grpc.CallOption) (*GenerateClaimProofResponse, error)
	GetAvailableRewards(ctx context.Context, in *GetAvailableRewardsRequest, opts ...grpc.CallOption) (*GetAvailableRewardsResponse, error)
	GetAvailableRewardsTokens(ctx context.Context, in *GetAvailableRewardsTokensRequest, opts ...grpc.CallOption) (*GetAvailableRewardsTokensResponse, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlockHeightResponse)
	err := c.cc.Invoke(ctx, Rpc_GetBlockHeight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetStateRoot(ctx context.Context, in *GetStateRootRequest, opts ...grpc.CallOption) (*GetStateRootResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStateRootResponse)
	err := c.cc.Invoke(ctx, Rpc_GetStateRoot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetRewardsRoot(ctx context.Context, in *GetRewardsRootRequest, opts ...grpc.CallOption) (*GetRewardsRootResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRewardsRootResponse)
	err := c.cc.Invoke(ctx, Rpc_GetRewardsRoot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GenerateClaimProof(ctx context.Context, in *GenerateClaimProofRequest, opts ...grpc.CallOption) (*GenerateClaimProofResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateClaimProofResponse)
	err := c.cc.Invoke(ctx, Rpc_GenerateClaimProof_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetAvailableRewards(ctx context.Context, in *GetAvailableRewardsRequest, opts ...grpc.CallOption) (*GetAvailableRewardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAvailableRewardsResponse)
	err := c.cc.Invoke(ctx, Rpc_GetAvailableRewards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetAvailableRewardsTokens(ctx context.Context, in *GetAvailableRewardsTokensRequest, opts ...grpc.CallOption) (*GetAvailableRewardsTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAvailableRewardsTokensResponse)
	err := c.cc.Invoke(ctx, Rpc_GetAvailableRewardsTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations should embed UnimplementedRpcServer
// for forward compatibility.
type RpcServer interface {
	GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error)
	GetStateRoot(context.Context, *GetStateRootRequest) (*GetStateRootResponse, error)
	GetRewardsRoot(context.Context, *GetRewardsRootRequest) (*GetRewardsRootResponse, error)
	GenerateClaimProof(context.Context, *GenerateClaimProofRequest) (*GenerateClaimProofResponse, error)
	GetAvailableRewards(context.Context, *GetAvailableRewardsRequest) (*GetAvailableRewardsResponse, error)
	GetAvailableRewardsTokens(context.Context, *GetAvailableRewardsTokensRequest) (*GetAvailableRewardsTokensResponse, error)
}

// UnimplementedRpcServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRpcServer struct{}

func (UnimplementedRpcServer) GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeight not implemented")
}
func (UnimplementedRpcServer) GetStateRoot(context.Context, *GetStateRootRequest) (*GetStateRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateRoot not implemented")
}
func (UnimplementedRpcServer) GetRewardsRoot(context.Context, *GetRewardsRootRequest) (*GetRewardsRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewardsRoot not implemented")
}
func (UnimplementedRpcServer) GenerateClaimProof(context.Context, *GenerateClaimProofRequest) (*GenerateClaimProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateClaimProof not implemented")
}
func (UnimplementedRpcServer) GetAvailableRewards(context.Context, *GetAvailableRewardsRequest) (*GetAvailableRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableRewards not implemented")
}
func (UnimplementedRpcServer) GetAvailableRewardsTokens(context.Context, *GetAvailableRewardsTokensRequest) (*GetAvailableRewardsTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableRewardsTokens not implemented")
}
func (UnimplementedRpcServer) testEmbeddedByValue() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	// If the following call pancis, it indicates UnimplementedRpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_GetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetBlockHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetBlockHeight(ctx, req.(*GetBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetStateRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetStateRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetStateRoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetStateRoot(ctx, req.(*GetStateRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetRewardsRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRewardsRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetRewardsRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetRewardsRoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetRewardsRoot(ctx, req.(*GetRewardsRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GenerateClaimProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateClaimProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GenerateClaimProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GenerateClaimProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GenerateClaimProof(ctx, req.(*GenerateClaimProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetAvailableRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetAvailableRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetAvailableRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetAvailableRewards(ctx, req.(*GetAvailableRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetAvailableRewardsTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableRewardsTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetAvailableRewardsTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetAvailableRewardsTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetAvailableRewardsTokens(ctx, req.(*GetAvailableRewardsTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eigenlayer.sidecar.api.v1.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockHeight",
			Handler:    _Rpc_GetBlockHeight_Handler,
		},
		{
			MethodName: "GetStateRoot",
			Handler:    _Rpc_GetStateRoot_Handler,
		},
		{
			MethodName: "GetRewardsRoot",
			Handler:    _Rpc_GetRewardsRoot_Handler,
		},
		{
			MethodName: "GenerateClaimProof",
			Handler:    _Rpc_GenerateClaimProof_Handler,
		},
		{
			MethodName: "GetAvailableRewards",
			Handler:    _Rpc_GetAvailableRewards_Handler,
		},
		{
			MethodName: "GetAvailableRewardsTokens",
			Handler:    _Rpc_GetAvailableRewardsTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eigenlayer/sidecar/v1/api.proto",
}
