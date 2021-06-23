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

// AccountBudgetServiceClient is the client API for AccountBudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountBudgetServiceClient interface {
	// Returns an account-level budget in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAccountBudget(ctx context.Context, in *GetAccountBudgetRequest, opts ...grpc.CallOption) (*resources.AccountBudget, error)
}

type accountBudgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountBudgetServiceClient(cc grpc.ClientConnInterface) AccountBudgetServiceClient {
	return &accountBudgetServiceClient{cc}
}

func (c *accountBudgetServiceClient) GetAccountBudget(ctx context.Context, in *GetAccountBudgetRequest, opts ...grpc.CallOption) (*resources.AccountBudget, error) {
	out := new(resources.AccountBudget)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.AccountBudgetService/GetAccountBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetServiceServer is the server API for AccountBudgetService service.
// All implementations must embed UnimplementedAccountBudgetServiceServer
// for forward compatibility
type AccountBudgetServiceServer interface {
	// Returns an account-level budget in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAccountBudget(context.Context, *GetAccountBudgetRequest) (*resources.AccountBudget, error)
	mustEmbedUnimplementedAccountBudgetServiceServer()
}

// UnimplementedAccountBudgetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountBudgetServiceServer struct {
}

func (UnimplementedAccountBudgetServiceServer) GetAccountBudget(context.Context, *GetAccountBudgetRequest) (*resources.AccountBudget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBudget not implemented")
}
func (UnimplementedAccountBudgetServiceServer) mustEmbedUnimplementedAccountBudgetServiceServer() {}

// UnsafeAccountBudgetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountBudgetServiceServer will
// result in compilation errors.
type UnsafeAccountBudgetServiceServer interface {
	mustEmbedUnimplementedAccountBudgetServiceServer()
}

func RegisterAccountBudgetServiceServer(s grpc.ServiceRegistrar, srv AccountBudgetServiceServer) {
	s.RegisterService(&AccountBudgetService_ServiceDesc, srv)
}

func _AccountBudgetService_GetAccountBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetServiceServer).GetAccountBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.AccountBudgetService/GetAccountBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetServiceServer).GetAccountBudget(ctx, req.(*GetAccountBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountBudgetService_ServiceDesc is the grpc.ServiceDesc for AccountBudgetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountBudgetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.AccountBudgetService",
	HandlerType: (*AccountBudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountBudget",
			Handler:    _AccountBudgetService_GetAccountBudget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/account_budget_service.proto",
}
