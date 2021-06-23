// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/search_term_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/ycd/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [SearchTermViewService.GetSearchTermView][google.ads.googleads.v0.services.SearchTermViewService.GetSearchTermView].
type GetSearchTermViewRequest struct {
	// The resource name of the search term view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSearchTermViewRequest) Reset()         { *m = GetSearchTermViewRequest{} }
func (m *GetSearchTermViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetSearchTermViewRequest) ProtoMessage()    {}
func (*GetSearchTermViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2606b8baa50c1e2, []int{0}
}

func (m *GetSearchTermViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSearchTermViewRequest.Unmarshal(m, b)
}
func (m *GetSearchTermViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSearchTermViewRequest.Marshal(b, m, deterministic)
}
func (m *GetSearchTermViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSearchTermViewRequest.Merge(m, src)
}
func (m *GetSearchTermViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetSearchTermViewRequest.Size(m)
}
func (m *GetSearchTermViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSearchTermViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSearchTermViewRequest proto.InternalMessageInfo

func (m *GetSearchTermViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSearchTermViewRequest)(nil), "google.ads.googleads.v0.services.GetSearchTermViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/search_term_view_service.proto", fileDescriptor_b2606b8baa50c1e2)
}

var fileDescriptor_b2606b8baa50c1e2 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x6a, 0xe3, 0x40,
	0x14, 0x45, 0x5a, 0x58, 0x58, 0xb1, 0x5b, 0xac, 0x60, 0xc1, 0x88, 0x2d, 0x8c, 0xd7, 0xc5, 0xe2,
	0x62, 0x46, 0x5e, 0x37, 0xbb, 0xb3, 0x04, 0x23, 0x37, 0x4e, 0x15, 0x8c, 0x1d, 0x54, 0x04, 0x81,
	0x98, 0x48, 0x1f, 0x45, 0x60, 0x69, 0x9c, 0xf9, 0xb2, 0x5c, 0x84, 0x14, 0xc9, 0x15, 0x72, 0x83,
	0x94, 0xb9, 0x43, 0x2e, 0x90, 0x36, 0x45, 0x2e, 0x90, 0x2a, 0xa7, 0x08, 0xf2, 0x68, 0x04, 0x26,
	0x16, 0xee, 0x1e, 0xf3, 0xdf, 0x7b, 0xff, 0xfd, 0x27, 0x59, 0xe3, 0x44, 0x88, 0x64, 0x09, 0x94,
	0xc7, 0x48, 0x15, 0xac, 0x50, 0xe9, 0x52, 0x04, 0x59, 0xa6, 0x11, 0x20, 0x45, 0xe0, 0x32, 0xba,
	0x08, 0x0b, 0x90, 0x59, 0x58, 0xa6, 0xb0, 0x09, 0xeb, 0x09, 0x59, 0x49, 0x51, 0x08, 0xbb, 0xab,
	0x54, 0x84, 0xc7, 0x48, 0x1a, 0x03, 0x52, 0xba, 0x44, 0x1b, 0x38, 0x7f, 0xdb, 0x56, 0x48, 0x40,
	0xb1, 0x96, 0xfb, 0x76, 0x28, 0x6f, 0xe7, 0xa7, 0x56, 0xae, 0x52, 0xca, 0xf3, 0x5c, 0x14, 0xbc,
	0x48, 0x45, 0x8e, 0x6a, 0xda, 0x1b, 0x5b, 0x9d, 0x29, 0x14, 0x8b, 0xad, 0xf4, 0x14, 0x64, 0xe6,
	0xa7, 0xb0, 0x99, 0xc3, 0xe5, 0x1a, 0xb0, 0xb0, 0x7f, 0x59, 0xdf, 0xb4, 0x7b, 0x98, 0xf3, 0x0c,
	0x3a, 0x46, 0xd7, 0xf8, 0xfd, 0x65, 0xfe, 0x55, 0x3f, 0x9e, 0xf0, 0x0c, 0xfe, 0xbc, 0x18, 0xd6,
	0x8f, 0x5d, 0xf9, 0x42, 0x65, 0xb6, 0x1f, 0x0d, 0xeb, 0xfb, 0x07, 0x6f, 0x9b, 0x91, 0x43, 0xb7,
	0x92, 0xb6, 0x40, 0xce, 0xb0, 0x55, 0xdb, 0xb4, 0x40, 0x76, 0x95, 0xbd, 0x7f, 0xb7, 0xcf, 0xaf,
	0x77, 0xe6, 0xc8, 0x1e, 0x56, 0x5d, 0x5d, 0xed, 0x9c, 0x73, 0x14, 0xad, 0xb1, 0x10, 0x19, 0x48,
	0xa4, 0x83, 0xba, 0x3c, 0x2d, 0x43, 0x3a, 0xb8, 0x9e, 0xdc, 0x98, 0x56, 0x3f, 0x12, 0xd9, 0xc1,
	0xbc, 0x13, 0x67, 0xef, 0xfd, 0xb3, 0xaa, 0xdf, 0x99, 0x71, 0x76, 0x5c, 0xeb, 0x13, 0xb1, 0xe4,
	0x79, 0x42, 0x84, 0x4c, 0x68, 0x02, 0xf9, 0xb6, 0x7d, 0xfd, 0x25, 0x57, 0x29, 0xb6, 0xff, 0x3b,
	0xff, 0x35, 0xb8, 0x37, 0x3f, 0x4d, 0x3d, 0xef, 0xc1, 0xec, 0x4e, 0x95, 0xa1, 0x17, 0x23, 0x51,
	0xb0, 0x42, 0xbe, 0x4b, 0xea, 0xc5, 0xf8, 0xa4, 0x29, 0x81, 0x17, 0x63, 0xd0, 0x50, 0x02, 0xdf,
	0x0d, 0x34, 0xe5, 0xcd, 0xec, 0xab, 0x77, 0xc6, 0xbc, 0x18, 0x19, 0x6b, 0x48, 0x8c, 0xf9, 0x2e,
	0x63, 0x9a, 0x76, 0xfe, 0x79, 0x9b, 0x73, 0xf4, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x49, 0xc5, 0x5c,
	0x30, 0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchTermViewServiceClient is the client API for SearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchTermViewServiceClient interface {
	// Returns the attributes of the requested search term view.
	GetSearchTermView(ctx context.Context, in *GetSearchTermViewRequest, opts ...grpc.CallOption) (*resources.SearchTermView, error)
}

type searchTermViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchTermViewServiceClient(cc *grpc.ClientConn) SearchTermViewServiceClient {
	return &searchTermViewServiceClient{cc}
}

func (c *searchTermViewServiceClient) GetSearchTermView(ctx context.Context, in *GetSearchTermViewRequest, opts ...grpc.CallOption) (*resources.SearchTermView, error) {
	out := new(resources.SearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.SearchTermViewService/GetSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchTermViewServiceServer is the server API for SearchTermViewService service.
type SearchTermViewServiceServer interface {
	// Returns the attributes of the requested search term view.
	GetSearchTermView(context.Context, *GetSearchTermViewRequest) (*resources.SearchTermView, error)
}

func RegisterSearchTermViewServiceServer(s *grpc.Server, srv SearchTermViewServiceServer) {
	s.RegisterService(&_SearchTermViewService_serviceDesc, srv)
}

func _SearchTermViewService_GetSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTermViewServiceServer).GetSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.SearchTermViewService/GetSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTermViewServiceServer).GetSearchTermView(ctx, req.(*GetSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchTermViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.SearchTermViewService",
	HandlerType: (*SearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSearchTermView",
			Handler:    _SearchTermViewService_GetSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/search_term_view_service.proto",
}
