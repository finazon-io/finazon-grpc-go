// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api_usage.proto

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
	ApiUsageService_GetApiUsage_FullMethodName = "/finazon.ApiUsageService/GetApiUsage"
)

// ApiUsageServiceClient is the client API for ApiUsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiUsageServiceClient interface {
	// Get a list of products with quota limit/usage
	GetApiUsage(ctx context.Context, in *GetApiUsageRequest, opts ...grpc.CallOption) (*GetApiUsageResponse, error)
}

type apiUsageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiUsageServiceClient(cc grpc.ClientConnInterface) ApiUsageServiceClient {
	return &apiUsageServiceClient{cc}
}

func (c *apiUsageServiceClient) GetApiUsage(ctx context.Context, in *GetApiUsageRequest, opts ...grpc.CallOption) (*GetApiUsageResponse, error) {
	out := new(GetApiUsageResponse)
	err := c.cc.Invoke(ctx, ApiUsageService_GetApiUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiUsageServiceServer is the server API for ApiUsageService service.
// All implementations must embed UnimplementedApiUsageServiceServer
// for forward compatibility
type ApiUsageServiceServer interface {
	// Get a list of products with quota limit/usage
	GetApiUsage(context.Context, *GetApiUsageRequest) (*GetApiUsageResponse, error)
	mustEmbedUnimplementedApiUsageServiceServer()
}

// UnimplementedApiUsageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiUsageServiceServer struct {
}

func (UnimplementedApiUsageServiceServer) GetApiUsage(context.Context, *GetApiUsageRequest) (*GetApiUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiUsage not implemented")
}
func (UnimplementedApiUsageServiceServer) mustEmbedUnimplementedApiUsageServiceServer() {}

// UnsafeApiUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiUsageServiceServer will
// result in compilation errors.
type UnsafeApiUsageServiceServer interface {
	mustEmbedUnimplementedApiUsageServiceServer()
}

func RegisterApiUsageServiceServer(s grpc.ServiceRegistrar, srv ApiUsageServiceServer) {
	s.RegisterService(&ApiUsageService_ServiceDesc, srv)
}

func _ApiUsageService_GetApiUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiUsageServiceServer).GetApiUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiUsageService_GetApiUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiUsageServiceServer).GetApiUsage(ctx, req.(*GetApiUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiUsageService_ServiceDesc is the grpc.ServiceDesc for ApiUsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiUsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finazon.ApiUsageService",
	HandlerType: (*ApiUsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApiUsage",
			Handler:    _ApiUsageService_GetApiUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_usage.proto",
}