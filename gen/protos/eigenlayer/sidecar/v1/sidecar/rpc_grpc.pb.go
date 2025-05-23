// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: eigenlayer/sidecar/v1/sidecar/rpc.proto

package sidecar

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
	Rpc_GetBlockHeight_FullMethodName = "/eigenlayer.sidecar.v1.sidecar.Rpc/GetBlockHeight"
	Rpc_GetStateRoot_FullMethodName   = "/eigenlayer.sidecar.v1.sidecar.Rpc/GetStateRoot"
	Rpc_About_FullMethodName          = "/eigenlayer.sidecar.v1.sidecar.Rpc/About"
	Rpc_LoadContract_FullMethodName   = "/eigenlayer.sidecar.v1.sidecar.Rpc/LoadContract"
	Rpc_LoadContracts_FullMethodName  = "/eigenlayer.sidecar.v1.sidecar.Rpc/LoadContracts"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error)
	GetStateRoot(ctx context.Context, in *GetStateRootRequest, opts ...grpc.CallOption) (*GetStateRootResponse, error)
	// About returns information about the running sidecar process
	About(ctx context.Context, in *AboutRequest, opts ...grpc.CallOption) (*AboutResponse, error)
	LoadContract(ctx context.Context, in *LoadContractRequest, opts ...grpc.CallOption) (*LoadContractResponse, error)
	LoadContracts(ctx context.Context, in *LoadContractsRequest, opts ...grpc.CallOption) (*LoadContractsResponse, error)
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

func (c *rpcClient) About(ctx context.Context, in *AboutRequest, opts ...grpc.CallOption) (*AboutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AboutResponse)
	err := c.cc.Invoke(ctx, Rpc_About_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) LoadContract(ctx context.Context, in *LoadContractRequest, opts ...grpc.CallOption) (*LoadContractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadContractResponse)
	err := c.cc.Invoke(ctx, Rpc_LoadContract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) LoadContracts(ctx context.Context, in *LoadContractsRequest, opts ...grpc.CallOption) (*LoadContractsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadContractsResponse)
	err := c.cc.Invoke(ctx, Rpc_LoadContracts_FullMethodName, in, out, cOpts...)
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
	// About returns information about the running sidecar process
	About(context.Context, *AboutRequest) (*AboutResponse, error)
	LoadContract(context.Context, *LoadContractRequest) (*LoadContractResponse, error)
	LoadContracts(context.Context, *LoadContractsRequest) (*LoadContractsResponse, error)
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
func (UnimplementedRpcServer) About(context.Context, *AboutRequest) (*AboutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method About not implemented")
}
func (UnimplementedRpcServer) LoadContract(context.Context, *LoadContractRequest) (*LoadContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadContract not implemented")
}
func (UnimplementedRpcServer) LoadContracts(context.Context, *LoadContractsRequest) (*LoadContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadContracts not implemented")
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

func _Rpc_About_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AboutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).About(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_About_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).About(ctx, req.(*AboutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_LoadContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).LoadContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_LoadContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).LoadContract(ctx, req.(*LoadContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_LoadContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).LoadContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_LoadContracts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).LoadContracts(ctx, req.(*LoadContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eigenlayer.sidecar.v1.sidecar.Rpc",
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
			MethodName: "About",
			Handler:    _Rpc_About_Handler,
		},
		{
			MethodName: "LoadContract",
			Handler:    _Rpc_LoadContract_Handler,
		},
		{
			MethodName: "LoadContracts",
			Handler:    _Rpc_LoadContracts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eigenlayer/sidecar/v1/sidecar/rpc.proto",
}
