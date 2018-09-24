// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/new_resource_creation_error.proto

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

// Enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum_NewResourceCreationError int32

const (
	// Enum unspecified.
	NewResourceCreationErrorEnum_UNSPECIFIED NewResourceCreationErrorEnum_NewResourceCreationError = 0
	// The received error code is not known in this version.
	NewResourceCreationErrorEnum_UNKNOWN NewResourceCreationErrorEnum_NewResourceCreationError = 1
	// Do not set the id field while creating new entities.
	NewResourceCreationErrorEnum_CANNOT_SET_ID_FOR_ADD NewResourceCreationErrorEnum_NewResourceCreationError = 2
	// Creating more than one resource with the same temp ID is not allowed.
	NewResourceCreationErrorEnum_DUPLICATE_TEMP_IDS NewResourceCreationErrorEnum_NewResourceCreationError = 3
	// Parent object with specified temp id failed validation, so no deep
	// validation will be done for this child resource.
	NewResourceCreationErrorEnum_TEMP_ID_RESOURCE_HAD_ERRORS NewResourceCreationErrorEnum_NewResourceCreationError = 4
)

var NewResourceCreationErrorEnum_NewResourceCreationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_SET_ID_FOR_ADD",
	3: "DUPLICATE_TEMP_IDS",
	4: "TEMP_ID_RESOURCE_HAD_ERRORS",
}

var NewResourceCreationErrorEnum_NewResourceCreationError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"CANNOT_SET_ID_FOR_ADD":       2,
	"DUPLICATE_TEMP_IDS":          3,
	"TEMP_ID_RESOURCE_HAD_ERRORS": 4,
}

func (x NewResourceCreationErrorEnum_NewResourceCreationError) String() string {
	return proto.EnumName(NewResourceCreationErrorEnum_NewResourceCreationError_name, int32(x))
}

func (NewResourceCreationErrorEnum_NewResourceCreationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44d42568dcd14b02, []int{0, 0}
}

// Container for enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResourceCreationErrorEnum) Reset()         { *m = NewResourceCreationErrorEnum{} }
func (m *NewResourceCreationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NewResourceCreationErrorEnum) ProtoMessage()    {}
func (*NewResourceCreationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_44d42568dcd14b02, []int{0}
}
func (m *NewResourceCreationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Unmarshal(m, b)
}
func (m *NewResourceCreationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Marshal(b, m, deterministic)
}
func (m *NewResourceCreationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResourceCreationErrorEnum.Merge(m, src)
}
func (m *NewResourceCreationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Size(m)
}
func (m *NewResourceCreationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResourceCreationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NewResourceCreationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewResourceCreationErrorEnum)(nil), "google.ads.googleads.v0.errors.NewResourceCreationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.NewResourceCreationErrorEnum_NewResourceCreationError", NewResourceCreationErrorEnum_NewResourceCreationError_name, NewResourceCreationErrorEnum_NewResourceCreationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/new_resource_creation_error.proto", fileDescriptor_44d42568dcd14b02)
}

var fileDescriptor_44d42568dcd14b02 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xdd, 0x4a, 0xf3, 0x30,
	0x18, 0x80, 0xbf, 0x6e, 0x1f, 0x0a, 0xd9, 0x81, 0x25, 0xa0, 0x28, 0xea, 0x84, 0x5d, 0x40, 0x5a,
	0xf0, 0x06, 0xcc, 0x9a, 0x6c, 0x16, 0x35, 0x2d, 0x69, 0x3b, 0x4f, 0x0a, 0x2f, 0x73, 0x0d, 0x43,
	0x70, 0x8b, 0x24, 0xfb, 0xb9, 0x09, 0xaf, 0xc2, 0x43, 0x41, 0x2f, 0xc4, 0xcb, 0xf0, 0x4a, 0xa4,
	0x4b, 0xdd, 0xd9, 0xf4, 0xec, 0x09, 0xef, 0x93, 0x87, 0xfc, 0xa0, 0xab, 0xa9, 0xd6, 0xd3, 0x27,
	0x15, 0x8c, 0x2b, 0x1b, 0x38, 0xac, 0x69, 0x15, 0x06, 0xca, 0x18, 0x6d, 0x6c, 0x30, 0x57, 0x6b,
	0x30, 0xca, 0xea, 0xa5, 0x99, 0x28, 0x98, 0x18, 0x35, 0x5e, 0x3c, 0xea, 0x39, 0x6c, 0x86, 0xe4,
	0xd9, 0xe8, 0x85, 0xc6, 0x5d, 0xb7, 0x8d, 0x8c, 0x2b, 0x4b, 0xb6, 0x05, 0xb2, 0x0a, 0x89, 0x2b,
	0xf4, 0xde, 0x3d, 0x74, 0x26, 0xd4, 0x5a, 0x36, 0x91, 0xa8, 0x69, 0xf0, 0x7a, 0xca, 0xe7, 0xcb,
	0x59, 0xef, 0xc5, 0x43, 0xc7, 0xbb, 0x04, 0x7c, 0x80, 0x3a, 0x85, 0xc8, 0x52, 0x1e, 0xc5, 0x83,
	0x98, 0x33, 0xff, 0x1f, 0xee, 0xa0, 0xfd, 0x42, 0xdc, 0x88, 0xe4, 0x5e, 0xf8, 0x1e, 0x3e, 0x41,
	0x87, 0x11, 0x15, 0x22, 0xc9, 0x21, 0xe3, 0x39, 0xc4, 0x0c, 0x06, 0x89, 0x04, 0xca, 0x98, 0xdf,
	0xc2, 0x47, 0x08, 0xb3, 0x22, 0xbd, 0x8d, 0x23, 0x9a, 0x73, 0xc8, 0xf9, 0x5d, 0x0a, 0x31, 0xcb,
	0xfc, 0x36, 0xbe, 0x40, 0xa7, 0xcd, 0x0a, 0x24, 0xcf, 0x92, 0x42, 0x46, 0x1c, 0xae, 0x29, 0x03,
	0x2e, 0x65, 0x22, 0x33, 0xff, 0x7f, 0xff, 0xc3, 0x43, 0xbd, 0x89, 0x9e, 0x91, 0xdf, 0xaf, 0xd5,
	0x3f, 0xdf, 0x75, 0xe4, 0xb4, 0x7e, 0x95, 0xd4, 0x7b, 0x6d, 0xb5, 0x87, 0x94, 0xbe, 0xb5, 0xba,
	0x43, 0xd7, 0xa1, 0x95, 0x25, 0x0e, 0x6b, 0x1a, 0x85, 0x64, 0x23, 0xdb, 0xcf, 0x1f, 0xa1, 0xa4,
	0x95, 0x2d, 0xb7, 0x42, 0x39, 0x0a, 0x4b, 0x27, 0x7c, 0xfd, 0x25, 0x3c, 0xec, 0x6d, 0xfe, 0xe1,
	0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x10, 0x3a, 0xf3, 0xcb, 0x01, 0x00, 0x00,
}
