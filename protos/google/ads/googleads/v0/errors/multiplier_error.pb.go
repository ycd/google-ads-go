// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/multiplier_error.proto

package google_ads_googleads_v0_errors

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

// Enum describing possible multiplier errors.
type MultiplierErrorEnum_MultiplierError int32

const (
	// Enum unspecified.
	MultiplierErrorEnum_UNSPECIFIED MultiplierErrorEnum_MultiplierError = 0
	// The received error code is not known in this version.
	MultiplierErrorEnum_UNKNOWN MultiplierErrorEnum_MultiplierError = 1
	// Multiplier value is too high
	MultiplierErrorEnum_MULTIPLIER_TOO_HIGH MultiplierErrorEnum_MultiplierError = 2
	// Multiplier value is too low
	MultiplierErrorEnum_MULTIPLIER_TOO_LOW MultiplierErrorEnum_MultiplierError = 3
	// Too many fractional digits
	MultiplierErrorEnum_TOO_MANY_FRACTIONAL_DIGITS MultiplierErrorEnum_MultiplierError = 4
	// A multiplier cannot be set for this bidding strategy
	MultiplierErrorEnum_MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY MultiplierErrorEnum_MultiplierError = 5
	// A multiplier cannot be set when there is no base bid (e.g., content max
	// cpc)
	MultiplierErrorEnum_MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING MultiplierErrorEnum_MultiplierError = 6
	// A bid multiplier must be specified
	MultiplierErrorEnum_NO_MULTIPLIER_SPECIFIED MultiplierErrorEnum_MultiplierError = 7
	// Multiplier causes bid to exceed daily budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET MultiplierErrorEnum_MultiplierError = 8
	// Multiplier causes bid to exceed monthly budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET MultiplierErrorEnum_MultiplierError = 9
	// Multiplier causes bid to exceed custom budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET MultiplierErrorEnum_MultiplierError = 10
	// Multiplier causes bid to exceed maximum allowed bid
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID MultiplierErrorEnum_MultiplierError = 11
	// Multiplier causes bid to become less than the minimum bid allowed
	MultiplierErrorEnum_BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER MultiplierErrorEnum_MultiplierError = 12
	// Multiplier type (cpc vs. cpm) needs to match campaign's bidding strategy
	MultiplierErrorEnum_MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH MultiplierErrorEnum_MultiplierError = 13
)

var MultiplierErrorEnum_MultiplierError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "MULTIPLIER_TOO_HIGH",
	3:  "MULTIPLIER_TOO_LOW",
	4:  "TOO_MANY_FRACTIONAL_DIGITS",
	5:  "MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY",
	6:  "MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING",
	7:  "NO_MULTIPLIER_SPECIFIED",
	8:  "MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET",
	9:  "MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET",
	10: "MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET",
	11: "MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID",
	12: "BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER",
	13: "MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH",
}

var MultiplierErrorEnum_MultiplierError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"MULTIPLIER_TOO_HIGH":        2,
	"MULTIPLIER_TOO_LOW":         3,
	"TOO_MANY_FRACTIONAL_DIGITS": 4,
	"MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY":     5,
	"MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING": 6,
	"NO_MULTIPLIER_SPECIFIED":                         7,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET":    8,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET":  9,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET":   10,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID": 11,
	"BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER":   12,
	"MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH":   13,
}

func (x MultiplierErrorEnum_MultiplierError) String() string {
	return proto.EnumName(MultiplierErrorEnum_MultiplierError_name, int32(x))
}

func (MultiplierErrorEnum_MultiplierError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb6c3c8aa121bbde, []int{0, 0}
}

// Container for enum describing possible multiplier errors.
type MultiplierErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplierErrorEnum) Reset()         { *m = MultiplierErrorEnum{} }
func (m *MultiplierErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MultiplierErrorEnum) ProtoMessage()    {}
func (*MultiplierErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6c3c8aa121bbde, []int{0}
}
func (m *MultiplierErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplierErrorEnum.Unmarshal(m, b)
}
func (m *MultiplierErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplierErrorEnum.Marshal(b, m, deterministic)
}
func (m *MultiplierErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplierErrorEnum.Merge(m, src)
}
func (m *MultiplierErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MultiplierErrorEnum.Size(m)
}
func (m *MultiplierErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplierErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplierErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MultiplierErrorEnum)(nil), "google.ads.googleads.v0.errors.MultiplierErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.MultiplierErrorEnum_MultiplierError", MultiplierErrorEnum_MultiplierError_name, MultiplierErrorEnum_MultiplierError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/multiplier_error.proto", fileDescriptor_bb6c3c8aa121bbde)
}

var fileDescriptor_bb6c3c8aa121bbde = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0x69, 0x57, 0x36, 0x70, 0x41, 0xb3, 0x3c, 0xc4, 0x24, 0x90, 0x7a, 0xd1, 0x5b, 0xc0,
	0x09, 0x4c, 0x3c, 0x80, 0x13, 0xbb, 0x89, 0x45, 0x62, 0x47, 0xb5, 0xb3, 0xac, 0x52, 0xa5, 0xa3,
	0x41, 0xab, 0x69, 0x52, 0x4b, 0xa6, 0x64, 0xdb, 0x03, 0x21, 0x71, 0xc3, 0x0b, 0xf0, 0x0e, 0xdc,
	0xf0, 0x0e, 0x3c, 0x09, 0x72, 0xa3, 0x75, 0x51, 0xf8, 0xd3, 0xbb, 0x5f, 0x8e, 0xbf, 0x2f, 0xce,
	0x39, 0x3a, 0x41, 0xef, 0x2f, 0xca, 0xf2, 0x62, 0xb5, 0xf4, 0xce, 0x17, 0xb5, 0xd7, 0x44, 0x97,
	0x6e, 0x7d, 0x6f, 0x59, 0x55, 0x65, 0x55, 0x7b, 0xeb, 0x9b, 0xd5, 0xf5, 0xe5, 0xd5, 0xea, 0x72,
	0x59, 0xc1, 0xa6, 0x42, 0xaf, 0xaa, 0xf2, 0xba, 0x24, 0xa3, 0x86, 0xa5, 0xe7, 0x8b, 0x9a, 0x6e,
	0x35, 0x7a, 0xeb, 0xd3, 0x46, 0x1b, 0xff, 0x1c, 0xa0, 0xa3, 0x74, 0xab, 0x0a, 0x57, 0x14, 0x9f,
	0x6f, 0xd6, 0xe3, 0xef, 0x03, 0x74, 0xd8, 0xa9, 0x93, 0x43, 0x34, 0xcc, 0x95, 0xc9, 0x44, 0x28,
	0x27, 0x52, 0x70, 0xfc, 0x80, 0x0c, 0xd1, 0x41, 0xae, 0x3e, 0x28, 0x5d, 0x28, 0xdc, 0x23, 0xc7,
	0xe8, 0x28, 0xcd, 0x13, 0x2b, 0xb3, 0x44, 0x8a, 0x29, 0x58, 0xad, 0x21, 0x96, 0x51, 0x8c, 0xfb,
	0xe4, 0x39, 0x22, 0x9d, 0x83, 0x44, 0x17, 0x78, 0x8f, 0x8c, 0xd0, 0x0b, 0xf7, 0x90, 0x32, 0x35,
	0x83, 0xc9, 0x94, 0x85, 0x56, 0x6a, 0xc5, 0x12, 0xe0, 0x32, 0x92, 0xd6, 0xe0, 0x01, 0xf1, 0xd0,
	0xab, 0x96, 0xa7, 0xb4, 0x05, 0x96, 0x24, 0xba, 0x10, 0x1c, 0x26, 0x7a, 0x0a, 0x81, 0xe4, 0x5c,
	0xaa, 0x08, 0x8c, 0x9d, 0x32, 0x2b, 0xa2, 0x19, 0x7e, 0x48, 0x4e, 0x90, 0xf7, 0x0f, 0xa1, 0x88,
	0x85, 0x82, 0x80, 0x19, 0xe1, 0x34, 0x90, 0x06, 0x52, 0x69, 0x8c, 0x54, 0x11, 0xde, 0x27, 0x2f,
	0xd1, 0xb1, 0xd2, 0xd0, 0xf2, 0xee, 0x1b, 0x3c, 0x20, 0x3e, 0x7a, 0xdd, 0x3a, 0x09, 0x59, 0x6e,
	0x84, 0xd9, 0xbc, 0xc2, 0x6a, 0x10, 0x67, 0xa1, 0x10, 0x1c, 0x38, 0x93, 0xc9, 0x0c, 0x82, 0x9c,
	0x47, 0xc2, 0xe2, 0x47, 0xe4, 0x1d, 0xa2, 0xbb, 0x8c, 0x54, 0x2b, 0x1b, 0xdf, 0x3b, 0x8f, 0xc9,
	0x5b, 0xf4, 0x66, 0x97, 0x13, 0xe6, 0xc6, 0xea, 0xf4, 0x4e, 0x41, 0x9d, 0x56, 0xff, 0x7e, 0x0d,
	0x3b, 0xdb, 0xf6, 0x1f, 0x48, 0x8e, 0x87, 0xee, 0x1e, 0x87, 0x24, 0xc2, 0x18, 0xb0, 0x31, 0x53,
	0x90, 0x4a, 0xd5, 0x46, 0xa0, 0x90, 0x36, 0x6e, 0x8d, 0x02, 0x3f, 0xe9, 0x7c, 0x1a, 0x53, 0xfc,
	0x8f, 0xb9, 0x83, 0x9d, 0x65, 0xc2, 0xcd, 0x33, 0x65, 0x36, 0x8c, 0xf1, 0xd3, 0xe0, 0x6b, 0x0f,
	0x8d, 0x3f, 0x95, 0x6b, 0xfa, 0xff, 0xc5, 0x0b, 0x9e, 0x75, 0xb6, 0x2b, 0x73, 0xeb, 0x9a, 0xf5,
	0xbe, 0xf4, 0xf7, 0x22, 0xc6, 0xbe, 0xf5, 0x47, 0x51, 0xa3, 0xb3, 0x45, 0x4d, 0x9b, 0xe8, 0xd2,
	0xa9, 0x4f, 0x37, 0x70, 0xfd, 0xe3, 0x0e, 0x98, 0xb3, 0x45, 0x3d, 0xdf, 0x02, 0xf3, 0x53, 0x7f,
	0xde, 0x00, 0xbf, 0x76, 0x01, 0x1f, 0xf7, 0x37, 0x3f, 0xc8, 0xc9, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x94, 0x34, 0x48, 0x59, 0x03, 0x00, 0x00,
}
