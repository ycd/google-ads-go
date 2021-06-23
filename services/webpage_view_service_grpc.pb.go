// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/ycd/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WebpageViewServiceClient is the client API for WebpageViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebpageViewServiceClient interface {
	// Returns the requested webpage view in full detail.
	GetWebpageView(ctx context.Context, in *GetWebpageViewRequest, opts ...grpc.CallOption) (*resources.WebpageView, error)
}

type webpageViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebpageViewServiceClient(cc grpc.ClientConnInterface) WebpageViewServiceClient {
	return &webpageViewServiceClient{cc}
}

func (c *webpageViewServiceClient) GetWebpageView(ctx context.Context, in *GetWebpageViewRequest, opts ...grpc.CallOption) (*resources.WebpageView, error) {
	out := new(resources.WebpageView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.WebpageViewService/GetWebpageView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebpageViewServiceServer is the server API for WebpageViewService service.
// All implementations must embed UnimplementedWebpageViewServiceServer
// for forward compatibility
type WebpageViewServiceServer interface {
	// Returns the requested webpage view in full detail.
	GetWebpageView(context.Context, *GetWebpageViewRequest) (*resources.WebpageView, error)
	mustEmbedUnimplementedWebpageViewServiceServer()
}

// UnimplementedWebpageViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebpageViewServiceServer struct {
}

func (UnimplementedWebpageViewServiceServer) GetWebpageView(context.Context, *GetWebpageViewRequest) (*resources.WebpageView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebpageView not implemented")
}
func (UnimplementedWebpageViewServiceServer) mustEmbedUnimplementedWebpageViewServiceServer() {}

// UnsafeWebpageViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebpageViewServiceServer will
// result in compilation errors.
type UnsafeWebpageViewServiceServer interface {
	mustEmbedUnimplementedWebpageViewServiceServer()
}

func RegisterWebpageViewServiceServer(s grpc.ServiceRegistrar, srv WebpageViewServiceServer) {
	s.RegisterService(&WebpageViewService_ServiceDesc, srv)
}

func _WebpageViewService_GetWebpageView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWebpageViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebpageViewServiceServer).GetWebpageView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.WebpageViewService/GetWebpageView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebpageViewServiceServer).GetWebpageView(ctx, req.(*GetWebpageViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebpageViewService_ServiceDesc is the grpc.ServiceDesc for WebpageViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebpageViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.WebpageViewService",
	HandlerType: (*WebpageViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWebpageView",
			Handler:    _WebpageViewService_GetWebpageView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/webpage_view_service.proto",
}
