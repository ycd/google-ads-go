// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/mobile_device_constant.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "github.com/ycd/google-ads-go/enums"
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

// A mobile device constant.
type MobileDeviceConstant struct {
	// The resource name of the mobile device constant.
	// Mobile device constant resource names have the form:
	//
	// `mobileDeviceConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile device constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the mobile device.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The manufacturer of the mobile device.
	ManufacturerName *wrappers.StringValue `protobuf:"bytes,4,opt,name=manufacturer_name,json=manufacturerName,proto3" json:"manufacturer_name,omitempty"`
	// The operating system of the mobile device.
	OperatingSystemName *wrappers.StringValue `protobuf:"bytes,5,opt,name=operating_system_name,json=operatingSystemName,proto3" json:"operating_system_name,omitempty"`
	// The type of mobile device.
	Type                 enums.MobileDeviceTypeEnum_MobileDeviceType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.MobileDeviceTypeEnum_MobileDeviceType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *MobileDeviceConstant) Reset()         { *m = MobileDeviceConstant{} }
func (m *MobileDeviceConstant) String() string { return proto.CompactTextString(m) }
func (*MobileDeviceConstant) ProtoMessage()    {}
func (*MobileDeviceConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_144dccbaa67cc327, []int{0}
}

func (m *MobileDeviceConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileDeviceConstant.Unmarshal(m, b)
}
func (m *MobileDeviceConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileDeviceConstant.Marshal(b, m, deterministic)
}
func (m *MobileDeviceConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileDeviceConstant.Merge(m, src)
}
func (m *MobileDeviceConstant) XXX_Size() int {
	return xxx_messageInfo_MobileDeviceConstant.Size(m)
}
func (m *MobileDeviceConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileDeviceConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileDeviceConstant proto.InternalMessageInfo

func (m *MobileDeviceConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileDeviceConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileDeviceConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MobileDeviceConstant) GetManufacturerName() *wrappers.StringValue {
	if m != nil {
		return m.ManufacturerName
	}
	return nil
}

func (m *MobileDeviceConstant) GetOperatingSystemName() *wrappers.StringValue {
	if m != nil {
		return m.OperatingSystemName
	}
	return nil
}

func (m *MobileDeviceConstant) GetType() enums.MobileDeviceTypeEnum_MobileDeviceType {
	if m != nil {
		return m.Type
	}
	return enums.MobileDeviceTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MobileDeviceConstant)(nil), "google.ads.googleads.v0.resources.MobileDeviceConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/mobile_device_constant.proto", fileDescriptor_144dccbaa67cc327)
}

var fileDescriptor_144dccbaa67cc327 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0x69, 0x67, 0x5c, 0x30, 0xfe, 0x41, 0xab, 0xc2, 0xb8, 0x8a, 0xcc, 0x2a, 0x0b, 0x03,
	0x42, 0x5a, 0x56, 0xd9, 0x8b, 0x08, 0x42, 0xd7, 0x95, 0x65, 0x05, 0x65, 0x98, 0x95, 0x41, 0x64,
	0x60, 0xc8, 0x34, 0x67, 0x43, 0xa1, 0x49, 0x4a, 0x92, 0x8e, 0xcc, 0x0b, 0x78, 0xef, 0x2b, 0x78,
	0xe9, 0xa3, 0xf8, 0x28, 0x3e, 0x85, 0x34, 0x69, 0xc3, 0xe2, 0xba, 0x3a, 0x77, 0xa7, 0xcd, 0xf7,
	0xfb, 0xce, 0x97, 0x9c, 0x83, 0x5e, 0x73, 0xa5, 0x78, 0x05, 0x29, 0x65, 0x26, 0xf5, 0x65, 0x5b,
	0xad, 0xb3, 0x54, 0x83, 0x51, 0x8d, 0x2e, 0xc0, 0xa4, 0x42, 0xad, 0xca, 0x0a, 0x96, 0x0c, 0xd6,
	0x65, 0x01, 0xcb, 0x42, 0x49, 0x63, 0xa9, 0xb4, 0xb8, 0xd6, 0xca, 0xaa, 0x64, 0xcf, 0x43, 0x98,
	0x32, 0x83, 0x03, 0x8f, 0xd7, 0x19, 0x0e, 0xfc, 0xee, 0xe1, 0x55, 0x2d, 0x40, 0x36, 0xe2, 0x4f,
	0x7b, 0xbb, 0xa9, 0xc1, 0x5b, 0xef, 0x3e, 0xe9, 0x38, 0xf7, 0xb5, 0x6a, 0xce, 0xd3, 0x2f, 0x9a,
	0xd6, 0x35, 0x68, 0xe3, 0xcf, 0x9f, 0x7e, 0x1b, 0xa0, 0xfb, 0xef, 0x1d, 0x7c, 0xec, 0xd8, 0x37,
	0x5d, 0xb2, 0xe4, 0x19, 0xba, 0xd5, 0x77, 0x5f, 0x4a, 0x2a, 0x60, 0x14, 0x8d, 0xa3, 0xc9, 0xf5,
	0xd9, 0xcd, 0xfe, 0xe7, 0x07, 0x2a, 0x20, 0x79, 0x8e, 0xe2, 0x92, 0x8d, 0xe2, 0x71, 0x34, 0xb9,
	0x71, 0xf0, 0xa8, 0x8b, 0x8e, 0xfb, 0x56, 0xf8, 0x54, 0xda, 0xc3, 0x97, 0x73, 0x5a, 0x35, 0x30,
	0x8b, 0x4b, 0x96, 0x64, 0x68, 0xe8, 0x8c, 0x06, 0x4e, 0xfe, 0xf8, 0x92, 0xfc, 0xcc, 0xea, 0x52,
	0x72, 0xaf, 0x77, 0xca, 0xe4, 0x14, 0xdd, 0x15, 0x54, 0x36, 0xe7, 0xb4, 0xb0, 0x8d, 0x06, 0xed,
	0x73, 0x0c, 0xb7, 0xc0, 0xef, 0x5c, 0xc4, 0x5c, 0xd2, 0x29, 0x7a, 0xa0, 0x6a, 0xd0, 0xd4, 0x96,
	0x92, 0x2f, 0xcd, 0xc6, 0x58, 0x10, 0xde, 0xee, 0xda, 0x16, 0x76, 0xf7, 0x02, 0x7a, 0xe6, 0x48,
	0xe7, 0xf8, 0x09, 0x0d, 0xdb, 0x77, 0x1e, 0xed, 0x8c, 0xa3, 0xc9, 0xed, 0x83, 0x63, 0x7c, 0xd5,
	0x0c, 0xdd, 0x80, 0xf0, 0xc5, 0x37, 0xfe, 0xb8, 0xa9, 0xe1, 0xad, 0x6c, 0xc4, 0xa5, 0x9f, 0x33,
	0xe7, 0x78, 0xf4, 0x35, 0x46, 0xfb, 0x85, 0x12, 0xf8, 0xbf, 0x5b, 0x71, 0xf4, 0xf0, 0x6f, 0xa3,
	0x9b, 0xb6, 0x57, 0x98, 0x46, 0x9f, 0xdf, 0x75, 0x3c, 0x57, 0x15, 0x95, 0x1c, 0x2b, 0xcd, 0x53,
	0x0e, 0xd2, 0x5d, 0xb0, 0x5f, 0xa1, 0xba, 0x34, 0xff, 0x58, 0xda, 0x57, 0xa1, 0xfa, 0x1e, 0x0f,
	0x4e, 0xf2, 0xfc, 0x47, 0xbc, 0x77, 0xe2, 0x2d, 0x73, 0x66, 0xb0, 0x2f, 0xdb, 0x6a, 0x9e, 0xe1,
	0x59, 0xaf, 0xfc, 0xd9, 0x6b, 0x16, 0x39, 0x33, 0x8b, 0xa0, 0x59, 0xcc, 0xb3, 0x45, 0xd0, 0xfc,
	0x8a, 0xf7, 0xfd, 0x01, 0x21, 0x39, 0x33, 0x84, 0x04, 0x15, 0x21, 0xf3, 0x8c, 0x90, 0xa0, 0x5b,
	0xed, 0xb8, 0xb0, 0x2f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xc5, 0xfd, 0xf2, 0x60, 0x03,
	0x00, 0x00,
}
