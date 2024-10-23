// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: eigenlayer/sidecar/v1/protocol.proto

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
	Protocol_GetRegisteredAvsForOperator_FullMethodName          = "/eigenlayer.sidecar.protocol.v1.Protocol/GetRegisteredAvsForOperator"
	Protocol_GetDelegatedStrategiesForOperator_FullMethodName    = "/eigenlayer.sidecar.protocol.v1.Protocol/GetDelegatedStrategiesForOperator"
	Protocol_GetOperatorDelegatedStakeForStrategy_FullMethodName = "/eigenlayer.sidecar.protocol.v1.Protocol/GetOperatorDelegatedStakeForStrategy"
	Protocol_GetDelegatedStakersForOperator_FullMethodName       = "/eigenlayer.sidecar.protocol.v1.Protocol/GetDelegatedStakersForOperator"
	Protocol_GetStakerShares_FullMethodName                      = "/eigenlayer.sidecar.protocol.v1.Protocol/GetStakerShares"
)

// ProtocolClient is the client API for Protocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtocolClient interface {
	// GetRegisteredAvsForOperator returns the list of registered AVs for a given operator
	// BlockHeight is optional, otherwise latest is used.
	GetRegisteredAvsForOperator(ctx context.Context, in *GetRegisteredAvsForOperatorRequest, opts ...grpc.CallOption) (*GetRegisteredAvsForOperatorResponse, error)
	// GetDelegatedStrategiesForOperator returns strategies an Operator has delegated
	GetDelegatedStrategiesForOperator(ctx context.Context, in *GetDelegatedStrategiesForOperatorRequest, opts ...grpc.CallOption) (*GetDelegatedStrategiesForOperatorResponse, error)
	// GetOperatorDelegatedStakeForStrategy returns the amount of delegated stake for a given strategy for an operator
	GetOperatorDelegatedStakeForStrategy(ctx context.Context, in *GetOperatorDelegatedStakeForStrategyRequest, opts ...grpc.CallOption) (*GetOperatorDelegatedStakeForStrategyResponse, error)
	// GetDelegatedStakersForOperator returns the list of stakers that have delegated to an operator.
	// BlockHeight is optional, otherwise latest is used.
	GetDelegatedStakersForOperator(ctx context.Context, in *GetDelegatedStakersForOperatorRequest, opts ...grpc.CallOption) (*GetDelegatedStakersForOperatorResponse, error)
	// GetStakerShares returns the shares of a staker, and optionally, the operator operator they've delegated to and
	// the AVS the operator has registered with.
	GetStakerShares(ctx context.Context, in *GetStakerSharesRequest, opts ...grpc.CallOption) (*GetStakerSharesResponse, error)
}

type protocolClient struct {
	cc grpc.ClientConnInterface
}

func NewProtocolClient(cc grpc.ClientConnInterface) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) GetRegisteredAvsForOperator(ctx context.Context, in *GetRegisteredAvsForOperatorRequest, opts ...grpc.CallOption) (*GetRegisteredAvsForOperatorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRegisteredAvsForOperatorResponse)
	err := c.cc.Invoke(ctx, Protocol_GetRegisteredAvsForOperator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) GetDelegatedStrategiesForOperator(ctx context.Context, in *GetDelegatedStrategiesForOperatorRequest, opts ...grpc.CallOption) (*GetDelegatedStrategiesForOperatorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDelegatedStrategiesForOperatorResponse)
	err := c.cc.Invoke(ctx, Protocol_GetDelegatedStrategiesForOperator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) GetOperatorDelegatedStakeForStrategy(ctx context.Context, in *GetOperatorDelegatedStakeForStrategyRequest, opts ...grpc.CallOption) (*GetOperatorDelegatedStakeForStrategyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOperatorDelegatedStakeForStrategyResponse)
	err := c.cc.Invoke(ctx, Protocol_GetOperatorDelegatedStakeForStrategy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) GetDelegatedStakersForOperator(ctx context.Context, in *GetDelegatedStakersForOperatorRequest, opts ...grpc.CallOption) (*GetDelegatedStakersForOperatorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDelegatedStakersForOperatorResponse)
	err := c.cc.Invoke(ctx, Protocol_GetDelegatedStakersForOperator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) GetStakerShares(ctx context.Context, in *GetStakerSharesRequest, opts ...grpc.CallOption) (*GetStakerSharesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStakerSharesResponse)
	err := c.cc.Invoke(ctx, Protocol_GetStakerShares_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtocolServer is the server API for Protocol service.
// All implementations should embed UnimplementedProtocolServer
// for forward compatibility.
type ProtocolServer interface {
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
}

// UnimplementedProtocolServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProtocolServer struct{}

func (UnimplementedProtocolServer) GetRegisteredAvsForOperator(context.Context, *GetRegisteredAvsForOperatorRequest) (*GetRegisteredAvsForOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredAvsForOperator not implemented")
}
func (UnimplementedProtocolServer) GetDelegatedStrategiesForOperator(context.Context, *GetDelegatedStrategiesForOperatorRequest) (*GetDelegatedStrategiesForOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegatedStrategiesForOperator not implemented")
}
func (UnimplementedProtocolServer) GetOperatorDelegatedStakeForStrategy(context.Context, *GetOperatorDelegatedStakeForStrategyRequest) (*GetOperatorDelegatedStakeForStrategyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorDelegatedStakeForStrategy not implemented")
}
func (UnimplementedProtocolServer) GetDelegatedStakersForOperator(context.Context, *GetDelegatedStakersForOperatorRequest) (*GetDelegatedStakersForOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegatedStakersForOperator not implemented")
}
func (UnimplementedProtocolServer) GetStakerShares(context.Context, *GetStakerSharesRequest) (*GetStakerSharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStakerShares not implemented")
}
func (UnimplementedProtocolServer) testEmbeddedByValue() {}

// UnsafeProtocolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtocolServer will
// result in compilation errors.
type UnsafeProtocolServer interface {
	mustEmbedUnimplementedProtocolServer()
}

func RegisterProtocolServer(s grpc.ServiceRegistrar, srv ProtocolServer) {
	// If the following call pancis, it indicates UnimplementedProtocolServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Protocol_ServiceDesc, srv)
}

func _Protocol_GetRegisteredAvsForOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisteredAvsForOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).GetRegisteredAvsForOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_GetRegisteredAvsForOperator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).GetRegisteredAvsForOperator(ctx, req.(*GetRegisteredAvsForOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_GetDelegatedStrategiesForOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegatedStrategiesForOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).GetDelegatedStrategiesForOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_GetDelegatedStrategiesForOperator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).GetDelegatedStrategiesForOperator(ctx, req.(*GetDelegatedStrategiesForOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_GetOperatorDelegatedStakeForStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorDelegatedStakeForStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).GetOperatorDelegatedStakeForStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_GetOperatorDelegatedStakeForStrategy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).GetOperatorDelegatedStakeForStrategy(ctx, req.(*GetOperatorDelegatedStakeForStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_GetDelegatedStakersForOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegatedStakersForOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).GetDelegatedStakersForOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_GetDelegatedStakersForOperator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).GetDelegatedStakersForOperator(ctx, req.(*GetDelegatedStakersForOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_GetStakerShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStakerSharesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).GetStakerShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_GetStakerShares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).GetStakerShares(ctx, req.(*GetStakerSharesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Protocol_ServiceDesc is the grpc.ServiceDesc for Protocol service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Protocol_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eigenlayer.sidecar.protocol.v1.Protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegisteredAvsForOperator",
			Handler:    _Protocol_GetRegisteredAvsForOperator_Handler,
		},
		{
			MethodName: "GetDelegatedStrategiesForOperator",
			Handler:    _Protocol_GetDelegatedStrategiesForOperator_Handler,
		},
		{
			MethodName: "GetOperatorDelegatedStakeForStrategy",
			Handler:    _Protocol_GetOperatorDelegatedStakeForStrategy_Handler,
		},
		{
			MethodName: "GetDelegatedStakersForOperator",
			Handler:    _Protocol_GetDelegatedStakersForOperator_Handler,
		},
		{
			MethodName: "GetStakerShares",
			Handler:    _Protocol_GetStakerShares_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eigenlayer/sidecar/v1/protocol.proto",
}