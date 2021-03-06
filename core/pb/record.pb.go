// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: record.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	record.proto

It has these top-level messages:
	Record
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Record struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Owner     []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorRecord, []int{0} }

func (m *Record) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Record) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "corepb.Record")
}

func init() { proto.RegisterFile("record.proto", fileDescriptorRecord) }

var fileDescriptorRecord = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0x2f, 0x4a, 0x2d, 0x48,
	0x52, 0x0a, 0xe0, 0x62, 0x0b, 0x02, 0x8b, 0x0b, 0x09, 0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0xf9, 0xe5, 0x79, 0xa9,
	0x45, 0x12, 0x4c, 0x60, 0x41, 0x08, 0x47, 0x48, 0x86, 0x8b, 0xb3, 0x24, 0x33, 0x37, 0xb5, 0xb8,
	0x24, 0x31, 0xb7, 0x40, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x08, 0x21, 0x90, 0xc4, 0x06, 0xb6,
	0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x1e, 0x36, 0x02, 0x70, 0x00, 0x00, 0x00,
}
