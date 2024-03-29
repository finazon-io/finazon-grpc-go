// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: sec.proto

package finazon

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
	SecService_GetFilings_FullMethodName = "/finazon.SecService/GetFilings"
)

// SecServiceClient is the client API for SecService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecServiceClient interface {
	// Real-time and historical access to all forms, filings, and exhibits directly from the SEC's EDGAR system.
	GetFilings(ctx context.Context, in *GetSecFillingsRequest, opts ...grpc.CallOption) (*GetSecFillingsResponse, error)
}

type secServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecServiceClient(cc grpc.ClientConnInterface) SecServiceClient {
	return &secServiceClient{cc}
}

func (c *secServiceClient) GetFilings(ctx context.Context, in *GetSecFillingsRequest, opts ...grpc.CallOption) (*GetSecFillingsResponse, error) {
	out := new(GetSecFillingsResponse)
	err := c.cc.Invoke(ctx, SecService_GetFilings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecServiceServer is the server API for SecService service.
// All implementations must embed UnimplementedSecServiceServer
// for forward compatibility
type SecServiceServer interface {
	// Real-time and historical access to all forms, filings, and exhibits directly from the SEC's EDGAR system.
	GetFilings(context.Context, *GetSecFillingsRequest) (*GetSecFillingsResponse, error)
	mustEmbedUnimplementedSecServiceServer()
}

// UnimplementedSecServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecServiceServer struct {
}

func (UnimplementedSecServiceServer) GetFilings(context.Context, *GetSecFillingsRequest) (*GetSecFillingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilings not implemented")
}
func (UnimplementedSecServiceServer) mustEmbedUnimplementedSecServiceServer() {}

// UnsafeSecServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecServiceServer will
// result in compilation errors.
type UnsafeSecServiceServer interface {
	mustEmbedUnimplementedSecServiceServer()
}

func RegisterSecServiceServer(s grpc.ServiceRegistrar, srv SecServiceServer) {
	s.RegisterService(&SecService_ServiceDesc, srv)
}

func _SecService_GetFilings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecFillingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecServiceServer).GetFilings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecService_GetFilings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecServiceServer).GetFilings(ctx, req.(*GetSecFillingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecService_ServiceDesc is the grpc.ServiceDesc for SecService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finazon.SecService",
	HandlerType: (*SecServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFilings",
			Handler:    _SecService_GetFilings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sec.proto",
}
