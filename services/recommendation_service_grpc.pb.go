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

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendationServiceClient interface {
	// Returns the requested recommendation in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetRecommendation(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (*resources.Recommendation, error)
	// Applies given recommendations with corresponding apply parameters.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RecommendationError]()
	//   [RequestError]()
	//   [UrlFieldError]()
	ApplyRecommendation(ctx context.Context, in *ApplyRecommendationRequest, opts ...grpc.CallOption) (*ApplyRecommendationResponse, error)
	// Dismisses given recommendations.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RecommendationError]()
	//   [RequestError]()
	DismissRecommendation(ctx context.Context, in *DismissRecommendationRequest, opts ...grpc.CallOption) (*DismissRecommendationResponse, error)
}

type recommendationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationServiceClient(cc grpc.ClientConnInterface) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) GetRecommendation(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (*resources.Recommendation, error) {
	out := new(resources.Recommendation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.RecommendationService/GetRecommendation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceClient) ApplyRecommendation(ctx context.Context, in *ApplyRecommendationRequest, opts ...grpc.CallOption) (*ApplyRecommendationResponse, error) {
	out := new(ApplyRecommendationResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.RecommendationService/ApplyRecommendation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceClient) DismissRecommendation(ctx context.Context, in *DismissRecommendationRequest, opts ...grpc.CallOption) (*DismissRecommendationResponse, error) {
	out := new(DismissRecommendationResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.RecommendationService/DismissRecommendation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
// All implementations must embed UnimplementedRecommendationServiceServer
// for forward compatibility
type RecommendationServiceServer interface {
	// Returns the requested recommendation in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetRecommendation(context.Context, *GetRecommendationRequest) (*resources.Recommendation, error)
	// Applies given recommendations with corresponding apply parameters.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RecommendationError]()
	//   [RequestError]()
	//   [UrlFieldError]()
	ApplyRecommendation(context.Context, *ApplyRecommendationRequest) (*ApplyRecommendationResponse, error)
	// Dismisses given recommendations.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RecommendationError]()
	//   [RequestError]()
	DismissRecommendation(context.Context, *DismissRecommendationRequest) (*DismissRecommendationResponse, error)
	mustEmbedUnimplementedRecommendationServiceServer()
}

// UnimplementedRecommendationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecommendationServiceServer struct {
}

func (UnimplementedRecommendationServiceServer) GetRecommendation(context.Context, *GetRecommendationRequest) (*resources.Recommendation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendation not implemented")
}
func (UnimplementedRecommendationServiceServer) ApplyRecommendation(context.Context, *ApplyRecommendationRequest) (*ApplyRecommendationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRecommendation not implemented")
}
func (UnimplementedRecommendationServiceServer) DismissRecommendation(context.Context, *DismissRecommendationRequest) (*DismissRecommendationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DismissRecommendation not implemented")
}
func (UnimplementedRecommendationServiceServer) mustEmbedUnimplementedRecommendationServiceServer() {}

// UnsafeRecommendationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationServiceServer will
// result in compilation errors.
type UnsafeRecommendationServiceServer interface {
	mustEmbedUnimplementedRecommendationServiceServer()
}

func RegisterRecommendationServiceServer(s grpc.ServiceRegistrar, srv RecommendationServiceServer) {
	s.RegisterService(&RecommendationService_ServiceDesc, srv)
}

func _RecommendationService_GetRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).GetRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.RecommendationService/GetRecommendation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).GetRecommendation(ctx, req.(*GetRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationService_ApplyRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).ApplyRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.RecommendationService/ApplyRecommendation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).ApplyRecommendation(ctx, req.(*ApplyRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationService_DismissRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DismissRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).DismissRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.RecommendationService/DismissRecommendation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).DismissRecommendation(ctx, req.(*DismissRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendationService_ServiceDesc is the grpc.ServiceDesc for RecommendationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecommendation",
			Handler:    _RecommendationService_GetRecommendation_Handler,
		},
		{
			MethodName: "ApplyRecommendation",
			Handler:    _RecommendationService_ApplyRecommendation_Handler,
		},
		{
			MethodName: "DismissRecommendation",
			Handler:    _RecommendationService_DismissRecommendation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/recommendation_service.proto",
}
