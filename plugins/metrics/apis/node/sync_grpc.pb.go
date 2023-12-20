// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: github.com/rancher/opni/plugins/metrics/apis/node/sync.proto

package node

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

const (
	NodeMetricsCapability_Sync_FullMethodName = "/node.metrics.NodeMetricsCapability/Sync"
)

// NodeMetricsCapabilityClient is the client API for NodeMetricsCapability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeMetricsCapabilityClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type nodeMetricsCapabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeMetricsCapabilityClient(cc grpc.ClientConnInterface) NodeMetricsCapabilityClient {
	return &nodeMetricsCapabilityClient{cc}
}

func (c *nodeMetricsCapabilityClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, NodeMetricsCapability_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeMetricsCapabilityServer is the server API for NodeMetricsCapability service.
// All implementations should embed UnimplementedNodeMetricsCapabilityServer
// for forward compatibility
type NodeMetricsCapabilityServer interface {
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
}

// UnimplementedNodeMetricsCapabilityServer should be embedded to have forward compatible implementations.
type UnimplementedNodeMetricsCapabilityServer struct {
}

func (UnimplementedNodeMetricsCapabilityServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

// UnsafeNodeMetricsCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeMetricsCapabilityServer will
// result in compilation errors.
type UnsafeNodeMetricsCapabilityServer interface {
	mustEmbedUnimplementedNodeMetricsCapabilityServer()
}

func RegisterNodeMetricsCapabilityServer(s grpc.ServiceRegistrar, srv NodeMetricsCapabilityServer) {
	s.RegisterService(&NodeMetricsCapability_ServiceDesc, srv)
}

func _NodeMetricsCapability_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeMetricsCapabilityServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeMetricsCapability_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeMetricsCapabilityServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeMetricsCapability_ServiceDesc is the grpc.ServiceDesc for NodeMetricsCapability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeMetricsCapability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.metrics.NodeMetricsCapability",
	HandlerType: (*NodeMetricsCapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _NodeMetricsCapability_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/apis/node/sync.proto",
}
