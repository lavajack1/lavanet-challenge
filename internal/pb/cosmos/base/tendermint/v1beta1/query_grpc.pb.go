// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cosmos/base/tendermint/v1beta1/query.proto

package tmservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// GetNodeInfo queries the current node info.
	GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error)
	// GetSyncing queries node syncing.
	GetSyncing(ctx context.Context, in *GetSyncingRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error)
	// GetLatestBlock returns the latest block.
	GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error)
	// GetBlockByHeight queries block for given height.
	GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error)
	// GetLatestValidatorSet queries latest validator-set.
	GetLatestValidatorSet(ctx context.Context, in *GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error)
	// GetValidatorSetByHeight queries validator-set at a given height.
	GetValidatorSetByHeight(ctx context.Context, in *GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error)
	// ABCIQuery defines a query handler that supports ABCI queries directly to the
	// application, bypassing Tendermint completely. The ABCI query must contain
	// a valid and supported path, including app, custom, p2p, and store.
	//
	// Since: cosmos-sdk 0.46
	ABCIQuery(ctx context.Context, in *ABCIQueryRequest, opts ...grpc.CallOption) (*ABCIQueryResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error) {
	out := new(GetNodeInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetSyncing(ctx context.Context, in *GetSyncingRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error) {
	out := new(GetSyncingResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetSyncing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error) {
	out := new(GetLatestBlockResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error) {
	out := new(GetBlockByHeightResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestValidatorSet(ctx context.Context, in *GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error) {
	out := new(GetLatestValidatorSetResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetValidatorSetByHeight(ctx context.Context, in *GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error) {
	out := new(GetValidatorSetByHeightResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ABCIQuery(ctx context.Context, in *ABCIQueryRequest, opts ...grpc.CallOption) (*ABCIQueryResponse, error) {
	out := new(ABCIQueryResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/ABCIQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations should embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// GetNodeInfo queries the current node info.
	GetNodeInfo(context.Context, *GetNodeInfoRequest) (*GetNodeInfoResponse, error)
	// GetSyncing queries node syncing.
	GetSyncing(context.Context, *GetSyncingRequest) (*GetSyncingResponse, error)
	// GetLatestBlock returns the latest block.
	GetLatestBlock(context.Context, *GetLatestBlockRequest) (*GetLatestBlockResponse, error)
	// GetBlockByHeight queries block for given height.
	GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error)
	// GetLatestValidatorSet queries latest validator-set.
	GetLatestValidatorSet(context.Context, *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error)
	// GetValidatorSetByHeight queries validator-set at a given height.
	GetValidatorSetByHeight(context.Context, *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error)
	// ABCIQuery defines a query handler that supports ABCI queries directly to the
	// application, bypassing Tendermint completely. The ABCI query must contain
	// a valid and supported path, including app, custom, p2p, and store.
	//
	// Since: cosmos-sdk 0.46
	ABCIQuery(context.Context, *ABCIQueryRequest) (*ABCIQueryResponse, error)
}

// UnimplementedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetNodeInfo(context.Context, *GetNodeInfoRequest) (*GetNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (UnimplementedServiceServer) GetSyncing(context.Context, *GetSyncingRequest) (*GetSyncingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncing not implemented")
}
func (UnimplementedServiceServer) GetLatestBlock(context.Context, *GetLatestBlockRequest) (*GetLatestBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (UnimplementedServiceServer) GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHeight not implemented")
}
func (UnimplementedServiceServer) GetLatestValidatorSet(context.Context, *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValidatorSet not implemented")
}
func (UnimplementedServiceServer) GetValidatorSetByHeight(context.Context, *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorSetByHeight not implemented")
}
func (UnimplementedServiceServer) ABCIQuery(context.Context, *ABCIQueryRequest) (*ABCIQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ABCIQuery not implemented")
}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetNodeInfo(ctx, req.(*GetNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetSyncing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSyncingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetSyncing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetSyncing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetSyncing(ctx, req.(*GetSyncingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestBlock(ctx, req.(*GetLatestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBlockByHeight(ctx, req.(*GetBlockByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestValidatorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, req.(*GetLatestValidatorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetValidatorSetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorSetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, req.(*GetValidatorSetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ABCIQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABCIQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ABCIQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/ABCIQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ABCIQuery(ctx, req.(*ABCIQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.tendermint.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeInfo",
			Handler:    _Service_GetNodeInfo_Handler,
		},
		{
			MethodName: "GetSyncing",
			Handler:    _Service_GetSyncing_Handler,
		},
		{
			MethodName: "GetLatestBlock",
			Handler:    _Service_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlockByHeight",
			Handler:    _Service_GetBlockByHeight_Handler,
		},
		{
			MethodName: "GetLatestValidatorSet",
			Handler:    _Service_GetLatestValidatorSet_Handler,
		},
		{
			MethodName: "GetValidatorSetByHeight",
			Handler:    _Service_GetValidatorSetByHeight_Handler,
		},
		{
			MethodName: "ABCIQuery",
			Handler:    _Service_ABCIQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/tendermint/v1beta1/query.proto",
}
