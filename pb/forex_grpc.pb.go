// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: forex.proto

package finazon_grpc_go

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
	ForexService_GetTimeSeries_FullMethodName = "/finazon.ForexService/GetTimeSeries"
)

// ForexServiceClient is the client API for ForexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForexServiceClient interface {
	// Get time series data for any given ticker
	GetTimeSeries(ctx context.Context, in *GetForexTimeSeriesRequest, opts ...grpc.CallOption) (*GetForexTimeSeriesResponse, error)
}

type forexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewForexServiceClient(cc grpc.ClientConnInterface) ForexServiceClient {
	return &forexServiceClient{cc}
}

func (c *forexServiceClient) GetTimeSeries(ctx context.Context, in *GetForexTimeSeriesRequest, opts ...grpc.CallOption) (*GetForexTimeSeriesResponse, error) {
	out := new(GetForexTimeSeriesResponse)
	err := c.cc.Invoke(ctx, ForexService_GetTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForexServiceServer is the server API for ForexService service.
// All implementations must embed UnimplementedForexServiceServer
// for forward compatibility
type ForexServiceServer interface {
	// Get time series data for any given ticker
	GetTimeSeries(context.Context, *GetForexTimeSeriesRequest) (*GetForexTimeSeriesResponse, error)
	mustEmbedUnimplementedForexServiceServer()
}

// UnimplementedForexServiceServer must be embedded to have forward compatible implementations.
type UnimplementedForexServiceServer struct {
}

func (UnimplementedForexServiceServer) GetTimeSeries(context.Context, *GetForexTimeSeriesRequest) (*GetForexTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeries not implemented")
}
func (UnimplementedForexServiceServer) mustEmbedUnimplementedForexServiceServer() {}

// UnsafeForexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForexServiceServer will
// result in compilation errors.
type UnsafeForexServiceServer interface {
	mustEmbedUnimplementedForexServiceServer()
}

func RegisterForexServiceServer(s grpc.ServiceRegistrar, srv ForexServiceServer) {
	s.RegisterService(&ForexService_ServiceDesc, srv)
}

func _ForexService_GetTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForexTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForexServiceServer).GetTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForexService_GetTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForexServiceServer).GetTimeSeries(ctx, req.(*GetForexTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ForexService_ServiceDesc is the grpc.ServiceDesc for ForexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ForexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finazon.ForexService",
	HandlerType: (*ForexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeSeries",
			Handler:    _ForexService_GetTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forex.proto",
}
