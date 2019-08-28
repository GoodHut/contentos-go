// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_blockProducer.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SoBlockProducer struct {
	Owner                *prototype.AccountName   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	CreatedTime          *prototype.TimePointSec  `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Url                  string                   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	BpVest               *prototype.BpVestId      `protobuf:"bytes,4,opt,name=bp_vest,json=bpVest,proto3" json:"bp_vest,omitempty"`
	SigningKey           *prototype.PublicKeyType `protobuf:"bytes,5,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	ProposedStaminaFree  uint64                   `protobuf:"varint,6,opt,name=proposed_stamina_free,json=proposedStaminaFree,proto3" json:"proposed_stamina_free,omitempty"`
	TpsExpected          uint64                   `protobuf:"varint,7,opt,name=tps_expected,json=tpsExpected,proto3" json:"tps_expected,omitempty"`
	AccountCreateFee     *prototype.Coin          `protobuf:"bytes,8,opt,name=account_create_fee,json=accountCreateFee,proto3" json:"account_create_fee,omitempty"`
	TopNAcquireFreeToken uint32                   `protobuf:"varint,9,opt,name=top_n_acquire_free_token,json=topNAcquireFreeToken,proto3" json:"top_n_acquire_free_token,omitempty"`
	EpochDuration        uint64                   `protobuf:"varint,10,opt,name=epoch_duration,json=epochDuration,proto3" json:"epoch_duration,omitempty"`
	PerTicketPrice       *prototype.Coin          `protobuf:"bytes,11,opt,name=per_ticket_price,json=perTicketPrice,proto3" json:"per_ticket_price,omitempty"`
	PerTicketWeight      uint64                   `protobuf:"varint,12,opt,name=per_ticket_weight,json=perTicketWeight,proto3" json:"per_ticket_weight,omitempty"`
	VoterCount           uint64                   `protobuf:"varint,13,opt,name=voter_count,json=voterCount,proto3" json:"voter_count,omitempty"`
	GenBlockCount        uint64                   `protobuf:"varint,14,opt,name=gen_block_count,json=genBlockCount,proto3" json:"gen_block_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SoBlockProducer) Reset()         { *m = SoBlockProducer{} }
func (m *SoBlockProducer) String() string { return proto.CompactTextString(m) }
func (*SoBlockProducer) ProtoMessage()    {}
func (*SoBlockProducer) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b6b3845d93ede7, []int{0}
}

func (m *SoBlockProducer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoBlockProducer.Unmarshal(m, b)
}
func (m *SoBlockProducer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoBlockProducer.Marshal(b, m, deterministic)
}
func (m *SoBlockProducer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoBlockProducer.Merge(m, src)
}
func (m *SoBlockProducer) XXX_Size() int {
	return xxx_messageInfo_SoBlockProducer.Size(m)
}
func (m *SoBlockProducer) XXX_DiscardUnknown() {
	xxx_messageInfo_SoBlockProducer.DiscardUnknown(m)
}

var xxx_messageInfo_SoBlockProducer proto.InternalMessageInfo

func (m *SoBlockProducer) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *SoBlockProducer) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoBlockProducer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SoBlockProducer) GetBpVest() *prototype.BpVestId {
	if m != nil {
		return m.BpVest
	}
	return nil
}

func (m *SoBlockProducer) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

func (m *SoBlockProducer) GetProposedStaminaFree() uint64 {
	if m != nil {
		return m.ProposedStaminaFree
	}
	return 0
}

func (m *SoBlockProducer) GetTpsExpected() uint64 {
	if m != nil {
		return m.TpsExpected
	}
	return 0
}

func (m *SoBlockProducer) GetAccountCreateFee() *prototype.Coin {
	if m != nil {
		return m.AccountCreateFee
	}
	return nil
}

func (m *SoBlockProducer) GetTopNAcquireFreeToken() uint32 {
	if m != nil {
		return m.TopNAcquireFreeToken
	}
	return 0
}

func (m *SoBlockProducer) GetEpochDuration() uint64 {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *SoBlockProducer) GetPerTicketPrice() *prototype.Coin {
	if m != nil {
		return m.PerTicketPrice
	}
	return nil
}

func (m *SoBlockProducer) GetPerTicketWeight() uint64 {
	if m != nil {
		return m.PerTicketWeight
	}
	return 0
}

func (m *SoBlockProducer) GetVoterCount() uint64 {
	if m != nil {
		return m.VoterCount
	}
	return 0
}

func (m *SoBlockProducer) GetGenBlockCount() uint64 {
	if m != nil {
		return m.GenBlockCount
	}
	return 0
}

type SoListBlockProducerByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListBlockProducerByOwner) Reset()         { *m = SoListBlockProducerByOwner{} }
func (m *SoListBlockProducerByOwner) String() string { return proto.CompactTextString(m) }
func (*SoListBlockProducerByOwner) ProtoMessage()    {}
func (*SoListBlockProducerByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b6b3845d93ede7, []int{1}
}

func (m *SoListBlockProducerByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListBlockProducerByOwner.Unmarshal(m, b)
}
func (m *SoListBlockProducerByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListBlockProducerByOwner.Marshal(b, m, deterministic)
}
func (m *SoListBlockProducerByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListBlockProducerByOwner.Merge(m, src)
}
func (m *SoListBlockProducerByOwner) XXX_Size() int {
	return xxx_messageInfo_SoListBlockProducerByOwner.Size(m)
}
func (m *SoListBlockProducerByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListBlockProducerByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoListBlockProducerByOwner proto.InternalMessageInfo

func (m *SoListBlockProducerByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoListBlockProducerByBpVest struct {
	BpVest               *prototype.BpVestId    `protobuf:"bytes,1,opt,name=bp_vest,json=bpVest,proto3" json:"bp_vest,omitempty"`
	Owner                *prototype.AccountName `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListBlockProducerByBpVest) Reset()         { *m = SoListBlockProducerByBpVest{} }
func (m *SoListBlockProducerByBpVest) String() string { return proto.CompactTextString(m) }
func (*SoListBlockProducerByBpVest) ProtoMessage()    {}
func (*SoListBlockProducerByBpVest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b6b3845d93ede7, []int{2}
}

func (m *SoListBlockProducerByBpVest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListBlockProducerByBpVest.Unmarshal(m, b)
}
func (m *SoListBlockProducerByBpVest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListBlockProducerByBpVest.Marshal(b, m, deterministic)
}
func (m *SoListBlockProducerByBpVest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListBlockProducerByBpVest.Merge(m, src)
}
func (m *SoListBlockProducerByBpVest) XXX_Size() int {
	return xxx_messageInfo_SoListBlockProducerByBpVest.Size(m)
}
func (m *SoListBlockProducerByBpVest) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListBlockProducerByBpVest.DiscardUnknown(m)
}

var xxx_messageInfo_SoListBlockProducerByBpVest proto.InternalMessageInfo

func (m *SoListBlockProducerByBpVest) GetBpVest() *prototype.BpVestId {
	if m != nil {
		return m.BpVest
	}
	return nil
}

func (m *SoListBlockProducerByBpVest) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoUniqueBlockProducerByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoUniqueBlockProducerByOwner) Reset()         { *m = SoUniqueBlockProducerByOwner{} }
func (m *SoUniqueBlockProducerByOwner) String() string { return proto.CompactTextString(m) }
func (*SoUniqueBlockProducerByOwner) ProtoMessage()    {}
func (*SoUniqueBlockProducerByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b6b3845d93ede7, []int{3}
}

func (m *SoUniqueBlockProducerByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueBlockProducerByOwner.Unmarshal(m, b)
}
func (m *SoUniqueBlockProducerByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueBlockProducerByOwner.Marshal(b, m, deterministic)
}
func (m *SoUniqueBlockProducerByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueBlockProducerByOwner.Merge(m, src)
}
func (m *SoUniqueBlockProducerByOwner) XXX_Size() int {
	return xxx_messageInfo_SoUniqueBlockProducerByOwner.Size(m)
}
func (m *SoUniqueBlockProducerByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueBlockProducerByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueBlockProducerByOwner proto.InternalMessageInfo

func (m *SoUniqueBlockProducerByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*SoBlockProducer)(nil), "table.so_blockProducer")
	proto.RegisterType((*SoListBlockProducerByOwner)(nil), "table.so_list_blockProducer_by_owner")
	proto.RegisterType((*SoListBlockProducerByBpVest)(nil), "table.so_list_blockProducer_by_bp_vest")
	proto.RegisterType((*SoUniqueBlockProducerByOwner)(nil), "table.so_unique_blockProducer_by_owner")
}

func init() { proto.RegisterFile("app/table/so_blockProducer.proto", fileDescriptor_44b6b3845d93ede7) }

var fileDescriptor_44b6b3845d93ede7 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6b, 0x13, 0x4f,
	0x14, 0xc5, 0xd9, 0xb6, 0x69, 0xbf, 0x9d, 0xf4, 0x47, 0xbe, 0x63, 0x8b, 0x63, 0x1f, 0x34, 0x16,
	0x94, 0x50, 0x6c, 0x02, 0x15, 0x04, 0x51, 0x1f, 0x6c, 0xb5, 0x2f, 0x82, 0xd6, 0x35, 0x28, 0xf8,
	0x72, 0xd9, 0x9d, 0xbd, 0xdd, 0x0c, 0xc9, 0xce, 0x4c, 0x67, 0xee, 0xb6, 0xe6, 0xd1, 0xbf, 0xc5,
	0x7f, 0x54, 0x76, 0x76, 0x1b, 0x57, 0x51, 0xa8, 0xe0, 0x4b, 0xc8, 0x9c, 0xf3, 0xb9, 0x77, 0xcf,
	0x2c, 0x87, 0x65, 0xfd, 0xc4, 0xda, 0x11, 0x25, 0xe9, 0x0c, 0x47, 0xde, 0x40, 0x3a, 0x33, 0x72,
	0x7a, 0xe6, 0x4c, 0x56, 0x4a, 0x74, 0x43, 0xeb, 0x0c, 0x19, 0xde, 0x09, 0xee, 0xde, 0x4e, 0x38,
	0xd1, 0xdc, 0xe2, 0xa8, 0xfa, 0xa9, 0xcd, 0x3d, 0xf1, 0x43, 0x2d, 0xca, 0x19, 0x29, 0x50, 0x59,
	0xed, 0xec, 0x7f, 0xeb, 0xb0, 0xde, 0xaf, 0x1b, 0xf9, 0x21, 0xeb, 0x98, 0x2b, 0x8d, 0x4e, 0x44,
	0xfd, 0x68, 0xd0, 0x3d, 0xba, 0x3d, 0x5c, 0x8c, 0x0f, 0x13, 0x29, 0x4d, 0xa9, 0x09, 0x74, 0x52,
	0x60, 0x5c, 0x53, 0xfc, 0x39, 0xdb, 0x90, 0x0e, 0x13, 0xc2, 0x0c, 0x48, 0x15, 0x28, 0x96, 0xc2,
	0xd4, 0x9d, 0xd6, 0x54, 0x25, 0x83, 0x35, 0x4a, 0x13, 0x78, 0x94, 0x71, 0xb7, 0xc1, 0xc7, 0xaa,
	0x40, 0xde, 0x63, 0xcb, 0xa5, 0x9b, 0x89, 0xe5, 0x7e, 0x34, 0x58, 0x8f, 0xab, 0xbf, 0x7c, 0xc8,
	0xd6, 0x52, 0x0b, 0x97, 0xe8, 0x49, 0xac, 0x84, 0x55, 0xbb, 0xad, 0x55, 0x8d, 0x03, 0x2a, 0x8b,
	0x57, 0x53, 0xfb, 0x11, 0x3d, 0xf1, 0x67, 0xac, 0xeb, 0x55, 0xae, 0x95, 0xce, 0x61, 0x8a, 0x73,
	0xd1, 0x09, 0x33, 0x7b, 0xad, 0x19, 0x5b, 0xa6, 0x33, 0x25, 0x2b, 0x13, 0xaa, 0x73, 0xcc, 0x1a,
	0xfc, 0x0d, 0xce, 0xf9, 0x11, 0xdb, 0xb5, 0xce, 0x58, 0xe3, 0x31, 0x03, 0x4f, 0x49, 0xa1, 0x74,
	0x02, 0xe7, 0x0e, 0x51, 0xac, 0xf6, 0xa3, 0xc1, 0x4a, 0x7c, 0xeb, 0xda, 0xfc, 0x50, 0x7b, 0xa7,
	0x0e, 0x91, 0xdf, 0x67, 0x1b, 0x64, 0x3d, 0xe0, 0x17, 0x8b, 0x92, 0x30, 0x13, 0x6b, 0x01, 0xed,
	0x92, 0xf5, 0xaf, 0x1b, 0x89, 0xbf, 0x60, 0xfc, 0xfa, 0x55, 0xd5, 0x97, 0x85, 0x73, 0x44, 0xf1,
	0x5f, 0x88, 0xb6, 0xdd, 0x8a, 0x26, 0x8d, 0xd2, 0x71, 0xaf, 0x41, 0x4f, 0x02, 0x79, 0x8a, 0xc8,
	0x9f, 0x30, 0x41, 0xc6, 0x82, 0x86, 0x44, 0x5e, 0x94, 0xca, 0x61, 0x88, 0x04, 0x64, 0xa6, 0xa8,
	0xc5, 0x7a, 0x3f, 0x1a, 0x6c, 0xc6, 0x3b, 0x64, 0xec, 0xdb, 0x97, 0xb5, 0x5b, 0x85, 0x1a, 0x57,
	0x1e, 0x7f, 0xc0, 0xb6, 0xd0, 0x1a, 0x39, 0x81, 0xac, 0x74, 0x09, 0x29, 0xa3, 0x05, 0x0b, 0xd9,
	0x36, 0x83, 0xfa, 0xaa, 0x11, 0xf9, 0x53, 0xd6, 0xb3, 0xe8, 0x80, 0x94, 0x9c, 0x22, 0x81, 0x75,
	0x4a, 0xa2, 0xe8, 0xfe, 0x3e, 0xdb, 0x96, 0x45, 0x37, 0x0e, 0xdc, 0x59, 0x85, 0xf1, 0x03, 0xf6,
	0x7f, 0x6b, 0xf4, 0x0a, 0x55, 0x3e, 0x21, 0xb1, 0x11, 0x1e, 0xb2, 0xbd, 0x40, 0x3f, 0x05, 0x99,
	0xdf, 0x63, 0xdd, 0x4b, 0x43, 0xe8, 0x20, 0xdc, 0x4e, 0x6c, 0x06, 0x8a, 0x05, 0xe9, 0xa4, 0x52,
	0xf8, 0x43, 0xb6, 0x9d, 0xa3, 0xae, 0xdb, 0xd7, 0x40, 0x5b, 0x75, 0xde, 0x1c, 0xf5, 0x71, 0xa5,
	0x06, 0x6e, 0xff, 0x1d, 0xbb, 0xeb, 0x0d, 0xcc, 0x94, 0xa7, 0x9f, 0x9b, 0x0a, 0xe9, 0x1c, 0xea,
	0x0e, 0xfe, 0x5d, 0x65, 0xf7, 0xbf, 0x46, 0xac, 0xff, 0xc7, 0x8d, 0x4d, 0xc5, 0xda, 0x3d, 0x8c,
	0x6e, 0xd2, 0xc3, 0x45, 0x86, 0xa5, 0x1b, 0x65, 0x78, 0x1f, 0x22, 0x94, 0x5a, 0x5d, 0x94, 0xf8,
	0x6f, 0xae, 0x75, 0xfc, 0xe8, 0xf3, 0x41, 0xae, 0x68, 0x52, 0xa6, 0x43, 0x69, 0x8a, 0x91, 0x34,
	0x5e, 0x4e, 0x12, 0xa5, 0x47, 0xd2, 0x68, 0x42, 0x4d, 0xc6, 0x1f, 0xe6, 0x66, 0xb4, 0xf8, 0x92,
	0xa4, 0xab, 0x61, 0xd9, 0xe3, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x11, 0x09, 0x65, 0x5d,
	0x04, 0x00, 0x00,
}
