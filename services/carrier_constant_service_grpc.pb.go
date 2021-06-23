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

// CarrierConstantServiceClient is the client API for CarrierConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarrierConstantServiceClient interface {
	// Returns the requested carrier constant in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error)
}

type carrierConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarrierConstantServiceClient(cc grpc.ClientConnInterface) CarrierConstantServiceClient {
	return &carrierConstantServiceClient{cc}
}

func (c *carrierConstantServiceClient) GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error) {
	out := new(resources.CarrierConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.CarrierConstantService/GetCarrierConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarrierConstantServiceServer is the server API for CarrierConstantService service.
// All implementations must embed UnimplementedCarrierConstantServiceServer
// for forward compatibility
type CarrierConstantServiceServer interface {
	// Returns the requested carrier constant in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCarrierConstant(context.Context, *GetCarrierConstantRequest) (*resources.CarrierConstant, error)
	mustEmbedUnimplementedCarrierConstantServiceServer()
}

// UnimplementedCarrierConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarrierConstantServiceServer struct {
}

func (UnimplementedCarrierConstantServiceServer) GetCarrierConstant(context.Context, *GetCarrierConstantRequest) (*resources.CarrierConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarrierConstant not implemented")
}
func (UnimplementedCarrierConstantServiceServer) mustEmbedUnimplementedCarrierConstantServiceServer() {
}

// UnsafeCarrierConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarrierConstantServiceServer will
// result in compilation errors.
type UnsafeCarrierConstantServiceServer interface {
	mustEmbedUnimplementedCarrierConstantServiceServer()
}

func RegisterCarrierConstantServiceServer(s grpc.ServiceRegistrar, srv CarrierConstantServiceServer) {
	s.RegisterService(&CarrierConstantService_ServiceDesc, srv)
}

func _CarrierConstantService_GetCarrierConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarrierConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.CarrierConstantService/GetCarrierConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, req.(*GetCarrierConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CarrierConstantService_ServiceDesc is the grpc.ServiceDesc for CarrierConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarrierConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.CarrierConstantService",
	HandlerType: (*CarrierConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarrierConstant",
			Handler:    _CarrierConstantService_GetCarrierConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/carrier_constant_service.proto",
}
