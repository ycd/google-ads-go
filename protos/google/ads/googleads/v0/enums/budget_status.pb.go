// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/budget_status.proto

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

// Possible statuses of a Budget.
type BudgetStatusEnum_BudgetStatus int32

const (
	// Not specified.
	BudgetStatusEnum_UNSPECIFIED BudgetStatusEnum_BudgetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	BudgetStatusEnum_UNKNOWN BudgetStatusEnum_BudgetStatus = 1
	// Budget is enabled.
	BudgetStatusEnum_ENABLED BudgetStatusEnum_BudgetStatus = 2
	// Budget is removed.
	BudgetStatusEnum_REMOVED BudgetStatusEnum_BudgetStatus = 3
)

var BudgetStatusEnum_BudgetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var BudgetStatusEnum_BudgetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x BudgetStatusEnum_BudgetStatus) String() string {
	return proto.EnumName(BudgetStatusEnum_BudgetStatus_name, int32(x))
}

func (BudgetStatusEnum_BudgetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d40ae52141cda825, []int{0, 0}
}

// Message describing a Budget status
type BudgetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BudgetStatusEnum) Reset()         { *m = BudgetStatusEnum{} }
func (m *BudgetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BudgetStatusEnum) ProtoMessage()    {}
func (*BudgetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d40ae52141cda825, []int{0}
}
func (m *BudgetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BudgetStatusEnum.Unmarshal(m, b)
}
func (m *BudgetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BudgetStatusEnum.Marshal(b, m, deterministic)
}
func (m *BudgetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BudgetStatusEnum.Merge(m, src)
}
func (m *BudgetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BudgetStatusEnum.Size(m)
}
func (m *BudgetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BudgetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BudgetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BudgetStatusEnum)(nil), "google.ads.googleads.v0.enums.BudgetStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.BudgetStatusEnum_BudgetStatus", BudgetStatusEnum_BudgetStatus_name, BudgetStatusEnum_BudgetStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/budget_status.proto", fileDescriptor_d40ae52141cda825)
}

var fileDescriptor_d40ae52141cda825 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xa4, 0xd2, 0x94, 0xf4, 0xd4, 0x92, 0xf8, 0xe2, 0x92, 0xc4, 0x92, 0xd2,
	0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x3a, 0xbd, 0xc4, 0x94, 0x62, 0x3d,
	0xb8, 0x16, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x16, 0xa5, 0x28, 0x2e, 0x01, 0x27, 0xb0, 0xae, 0x60,
	0xb0, 0x26, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x37, 0x2e, 0x1e, 0x64, 0x31, 0x21, 0x7e, 0x2e, 0xee,
	0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e,
	0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x10, 0xc7, 0xd5, 0xcf, 0xd1, 0xc9,
	0xc7, 0xd5, 0x45, 0x80, 0x09, 0xc4, 0x09, 0x72, 0xf5, 0xf5, 0x0f, 0x73, 0x75, 0x11, 0x60, 0x76,
	0x9a, 0xc7, 0xc8, 0xa5, 0x98, 0x9c, 0x9f, 0xab, 0x87, 0xd7, 0x05, 0x4e, 0x82, 0xc8, 0x76, 0x05,
	0x80, 0xdc, 0x1c, 0xc0, 0xb8, 0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15, 0x93, 0xac, 0x3b, 0x44,
	0xab, 0x63, 0x4a, 0xb1, 0x1e, 0x84, 0x09, 0x62, 0x85, 0x19, 0xe8, 0x81, 0xdc, 0x59, 0x7c, 0x0a,
	0x26, 0x1f, 0xe3, 0x98, 0x52, 0x1c, 0x03, 0x97, 0x8f, 0x09, 0x33, 0x88, 0x01, 0xcb, 0x3f, 0x22,
	0x20, 0x9f, 0xc4, 0x06, 0x0e, 0x22, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x8e, 0x68,
	0x93, 0x57, 0x01, 0x00, 0x00,
}
