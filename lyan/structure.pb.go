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
	ASgyProposeInfo
	ASgyPropose
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

// 用以向系统提出新的分配策略
type ASgyProposeInfo struct {
	ContractAddr string      `protobuf:"bytes,1,opt,name=contractAddr" json:"contractAddr,omitempty"`
	PriorID      string      `protobuf:"bytes,2,opt,name=priorID" json:"priorID,omitempty"`
	UserAddr     string      `protobuf:"bytes,3,opt,name=userAddr" json:"userAddr,omitempty"`
	Mi           *MemberInfo `protobuf:"bytes,4,opt,name=mi" json:"mi,omitempty"`
	Timestamp    int64       `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Seq          int64       `protobuf:"varint,6,opt,name=seq" json:"seq,omitempty"`
}

func (m *ASgyProposeInfo) Reset()                    { *m = ASgyProposeInfo{} }
func (m *ASgyProposeInfo) String() string            { return proto.CompactTextString(m) }
func (*ASgyProposeInfo) ProtoMessage()               {}
func (*ASgyProposeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ASgyProposeInfo) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *ASgyProposeInfo) GetPriorID() string {
	if m != nil {
		return m.PriorID
	}
	return ""
}

func (m *ASgyProposeInfo) GetUserAddr() string {
	if m != nil {
		return m.UserAddr
	}
	return ""
}

func (m *ASgyProposeInfo) GetMi() *MemberInfo {
	if m != nil {
		return m.Mi
	}
	return nil
}

func (m *ASgyProposeInfo) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ASgyProposeInfo) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type ASgyPropose struct {
	Asgip  *ASgyProposeInfo `protobuf:"bytes,1,opt,name=asgip" json:"asgip,omitempty"`
	Script string           `protobuf:"bytes,2,opt,name=script" json:"script,omitempty"`
}

func (m *ASgyPropose) Reset()                    { *m = ASgyPropose{} }
func (m *ASgyPropose) String() string            { return proto.CompactTextString(m) }
func (*ASgyPropose) ProtoMessage()               {}
func (*ASgyPropose) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ASgyPropose) GetAsgip() *ASgyProposeInfo {
	if m != nil {
		return m.Asgip
	}
	return nil
}

func (m *ASgyPropose) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

type TranferOutput struct {
	OutAddr string  `protobuf:"bytes,1,opt,name=outAddr" json:"outAddr,omitempty"`
	OutNum  float32 `protobuf:"fixed32,2,opt,name=outNum" json:"outNum,omitempty"`
}

func (m *TranferOutput) Reset()                    { *m = TranferOutput{} }
func (m *TranferOutput) String() string            { return proto.CompactTextString(m) }
func (*TranferOutput) ProtoMessage()               {}
func (*TranferOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
func (*TxInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
func (*TX) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
	proto.RegisterType((*ASgyProposeInfo)(nil), "lyan.ASgyProposeInfo")
	proto.RegisterType((*ASgyPropose)(nil), "lyan.ASgyPropose")
	proto.RegisterType((*TranferOutput)(nil), "lyan.TranferOutput")
	proto.RegisterType((*TxInfo)(nil), "lyan.TxInfo")
	proto.RegisterType((*TX)(nil), "lyan.TX")
}

func init() { proto.RegisterFile("structure.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x72, 0x92, 0x36, 0x37, 0x5d, 0x1b, 0xb4, 0x0f, 0x4c, 0xe9, 0x43, 0x30, 0x7d, 0x08,
	0x14, 0xca, 0xc8, 0xf6, 0xb4, 0xb7, 0x8c, 0x32, 0xc8, 0x60, 0x5b, 0x91, 0xfb, 0xb0, 0x57, 0xc7,
	0x51, 0x5d, 0xd1, 0x58, 0x72, 0xf5, 0x31, 0x92, 0xdf, 0xb3, 0xdf, 0xb1, 0x87, 0xfd, 0xb3, 0x21,
	0x59, 0xfe, 0x2a, 0xcd, 0xde, 0xee, 0xb9, 0x32, 0xe7, 0x9e, 0x73, 0xae, 0x2c, 0x38, 0x53, 0x5a,
	0x9a, 0x4c, 0x1b, 0x49, 0xaf, 0x4b, 0x29, 0xb4, 0xc0, 0x83, 0xed, 0x3e, 0xe5, 0xb1, 0x82, 0xa3,
	0x65, 0x96, 0x09, 0xc3, 0x35, 0xc6, 0x30, 0x48, 0x37, 0x1b, 0x19, 0x05, 0xb3, 0x60, 0x3e, 0x26,
	0xae, 0xb6, 0xbd, 0x47, 0xc6, 0x37, 0x11, 0x9a, 0x05, 0xf3, 0x90, 0xb8, 0x1a, 0x47, 0x70, 0xb4,
	0x4e, 0xb7, 0x29, 0xcf, 0x68, 0x14, 0xce, 0x82, 0x39, 0x22, 0x35, 0xc4, 0xef, 0x60, 0x54, 0x9a,
	0xf5, 0x23, 0xdd, 0x47, 0x03, 0xc7, 0xe1, 0x11, 0x3e, 0x05, 0xb4, 0xba, 0x89, 0x86, 0xae, 0x87,
	0x56, 0x37, 0xf1, 0xef, 0x00, 0x26, 0xcb, 0xed, 0x56, 0x64, 0xa9, 0xa6, 0x49, 0x5e, 0x9f, 0x07,
	0xf5, 0xb9, 0x9d, 0x50, 0x4a, 0x26, 0xe4, 0xea, 0xc6, 0x0d, 0x1e, 0x93, 0x1a, 0x36, 0x1a, 0xc3,
	0x8e, 0xc6, 0x4b, 0x18, 0xa4, 0x49, 0x5e, 0xcd, 0x9c, 0x2c, 0xa6, 0xd7, 0xd6, 0xd7, 0xf5, 0x37,
	0x5a, 0xac, 0xa9, 0x5c, 0xf1, 0x7b, 0x41, 0xdc, 0x29, 0xbe, 0x80, 0xf1, 0x43, 0xaa, 0x12, 0x96,
	0x73, 0xba, 0x89, 0x86, 0xb3, 0x70, 0x7e, 0x4c, 0xda, 0x86, 0xe5, 0x7d, 0x48, 0xd5, 0x43, 0x34,
	0xaa, 0x78, 0x6d, 0x1d, 0xbf, 0x07, 0x68, 0x59, 0x70, 0x0c, 0xa8, 0x50, 0x51, 0x30, 0x0b, 0xe7,
	0x93, 0x05, 0xee, 0xce, 0x48, 0x18, 0xcf, 0xb7, 0x94, 0xa0, 0x42, 0xc5, 0xb7, 0x70, 0xd2, 0xed,
	0xe1, 0x73, 0x38, 0x56, 0xf4, 0xc9, 0x50, 0x1b, 0x55, 0xe0, 0x12, 0x6c, 0x70, 0xe3, 0x04, 0x75,
	0x9c, 0x4c, 0x21, 0x14, 0x52, 0xfb, 0x54, 0x6d, 0x19, 0xff, 0x82, 0xd3, 0x84, 0xe5, 0xf4, 0x8b,
	0x90, 0x49, 0xbe, 0x77, 0x3a, 0x5e, 0xda, 0x52, 0x95, 0x1f, 0x6a, 0xf2, 0x9b, 0x42, 0xa8, 0xe8,
	0x93, 0xe3, 0x09, 0x89, 0x2d, 0xad, 0x7b, 0xcd, 0x0a, 0xaa, 0x74, 0x5a, 0x94, 0x2e, 0xa8, 0x90,
	0xb4, 0x8d, 0x7a, 0xee, 0xb0, 0x9d, 0xfb, 0x15, 0xa0, 0x9d, 0x8b, 0x2f, 0x01, 0x29, 0xe6, 0x26,
	0x4e, 0x16, 0x6f, 0x2a, 0xef, 0x7d, 0x55, 0x04, 0x29, 0x66, 0xb7, 0xaf, 0x32, 0xc9, 0x4a, 0xed,
	0x95, 0x78, 0x14, 0x7f, 0xec, 0x7a, 0x48, 0xb4, 0x59, 0xe3, 0xd8, 0xea, 0xd3, 0x3e, 0xcc, 0xe9,
	0x73, 0x42, 0xab, 0x58, 0xc7, 0x7f, 0x02, 0x38, 0x5b, 0x26, 0xf9, 0xfe, 0x56, 0x8a, 0x52, 0x28,
	0xea, 0x77, 0x70, 0x92, 0x09, 0xae, 0x65, 0x9a, 0xe9, 0x65, 0x9b, 0x41, 0xaf, 0xf7, 0x9f, 0xbb,
	0x73, 0x0e, 0xc7, 0x46, 0x51, 0xb9, 0x6c, 0xef, 0x4f, 0x83, 0xf1, 0x0c, 0x50, 0xc1, 0x0e, 0xde,
	0x20, 0x54, 0xb0, 0x7e, 0x82, 0xc3, 0x17, 0x12, 0xb4, 0x89, 0x8f, 0x9a, 0xc4, 0x63, 0x02, 0x93,
	0x8e, 0x7c, 0x7c, 0x05, 0xc3, 0x54, 0xe5, 0xac, 0xf4, 0x29, 0xbe, 0xad, 0x66, 0x3c, 0x33, 0x48,
	0xaa, 0x6f, 0x0e, 0x26, 0xb9, 0x84, 0x57, 0x77, 0x32, 0xe5, 0xf7, 0x54, 0xfe, 0x30, 0xba, 0x34,
	0xda, 0x9a, 0x15, 0xa6, 0x9b, 0x45, 0x0d, 0x2d, 0x85, 0x30, 0xfa, 0xbb, 0x29, 0x1c, 0x05, 0x22,
	0x1e, 0xc5, 0x7f, 0x03, 0x18, 0xdd, 0xed, 0x5c, 0x9a, 0x3d, 0x47, 0xc1, 0x73, 0x47, 0x17, 0x30,
	0x66, 0xbc, 0xf4, 0xe4, 0x95, 0x8c, 0xb6, 0x61, 0x37, 0xe1, 0xc0, 0xe7, 0xde, 0x43, 0xd0, 0xeb,
	0x59, 0x09, 0x5c, 0x18, 0x7b, 0x3a, 0xa8, 0x24, 0x54, 0x08, 0x5f, 0x39, 0x69, 0xa5, 0xd1, 0xee,
	0x37, 0x9c, 0x2c, 0x5e, 0x57, 0x59, 0xf4, 0x9c, 0x11, 0xff, 0x89, 0xbd, 0xee, 0x8c, 0xdf, 0x8b,
	0xfa, 0xc7, 0xb4, 0x75, 0xfc, 0x09, 0xd0, 0xdd, 0x4f, 0x7c, 0x01, 0x48, 0xef, 0x7c, 0x9c, 0x27,
	0x9e, 0x62, 0x57, 0xad, 0x4b, 0xef, 0x0e, 0x45, 0xb8, 0x1e, 0xb9, 0xc7, 0xef, 0xc3, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x48, 0x21, 0x8e, 0x0f, 0x05, 0x00, 0x00,
}
