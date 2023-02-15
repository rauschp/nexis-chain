// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/block.proto

package pb

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

type Block struct {
	Header               *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb532b7e63676ceb, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Header struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	PreviousHash         []byte   `protobuf:"bytes,4,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb532b7e63676ceb, []int{1}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Header) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *Header) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*Header)(nil), "Header")
}

func init() {
	proto.RegisterFile("proto/block.proto", fileDescriptor_bb532b7e63676ceb)
}

var fileDescriptor_bb532b7e63676ceb = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4e, 0x03, 0x21,
	0x10, 0x86, 0x43, 0x77, 0x97, 0xc6, 0x29, 0x17, 0xe7, 0xa0, 0xc4, 0x98, 0x48, 0xf6, 0xc4, 0x89,
	0x9a, 0xfa, 0x06, 0x3d, 0xf5, 0x4c, 0x3c, 0xf5, 0xc6, 0x56, 0x22, 0x44, 0xbb, 0x10, 0xc0, 0x3e,
	0x87, 0x8f, 0x6c, 0x96, 0xad, 0x69, 0x7b, 0x9b, 0xff, 0xfb, 0x86, 0x9f, 0x0c, 0xdc, 0xc7, 0x14,
	0x4a, 0x58, 0x0f, 0xdf, 0xe1, 0xf0, 0xa5, 0xea, 0xfc, 0xf4, 0x38, 0xa3, 0x92, 0xcc, 0x98, 0xcd,
	0xa1, 0xf8, 0x30, 0xce, 0xa2, 0xdf, 0x43, 0xb7, 0x9d, 0xf6, 0xf0, 0x05, 0xa8, 0xb3, 0xe6, 0xc3,
	0x26, 0x4e, 0x04, 0x91, 0xab, 0xcd, 0x52, 0xed, 0x6a, 0xd4, 0x67, 0x8c, 0xaf, 0xc0, 0xae, 0x9e,
	0x67, 0xbe, 0x10, 0x8d, 0x5c, 0x6d, 0x98, 0x7a, 0xbf, 0x40, 0x7d, 0xb3, 0xd1, 0xff, 0x12, 0xa0,
	0x73, 0x09, 0x72, 0x58, 0x9e, 0x6c, 0xca, 0x3e, 0x8c, 0xb5, 0xbe, 0xd3, 0xff, 0x11, 0x1f, 0xa6,
	0x7f, 0xfd, 0xa7, 0x2b, 0x7c, 0x51, 0xc5, 0x39, 0x21, 0x42, 0xeb, 0x4c, 0x76, 0xbc, 0x11, 0x44,
	0x32, 0x5d, 0x67, 0xec, 0x81, 0xc5, 0x64, 0x4f, 0x3e, 0xfc, 0xe4, 0xdd, 0xe4, 0xda, 0xea, 0x6e,
	0x18, 0x3e, 0xc3, 0x5d, 0xf1, 0x47, 0x9b, 0x8b, 0x39, 0x46, 0xde, 0x09, 0x22, 0x1b, 0x7d, 0x01,
	0x5b, 0xba, 0x6f, 0xd5, 0x3a, 0x0e, 0x03, 0xad, 0xd7, 0xbf, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x89, 0xf4, 0xc9, 0x2b, 0x01, 0x00, 0x00,
}