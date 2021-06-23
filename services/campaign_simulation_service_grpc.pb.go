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

// CampaignSimulationServiceClient is the client API for CampaignSimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignSimulationServiceClient interface {
	// Returns the requested campaign simulation in full detail.
	GetCampaignSimulation(ctx context.Context, in *GetCampaignSimulationRequest, opts ...grpc.CallOption) (*resources.CampaignSimulation, error)
}

type campaignSimulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignSimulationServiceClient(cc grpc.ClientConnInterface) CampaignSimulationServiceClient {
	return &campaignSimulationServiceClient{cc}
}

func (c *campaignSimulationServiceClient) GetCampaignSimulation(ctx context.Context, in *GetCampaignSimulationRequest, opts ...grpc.CallOption) (*resources.CampaignSimulation, error) {
	out := new(resources.CampaignSimulation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.CampaignSimulationService/GetCampaignSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignSimulationServiceServer is the server API for CampaignSimulationService service.
// All implementations must embed UnimplementedCampaignSimulationServiceServer
// for forward compatibility
type CampaignSimulationServiceServer interface {
	// Returns the requested campaign simulation in full detail.
	GetCampaignSimulation(context.Context, *GetCampaignSimulationRequest) (*resources.CampaignSimulation, error)
	mustEmbedUnimplementedCampaignSimulationServiceServer()
}

// UnimplementedCampaignSimulationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignSimulationServiceServer struct {
}

func (UnimplementedCampaignSimulationServiceServer) GetCampaignSimulation(context.Context, *GetCampaignSimulationRequest) (*resources.CampaignSimulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignSimulation not implemented")
}
func (UnimplementedCampaignSimulationServiceServer) mustEmbedUnimplementedCampaignSimulationServiceServer() {
}

// UnsafeCampaignSimulationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignSimulationServiceServer will
// result in compilation errors.
type UnsafeCampaignSimulationServiceServer interface {
	mustEmbedUnimplementedCampaignSimulationServiceServer()
}

func RegisterCampaignSimulationServiceServer(s grpc.ServiceRegistrar, srv CampaignSimulationServiceServer) {
	s.RegisterService(&CampaignSimulationService_ServiceDesc, srv)
}

func _CampaignSimulationService_GetCampaignSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignSimulationServiceServer).GetCampaignSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.CampaignSimulationService/GetCampaignSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignSimulationServiceServer).GetCampaignSimulation(ctx, req.(*GetCampaignSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignSimulationService_ServiceDesc is the grpc.ServiceDesc for CampaignSimulationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignSimulationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.CampaignSimulationService",
	HandlerType: (*CampaignSimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignSimulation",
			Handler:    _CampaignSimulationService_GetCampaignSimulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/campaign_simulation_service.proto",
}
