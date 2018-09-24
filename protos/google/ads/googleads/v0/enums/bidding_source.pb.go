// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/bidding_source.proto

package google_ads_googleads_v0_enums

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible bidding sources.
type BiddingSourceEnum_BiddingSource int32

const (
	// Not specified.
	BiddingSourceEnum_UNSPECIFIED BiddingSourceEnum_BiddingSource = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingSourceEnum_UNKNOWN BiddingSourceEnum_BiddingSource = 1
	// Bidding entity is defined on the ad group.
	BiddingSourceEnum_ADGROUP BiddingSourceEnum_BiddingSource = 2
	// Bidding entity is defined on the ad group criterion.
	BiddingSourceEnum_CRITERION BiddingSourceEnum_BiddingSource = 3
	// Effective bidding entity is inherited from campaign bidding strategy.
	BiddingSourceEnum_CAMPAIGN_BIDDING_STRATEGY BiddingSourceEnum_BiddingSource = 5
)

var BiddingSourceEnum_BiddingSource_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADGROUP",
	3: "CRITERION",
	5: "CAMPAIGN_BIDDING_STRATEGY",
}

var BiddingSourceEnum_BiddingSource_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"ADGROUP":                   2,
	"CRITERION":                 3,
	"CAMPAIGN_BIDDING_STRATEGY": 5,
}

func (x BiddingSourceEnum_BiddingSource) String() string {
	return proto.EnumName(BiddingSourceEnum_BiddingSource_name, int32(x))
}

func (BiddingSourceEnum_BiddingSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_638500c1459ae2f0, []int{0, 0}
}

// Container for enum describing possible bidding sources.
type BiddingSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingSourceEnum) Reset()         { *m = BiddingSourceEnum{} }
func (m *BiddingSourceEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingSourceEnum) ProtoMessage()    {}
func (*BiddingSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_638500c1459ae2f0, []int{0}
}
func (m *BiddingSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingSourceEnum.Unmarshal(m, b)
}
func (m *BiddingSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingSourceEnum.Marshal(b, m, deterministic)
}
func (m *BiddingSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingSourceEnum.Merge(m, src)
}
func (m *BiddingSourceEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingSourceEnum.Size(m)
}
func (m *BiddingSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BiddingSourceEnum)(nil), "google.ads.googleads.v0.enums.BiddingSourceEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.BiddingSourceEnum_BiddingSource", BiddingSourceEnum_BiddingSource_name, BiddingSourceEnum_BiddingSource_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/bidding_source.proto", fileDescriptor_638500c1459ae2f0)
}

var fileDescriptor_638500c1459ae2f0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xa4, 0xcc, 0x94, 0x94, 0xcc, 0xbc, 0xf4, 0xf8, 0xe2, 0xfc, 0xd2, 0xa2,
	0xe4, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x42, 0xbd, 0xc4, 0x94, 0x62,
	0x3d, 0xb8, 0x1e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x1e, 0xa5, 0x5a, 0x2e, 0x41, 0x27, 0x88, 0xb6,
	0x60, 0xb0, 0x2e, 0xd7, 0xbc, 0xd2, 0x5c, 0xa5, 0x0c, 0x2e, 0x5e, 0x14, 0x41, 0x21, 0x7e, 0x2e,
	0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e,
	0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x10, 0xc7, 0xd1, 0xc5, 0x3d,
	0xc8, 0x3f, 0x34, 0x40, 0x80, 0x49, 0x88, 0x97, 0x8b, 0xd3, 0x39, 0xc8, 0x33, 0xc4, 0x35, 0xc8,
	0xd3, 0xdf, 0x4f, 0x80, 0x59, 0x48, 0x96, 0x4b, 0xd2, 0xd9, 0xd1, 0x37, 0xc0, 0xd1, 0xd3, 0xdd,
	0x2f, 0xde, 0xc9, 0xd3, 0xc5, 0xc5, 0xd3, 0xcf, 0x3d, 0x3e, 0x38, 0x24, 0xc8, 0x31, 0xc4, 0xd5,
	0x3d, 0x52, 0x80, 0xd5, 0x69, 0x3e, 0x23, 0x97, 0x62, 0x72, 0x7e, 0xae, 0x1e, 0x5e, 0x47, 0x3a,
	0x09, 0xa1, 0xb8, 0x26, 0x00, 0xe4, 0xaf, 0x00, 0xc6, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e, 0xab,
	0x98, 0x64, 0xdd, 0x21, 0x7a, 0x1d, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f,
	0xe4, 0x95, 0xe2, 0x53, 0x30, 0xf9, 0x18, 0xc7, 0x94, 0xe2, 0x18, 0xb8, 0x7c, 0x4c, 0x98, 0x41,
	0x0c, 0x58, 0xfe, 0x11, 0x01, 0xf9, 0x24, 0x36, 0x70, 0x30, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x76, 0x97, 0xa3, 0xc9, 0x7c, 0x01, 0x00, 0x00,
}
