// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/keyword_view_service.proto

package google_ads_googleads_v0_services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	resources "google/ads/googleads/v0/resources"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [KeywordViewService.GetKeywordView][].
type GetKeywordViewRequest struct {
	// The resource name of the keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordViewRequest) Reset()         { *m = GetKeywordViewRequest{} }
func (m *GetKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordViewRequest) ProtoMessage()    {}
func (*GetKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f154d6d131e5f762, []int{0}
}
func (m *GetKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordViewRequest.Merge(m, src)
}
func (m *GetKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordViewRequest.Size(m)
}
func (m *GetKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordViewRequest proto.InternalMessageInfo

func (m *GetKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordViewRequest)(nil), "google.ads.googleads.v0.services.GetKeywordViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordViewServiceClient is the client API for KeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordViewServiceClient interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error)
}

type keywordViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordViewServiceClient(cc *grpc.ClientConn) KeywordViewServiceClient {
	return &keywordViewServiceClient{cc}
}

func (c *keywordViewServiceClient) GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error) {
	out := new(resources.KeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordViewService/GetKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordViewServiceServer is the server API for KeywordViewService service.
type KeywordViewServiceServer interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(context.Context, *GetKeywordViewRequest) (*resources.KeywordView, error)
}

func RegisterKeywordViewServiceServer(s *grpc.Server, srv KeywordViewServiceServer) {
	s.RegisterService(&_KeywordViewService_serviceDesc, srv)
}

func _KeywordViewService_GetKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordViewService/GetKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, req.(*GetKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.KeywordViewService",
	HandlerType: (*KeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordView",
			Handler:    _KeywordViewService_GetKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/keyword_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/keyword_view_service.proto", fileDescriptor_f154d6d131e5f762)
}

var fileDescriptor_f154d6d131e5f762 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0xec, 0xd4, 0xca, 0xf2, 0xfc, 0xa2, 0x94, 0xf8, 0xb2,
	0xcc, 0xd4, 0xf2, 0x78, 0xa8, 0xa8, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x02, 0x44, 0x87,
	0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xb3, 0x5e, 0x99, 0x81, 0x1e, 0x4c, 0xb3, 0x94, 0x09, 0x2e,
	0xe3, 0x8b, 0x52, 0x8b, 0xf3, 0x4b, 0x8b, 0xd0, 0xcd, 0x87, 0x98, 0x2b, 0x25, 0x03, 0xd3, 0x55,
	0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x55,
	0xb2, 0xe1, 0x12, 0x75, 0x4f, 0x2d, 0xf1, 0x86, 0x68, 0x0b, 0xcb, 0x4c, 0x2d, 0x0f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe6, 0xe2, 0x85, 0x19, 0x1b, 0x9f, 0x97, 0x98, 0x9b, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x03, 0x13, 0xf4, 0x4b, 0xcc, 0x4d, 0x35, 0x3a, 0xc1,
	0xc8, 0x25, 0x84, 0xa4, 0x37, 0x18, 0xe2, 0x52, 0xa1, 0x8d, 0x8c, 0x5c, 0x7c, 0xa8, 0xa6, 0x0a,
	0x99, 0xeb, 0x11, 0xf2, 0x9e, 0x1e, 0x56, 0x77, 0x48, 0xe9, 0xe1, 0xd4, 0x08, 0xf7, 0xb5, 0x1e,
	0x92, 0x36, 0x25, 0xb3, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x19, 0x08, 0xe9, 0x81, 0x02, 0xa6, 0x1a,
	0xc5, 0x0b, 0xb6, 0xc9, 0xa5, 0xc5, 0x25, 0xf9, 0xb9, 0xa9, 0x45, 0xc5, 0xfa, 0x5a, 0xb0, 0x90,
	0x02, 0xe9, 0x29, 0xd6, 0xd7, 0xaa, 0x75, 0xda, 0xc0, 0xc8, 0xa5, 0x92, 0x9c, 0x9f, 0x4b, 0xd0,
	0x99, 0x4e, 0xe2, 0x98, 0x1e, 0x0e, 0x00, 0x05, 0x65, 0x00, 0xe3, 0x22, 0x26, 0x66, 0x77, 0x47,
	0xc7, 0x55, 0x4c, 0x0a, 0xee, 0x10, 0x33, 0x1c, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc,
	0x40, 0x0f, 0xaa, 0xbc, 0xf8, 0x14, 0x4c, 0x49, 0x8c, 0x63, 0x4a, 0x71, 0x0c, 0x5c, 0x49, 0x4c,
	0x98, 0x41, 0x0c, 0x4c, 0xc9, 0x23, 0xc2, 0x4a, 0x92, 0xd8, 0xc0, 0x51, 0x68, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x36, 0x06, 0xf1, 0x43, 0x77, 0x02, 0x00, 0x00,
}
