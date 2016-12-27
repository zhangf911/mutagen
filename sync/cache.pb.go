// Code generated by protoc-gen-go.
// source: cache.proto
// DO NOT EDIT!

/*
Package sync is a generated protocol buffer package.

It is generated from these files:
	cache.proto
	entry.proto

It has these top-level messages:
	CacheEntry
	Cache
	Entry
*/
package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CacheEntry struct {
	Mode             uint32                     `protobuf:"varint,1,opt,name=mode" json:"mode,omitempty"`
	ModificationTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=modificationTime" json:"modificationTime,omitempty"`
	Size             uint64                     `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Digest           []byte                     `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CacheEntry) Reset()                    { *m = CacheEntry{} }
func (m *CacheEntry) String() string            { return proto.CompactTextString(m) }
func (*CacheEntry) ProtoMessage()               {}
func (*CacheEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CacheEntry) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *CacheEntry) GetModificationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ModificationTime
	}
	return nil
}

func (m *CacheEntry) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CacheEntry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Cache struct {
	Entries map[string]*CacheEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Cache) Reset()                    { *m = Cache{} }
func (m *Cache) String() string            { return proto.CompactTextString(m) }
func (*Cache) ProtoMessage()               {}
func (*Cache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Cache) GetEntries() map[string]*CacheEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*CacheEntry)(nil), "sync.CacheEntry")
	proto.RegisterType((*Cache)(nil), "sync.Cache")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0x51, 0xec, 0xa4, 0xf4, 0x9c, 0x82, 0xd1, 0x50, 0x84, 0x97, 0x8a, 0x0c, 0x45, 0x93,
	0x02, 0xee, 0x52, 0xba, 0x96, 0x74, 0xea, 0x24, 0xf2, 0x02, 0x8e, 0x7d, 0x71, 0x45, 0x23, 0x2b,
	0x58, 0x4a, 0xc1, 0x7d, 0x83, 0xee, 0x7d, 0xe0, 0x62, 0xa9, 0x06, 0x43, 0xb6, 0x5f, 0xfa, 0xef,
	0x3e, 0xee, 0x83, 0xac, 0xae, 0xea, 0x0f, 0x94, 0xe7, 0xde, 0x7a, 0x4b, 0x53, 0x37, 0x74, 0x75,
	0xf1, 0xd0, 0x5a, 0xdb, 0x9e, 0x70, 0x1b, 0xfe, 0x0e, 0x97, 0xe3, 0xd6, 0x6b, 0x83, 0xce, 0x57,
	0xe6, 0x1c, 0xc7, 0x36, 0xbf, 0x04, 0xe0, 0x75, 0x5c, 0xdb, 0x75, 0xbe, 0x1f, 0x28, 0x85, 0xd4,
	0xd8, 0x06, 0x19, 0xe1, 0x44, 0xdc, 0xa9, 0x90, 0xe9, 0x1b, 0xe4, 0xc6, 0x36, 0xfa, 0xa8, 0xeb,
	0xca, 0x6b, 0xdb, 0xed, 0xb5, 0x41, 0xb6, 0xe0, 0x44, 0x64, 0x65, 0x21, 0x23, 0x5e, 0x4e, 0x78,
	0xb9, 0x9f, 0xf0, 0xea, 0x6a, 0x67, 0x64, 0x3b, 0xfd, 0x8d, 0x2c, 0xe1, 0x44, 0xa4, 0x2a, 0x64,
	0x7a, 0x0f, 0xab, 0x46, 0xb7, 0xe8, 0x3c, 0x4b, 0x39, 0x11, 0x6b, 0xf5, 0xff, 0xda, 0xfc, 0x10,
	0x58, 0x86, 0xb3, 0x68, 0x09, 0x37, 0xd8, 0xf9, 0x5e, 0xa3, 0x63, 0x84, 0x27, 0x22, 0x2b, 0x99,
	0x1c, 0xcd, 0x64, 0x68, 0xe5, 0x2e, 0x56, 0xe1, 0x78, 0x35, 0x0d, 0x16, 0xef, 0xb0, 0x9e, 0x17,
	0x34, 0x87, 0xe4, 0x13, 0x87, 0x20, 0x75, 0xab, 0xc6, 0x48, 0x1f, 0x61, 0xf9, 0x55, 0x9d, 0x2e,
	0x93, 0x48, 0x3e, 0x63, 0x46, 0x56, 0xac, 0x5f, 0x16, 0xcf, 0xe4, 0xb0, 0x0a, 0x76, 0x4f, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xd2, 0x3d, 0x70, 0x5f, 0x01, 0x00, 0x00,
}
