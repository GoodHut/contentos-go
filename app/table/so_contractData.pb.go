// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_contractData.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoContractData struct {
	Id                   *prototype.ContractDataId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                []byte                    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SoContractData) Reset()         { *m = SoContractData{} }
func (m *SoContractData) String() string { return proto.CompactTextString(m) }
func (*SoContractData) ProtoMessage()    {}
func (*SoContractData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3e7d262658ec8d, []int{0}
}

func (m *SoContractData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoContractData.Unmarshal(m, b)
}
func (m *SoContractData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoContractData.Marshal(b, m, deterministic)
}
func (m *SoContractData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoContractData.Merge(m, src)
}
func (m *SoContractData) XXX_Size() int {
	return xxx_messageInfo_SoContractData.Size(m)
}
func (m *SoContractData) XXX_DiscardUnknown() {
	xxx_messageInfo_SoContractData.DiscardUnknown(m)
}

var xxx_messageInfo_SoContractData proto.InternalMessageInfo

func (m *SoContractData) GetId() *prototype.ContractDataId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SoContractData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SoUniqueContractDataById struct {
	Id                   *prototype.ContractDataId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SoUniqueContractDataById) Reset()         { *m = SoUniqueContractDataById{} }
func (m *SoUniqueContractDataById) String() string { return proto.CompactTextString(m) }
func (*SoUniqueContractDataById) ProtoMessage()    {}
func (*SoUniqueContractDataById) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3e7d262658ec8d, []int{1}
}

func (m *SoUniqueContractDataById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueContractDataById.Unmarshal(m, b)
}
func (m *SoUniqueContractDataById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueContractDataById.Marshal(b, m, deterministic)
}
func (m *SoUniqueContractDataById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueContractDataById.Merge(m, src)
}
func (m *SoUniqueContractDataById) XXX_Size() int {
	return xxx_messageInfo_SoUniqueContractDataById.Size(m)
}
func (m *SoUniqueContractDataById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueContractDataById.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueContractDataById proto.InternalMessageInfo

func (m *SoUniqueContractDataById) GetId() *prototype.ContractDataId {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*SoContractData)(nil), "table.so_contractData")
	proto.RegisterType((*SoUniqueContractDataById)(nil), "table.so_unique_contractData_by_id")
}

func init() { proto.RegisterFile("app/table/so_contractData.proto", fileDescriptor_7d3e7d262658ec8d) }

var fileDescriptor_7d3e7d262658ec8d = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0x28, 0xd0,
	0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0xce, 0x8f, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c,
	0x2e, 0x71, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x4b, 0x4a,
	0x89, 0x80, 0x79, 0x25, 0x95, 0x05, 0xa9, 0xfa, 0x20, 0x02, 0x22, 0xa9, 0x14, 0xc2, 0xc5, 0x8f,
	0xa6, 0x4b, 0x48, 0x9b, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5a,
	0x0f, 0xae, 0x4b, 0x0f, 0xa6, 0x28, 0x3e, 0x25, 0xb1, 0x24, 0x31, 0x3e, 0x33, 0x25, 0x88, 0x29,
	0x33, 0x45, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x27, 0x08, 0xc2, 0x51, 0xf2, 0xe6, 0x92, 0x29, 0xce, 0x8f, 0x2f, 0xcd, 0xcb, 0x2c, 0x2c, 0x4d,
	0x45, 0x31, 0x3c, 0x3e, 0xa9, 0x32, 0x3e, 0x33, 0x85, 0x24, 0x2b, 0x9c, 0x34, 0xa2, 0xd4, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x93, 0x33, 0x12,
	0x33, 0xf3, 0xf4, 0x41, 0x6a, 0x53, 0xf3, 0x4a, 0xf2, 0x8b, 0x75, 0xd3, 0xf3, 0x21, 0xfe, 0x4f,
	0x62, 0x03, 0x9b, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x78, 0x12, 0x07, 0xee, 0x13, 0x01,
	0x00, 0x00,
}
