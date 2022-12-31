// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proxy/proxy.proto

package proxy

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

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyServiceClient interface {
	// stream latest N blocks to client
	GetLatestBlocks(ctx context.Context, in *GetLatestBlocksRequest, opts ...grpc.CallOption) (ProxyService_GetLatestBlocksClient, error)
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) GetLatestBlocks(ctx context.Context, in *GetLatestBlocksRequest, opts ...grpc.CallOption) (ProxyService_GetLatestBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[0], "/proxy.ProxyService/GetLatestBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServiceGetLatestBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProxyService_GetLatestBlocksClient interface {
	Recv() (*GetLatestBlocksResponse, error)
	grpc.ClientStream
}

type proxyServiceGetLatestBlocksClient struct {
	grpc.ClientStream
}

func (x *proxyServiceGetLatestBlocksClient) Recv() (*GetLatestBlocksResponse, error) {
	m := new(GetLatestBlocksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServiceServer is the server API for ProxyService service.
// All implementations should embed UnimplementedProxyServiceServer
// for forward compatibility
type ProxyServiceServer interface {
	// stream latest N blocks to client
	GetLatestBlocks(*GetLatestBlocksRequest, ProxyService_GetLatestBlocksServer) error
}

// UnimplementedProxyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProxyServiceServer struct {
}

func (UnimplementedProxyServiceServer) GetLatestBlocks(*GetLatestBlocksRequest, ProxyService_GetLatestBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLatestBlocks not implemented")
}

// UnsafeProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServiceServer will
// result in compilation errors.
type UnsafeProxyServiceServer interface {
	mustEmbedUnimplementedProxyServiceServer()
}

func RegisterProxyServiceServer(s grpc.ServiceRegistrar, srv ProxyServiceServer) {
	s.RegisterService(&ProxyService_ServiceDesc, srv)
}

func _ProxyService_GetLatestBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLatestBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServiceServer).GetLatestBlocks(m, &proxyServiceGetLatestBlocksServer{stream})
}

type ProxyService_GetLatestBlocksServer interface {
	Send(*GetLatestBlocksResponse) error
	grpc.ServerStream
}

type proxyServiceGetLatestBlocksServer struct {
	grpc.ServerStream
}

func (x *proxyServiceGetLatestBlocksServer) Send(m *GetLatestBlocksResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ProxyService_ServiceDesc is the grpc.ServiceDesc for ProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLatestBlocks",
			Handler:       _ProxyService_GetLatestBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proxy/proxy.proto",
}
