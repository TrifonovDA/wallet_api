// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet_api.proto

package wallet_api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetWalletsStatus int32

const (
	GetWalletsStatus_INTERNAL_ERROR GetWalletsStatus = 0
	GetWalletsStatus_EMPTY          GetWalletsStatus = 1
	GetWalletsStatus_OK             GetWalletsStatus = 2
)

var GetWalletsStatus_name = map[int32]string{
	0: "INTERNAL_ERROR",
	1: "EMPTY",
	2: "OK",
}

var GetWalletsStatus_value = map[string]int32{
	"INTERNAL_ERROR": 0,
	"EMPTY":          1,
	"OK":             2,
}

func (x GetWalletsStatus) String() string {
	return proto.EnumName(GetWalletsStatus_name, int32(x))
}

func (GetWalletsStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{0}
}

type SaveWalletStates int32

const (
	SaveWalletStates_ERROR     SaveWalletStates = 0
	SaveWalletStates_SUCCESS   SaveWalletStates = 1
	SaveWalletStates_WAS_SAVED SaveWalletStates = 2
)

var SaveWalletStates_name = map[int32]string{
	0: "ERROR",
	1: "SUCCESS",
	2: "WAS_SAVED",
}

var SaveWalletStates_value = map[string]int32{
	"ERROR":     0,
	"SUCCESS":   1,
	"WAS_SAVED": 2,
}

func (x SaveWalletStates) String() string {
	return proto.EnumName(SaveWalletStates_name, int32(x))
}

func (SaveWalletStates) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{1}
}

type ValidateStates int32

const (
	ValidateStates_VALIDATOR_ERROR ValidateStates = 0
	ValidateStates_CHECKED         ValidateStates = 1
)

var ValidateStates_name = map[int32]string{
	0: "VALIDATOR_ERROR",
	1: "CHECKED",
}

var ValidateStates_value = map[string]int32{
	"VALIDATOR_ERROR": 0,
	"CHECKED":         1,
}

func (x ValidateStates) String() string {
	return proto.EnumName(ValidateStates_name, int32(x))
}

func (ValidateStates) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{2}
}

type UserId struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{0}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetWalletsResult struct {
	Status               GetWalletsStatus `protobuf:"varint,1,opt,name=status,proto3,enum=wallet_api.GetWalletsStatus" json:"status,omitempty"`
	Coins                []*Coins         `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetWalletsResult) Reset()         { *m = GetWalletsResult{} }
func (m *GetWalletsResult) String() string { return proto.CompactTextString(m) }
func (*GetWalletsResult) ProtoMessage()    {}
func (*GetWalletsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{1}
}

func (m *GetWalletsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletsResult.Unmarshal(m, b)
}
func (m *GetWalletsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletsResult.Marshal(b, m, deterministic)
}
func (m *GetWalletsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletsResult.Merge(m, src)
}
func (m *GetWalletsResult) XXX_Size() int {
	return xxx_messageInfo_GetWalletsResult.Size(m)
}
func (m *GetWalletsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletsResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletsResult proto.InternalMessageInfo

func (m *GetWalletsResult) GetStatus() GetWalletsStatus {
	if m != nil {
		return m.Status
	}
	return GetWalletsStatus_INTERNAL_ERROR
}

func (m *GetWalletsResult) GetCoins() []*Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

type Coins struct {
	Coin                 string      `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
	Networks             []*Networks `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Coins) Reset()         { *m = Coins{} }
func (m *Coins) String() string { return proto.CompactTextString(m) }
func (*Coins) ProtoMessage()    {}
func (*Coins) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{2}
}

func (m *Coins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coins.Unmarshal(m, b)
}
func (m *Coins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coins.Marshal(b, m, deterministic)
}
func (m *Coins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coins.Merge(m, src)
}
func (m *Coins) XXX_Size() int {
	return xxx_messageInfo_Coins.Size(m)
}
func (m *Coins) XXX_DiscardUnknown() {
	xxx_messageInfo_Coins.DiscardUnknown(m)
}

var xxx_messageInfo_Coins proto.InternalMessageInfo

func (m *Coins) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *Coins) GetNetworks() []*Networks {
	if m != nil {
		return m.Networks
	}
	return nil
}

type Networks struct {
	Network              string   `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Wallets              []string `protobuf:"bytes,2,rep,name=wallets,proto3" json:"wallets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Networks) Reset()         { *m = Networks{} }
func (m *Networks) String() string { return proto.CompactTextString(m) }
func (*Networks) ProtoMessage()    {}
func (*Networks) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{3}
}

func (m *Networks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Networks.Unmarshal(m, b)
}
func (m *Networks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Networks.Marshal(b, m, deterministic)
}
func (m *Networks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Networks.Merge(m, src)
}
func (m *Networks) XXX_Size() int {
	return xxx_messageInfo_Networks.Size(m)
}
func (m *Networks) XXX_DiscardUnknown() {
	xxx_messageInfo_Networks.DiscardUnknown(m)
}

var xxx_messageInfo_Networks proto.InternalMessageInfo

func (m *Networks) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Networks) GetWallets() []string {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type WalletInfo struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Wallet               string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Coin                 string   `protobuf:"bytes,3,opt,name=coin,proto3" json:"coin,omitempty"`
	Network              string   `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletInfo) Reset()         { *m = WalletInfo{} }
func (m *WalletInfo) String() string { return proto.CompactTextString(m) }
func (*WalletInfo) ProtoMessage()    {}
func (*WalletInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{4}
}

func (m *WalletInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletInfo.Unmarshal(m, b)
}
func (m *WalletInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletInfo.Marshal(b, m, deterministic)
}
func (m *WalletInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletInfo.Merge(m, src)
}
func (m *WalletInfo) XXX_Size() int {
	return xxx_messageInfo_WalletInfo.Size(m)
}
func (m *WalletInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WalletInfo proto.InternalMessageInfo

func (m *WalletInfo) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WalletInfo) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *WalletInfo) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *WalletInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type SaveWalletStatus struct {
	Status               SaveWalletStates `protobuf:"varint,1,opt,name=status,proto3,enum=wallet_api.SaveWalletStates" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SaveWalletStatus) Reset()         { *m = SaveWalletStatus{} }
func (m *SaveWalletStatus) String() string { return proto.CompactTextString(m) }
func (*SaveWalletStatus) ProtoMessage()    {}
func (*SaveWalletStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{5}
}

func (m *SaveWalletStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveWalletStatus.Unmarshal(m, b)
}
func (m *SaveWalletStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveWalletStatus.Marshal(b, m, deterministic)
}
func (m *SaveWalletStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveWalletStatus.Merge(m, src)
}
func (m *SaveWalletStatus) XXX_Size() int {
	return xxx_messageInfo_SaveWalletStatus.Size(m)
}
func (m *SaveWalletStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveWalletStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SaveWalletStatus proto.InternalMessageInfo

func (m *SaveWalletStatus) GetStatus() SaveWalletStates {
	if m != nil {
		return m.Status
	}
	return SaveWalletStates_ERROR
}

type ValidateWalletRequest struct {
	Wallet               string   `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Network              string   `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateWalletRequest) Reset()         { *m = ValidateWalletRequest{} }
func (m *ValidateWalletRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateWalletRequest) ProtoMessage()    {}
func (*ValidateWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{6}
}

func (m *ValidateWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateWalletRequest.Unmarshal(m, b)
}
func (m *ValidateWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateWalletRequest.Marshal(b, m, deterministic)
}
func (m *ValidateWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateWalletRequest.Merge(m, src)
}
func (m *ValidateWalletRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateWalletRequest.Size(m)
}
func (m *ValidateWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateWalletRequest proto.InternalMessageInfo

func (m *ValidateWalletRequest) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *ValidateWalletRequest) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type ValidateResult struct {
	InternalStatus       ValidateStates `protobuf:"varint,1,opt,name=internal_status,json=internalStatus,proto3,enum=wallet_api.ValidateStates" json:"internal_status,omitempty"`
	Result               bool           `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ValidateResult) Reset()         { *m = ValidateResult{} }
func (m *ValidateResult) String() string { return proto.CompactTextString(m) }
func (*ValidateResult) ProtoMessage()    {}
func (*ValidateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_34bd94de71a6388e, []int{7}
}

func (m *ValidateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResult.Unmarshal(m, b)
}
func (m *ValidateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResult.Marshal(b, m, deterministic)
}
func (m *ValidateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResult.Merge(m, src)
}
func (m *ValidateResult) XXX_Size() int {
	return xxx_messageInfo_ValidateResult.Size(m)
}
func (m *ValidateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResult.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResult proto.InternalMessageInfo

func (m *ValidateResult) GetInternalStatus() ValidateStates {
	if m != nil {
		return m.InternalStatus
	}
	return ValidateStates_VALIDATOR_ERROR
}

func (m *ValidateResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterEnum("wallet_api.GetWalletsStatus", GetWalletsStatus_name, GetWalletsStatus_value)
	proto.RegisterEnum("wallet_api.SaveWalletStates", SaveWalletStates_name, SaveWalletStates_value)
	proto.RegisterEnum("wallet_api.ValidateStates", ValidateStates_name, ValidateStates_value)
	proto.RegisterType((*UserId)(nil), "wallet_api.User_id")
	proto.RegisterType((*GetWalletsResult)(nil), "wallet_api.Get_wallets_result")
	proto.RegisterType((*Coins)(nil), "wallet_api.Coins")
	proto.RegisterType((*Networks)(nil), "wallet_api.Networks")
	proto.RegisterType((*WalletInfo)(nil), "wallet_api.Wallet_info")
	proto.RegisterType((*SaveWalletStatus)(nil), "wallet_api.Save_wallet_status")
	proto.RegisterType((*ValidateWalletRequest)(nil), "wallet_api.Validate_wallet_request")
	proto.RegisterType((*ValidateResult)(nil), "wallet_api.Validate_result")
}

func init() { proto.RegisterFile("wallet_api.proto", fileDescriptor_34bd94de71a6388e) }

var fileDescriptor_34bd94de71a6388e = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xad, 0x94, 0x58, 0xb6, 0xc7, 0xd4, 0x56, 0x27, 0x25, 0x36, 0x2e, 0x94, 0xa0, 0x1e, 0x1a,
	0x72, 0x08, 0xc5, 0x81, 0x5e, 0x0a, 0x05, 0x21, 0x8b, 0xd4, 0xd8, 0xb1, 0xcb, 0xca, 0xb1, 0xe9,
	0x49, 0xa8, 0xf5, 0x16, 0x44, 0x85, 0x94, 0x4a, 0xab, 0xe6, 0x9f, 0xfb, 0x15, 0x45, 0xbb, 0x23,
	0x75, 0xd5, 0x26, 0xb9, 0x69, 0x66, 0xdf, 0x7b, 0xfb, 0xe6, 0x69, 0x24, 0xb0, 0xef, 0xa3, 0x24,
	0xe1, 0x22, 0x8c, 0xee, 0xe2, 0xcb, 0xbb, 0x3c, 0x13, 0x19, 0xc2, 0xdf, 0x8e, 0xe3, 0x40, 0xf7,
	0xb6, 0xe0, 0x79, 0x18, 0x1f, 0x70, 0x0c, 0xdd, 0x52, 0x3d, 0x4e, 0x8c, 0x33, 0xe3, 0xbc, 0xc3,
	0xac, 0xaa, 0x5c, 0x1c, 0x9c, 0x12, 0xf0, 0x9a, 0x8b, 0x50, 0xb1, 0x8a, 0x30, 0xe7, 0x45, 0x99,
	0x08, 0x7c, 0x0f, 0x56, 0x21, 0x22, 0x51, 0x16, 0x12, 0x3d, 0x9c, 0xbd, 0xbe, 0xd4, 0x2e, 0xd2,
	0xf1, 0x0a, 0xc5, 0x08, 0x8d, 0x6f, 0xa1, 0xf3, 0x2d, 0x8b, 0xd3, 0x62, 0x62, 0x9e, 0x1d, 0x9d,
	0x0f, 0x66, 0x2f, 0x74, 0x9a, 0x57, 0x1d, 0x30, 0x75, 0xee, 0xdc, 0x40, 0x47, 0xd6, 0x88, 0x70,
	0x5c, 0x75, 0xe4, 0x3d, 0x7d, 0x26, 0x9f, 0xf1, 0x1d, 0xf4, 0x52, 0x2e, 0xee, 0xb3, 0xfc, 0x47,
	0x2d, 0xf4, 0x52, 0x17, 0x5a, 0xd3, 0x19, 0x6b, 0x50, 0xce, 0x47, 0xe8, 0xd5, 0x5d, 0x9c, 0x40,
	0x97, 0xfa, 0x24, 0x5a, 0x97, 0xd5, 0x09, 0xf9, 0x96, 0xb2, 0x7d, 0x56, 0x97, 0x4e, 0x02, 0x83,
	0xbd, 0xba, 0x20, 0x4e, 0xbf, 0x67, 0x8f, 0xa6, 0x85, 0xa7, 0x60, 0x29, 0xca, 0xc4, 0x94, 0xd2,
	0x54, 0x35, 0x53, 0x1c, 0x69, 0x53, 0x68, 0x3e, 0x8e, 0x5b, 0x3e, 0x9c, 0x15, 0x60, 0x10, 0xfd,
	0xe2, 0x14, 0x22, 0x65, 0xf8, 0x74, 0xe6, 0xff, 0xe2, 0x79, 0x93, 0xb9, 0xb3, 0x84, 0xf1, 0x2e,
	0x4a, 0xe2, 0x43, 0x24, 0x1a, 0x44, 0xce, 0x7f, 0x96, 0xbc, 0x10, 0x9a, 0x5d, 0xa3, 0x65, 0x57,
	0xb3, 0x66, 0xb6, 0xad, 0x65, 0x30, 0x6a, 0xc4, 0x68, 0x17, 0xe6, 0x30, 0x8a, 0x53, 0xc1, 0xf3,
	0x34, 0x4a, 0xc2, 0x96, 0xc1, 0x57, 0xba, 0xc1, 0x86, 0x45, 0xee, 0x86, 0x35, 0x27, 0x50, 0xd3,
	0x9d, 0x82, 0xa5, 0xf4, 0xe4, 0x8d, 0x3d, 0x46, 0xd5, 0xc5, 0x87, 0xf6, 0xfe, 0x51, 0x16, 0x08,
	0xc3, 0xc5, 0x7a, 0xeb, 0xb3, 0xb5, 0xbb, 0x0a, 0x7d, 0xc6, 0x36, 0xcc, 0x7e, 0x86, 0x7d, 0xe8,
	0xf8, 0x37, 0x9f, 0xb7, 0x5f, 0x6c, 0x03, 0x2d, 0x30, 0x37, 0x4b, 0xdb, 0xac, 0xc8, 0xff, 0x07,
	0x23, 0x81, 0xc4, 0x19, 0x40, 0x37, 0xb8, 0xf5, 0x3c, 0x3f, 0x08, 0x6c, 0x03, 0x9f, 0x43, 0x7f,
	0xef, 0x06, 0x61, 0xe0, 0xee, 0xfc, 0xb9, 0x6d, 0x5e, 0x5c, 0x69, 0xa3, 0x12, 0xf3, 0x04, 0x46,
	0x3b, 0x77, 0xb5, 0x98, 0xbb, 0xdb, 0x0d, 0x0b, 0x35, 0x0d, 0xef, 0x93, 0xef, 0x2d, 0xfd, 0xb9,
	0x6d, 0xcc, 0x7e, 0x1b, 0x00, 0xfb, 0x66, 0x6a, 0x74, 0x01, 0xae, 0xb9, 0x50, 0x8d, 0x02, 0x4f,
	0xf4, 0x40, 0xe8, 0xcb, 0x9b, 0x3e, 0xfa, 0xe9, 0x50, 0xbc, 0x3e, 0x40, 0x35, 0x83, 0xd2, 0xc0,
	0xb1, 0x8e, 0xd6, 0x56, 0x72, 0xfa, 0xe4, 0x36, 0x94, 0x05, 0x32, 0x18, 0xd6, 0xd3, 0x90, 0xd4,
	0x9b, 0x07, 0x5f, 0x4f, 0x7b, 0x43, 0xa6, 0x0f, 0xbf, 0x43, 0x65, 0xed, 0xab, 0x25, 0x7f, 0x29,
	0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xea, 0x85, 0xb3, 0xb0, 0x66, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletApiClient is the client API for WalletApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletApiClient interface {
	GetWallets(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetWalletsResult, error)
	SaveWallet(ctx context.Context, in *WalletInfo, opts ...grpc.CallOption) (*SaveWalletStatus, error)
	ValidateWallet(ctx context.Context, in *ValidateWalletRequest, opts ...grpc.CallOption) (*ValidateResult, error)
}

type walletApiClient struct {
	cc *grpc.ClientConn
}

func NewWalletApiClient(cc *grpc.ClientConn) WalletApiClient {
	return &walletApiClient{cc}
}

func (c *walletApiClient) GetWallets(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetWalletsResult, error) {
	out := new(GetWalletsResult)
	err := c.cc.Invoke(ctx, "/wallet_api.Wallet_api/GetWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletApiClient) SaveWallet(ctx context.Context, in *WalletInfo, opts ...grpc.CallOption) (*SaveWalletStatus, error) {
	out := new(SaveWalletStatus)
	err := c.cc.Invoke(ctx, "/wallet_api.Wallet_api/SaveWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletApiClient) ValidateWallet(ctx context.Context, in *ValidateWalletRequest, opts ...grpc.CallOption) (*ValidateResult, error) {
	out := new(ValidateResult)
	err := c.cc.Invoke(ctx, "/wallet_api.Wallet_api/ValidateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletApiServer is the server API for WalletApi service.
type WalletApiServer interface {
	GetWallets(context.Context, *UserId) (*GetWalletsResult, error)
	SaveWallet(context.Context, *WalletInfo) (*SaveWalletStatus, error)
	ValidateWallet(context.Context, *ValidateWalletRequest) (*ValidateResult, error)
}

// UnimplementedWalletApiServer can be embedded to have forward compatible implementations.
type UnimplementedWalletApiServer struct {
}

func (*UnimplementedWalletApiServer) GetWallets(ctx context.Context, req *UserId) (*GetWalletsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallets not implemented")
}
func (*UnimplementedWalletApiServer) SaveWallet(ctx context.Context, req *WalletInfo) (*SaveWalletStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveWallet not implemented")
}
func (*UnimplementedWalletApiServer) ValidateWallet(ctx context.Context, req *ValidateWalletRequest) (*ValidateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateWallet not implemented")
}

func RegisterWalletApiServer(s *grpc.Server, srv WalletApiServer) {
	s.RegisterService(&_WalletApi_serviceDesc, srv)
}

func _WalletApi_GetWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletApiServer).GetWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.Wallet_api/GetWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletApiServer).GetWallets(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletApi_SaveWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletApiServer).SaveWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.Wallet_api/SaveWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletApiServer).SaveWallet(ctx, req.(*WalletInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletApi_ValidateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletApiServer).ValidateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_api.Wallet_api/ValidateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletApiServer).ValidateWallet(ctx, req.(*ValidateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WalletApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wallet_api.Wallet_api",
	HandlerType: (*WalletApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWallets",
			Handler:    _WalletApi_GetWallets_Handler,
		},
		{
			MethodName: "SaveWallet",
			Handler:    _WalletApi_SaveWallet_Handler,
		},
		{
			MethodName: "ValidateWallet",
			Handler:    _WalletApi_ValidateWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_api.proto",
}