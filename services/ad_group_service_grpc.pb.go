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

// AdGroupServiceClient is the client API for AdGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupServiceClient interface {
	// Returns the requested ad group in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroup(ctx context.Context, in *GetAdGroupRequest, opts ...grpc.CallOption) (*resources.AdGroup, error)
	// Creates, updates, or removes ad groups. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AdGroupError]()
	//   [AdxError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BiddingError]()
	//   [BiddingStrategyError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MultiplierError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SettingError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateAdGroups(ctx context.Context, in *MutateAdGroupsRequest, opts ...grpc.CallOption) (*MutateAdGroupsResponse, error)
}

type adGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupServiceClient(cc grpc.ClientConnInterface) AdGroupServiceClient {
	return &adGroupServiceClient{cc}
}

func (c *adGroupServiceClient) GetAdGroup(ctx context.Context, in *GetAdGroupRequest, opts ...grpc.CallOption) (*resources.AdGroup, error) {
	out := new(resources.AdGroup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.AdGroupService/GetAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupServiceClient) MutateAdGroups(ctx context.Context, in *MutateAdGroupsRequest, opts ...grpc.CallOption) (*MutateAdGroupsResponse, error) {
	out := new(MutateAdGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.AdGroupService/MutateAdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupServiceServer is the server API for AdGroupService service.
// All implementations must embed UnimplementedAdGroupServiceServer
// for forward compatibility
type AdGroupServiceServer interface {
	// Returns the requested ad group in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroup(context.Context, *GetAdGroupRequest) (*resources.AdGroup, error)
	// Creates, updates, or removes ad groups. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AdGroupError]()
	//   [AdxError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BiddingError]()
	//   [BiddingStrategyError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MultiplierError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SettingError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateAdGroups(context.Context, *MutateAdGroupsRequest) (*MutateAdGroupsResponse, error)
	mustEmbedUnimplementedAdGroupServiceServer()
}

// UnimplementedAdGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupServiceServer struct {
}

func (UnimplementedAdGroupServiceServer) GetAdGroup(context.Context, *GetAdGroupRequest) (*resources.AdGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroup not implemented")
}
func (UnimplementedAdGroupServiceServer) MutateAdGroups(context.Context, *MutateAdGroupsRequest) (*MutateAdGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroups not implemented")
}
func (UnimplementedAdGroupServiceServer) mustEmbedUnimplementedAdGroupServiceServer() {}

// UnsafeAdGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupServiceServer will
// result in compilation errors.
type UnsafeAdGroupServiceServer interface {
	mustEmbedUnimplementedAdGroupServiceServer()
}

func RegisterAdGroupServiceServer(s grpc.ServiceRegistrar, srv AdGroupServiceServer) {
	s.RegisterService(&AdGroupService_ServiceDesc, srv)
}

func _AdGroupService_GetAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupServiceServer).GetAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.AdGroupService/GetAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupServiceServer).GetAdGroup(ctx, req.(*GetAdGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupService_MutateAdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupServiceServer).MutateAdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.AdGroupService/MutateAdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupServiceServer).MutateAdGroups(ctx, req.(*MutateAdGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupService_ServiceDesc is the grpc.ServiceDesc for AdGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.AdGroupService",
	HandlerType: (*AdGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroup",
			Handler:    _AdGroupService_GetAdGroup_Handler,
		},
		{
			MethodName: "MutateAdGroups",
			Handler:    _AdGroupService_MutateAdGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/ad_group_service.proto",
}
