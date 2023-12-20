// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: github.com/rancher/opni/pkg/metrics/collector/remote.proto

package collector

import (
	context "context"
	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RemoteCollector_GetMetrics_FullMethodName = "/collector.RemoteCollector/GetMetrics"
)

// RemoteCollectorClient is the client API for RemoteCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteCollectorClient interface {
	GetMetrics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.MetricsData, error)
}

type remoteCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteCollectorClient(cc grpc.ClientConnInterface) RemoteCollectorClient {
	return &remoteCollectorClient{cc}
}

func (c *remoteCollectorClient) GetMetrics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.MetricsData, error) {
	out := new(v1.MetricsData)
	err := c.cc.Invoke(ctx, RemoteCollector_GetMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteCollectorServer is the server API for RemoteCollector service.
// All implementations should embed UnimplementedRemoteCollectorServer
// for forward compatibility
type RemoteCollectorServer interface {
	GetMetrics(context.Context, *emptypb.Empty) (*v1.MetricsData, error)
}

// UnimplementedRemoteCollectorServer should be embedded to have forward compatible implementations.
type UnimplementedRemoteCollectorServer struct {
}

func (UnimplementedRemoteCollectorServer) GetMetrics(context.Context, *emptypb.Empty) (*v1.MetricsData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

// UnsafeRemoteCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteCollectorServer will
// result in compilation errors.
type UnsafeRemoteCollectorServer interface {
	mustEmbedUnimplementedRemoteCollectorServer()
}

func RegisterRemoteCollectorServer(s grpc.ServiceRegistrar, srv RemoteCollectorServer) {
	s.RegisterService(&RemoteCollector_ServiceDesc, srv)
}

func _RemoteCollector_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCollectorServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteCollector_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCollectorServer).GetMetrics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteCollector_ServiceDesc is the grpc.ServiceDesc for RemoteCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collector.RemoteCollector",
	HandlerType: (*RemoteCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _RemoteCollector_GetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/metrics/collector/remote.proto",
}
