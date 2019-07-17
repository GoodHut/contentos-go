// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototype/type.proto

package prototype

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccountName struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountName) Reset()         { *m = AccountName{} }
func (m *AccountName) String() string { return proto.CompactTextString(m) }
func (*AccountName) ProtoMessage()    {}
func (*AccountName) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{0}
}

func (m *AccountName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountName.Unmarshal(m, b)
}
func (m *AccountName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountName.Marshal(b, m, deterministic)
}
func (m *AccountName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountName.Merge(m, src)
}
func (m *AccountName) XXX_Size() int {
	return xxx_messageInfo_AccountName.Size(m)
}
func (m *AccountName) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountName.DiscardUnknown(m)
}

var xxx_messageInfo_AccountName proto.InternalMessageInfo

func (m *AccountName) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ChainId struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainId) Reset()         { *m = ChainId{} }
func (m *ChainId) String() string { return proto.CompactTextString(m) }
func (*ChainId) ProtoMessage()    {}
func (*ChainId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{1}
}

func (m *ChainId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainId.Unmarshal(m, b)
}
func (m *ChainId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainId.Marshal(b, m, deterministic)
}
func (m *ChainId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainId.Merge(m, src)
}
func (m *ChainId) XXX_Size() int {
	return xxx_messageInfo_ChainId.Size(m)
}
func (m *ChainId) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainId.DiscardUnknown(m)
}

var xxx_messageInfo_ChainId proto.InternalMessageInfo

func (m *ChainId) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Coin struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coin) Reset()         { *m = Coin{} }
func (m *Coin) String() string { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()    {}
func (*Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{2}
}

func (m *Coin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coin.Unmarshal(m, b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
}
func (m *Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coin.Merge(m, src)
}
func (m *Coin) XXX_Size() int {
	return xxx_messageInfo_Coin.Size(m)
}
func (m *Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

func (m *Coin) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Vest struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vest) Reset()         { *m = Vest{} }
func (m *Vest) String() string { return proto.CompactTextString(m) }
func (*Vest) ProtoMessage()    {}
func (*Vest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{3}
}

func (m *Vest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vest.Unmarshal(m, b)
}
func (m *Vest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vest.Marshal(b, m, deterministic)
}
func (m *Vest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vest.Merge(m, src)
}
func (m *Vest) XXX_Size() int {
	return xxx_messageInfo_Vest.Size(m)
}
func (m *Vest) XXX_DiscardUnknown() {
	xxx_messageInfo_Vest.DiscardUnknown(m)
}

var xxx_messageInfo_Vest proto.InternalMessageInfo

func (m *Vest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PublicKeyType struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKeyType) Reset()         { *m = PublicKeyType{} }
func (m *PublicKeyType) String() string { return proto.CompactTextString(m) }
func (*PublicKeyType) ProtoMessage()    {}
func (*PublicKeyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{4}
}

func (m *PublicKeyType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeyType.Unmarshal(m, b)
}
func (m *PublicKeyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeyType.Marshal(b, m, deterministic)
}
func (m *PublicKeyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeyType.Merge(m, src)
}
func (m *PublicKeyType) XXX_Size() int {
	return xxx_messageInfo_PublicKeyType.Size(m)
}
func (m *PublicKeyType) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeyType.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeyType proto.InternalMessageInfo

func (m *PublicKeyType) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PrivateKeyType struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateKeyType) Reset()         { *m = PrivateKeyType{} }
func (m *PrivateKeyType) String() string { return proto.CompactTextString(m) }
func (*PrivateKeyType) ProtoMessage()    {}
func (*PrivateKeyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{5}
}

func (m *PrivateKeyType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeyType.Unmarshal(m, b)
}
func (m *PrivateKeyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeyType.Marshal(b, m, deterministic)
}
func (m *PrivateKeyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeyType.Merge(m, src)
}
func (m *PrivateKeyType) XXX_Size() int {
	return xxx_messageInfo_PrivateKeyType.Size(m)
}
func (m *PrivateKeyType) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeyType.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeyType proto.InternalMessageInfo

func (m *PrivateKeyType) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TimePointSec struct {
	UtcSeconds           uint32   `protobuf:"varint,1,opt,name=utc_seconds,json=utcSeconds,proto3" json:"utc_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimePointSec) Reset()         { *m = TimePointSec{} }
func (m *TimePointSec) String() string { return proto.CompactTextString(m) }
func (*TimePointSec) ProtoMessage()    {}
func (*TimePointSec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{6}
}

func (m *TimePointSec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimePointSec.Unmarshal(m, b)
}
func (m *TimePointSec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimePointSec.Marshal(b, m, deterministic)
}
func (m *TimePointSec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimePointSec.Merge(m, src)
}
func (m *TimePointSec) XXX_Size() int {
	return xxx_messageInfo_TimePointSec.Size(m)
}
func (m *TimePointSec) XXX_DiscardUnknown() {
	xxx_messageInfo_TimePointSec.DiscardUnknown(m)
}

var xxx_messageInfo_TimePointSec proto.InternalMessageInfo

func (m *TimePointSec) GetUtcSeconds() uint32 {
	if m != nil {
		return m.UtcSeconds
	}
	return 0
}

type SignatureType struct {
	Sig                  []byte   `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureType) Reset()         { *m = SignatureType{} }
func (m *SignatureType) String() string { return proto.CompactTextString(m) }
func (*SignatureType) ProtoMessage()    {}
func (*SignatureType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{7}
}

func (m *SignatureType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureType.Unmarshal(m, b)
}
func (m *SignatureType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureType.Marshal(b, m, deterministic)
}
func (m *SignatureType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureType.Merge(m, src)
}
func (m *SignatureType) XXX_Size() int {
	return xxx_messageInfo_SignatureType.Size(m)
}
func (m *SignatureType) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureType.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureType proto.InternalMessageInfo

func (m *SignatureType) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type Sha256 struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sha256) Reset()         { *m = Sha256{} }
func (m *Sha256) String() string { return proto.CompactTextString(m) }
func (*Sha256) ProtoMessage()    {}
func (*Sha256) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{8}
}

func (m *Sha256) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sha256.Unmarshal(m, b)
}
func (m *Sha256) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sha256.Marshal(b, m, deterministic)
}
func (m *Sha256) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sha256.Merge(m, src)
}
func (m *Sha256) XXX_Size() int {
	return xxx_messageInfo_Sha256.Size(m)
}
func (m *Sha256) XXX_DiscardUnknown() {
	xxx_messageInfo_Sha256.DiscardUnknown(m)
}

var xxx_messageInfo_Sha256 proto.InternalMessageInfo

func (m *Sha256) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type GiftTicketKeyType struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	CreateBlock          uint64   `protobuf:"varint,4,opt,name=create_block,json=createBlock,proto3" json:"create_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GiftTicketKeyType) Reset()         { *m = GiftTicketKeyType{} }
func (m *GiftTicketKeyType) String() string { return proto.CompactTextString(m) }
func (*GiftTicketKeyType) ProtoMessage()    {}
func (*GiftTicketKeyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{9}
}

func (m *GiftTicketKeyType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GiftTicketKeyType.Unmarshal(m, b)
}
func (m *GiftTicketKeyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GiftTicketKeyType.Marshal(b, m, deterministic)
}
func (m *GiftTicketKeyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GiftTicketKeyType.Merge(m, src)
}
func (m *GiftTicketKeyType) XXX_Size() int {
	return xxx_messageInfo_GiftTicketKeyType.Size(m)
}
func (m *GiftTicketKeyType) XXX_DiscardUnknown() {
	xxx_messageInfo_GiftTicketKeyType.DiscardUnknown(m)
}

var xxx_messageInfo_GiftTicketKeyType proto.InternalMessageInfo

func (m *GiftTicketKeyType) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GiftTicketKeyType) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GiftTicketKeyType) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *GiftTicketKeyType) GetCreateBlock() uint64 {
	if m != nil {
		return m.CreateBlock
	}
	return 0
}

type ChainProperties struct {
	AccountCreationFee   *Coin    `protobuf:"bytes,1,opt,name=account_creation_fee,json=accountCreationFee,proto3" json:"account_creation_fee,omitempty"`
	MaximumBlockSize     uint32   `protobuf:"varint,2,opt,name=maximum_block_size,json=maximumBlockSize,proto3" json:"maximum_block_size,omitempty"`
	StaminaFree          uint64   `protobuf:"varint,3,opt,name=stamina_free,json=staminaFree,proto3" json:"stamina_free,omitempty"`
	TpsExpected          uint64   `protobuf:"varint,4,opt,name=tps_expected,json=tpsExpected,proto3" json:"tps_expected,omitempty"`
	TopNAcquireFreeToken uint32   `protobuf:"varint,5,opt,name=top_n_acquire_free_token,json=topNAcquireFreeToken,proto3" json:"top_n_acquire_free_token,omitempty"`
	EpochDuration        uint64   `protobuf:"varint,6,opt,name=epoch_duration,json=epochDuration,proto3" json:"epoch_duration,omitempty"`
	PerTicketPrice       *Coin    `protobuf:"bytes,7,opt,name=per_ticket_price,json=perTicketPrice,proto3" json:"per_ticket_price,omitempty"`
	PerTicketWeight      uint64   `protobuf:"varint,8,opt,name=per_ticket_weight,json=perTicketWeight,proto3" json:"per_ticket_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainProperties) Reset()         { *m = ChainProperties{} }
func (m *ChainProperties) String() string { return proto.CompactTextString(m) }
func (*ChainProperties) ProtoMessage()    {}
func (*ChainProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{10}
}

func (m *ChainProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainProperties.Unmarshal(m, b)
}
func (m *ChainProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainProperties.Marshal(b, m, deterministic)
}
func (m *ChainProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainProperties.Merge(m, src)
}
func (m *ChainProperties) XXX_Size() int {
	return xxx_messageInfo_ChainProperties.Size(m)
}
func (m *ChainProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainProperties.DiscardUnknown(m)
}

var xxx_messageInfo_ChainProperties proto.InternalMessageInfo

func (m *ChainProperties) GetAccountCreationFee() *Coin {
	if m != nil {
		return m.AccountCreationFee
	}
	return nil
}

func (m *ChainProperties) GetMaximumBlockSize() uint32 {
	if m != nil {
		return m.MaximumBlockSize
	}
	return 0
}

func (m *ChainProperties) GetStaminaFree() uint64 {
	if m != nil {
		return m.StaminaFree
	}
	return 0
}

func (m *ChainProperties) GetTpsExpected() uint64 {
	if m != nil {
		return m.TpsExpected
	}
	return 0
}

func (m *ChainProperties) GetTopNAcquireFreeToken() uint32 {
	if m != nil {
		return m.TopNAcquireFreeToken
	}
	return 0
}

func (m *ChainProperties) GetEpochDuration() uint64 {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *ChainProperties) GetPerTicketPrice() *Coin {
	if m != nil {
		return m.PerTicketPrice
	}
	return nil
}

func (m *ChainProperties) GetPerTicketWeight() uint64 {
	if m != nil {
		return m.PerTicketWeight
	}
	return 0
}

type DynamicProperties struct {
	HeadBlockId            *Sha256       `protobuf:"bytes,2,opt,name=head_block_id,json=headBlockId,proto3" json:"head_block_id,omitempty"`
	HeadBlockNumber        uint64        `protobuf:"varint,3,opt,name=head_block_number,json=headBlockNumber,proto3" json:"head_block_number,omitempty"`
	MaximumBlockSize       uint32        `protobuf:"varint,4,opt,name=maximum_block_size,json=maximumBlockSize,proto3" json:"maximum_block_size,omitempty"`
	TotalCos               *Coin         `protobuf:"bytes,5,opt,name=total_cos,json=totalCos,proto3" json:"total_cos,omitempty"`
	Time                   *TimePointSec `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	CurrentWitness         *AccountName  `protobuf:"bytes,7,opt,name=current_witness,json=currentWitness,proto3" json:"current_witness,omitempty"`
	Tps                    uint32        `protobuf:"varint,9,opt,name=tps,proto3" json:"tps,omitempty"`
	TotalVestingShares     *Vest         `protobuf:"bytes,10,opt,name=total_vesting_shares,json=totalVestingShares,proto3" json:"total_vesting_shares,omitempty"`
	CurrentSupply          *Coin         `protobuf:"bytes,11,opt,name=current_supply,json=currentSupply,proto3" json:"current_supply,omitempty"`
	PostWeightedVps        string        `protobuf:"bytes,13,opt,name=post_weighted_vps,json=postWeightedVps,proto3" json:"post_weighted_vps,omitempty"`
	PostRewards            *Vest         `protobuf:"bytes,14,opt,name=post_rewards,json=postRewards,proto3" json:"post_rewards,omitempty"`
	TotalTrxCnt            uint64        `protobuf:"varint,15,opt,name=total_trx_cnt,json=totalTrxCnt,proto3" json:"total_trx_cnt,omitempty"`
	TotalPostCnt           uint64        `protobuf:"varint,16,opt,name=total_post_cnt,json=totalPostCnt,proto3" json:"total_post_cnt,omitempty"`
	TotalUserCnt           uint64        `protobuf:"varint,17,opt,name=total_user_cnt,json=totalUserCnt,proto3" json:"total_user_cnt,omitempty"`
	MaxTps                 uint32        `protobuf:"varint,18,opt,name=max_tps,json=maxTps,proto3" json:"max_tps,omitempty"`
	MaxTpsBlockNum         uint64        `protobuf:"varint,19,opt,name=max_tps_block_num,json=maxTpsBlockNum,proto3" json:"max_tps_block_num,omitempty"`
	HeadBlockPrefix        uint32        `protobuf:"varint,20,opt,name=head_block_prefix,json=headBlockPrefix,proto3" json:"head_block_prefix,omitempty"`
	ReportRewards          *Vest         `protobuf:"bytes,21,opt,name=report_rewards,json=reportRewards,proto3" json:"report_rewards,omitempty"`
	IthYear                uint32        `protobuf:"varint,22,opt,name=ith_year,json=ithYear,proto3" json:"ith_year,omitempty"`
	AnnualBudget           *Vest         `protobuf:"bytes,23,opt,name=annual_budget,json=annualBudget,proto3" json:"annual_budget,omitempty"`
	AnnualMinted           *Vest         `protobuf:"bytes,24,opt,name=annual_minted,json=annualMinted,proto3" json:"annual_minted,omitempty"`
	PostDappRewards        *Vest         `protobuf:"bytes,25,opt,name=post_dapp_rewards,json=postDappRewards,proto3" json:"post_dapp_rewards,omitempty"`
	VoterRewards           *Vest         `protobuf:"bytes,26,opt,name=voter_rewards,json=voterRewards,proto3" json:"voter_rewards,omitempty"`
	ReplyRewards           *Vest         `protobuf:"bytes,27,opt,name=reply_rewards,json=replyRewards,proto3" json:"reply_rewards,omitempty"`
	ReplyWeightedVps       string        `protobuf:"bytes,28,opt,name=reply_weighted_vps,json=replyWeightedVps,proto3" json:"reply_weighted_vps,omitempty"`
	ReplyDappRewards       *Vest         `protobuf:"bytes,29,opt,name=reply_dapp_rewards,json=replyDappRewards,proto3" json:"reply_dapp_rewards,omitempty"`
	StakeVestingShares     *Vest         `protobuf:"bytes,30,opt,name=stake_vesting_shares,json=stakeVestingShares,proto3" json:"stake_vesting_shares,omitempty"`
	WitnessBootCompleted   bool          `protobuf:"varint,31,opt,name=witness_boot_completed,json=witnessBootCompleted,proto3" json:"witness_boot_completed,omitempty"`
	StaminaFree            uint64        `protobuf:"varint,32,opt,name=stamina_free,json=staminaFree,proto3" json:"stamina_free,omitempty"`
	TpsExpected            uint64        `protobuf:"varint,33,opt,name=tps_expected,json=tpsExpected,proto3" json:"tps_expected,omitempty"`
	AvgTpsUpdateBlock      uint64        `protobuf:"varint,34,opt,name=avg_tps_update_block,json=avgTpsUpdateBlock,proto3" json:"avg_tps_update_block,omitempty"`
	AvgTpsInWindow         uint64        `protobuf:"varint,35,opt,name=avg_tps_in_window,json=avgTpsInWindow,proto3" json:"avg_tps_in_window,omitempty"`
	OneDayStamina          uint64        `protobuf:"varint,36,opt,name=one_day_stamina,json=oneDayStamina,proto3" json:"one_day_stamina,omitempty"`
	AccountCreateFee       *Coin         `protobuf:"bytes,37,opt,name=account_create_fee,json=accountCreateFee,proto3" json:"account_create_fee,omitempty"`
	ReputationAdmin        *AccountName  `protobuf:"bytes,38,opt,name=reputation_admin,json=reputationAdmin,proto3" json:"reputation_admin,omitempty"`
	CurrentEpochStartBlock uint64        `protobuf:"varint,39,opt,name=current_epoch_start_block,json=currentEpochStartBlock,proto3" json:"current_epoch_start_block,omitempty"`
	EpochDuration          uint64        `protobuf:"varint,40,opt,name=epoch_duration,json=epochDuration,proto3" json:"epoch_duration,omitempty"`
	TopNAcquireFreeToken   uint32        `protobuf:"varint,41,opt,name=top_n_acquire_free_token,json=topNAcquireFreeToken,proto3" json:"top_n_acquire_free_token,omitempty"`
	PerTicketPrice         *Coin         `protobuf:"bytes,42,opt,name=per_ticket_price,json=perTicketPrice,proto3" json:"per_ticket_price,omitempty"`
	PerTicketWeight        uint64        `protobuf:"varint,43,opt,name=per_ticket_weight,json=perTicketWeight,proto3" json:"per_ticket_weight,omitempty"`
	TicketsIncome          *Vest         `protobuf:"bytes,44,opt,name=tickets_income,json=ticketsIncome,proto3" json:"tickets_income,omitempty"`
	ChargedTicketsNum      uint64        `protobuf:"varint,45,opt,name=charged_tickets_num,json=chargedTicketsNum,proto3" json:"charged_tickets_num,omitempty"`
	CopyrightAdmin         *AccountName  `protobuf:"bytes,46,opt,name=copyright_admin,json=copyrightAdmin,proto3" json:"copyright_admin,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}      `json:"-"`
	XXX_unrecognized       []byte        `json:"-"`
	XXX_sizecache          int32         `json:"-"`
}

func (m *DynamicProperties) Reset()         { *m = DynamicProperties{} }
func (m *DynamicProperties) String() string { return proto.CompactTextString(m) }
func (*DynamicProperties) ProtoMessage()    {}
func (*DynamicProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{11}
}

func (m *DynamicProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicProperties.Unmarshal(m, b)
}
func (m *DynamicProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicProperties.Marshal(b, m, deterministic)
}
func (m *DynamicProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicProperties.Merge(m, src)
}
func (m *DynamicProperties) XXX_Size() int {
	return xxx_messageInfo_DynamicProperties.Size(m)
}
func (m *DynamicProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicProperties.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicProperties proto.InternalMessageInfo

func (m *DynamicProperties) GetHeadBlockId() *Sha256 {
	if m != nil {
		return m.HeadBlockId
	}
	return nil
}

func (m *DynamicProperties) GetHeadBlockNumber() uint64 {
	if m != nil {
		return m.HeadBlockNumber
	}
	return 0
}

func (m *DynamicProperties) GetMaximumBlockSize() uint32 {
	if m != nil {
		return m.MaximumBlockSize
	}
	return 0
}

func (m *DynamicProperties) GetTotalCos() *Coin {
	if m != nil {
		return m.TotalCos
	}
	return nil
}

func (m *DynamicProperties) GetTime() *TimePointSec {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *DynamicProperties) GetCurrentWitness() *AccountName {
	if m != nil {
		return m.CurrentWitness
	}
	return nil
}

func (m *DynamicProperties) GetTps() uint32 {
	if m != nil {
		return m.Tps
	}
	return 0
}

func (m *DynamicProperties) GetTotalVestingShares() *Vest {
	if m != nil {
		return m.TotalVestingShares
	}
	return nil
}

func (m *DynamicProperties) GetCurrentSupply() *Coin {
	if m != nil {
		return m.CurrentSupply
	}
	return nil
}

func (m *DynamicProperties) GetPostWeightedVps() string {
	if m != nil {
		return m.PostWeightedVps
	}
	return ""
}

func (m *DynamicProperties) GetPostRewards() *Vest {
	if m != nil {
		return m.PostRewards
	}
	return nil
}

func (m *DynamicProperties) GetTotalTrxCnt() uint64 {
	if m != nil {
		return m.TotalTrxCnt
	}
	return 0
}

func (m *DynamicProperties) GetTotalPostCnt() uint64 {
	if m != nil {
		return m.TotalPostCnt
	}
	return 0
}

func (m *DynamicProperties) GetTotalUserCnt() uint64 {
	if m != nil {
		return m.TotalUserCnt
	}
	return 0
}

func (m *DynamicProperties) GetMaxTps() uint32 {
	if m != nil {
		return m.MaxTps
	}
	return 0
}

func (m *DynamicProperties) GetMaxTpsBlockNum() uint64 {
	if m != nil {
		return m.MaxTpsBlockNum
	}
	return 0
}

func (m *DynamicProperties) GetHeadBlockPrefix() uint32 {
	if m != nil {
		return m.HeadBlockPrefix
	}
	return 0
}

func (m *DynamicProperties) GetReportRewards() *Vest {
	if m != nil {
		return m.ReportRewards
	}
	return nil
}

func (m *DynamicProperties) GetIthYear() uint32 {
	if m != nil {
		return m.IthYear
	}
	return 0
}

func (m *DynamicProperties) GetAnnualBudget() *Vest {
	if m != nil {
		return m.AnnualBudget
	}
	return nil
}

func (m *DynamicProperties) GetAnnualMinted() *Vest {
	if m != nil {
		return m.AnnualMinted
	}
	return nil
}

func (m *DynamicProperties) GetPostDappRewards() *Vest {
	if m != nil {
		return m.PostDappRewards
	}
	return nil
}

func (m *DynamicProperties) GetVoterRewards() *Vest {
	if m != nil {
		return m.VoterRewards
	}
	return nil
}

func (m *DynamicProperties) GetReplyRewards() *Vest {
	if m != nil {
		return m.ReplyRewards
	}
	return nil
}

func (m *DynamicProperties) GetReplyWeightedVps() string {
	if m != nil {
		return m.ReplyWeightedVps
	}
	return ""
}

func (m *DynamicProperties) GetReplyDappRewards() *Vest {
	if m != nil {
		return m.ReplyDappRewards
	}
	return nil
}

func (m *DynamicProperties) GetStakeVestingShares() *Vest {
	if m != nil {
		return m.StakeVestingShares
	}
	return nil
}

func (m *DynamicProperties) GetWitnessBootCompleted() bool {
	if m != nil {
		return m.WitnessBootCompleted
	}
	return false
}

func (m *DynamicProperties) GetStaminaFree() uint64 {
	if m != nil {
		return m.StaminaFree
	}
	return 0
}

func (m *DynamicProperties) GetTpsExpected() uint64 {
	if m != nil {
		return m.TpsExpected
	}
	return 0
}

func (m *DynamicProperties) GetAvgTpsUpdateBlock() uint64 {
	if m != nil {
		return m.AvgTpsUpdateBlock
	}
	return 0
}

func (m *DynamicProperties) GetAvgTpsInWindow() uint64 {
	if m != nil {
		return m.AvgTpsInWindow
	}
	return 0
}

func (m *DynamicProperties) GetOneDayStamina() uint64 {
	if m != nil {
		return m.OneDayStamina
	}
	return 0
}

func (m *DynamicProperties) GetAccountCreateFee() *Coin {
	if m != nil {
		return m.AccountCreateFee
	}
	return nil
}

func (m *DynamicProperties) GetReputationAdmin() *AccountName {
	if m != nil {
		return m.ReputationAdmin
	}
	return nil
}

func (m *DynamicProperties) GetCurrentEpochStartBlock() uint64 {
	if m != nil {
		return m.CurrentEpochStartBlock
	}
	return 0
}

func (m *DynamicProperties) GetEpochDuration() uint64 {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *DynamicProperties) GetTopNAcquireFreeToken() uint32 {
	if m != nil {
		return m.TopNAcquireFreeToken
	}
	return 0
}

func (m *DynamicProperties) GetPerTicketPrice() *Coin {
	if m != nil {
		return m.PerTicketPrice
	}
	return nil
}

func (m *DynamicProperties) GetPerTicketWeight() uint64 {
	if m != nil {
		return m.PerTicketWeight
	}
	return 0
}

func (m *DynamicProperties) GetTicketsIncome() *Vest {
	if m != nil {
		return m.TicketsIncome
	}
	return nil
}

func (m *DynamicProperties) GetChargedTicketsNum() uint64 {
	if m != nil {
		return m.ChargedTicketsNum
	}
	return 0
}

func (m *DynamicProperties) GetCopyrightAdmin() *AccountName {
	if m != nil {
		return m.CopyrightAdmin
	}
	return nil
}

type BeneficiaryRouteType struct {
	Name                 *AccountName `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               uint32       `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BeneficiaryRouteType) Reset()         { *m = BeneficiaryRouteType{} }
func (m *BeneficiaryRouteType) String() string { return proto.CompactTextString(m) }
func (*BeneficiaryRouteType) ProtoMessage()    {}
func (*BeneficiaryRouteType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b10af7c504b1c5, []int{12}
}

func (m *BeneficiaryRouteType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeneficiaryRouteType.Unmarshal(m, b)
}
func (m *BeneficiaryRouteType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeneficiaryRouteType.Marshal(b, m, deterministic)
}
func (m *BeneficiaryRouteType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeneficiaryRouteType.Merge(m, src)
}
func (m *BeneficiaryRouteType) XXX_Size() int {
	return xxx_messageInfo_BeneficiaryRouteType.Size(m)
}
func (m *BeneficiaryRouteType) XXX_DiscardUnknown() {
	xxx_messageInfo_BeneficiaryRouteType.DiscardUnknown(m)
}

var xxx_messageInfo_BeneficiaryRouteType proto.InternalMessageInfo

func (m *BeneficiaryRouteType) GetName() *AccountName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BeneficiaryRouteType) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountName)(nil), "prototype.account_name")
	proto.RegisterType((*ChainId)(nil), "prototype.chain_id")
	proto.RegisterType((*Coin)(nil), "prototype.coin")
	proto.RegisterType((*Vest)(nil), "prototype.vest")
	proto.RegisterType((*PublicKeyType)(nil), "prototype.public_key_type")
	proto.RegisterType((*PrivateKeyType)(nil), "prototype.private_key_type")
	proto.RegisterType((*TimePointSec)(nil), "prototype.time_point_sec")
	proto.RegisterType((*SignatureType)(nil), "prototype.signature_type")
	proto.RegisterType((*Sha256)(nil), "prototype.sha256")
	proto.RegisterType((*GiftTicketKeyType)(nil), "prototype.gift_ticket_key_type")
	proto.RegisterType((*ChainProperties)(nil), "prototype.chain_properties")
	proto.RegisterType((*DynamicProperties)(nil), "prototype.dynamic_properties")
	proto.RegisterType((*BeneficiaryRouteType)(nil), "prototype.beneficiary_route_type")
}

func init() { proto.RegisterFile("prototype/type.proto", fileDescriptor_f1b10af7c504b1c5) }

var fileDescriptor_f1b10af7c504b1c5 = []byte{
	// 1321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xef, 0x4e, 0x5b, 0x3b,
	0x12, 0xc0, 0x05, 0x4d, 0x29, 0x38, 0xe4, 0x0f, 0x6e, 0x96, 0x1e, 0x76, 0xbb, 0x5b, 0x9a, 0x6d,
	0xbb, 0xfd, 0x43, 0xc3, 0x2e, 0xdb, 0x56, 0xaa, 0x56, 0x2b, 0x2d, 0xd0, 0x56, 0xea, 0x87, 0xad,
	0x50, 0xa0, 0x45, 0xf7, 0x4a, 0x57, 0x96, 0x73, 0xce, 0x90, 0x58, 0xe4, 0xd8, 0xae, 0xed, 0x03,
	0xa4, 0x2f, 0x71, 0x1f, 0xf0, 0xbe, 0xcc, 0x95, 0xc7, 0x3e, 0xe1, 0xa4, 0x04, 0xfa, 0xe1, 0x7e,
	0x01, 0x67, 0xe6, 0x37, 0x1e, 0x7b, 0x66, 0x3c, 0x73, 0x48, 0x47, 0x1b, 0xe5, 0x94, 0x9b, 0x68,
	0xd8, 0xf6, 0x7f, 0x7a, 0xf8, 0x93, 0xae, 0x4c, 0xa5, 0xdd, 0x47, 0x64, 0x95, 0xa7, 0xa9, 0x2a,
	0xa4, 0x63, 0x92, 0xe7, 0x40, 0x3b, 0xe4, 0xf6, 0x19, 0x1f, 0x17, 0x90, 0x2c, 0x6c, 0x2e, 0x3c,
	0x5d, 0xe9, 0x87, 0x1f, 0xdd, 0x4d, 0xb2, 0x9c, 0x8e, 0xb8, 0x90, 0x4c, 0x64, 0xb3, 0x44, 0xe3,
	0x92, 0xa8, 0xa5, 0x4a, 0x48, 0x9a, 0x54, 0xb5, 0xb5, 0xbd, 0xc5, 0x7f, 0x2e, 0x54, 0x88, 0x33,
	0xb0, 0xee, 0x06, 0xe2, 0x31, 0x69, 0xe9, 0x62, 0x30, 0x16, 0x29, 0x3b, 0x85, 0x09, 0xf3, 0xc7,
	0xa3, 0x94, 0xd4, 0x32, 0xee, 0x38, 0xb2, 0xab, 0x7d, 0x5c, 0x77, 0x9f, 0x90, 0xb6, 0x36, 0xe2,
	0x8c, 0x3b, 0xb8, 0x99, 0xfb, 0x17, 0x69, 0x3a, 0x91, 0x03, 0xd3, 0x4a, 0x48, 0xc7, 0x2c, 0xa4,
	0xf4, 0x01, 0xa9, 0x17, 0x2e, 0xf5, 0x4b, 0x25, 0x33, 0x1b, 0x2f, 0x40, 0x0a, 0x97, 0x1e, 0x06,
	0x49, 0xb7, 0x4b, 0x9a, 0x56, 0x0c, 0x25, 0x77, 0x85, 0x81, 0xb0, 0x71, 0x9b, 0xdc, 0xb2, 0x62,
	0x18, 0xf7, 0xf5, 0xcb, 0xee, 0x7d, 0xb2, 0x64, 0x47, 0x7c, 0xe7, 0xf5, 0x1b, 0xef, 0x74, 0xc4,
	0xed, 0xa8, 0x74, 0xea, 0xd7, 0xdd, 0xaf, 0xa4, 0x33, 0x14, 0x27, 0x8e, 0x39, 0x91, 0x9e, 0x82,
	0x9b, 0x39, 0xa0, 0xff, 0x1f, 0x7d, 0xd6, 0x4a, 0xd9, 0x89, 0x51, 0x79, 0xb2, 0x88, 0xa1, 0xc6,
	0x35, 0x6d, 0x92, 0x45, 0xa7, 0x92, 0x5b, 0x28, 0x59, 0x74, 0x8a, 0x3e, 0x24, 0xab, 0xa9, 0x01,
	0x7f, 0xd7, 0xc1, 0x58, 0xa5, 0xa7, 0x49, 0xcd, 0x07, 0xad, 0x5f, 0x0f, 0xb2, 0x3d, 0x2f, 0xea,
	0xfe, 0x7a, 0x8b, 0xb4, 0x43, 0x76, 0xb4, 0x51, 0x1a, 0x8c, 0x13, 0x60, 0xe9, 0x2e, 0xe9, 0x94,
	0x79, 0x45, 0x56, 0x28, 0xc9, 0x4e, 0x20, 0xf8, 0xaf, 0xef, 0xb4, 0x7a, 0xd3, 0x0a, 0xe8, 0xf9,
	0xb4, 0xf5, 0x69, 0x84, 0xf7, 0x23, 0xfb, 0x01, 0x80, 0x6e, 0x11, 0x9a, 0xf3, 0x0b, 0x91, 0x17,
	0x79, 0xf0, 0xcd, 0xac, 0xf8, 0x06, 0x78, 0xd8, 0x46, 0xbf, 0x1d, 0x35, 0x78, 0x82, 0x43, 0xf1,
	0x0d, 0xfc, 0x41, 0xad, 0xe3, 0xb9, 0x90, 0x9c, 0x9d, 0x18, 0x00, 0xbc, 0x42, 0xad, 0x5f, 0x8f,
	0xb2, 0x0f, 0x06, 0x10, 0x71, 0xda, 0x32, 0xb8, 0xd0, 0x90, 0x3a, 0xc8, 0xca, 0xbb, 0x38, 0x6d,
	0xdf, 0x47, 0x11, 0x7d, 0x43, 0x12, 0xa7, 0x34, 0x93, 0x8c, 0xa7, 0x5f, 0x0b, 0x61, 0x00, 0xf7,
	0x62, 0x4e, 0x9d, 0x82, 0x4c, 0x6e, 0xa3, 0xe7, 0x8e, 0x53, 0xfa, 0xd3, 0x6e, 0xd0, 0xfa, 0x5d,
	0x8f, 0xbc, 0x8e, 0x3e, 0x26, 0x4d, 0xd0, 0x2a, 0x1d, 0xb1, 0xac, 0x30, 0x78, 0x81, 0x64, 0x09,
	0x37, 0x6f, 0xa0, 0xf4, 0x5d, 0x14, 0xd2, 0xb7, 0xa4, 0xad, 0xc1, 0x94, 0xc9, 0xd1, 0x46, 0xa4,
	0x90, 0xdc, 0x99, 0x1f, 0x91, 0xa6, 0x06, 0x73, 0x84, 0xdc, 0x81, 0xc7, 0xe8, 0x73, 0xb2, 0x56,
	0x31, 0x3d, 0x07, 0x31, 0x1c, 0xb9, 0x64, 0x19, 0x9d, 0xb4, 0xa6, 0xe8, 0x31, 0x8a, 0xbb, 0xbf,
	0xb5, 0x09, 0xcd, 0x26, 0x92, 0xe7, 0x22, 0xad, 0xe6, 0xe4, 0x35, 0x69, 0x8c, 0x80, 0x67, 0x31,
	0x9a, 0x22, 0xc3, 0x58, 0xd6, 0x77, 0xd6, 0x2a, 0xae, 0x43, 0x65, 0xf5, 0xeb, 0x9e, 0xc3, 0xd8,
	0x7e, 0xcc, 0xbc, 0xe7, 0x8a, 0x99, 0x2c, 0xf2, 0x01, 0x98, 0x18, 0xde, 0xd6, 0x94, 0xfb, 0x84,
	0xe2, 0x6b, 0x72, 0x56, 0xbb, 0x26, 0x67, 0x5b, 0x64, 0xc5, 0x29, 0xc7, 0xc7, 0x2c, 0x55, 0x16,
	0xc3, 0x3b, 0x27, 0x0e, 0xcb, 0x48, 0xec, 0x2b, 0x4b, 0x5f, 0x92, 0x9a, 0x7f, 0x4f, 0x18, 0xd9,
	0xfa, 0xce, 0x46, 0x05, 0x9c, 0x7d, 0x66, 0x7d, 0xc4, 0xe8, 0xff, 0x48, 0x2b, 0x2d, 0x8c, 0x01,
	0xe9, 0xd8, 0xb9, 0x70, 0x12, 0xac, 0x8d, 0xa1, 0xbe, 0x57, 0xb1, 0xac, 0xf6, 0x9e, 0x7e, 0x33,
	0xf2, 0xc7, 0x01, 0xf7, 0x6f, 0xcf, 0x69, 0x9b, 0xac, 0xe0, 0xe9, 0xfd, 0xd2, 0x57, 0x75, 0x38,
	0xb0, 0xef, 0x24, 0x42, 0x0e, 0x99, 0x1d, 0x71, 0x03, 0x36, 0x21, 0x57, 0xce, 0xee, 0x81, 0x3e,
	0x45, 0xf8, 0x4b, 0x60, 0x0f, 0x11, 0xa5, 0x6f, 0x48, 0xe9, 0x86, 0xd9, 0x42, 0xeb, 0xf1, 0x24,
	0xa9, 0xcf, 0xbf, 0x78, 0x23, 0x62, 0x87, 0x48, 0x61, 0xfe, 0x95, 0x2d, 0x33, 0x0f, 0x19, 0x3b,
	0xd3, 0x36, 0x69, 0xe0, 0x3b, 0x6d, 0x79, 0xc5, 0x71, 0x94, 0x7f, 0xd1, 0x96, 0xee, 0x90, 0x55,
	0x64, 0x0d, 0x9c, 0x73, 0x93, 0xd9, 0xa4, 0x39, 0xff, 0x78, 0x75, 0x0f, 0xf5, 0x03, 0x43, 0xbb,
	0xa4, 0x11, 0xae, 0xe6, 0xcc, 0x05, 0x4b, 0xa5, 0x4b, 0x5a, 0xf1, 0x75, 0x78, 0xe1, 0x91, 0xb9,
	0xd8, 0x97, 0x8e, 0x3e, 0x22, 0xcd, 0xc0, 0xe0, 0xee, 0x1e, 0x6a, 0x23, 0xb4, 0x8a, 0xd2, 0x03,
	0x65, 0xdd, 0x0c, 0x55, 0x58, 0x30, 0x48, 0xad, 0x55, 0xa8, 0xcf, 0x16, 0x8c, 0xa7, 0xee, 0x91,
	0x3b, 0x39, 0xbf, 0x60, 0x3e, 0xc0, 0x14, 0x03, 0xbc, 0x94, 0xf3, 0x8b, 0x23, 0xed, 0xd3, 0xbc,
	0x16, 0x15, 0x97, 0x15, 0x97, 0xdc, 0x9d, 0xf6, 0xea, 0x66, 0xc0, 0xca, 0xa2, 0xfb, 0xae, 0x3a,
	0xb5, 0x81, 0x13, 0x71, 0x91, 0x74, 0x70, 0xc7, 0xcb, 0xea, 0x3c, 0x40, 0xb1, 0x8f, 0xbd, 0x01,
	0xad, 0xcc, 0x65, 0x64, 0xfe, 0x34, 0x3f, 0x32, 0x8d, 0x80, 0x95, 0xb1, 0xd9, 0x20, 0xcb, 0xc2,
	0x8d, 0xd8, 0x04, 0xb8, 0x49, 0xd6, 0x71, 0xeb, 0x3b, 0xc2, 0x8d, 0x7e, 0x02, 0x6e, 0xe8, 0x2b,
	0xd2, 0xe0, 0x52, 0x16, 0x7c, 0xcc, 0x06, 0x45, 0x36, 0x04, 0x97, 0xdc, 0x9b, 0xbf, 0xe3, 0x6a,
	0xa0, 0xf6, 0x10, 0xaa, 0x58, 0xe5, 0x42, 0xfa, 0x56, 0x94, 0xdc, 0x68, 0xf5, 0x7f, 0x84, 0xe8,
	0x7f, 0x62, 0x09, 0x64, 0x5c, 0xeb, 0xe9, 0x0d, 0x36, 0xe6, 0x5b, 0x62, 0x4d, 0xbc, 0xe3, 0x5a,
	0x97, 0x77, 0x78, 0x45, 0x1a, 0x67, 0xca, 0x81, 0x99, 0x1a, 0xfe, 0xf9, 0x1a, 0x97, 0x48, 0x55,
	0xac, 0x0c, 0xe8, 0xf1, 0x64, 0x6a, 0xf5, 0x97, 0x6b, 0xac, 0x90, 0x2a, 0xad, 0xb6, 0x08, 0x0d,
	0x56, 0x33, 0xc5, 0x7a, 0x1f, 0x8b, 0xb5, 0x8d, 0x9a, 0x6a, 0xb5, 0xfe, 0xb7, 0xa4, 0x67, 0xee,
	0xf5, 0xd7, 0xf9, 0x8e, 0x82, 0x79, 0xf5, 0x62, 0xbb, 0xa4, 0x63, 0x1d, 0x3f, 0x85, 0xef, 0xdf,
	0xe4, 0xdf, 0xae, 0x79, 0x93, 0x08, 0xcf, 0xbe, 0xc9, 0x57, 0x64, 0x3d, 0xb6, 0x08, 0x36, 0x50,
	0xca, 0xb1, 0x54, 0xe5, 0x7a, 0x0c, 0x3e, 0x2f, 0x0f, 0x36, 0x17, 0x9e, 0x2e, 0xf7, 0x3b, 0x51,
	0xbb, 0xa7, 0x94, 0xdb, 0x2f, 0x75, 0x57, 0x26, 0xce, 0xe6, 0x8f, 0x27, 0xce, 0xc3, 0xab, 0x13,
	0x67, 0x9b, 0x74, 0xf8, 0xd9, 0x10, 0xcb, 0xbd, 0xd0, 0xd9, 0xe5, 0xa0, 0xed, 0x22, 0xba, 0xc6,
	0xcf, 0x86, 0x47, 0xda, 0x7e, 0x46, 0x0d, 0x96, 0x32, 0x7d, 0x46, 0xd6, 0x4a, 0x03, 0x21, 0xd9,
	0xb9, 0x90, 0x99, 0x3a, 0x4f, 0xfe, 0x8e, 0x74, 0x33, 0xd0, 0x1f, 0xe5, 0x31, 0x4a, 0xe9, 0x13,
	0xd2, 0x52, 0x12, 0x58, 0xc6, 0x27, 0x2c, 0x9e, 0x2a, 0x79, 0x14, 0xc6, 0x92, 0x92, 0xf0, 0x8e,
	0x4f, 0x0e, 0x83, 0xd0, 0x67, 0x60, 0x66, 0x58, 0x03, 0x8e, 0xea, 0xc7, 0xf3, 0xfb, 0x52, 0xbb,
	0x3a, 0xaa, 0xc1, 0x0f, 0xea, 0x3d, 0xe2, 0xb3, 0x52, 0xb8, 0x30, 0xe5, 0x79, 0x96, 0x0b, 0x99,
	0x3c, 0xb9, 0xb9, 0xd5, 0xb6, 0x2e, 0x0d, 0x76, 0x3d, 0x4f, 0xdf, 0x92, 0x8d, 0xb2, 0x2d, 0x86,
	0x41, 0x6a, 0x1d, 0x37, 0x2e, 0xc6, 0xe2, 0x1f, 0x78, 0xe8, 0xf5, 0x08, 0xbc, 0xf7, 0xfa, 0x43,
	0xaf, 0x0e, 0x01, 0xb9, 0x3a, 0x7b, 0x9f, 0xce, 0x9b, 0xbd, 0x37, 0x8d, 0xf6, 0x67, 0x37, 0x8c,
	0xf6, 0x79, 0x33, 0xfb, 0xf9, 0x1f, 0x98, 0xd9, 0x2f, 0xe6, 0xce, 0x6c, 0xdf, 0x9b, 0x02, 0xe7,
	0xd3, 0x9a, 0xaa, 0x1c, 0x92, 0xad, 0x6b, 0x7a, 0x53, 0xc4, 0x3e, 0x22, 0x45, 0x7b, 0xe4, 0x6e,
	0x3a, 0xe2, 0x66, 0x08, 0x19, 0x2b, 0xed, 0x7d, 0xc3, 0x7c, 0x19, 0xca, 0x27, 0xaa, 0x82, 0x27,
	0xeb, 0xfb, 0xa5, 0x1f, 0x8b, 0x4a, 0x4f, 0x8c, 0x77, 0x1a, 0x73, 0xd5, 0xfb, 0xd1, 0x58, 0x2c,
	0x79, 0x4c, 0x55, 0xf7, 0x17, 0xb2, 0x3e, 0x00, 0x09, 0x27, 0x22, 0x15, 0xdc, 0x4c, 0x98, 0x51,
	0x85, 0x8b, 0x1f, 0xab, 0x2f, 0x48, 0xcd, 0x5b, 0xc4, 0x8f, 0xbc, 0x6b, 0x37, 0x44, 0x88, 0xae,
	0x93, 0xa5, 0x18, 0x91, 0xf0, 0x49, 0x17, 0x7f, 0xed, 0x1d, 0x90, 0xae, 0x50, 0xbd, 0x54, 0x49,
	0x07, 0xd2, 0x29, 0xdb, 0xe3, 0x32, 0x33, 0x4a, 0x64, 0x3d, 0x9b, 0x9d, 0x5e, 0x6e, 0xf8, 0xf3,
	0xf3, 0xa1, 0x70, 0xa3, 0x62, 0xd0, 0x4b, 0x55, 0xbe, 0x9d, 0x2a, 0x8b, 0xdf, 0x9f, 0xdb, 0x53,
	0xa3, 0x97, 0x43, 0xb5, 0x3d, 0x65, 0x07, 0x4b, 0xb8, 0xfc, 0xf7, 0xef, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0x9d, 0x7c, 0x3b, 0x8d, 0x0c, 0x00, 0x00,
}
