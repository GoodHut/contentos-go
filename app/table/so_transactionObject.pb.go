// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_transactionObject.proto

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

type SoTransactionObject struct {
	TrxId                *prototype.Sha256       `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	Expiration           *prototype.TimePointSec `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoTransactionObject) Reset()         { *m = SoTransactionObject{} }
func (m *SoTransactionObject) String() string { return proto.CompactTextString(m) }
func (*SoTransactionObject) ProtoMessage()    {}
func (*SoTransactionObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e5dbc6d4c15723, []int{0}
}

func (m *SoTransactionObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoTransactionObject.Unmarshal(m, b)
}
func (m *SoTransactionObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoTransactionObject.Marshal(b, m, deterministic)
}
func (m *SoTransactionObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoTransactionObject.Merge(m, src)
}
func (m *SoTransactionObject) XXX_Size() int {
	return xxx_messageInfo_SoTransactionObject.Size(m)
}
func (m *SoTransactionObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SoTransactionObject.DiscardUnknown(m)
}

var xxx_messageInfo_SoTransactionObject proto.InternalMessageInfo

func (m *SoTransactionObject) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func (m *SoTransactionObject) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type SoListTransactionObjectByExpiration struct {
	Expiration           *prototype.TimePointSec `protobuf:"bytes,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
	TrxId                *prototype.Sha256       `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListTransactionObjectByExpiration) Reset()         { *m = SoListTransactionObjectByExpiration{} }
func (m *SoListTransactionObjectByExpiration) String() string { return proto.CompactTextString(m) }
func (*SoListTransactionObjectByExpiration) ProtoMessage()    {}
func (*SoListTransactionObjectByExpiration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e5dbc6d4c15723, []int{1}
}

func (m *SoListTransactionObjectByExpiration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Unmarshal(m, b)
}
func (m *SoListTransactionObjectByExpiration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Marshal(b, m, deterministic)
}
func (m *SoListTransactionObjectByExpiration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListTransactionObjectByExpiration.Merge(m, src)
}
func (m *SoListTransactionObjectByExpiration) XXX_Size() int {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Size(m)
}
func (m *SoListTransactionObjectByExpiration) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListTransactionObjectByExpiration.DiscardUnknown(m)
}

var xxx_messageInfo_SoListTransactionObjectByExpiration proto.InternalMessageInfo

func (m *SoListTransactionObjectByExpiration) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *SoListTransactionObjectByExpiration) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoUniqueTransactionObjectByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoUniqueTransactionObjectByTrxId) Reset()         { *m = SoUniqueTransactionObjectByTrxId{} }
func (m *SoUniqueTransactionObjectByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueTransactionObjectByTrxId) ProtoMessage()    {}
func (*SoUniqueTransactionObjectByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e5dbc6d4c15723, []int{2}
}

func (m *SoUniqueTransactionObjectByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Unmarshal(m, b)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Marshal(b, m, deterministic)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Merge(m, src)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Size(m)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueTransactionObjectByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueTransactionObjectByTrxId proto.InternalMessageInfo

func (m *SoUniqueTransactionObjectByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func init() {
	proto.RegisterType((*SoTransactionObject)(nil), "table.so_transactionObject")
	proto.RegisterType((*SoListTransactionObjectByExpiration)(nil), "table.so_list_transactionObject_by_expiration")
	proto.RegisterType((*SoUniqueTransactionObjectByTrxId)(nil), "table.so_unique_transactionObject_by_trx_id")
}

func init() {
	proto.RegisterFile("app/table/so_transactionObject.proto", fileDescriptor_b0e5dbc6d4c15723)
}

var fileDescriptor_b0e5dbc6d4c15723 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x42, 0x7b, 0x88, 0x27, 0x97, 0x1e, 0xaa, 0x27, 0x59, 0xfc, 0xb3, 0x17, 0x37,
	0x50, 0x51, 0xf0, 0xea, 0xcd, 0x93, 0xd8, 0xa3, 0x97, 0x90, 0xa4, 0x43, 0x77, 0xa4, 0xcd, 0xc4,
	0xcc, 0x2c, 0x6c, 0xf1, 0x3b, 0xf8, 0x99, 0xa5, 0xbb, 0xa0, 0x8b, 0x16, 0xb1, 0x97, 0x40, 0xc8,
	0xef, 0xcd, 0x7b, 0x2f, 0xa3, 0xce, 0x6d, 0x8c, 0x5a, 0xac, 0x5b, 0x83, 0x66, 0x32, 0x92, 0x6c,
	0x60, 0xeb, 0x05, 0x29, 0x3c, 0xb9, 0x57, 0xf0, 0x52, 0xc5, 0x44, 0x42, 0xf9, 0xb8, 0x23, 0x4e,
	0xa7, 0xdd, 0x4d, 0xb6, 0x11, 0xf4, 0xee, 0xe8, 0x1f, 0x8b, 0x77, 0x35, 0xdd, 0x27, 0xcd, 0x4b,
	0x35, 0x91, 0xd4, 0x1a, 0x5c, 0xce, 0xb2, 0xb3, 0xac, 0x3c, 0x9a, 0x1f, 0x57, 0x5f, 0xf2, 0x8a,
	0x6b, 0x3b, 0xbf, 0xbd, 0x5b, 0x8c, 0x25, 0xb5, 0x8f, 0xcb, 0xfc, 0x5e, 0x29, 0x68, 0x23, 0x26,
	0xbb, 0x53, 0xcf, 0x46, 0x1d, 0x7d, 0x32, 0xa0, 0x05, 0x37, 0x60, 0x22, 0x61, 0x10, 0xc3, 0xe0,
	0x17, 0x03, 0xb8, 0xf8, 0xc8, 0xd4, 0x15, 0x93, 0x59, 0x23, 0xcb, 0xef, 0x08, 0xc6, 0x6d, 0xcd,
	0x37, 0xfb, 0xc3, 0x26, 0x3b, 0xc0, 0x66, 0xd0, 0x65, 0xf4, 0x77, 0x97, 0xe2, 0x59, 0x5d, 0x30,
	0x99, 0x26, 0xe0, 0x5b, 0x03, 0xfb, 0x13, 0xf5, 0x83, 0xfe, 0xff, 0x3d, 0x0f, 0xe5, 0xcb, 0xe5,
	0x0a, 0xa5, 0x6e, 0x5c, 0xe5, 0x69, 0xa3, 0x3d, 0xb1, 0xaf, 0x2d, 0x06, 0xed, 0x29, 0x08, 0x04,
	0x21, 0xbe, 0x5e, 0x51, 0xbf, 0x42, 0x37, 0xe9, 0x46, 0xdc, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x16, 0x14, 0x70, 0xd6, 0x01, 0x00, 0x00,
}