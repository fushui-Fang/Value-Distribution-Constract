// Code generated by protoc-gen-go. DO NOT EDIT.
// source: structure.proto

/*
Package lyan is a generated protocol buffer package.

It is generated from these files:
	structure.proto

It has these top-level messages:
	Account
	AllocateSgy
	MemberInfo
	MemberSingle
	SigeForSgyInfo
	SigeForSgy
	SigeForSgyStub
	TranferOutput
	TxInfo
	TX
*/
package lyan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Addr    string  `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Kind    int64   `protobuf:"varint,2,opt,name=kind" json:"kind,omitempty"`
	Balance float32 `protobuf:"fixed32,3,opt,name=balance" json:"balance,omitempty"`
	Pubkey  string  `protobuf:"bytes,4,opt,name=pubkey" json:"pubkey,omitempty"`
	ID      string  `protobuf:"bytes,5,opt,name=ID" json:"ID,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Account) GetKind() int64 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *Account) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *Account) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type AllocateSgy struct {
	ID        string      `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	PriorID   string      `protobuf:"bytes,2,opt,name=priorID" json:"priorID,omitempty"`
	Addr      string      `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	ASgy      *MemberInfo `protobuf:"bytes,4,opt,name=aSgy" json:"aSgy,omitempty"`
	HasSigned []bool      `protobuf:"varint,5,rep,packed,name=hasSigned" json:"hasSigned,omitempty"`
	Hash      string      `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
}

func (m *AllocateSgy) Reset()                    { *m = AllocateSgy{} }
func (m *AllocateSgy) String() string            { return proto.CompactTextString(m) }
func (*AllocateSgy) ProtoMessage()               {}
func (*AllocateSgy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllocateSgy) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AllocateSgy) GetPriorID() string {
	if m != nil {
		return m.PriorID
	}
	return ""
}

func (m *AllocateSgy) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *AllocateSgy) GetASgy() *MemberInfo {
	if m != nil {
		return m.ASgy
	}
	return nil
}

func (m *AllocateSgy) GetHasSigned() []bool {
	if m != nil {
		return m.HasSigned
	}
	return nil
}

func (m *AllocateSgy) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type MemberInfo struct {
	Ms []*MemberSingle `protobuf:"bytes,1,rep,name=ms" json:"ms,omitempty"`
}

func (m *MemberInfo) Reset()                    { *m = MemberInfo{} }
func (m *MemberInfo) String() string            { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()               {}
func (*MemberInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MemberInfo) GetMs() []*MemberSingle {
	if m != nil {
		return m.Ms
	}
	return nil
}

type MemberSingle struct {
	Sequence int64   `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Addr     string  `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	Ort      float32 `protobuf:"fixed32,3,opt,name=ort" json:"ort,omitempty"`
}

func (m *MemberSingle) Reset()                    { *m = MemberSingle{} }
func (m *MemberSingle) String() string            { return proto.CompactTextString(m) }
func (*MemberSingle) ProtoMessage()               {}
func (*MemberSingle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MemberSingle) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *MemberSingle) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *MemberSingle) GetOrt() float32 {
	if m != nil {
		return m.Ort
	}
	return 0
}

// 用户分配策略信息
type SigeForSgyInfo struct {
	Addr      string  `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	ID        string  `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	Seq       int64   `protobuf:"varint,3,opt,name=seq" json:"seq,omitempty"`
	Timestamp int64   `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Ort       float32 `protobuf:"fixed32,5,opt,name=ort" json:"ort,omitempty"`
}

func (m *SigeForSgyInfo) Reset()                    { *m = SigeForSgyInfo{} }
func (m *SigeForSgyInfo) String() string            { return proto.CompactTextString(m) }
func (*SigeForSgyInfo) ProtoMessage()               {}
func (*SigeForSgyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SigeForSgyInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SigeForSgyInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *SigeForSgyInfo) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *SigeForSgyInfo) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SigeForSgyInfo) GetOrt() float32 {
	if m != nil {
		return m.Ort
	}
	return 0
}

type SigeForSgy struct {
	Si     *SigeForSgyInfo `protobuf:"bytes,1,opt,name=si" json:"si,omitempty"`
	Script string          `protobuf:"bytes,2,opt,name=script" json:"script,omitempty"`
}

func (m *SigeForSgy) Reset()                    { *m = SigeForSgy{} }
func (m *SigeForSgy) String() string            { return proto.CompactTextString(m) }
func (*SigeForSgy) ProtoMessage()               {}
func (*SigeForSgy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SigeForSgy) GetSi() *SigeForSgyInfo {
	if m != nil {
		return m.Si
	}
	return nil
}

func (m *SigeForSgy) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

type SigeForSgyStub struct {
	Set []*SigeForSgy `protobuf:"bytes,1,rep,name=set" json:"set,omitempty"`
}

func (m *SigeForSgyStub) Reset()                    { *m = SigeForSgyStub{} }
func (m *SigeForSgyStub) String() string            { return proto.CompactTextString(m) }
func (*SigeForSgyStub) ProtoMessage()               {}
func (*SigeForSgyStub) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SigeForSgyStub) GetSet() []*SigeForSgy {
	if m != nil {
		return m.Set
	}
	return nil
}

type TranferOutput struct {
	OutAddr string  `protobuf:"bytes,1,opt,name=outAddr" json:"outAddr,omitempty"`
	OutNum  float32 `protobuf:"fixed32,2,opt,name=outNum" json:"outNum,omitempty"`
}

func (m *TranferOutput) Reset()                    { *m = TranferOutput{} }
func (m *TranferOutput) String() string            { return proto.CompactTextString(m) }
func (*TranferOutput) ProtoMessage()               {}
func (*TranferOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TranferOutput) GetOutAddr() string {
	if m != nil {
		return m.OutAddr
	}
	return ""
}

func (m *TranferOutput) GetOutNum() float32 {
	if m != nil {
		return m.OutNum
	}
	return 0
}

type TxInfo struct {
	Timestamp    int64            `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	InputAddr    string           `protobuf:"bytes,2,opt,name=inputAddr" json:"inputAddr,omitempty"`
	InputBalance float32          `protobuf:"fixed32,3,opt,name=inputBalance" json:"inputBalance,omitempty"`
	Nounce       float32          `protobuf:"fixed32,4,opt,name=nounce" json:"nounce,omitempty"`
	Output       []*TranferOutput `protobuf:"bytes,5,rep,name=output" json:"output,omitempty"`
	Info         string           `protobuf:"bytes,6,opt,name=info" json:"info,omitempty"`
}

func (m *TxInfo) Reset()                    { *m = TxInfo{} }
func (m *TxInfo) String() string            { return proto.CompactTextString(m) }
func (*TxInfo) ProtoMessage()               {}
func (*TxInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TxInfo) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TxInfo) GetInputAddr() string {
	if m != nil {
		return m.InputAddr
	}
	return ""
}

func (m *TxInfo) GetInputBalance() float32 {
	if m != nil {
		return m.InputBalance
	}
	return 0
}

func (m *TxInfo) GetNounce() float32 {
	if m != nil {
		return m.Nounce
	}
	return 0
}

func (m *TxInfo) GetOutput() []*TranferOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *TxInfo) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type TX struct {
	Tx     *TxInfo `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Script string  `protobuf:"bytes,2,opt,name=script" json:"script,omitempty"`
}

func (m *TX) Reset()                    { *m = TX{} }
func (m *TX) String() string            { return proto.CompactTextString(m) }
func (*TX) ProtoMessage()               {}
func (*TX) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TX) GetTx() *TxInfo {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TX) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "lyan.Account")
	proto.RegisterType((*AllocateSgy)(nil), "lyan.AllocateSgy")
	proto.RegisterType((*MemberInfo)(nil), "lyan.MemberInfo")
	proto.RegisterType((*MemberSingle)(nil), "lyan.MemberSingle")
	proto.RegisterType((*SigeForSgyInfo)(nil), "lyan.SigeForSgyInfo")
	proto.RegisterType((*SigeForSgy)(nil), "lyan.SigeForSgy")
	proto.RegisterType((*SigeForSgyStub)(nil), "lyan.SigeForSgyStub")
	proto.RegisterType((*TranferOutput)(nil), "lyan.TranferOutput")
	proto.RegisterType((*TxInfo)(nil), "lyan.TxInfo")
	proto.RegisterType((*TX)(nil), "lyan.TX")
}

func init() { proto.RegisterFile("structure.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x8b, 0x1a, 0x31,
	0x18, 0x66, 0x32, 0xea, 0xae, 0xaf, 0x76, 0x2b, 0x69, 0x29, 0x43, 0xf1, 0x20, 0x61, 0x0f, 0x42,
	0x41, 0x8a, 0xed, 0xa9, 0x37, 0x8b, 0x14, 0x2c, 0xf4, 0x83, 0x8c, 0x87, 0x5e, 0xc7, 0x31, 0x8e,
	0x61, 0x9d, 0x64, 0x36, 0x1f, 0x45, 0x7f, 0x4f, 0x7f, 0x49, 0xff, 0x59, 0x49, 0x26, 0xe3, 0xe8,
	0xd2, 0xbd, 0xbd, 0xcf, 0x9b, 0xe1, 0xf9, 0x4a, 0x06, 0x5e, 0x6a, 0xa3, 0x6c, 0x6e, 0xac, 0x62,
	0xb3, 0x4a, 0x49, 0x23, 0x71, 0xe7, 0x70, 0xca, 0x04, 0xd1, 0x70, 0xb3, 0xc8, 0x73, 0x69, 0x85,
	0xc1, 0x18, 0x3a, 0xd9, 0x76, 0xab, 0x92, 0x68, 0x12, 0x4d, 0xfb, 0xd4, 0xcf, 0x6e, 0xf7, 0xc0,
	0xc5, 0x36, 0x41, 0x93, 0x68, 0x1a, 0x53, 0x3f, 0xe3, 0x04, 0x6e, 0x36, 0xd9, 0x21, 0x13, 0x39,
	0x4b, 0xe2, 0x49, 0x34, 0x45, 0xb4, 0x81, 0xf8, 0x0d, 0xf4, 0x2a, 0xbb, 0x79, 0x60, 0xa7, 0xa4,
	0xe3, 0x39, 0x02, 0xc2, 0x77, 0x80, 0x56, 0xcb, 0xa4, 0xeb, 0x77, 0x68, 0xb5, 0x24, 0x7f, 0x22,
	0x18, 0x2c, 0x0e, 0x07, 0x99, 0x67, 0x86, 0xa5, 0x45, 0x73, 0x1e, 0x35, 0xe7, 0x4e, 0xa1, 0x52,
	0x5c, 0xaa, 0xd5, 0xd2, 0x0b, 0xf7, 0x69, 0x03, 0xcf, 0x1e, 0xe3, 0x0b, 0x8f, 0xf7, 0xd0, 0xc9,
	0xd2, 0xa2, 0xd6, 0x1c, 0xcc, 0x47, 0x33, 0x97, 0x6b, 0xf6, 0x8d, 0x95, 0x1b, 0xa6, 0x56, 0x62,
	0x27, 0xa9, 0x3f, 0xc5, 0x63, 0xe8, 0xef, 0x33, 0x9d, 0xf2, 0x42, 0xb0, 0x6d, 0xd2, 0x9d, 0xc4,
	0xd3, 0x5b, 0xda, 0x2e, 0x1c, 0xef, 0x3e, 0xd3, 0xfb, 0xa4, 0x57, 0xf3, 0xba, 0x99, 0xbc, 0x07,
	0x68, 0x59, 0x30, 0x01, 0x54, 0xea, 0x24, 0x9a, 0xc4, 0xd3, 0xc1, 0x1c, 0x5f, 0x6a, 0xa4, 0x5c,
	0x14, 0x07, 0x46, 0x51, 0xa9, 0xc9, 0x4f, 0x18, 0x5e, 0xee, 0xf0, 0x5b, 0xb8, 0xd5, 0xec, 0xd1,
	0x32, 0x57, 0x55, 0xe4, 0x1b, 0x3c, 0xe3, 0x73, 0x12, 0x74, 0x91, 0x64, 0x04, 0xb1, 0x54, 0x26,
	0xb4, 0xea, 0x46, 0xf2, 0x1b, 0xee, 0x52, 0x5e, 0xb0, 0x2f, 0x52, 0xa5, 0xc5, 0xc9, 0xfb, 0xf8,
	0xdf, 0x2d, 0xd5, 0xfd, 0xa1, 0x73, 0x7f, 0x23, 0x88, 0x35, 0x7b, 0xf4, 0x3c, 0x31, 0x75, 0xa3,
	0x4b, 0x6f, 0x78, 0xc9, 0xb4, 0xc9, 0xca, 0xca, 0x17, 0x15, 0xd3, 0x76, 0xd1, 0xe8, 0x76, 0x5b,
	0xdd, 0xaf, 0x00, 0xad, 0x2e, 0xbe, 0x07, 0xa4, 0xb9, 0x57, 0x1c, 0xcc, 0x5f, 0xd7, 0xd9, 0xaf,
	0x5d, 0x51, 0xa4, 0xb9, 0xbb, 0x7d, 0x9d, 0x2b, 0x5e, 0x99, 0xe0, 0x24, 0x20, 0xf2, 0xf1, 0x32,
	0x43, 0x6a, 0xec, 0x06, 0x13, 0xe7, 0xcf, 0x84, 0x32, 0x47, 0x4f, 0x09, 0x9d, 0x63, 0x43, 0x16,
	0xf0, 0x62, 0xad, 0x32, 0xb1, 0x63, 0xea, 0x87, 0x35, 0x95, 0x35, 0xee, 0x51, 0x48, 0x6b, 0x16,
	0x6d, 0xf6, 0x06, 0x3a, 0x61, 0x69, 0xcd, 0x77, 0x5b, 0x7a, 0x61, 0x44, 0x03, 0x22, 0x7f, 0x23,
	0xe8, 0xad, 0x8f, 0xbe, 0xb5, 0xab, 0xfc, 0xd1, 0xd3, 0xfc, 0x63, 0xe8, 0x73, 0x51, 0x05, 0xf2,
	0xda, 0x7c, 0xbb, 0xc0, 0x04, 0x86, 0x1e, 0x7c, 0xbe, 0x7a, 0xf4, 0x57, 0x3b, 0x67, 0x41, 0x48,
	0xeb, 0x4e, 0x3b, 0xb5, 0x85, 0x1a, 0xe1, 0x77, 0xde, 0x5a, 0x65, 0x8d, 0x7f, 0x72, 0x83, 0xf9,
	0xab, 0x3a, 0xec, 0x55, 0x32, 0x1a, 0x3e, 0x71, 0x57, 0xcb, 0xc5, 0x4e, 0x36, 0x8f, 0xd0, 0xcd,
	0xe4, 0x13, 0xa0, 0xf5, 0x2f, 0x3c, 0x06, 0x64, 0x8e, 0xe1, 0x02, 0x86, 0x81, 0xe2, 0x58, 0x17,
	0x6f, 0x8e, 0xcf, 0x15, 0xbf, 0xe9, 0xf9, 0x1f, 0xfd, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x11, 0x6f, 0x9d, 0xfe, 0xfb, 0x03, 0x00, 0x00,
}
