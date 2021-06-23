// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/ercling/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClickViewServiceClient is the client API for ClickViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClickViewServiceClient interface {
	// Returns the requested click view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error)
}

type clickViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClickViewServiceClient(cc grpc.ClientConnInterface) ClickViewServiceClient {
	return &clickViewServiceClient{cc}
}

func (c *clickViewServiceClient) GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error) {
	out := new(resources.ClickView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.ClickViewService/GetClickView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickViewServiceServer is the server API for ClickViewService service.
// All implementations must embed UnimplementedClickViewServiceServer
// for forward compatibility
type ClickViewServiceServer interface {
	// Returns the requested click view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error)
	mustEmbedUnimplementedClickViewServiceServer()
}

// UnimplementedClickViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClickViewServiceServer struct {
}

func (UnimplementedClickViewServiceServer) GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClickView not implemented")
}
func (UnimplementedClickViewServiceServer) mustEmbedUnimplementedClickViewServiceServer() {}

// UnsafeClickViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClickViewServiceServer will
// result in compilation errors.
type UnsafeClickViewServiceServer interface {
	mustEmbedUnimplementedClickViewServiceServer()
}

func RegisterClickViewServiceServer(s grpc.ServiceRegistrar, srv ClickViewServiceServer) {
	s.RegisterService(&ClickViewService_ServiceDesc, srv)
}

func _ClickViewService_GetClickView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClickViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickViewServiceServer).GetClickView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.ClickViewService/GetClickView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickViewServiceServer).GetClickView(ctx, req.(*GetClickViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClickViewService_ServiceDesc is the grpc.ServiceDesc for ClickViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClickViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.ClickViewService",
	HandlerType: (*ClickViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClickView",
			Handler:    _ClickViewService_GetClickView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/click_view_service.proto",
}
