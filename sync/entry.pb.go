// Code generated by protoc-gen-gogo.
// source: entry.proto
// DO NOT EDIT!

package sync

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EntryKind int32

const (
	EntryKind_Directory EntryKind = 0
	EntryKind_File      EntryKind = 1
)

var EntryKind_name = map[int32]string{
	0: "Directory",
	1: "File",
}
var EntryKind_value = map[string]int32{
	"Directory": 0,
	"File":      1,
}

func (x EntryKind) String() string {
	return proto.EnumName(EntryKind_name, int32(x))
}
func (EntryKind) EnumDescriptor() ([]byte, []int) { return fileDescriptorEntry, []int{0} }

type Entry struct {
	Kind       EntryKind         `protobuf:"varint,1,opt,name=kind,proto3,enum=sync.EntryKind" json:"kind,omitempty"`
	Executable bool              `protobuf:"varint,2,opt,name=executable,proto3" json:"executable,omitempty"`
	Digest     []byte            `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Contents   map[string]*Entry `protobuf:"bytes,4,rep,name=contents" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptorEntry, []int{0} }

func (m *Entry) GetKind() EntryKind {
	if m != nil {
		return m.Kind
	}
	return EntryKind_Directory
}

func (m *Entry) GetExecutable() bool {
	if m != nil {
		return m.Executable
	}
	return false
}

func (m *Entry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Entry) GetContents() map[string]*Entry {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "sync.Entry")
	proto.RegisterEnum("sync.EntryKind", EntryKind_name, EntryKind_value)
}
func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEntry(dAtA, i, uint64(m.Kind))
	}
	if m.Executable {
		dAtA[i] = 0x10
		i++
		if m.Executable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Digest) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEntry(dAtA, i, uint64(len(m.Digest)))
		i += copy(dAtA[i:], m.Digest)
	}
	if len(m.Contents) > 0 {
		for k, _ := range m.Contents {
			dAtA[i] = 0x22
			i++
			v := m.Contents[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovEntry(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovEntry(uint64(len(k))) + msgSize
			i = encodeVarintEntry(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintEntry(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintEntry(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func encodeFixed64Entry(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Entry(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEntry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Entry) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovEntry(uint64(m.Kind))
	}
	if m.Executable {
		n += 2
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovEntry(uint64(l))
	}
	if len(m.Contents) > 0 {
		for k, v := range m.Contents {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovEntry(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovEntry(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovEntry(uint64(mapEntrySize))
		}
	}
	return n
}

func sovEntry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEntry(x uint64) (n int) {
	return sovEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (EntryKind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Executable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Executable = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = append(m.Digest[:0], dAtA[iNdEx:postIndex]...)
			if m.Digest == nil {
				m.Digest = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthEntry
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Contents == nil {
				m.Contents = make(map[string]*Entry)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthEntry
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthEntry
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &Entry{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.Contents[mapkey] = mapvalue
			} else {
				var mapvalue *Entry
				m.Contents[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntry   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("entry.proto", fileDescriptorEntry) }

var fileDescriptorEntry = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x3d, 0xd3, 0xce, 0xd0, 0x9e, 0xce, 0x68, 0xc9, 0x42, 0xea, 0x2c, 0x4a, 0xbd, 0x2c,
	0x8a, 0x60, 0x07, 0x2a, 0x82, 0xb8, 0xf4, 0x86, 0xe0, 0x2e, 0x6f, 0x30, 0x4d, 0x63, 0x0d, 0x53,
	0x13, 0x69, 0x53, 0xb1, 0x6f, 0xe8, 0xd2, 0x47, 0x90, 0x82, 0xef, 0x21, 0x4d, 0x87, 0xa1, 0xee,
	0xce, 0x97, 0x2f, 0xf9, 0xf3, 0x27, 0xe8, 0x71, 0xa9, 0xab, 0x36, 0x79, 0xaf, 0x94, 0x56, 0xc4,
	0xae, 0x5b, 0xc9, 0x96, 0x17, 0x85, 0xd0, 0xaf, 0x4d, 0x96, 0x30, 0xf5, 0xb6, 0x2a, 0x54, 0xa1,
	0x56, 0x46, 0x66, 0xcd, 0x8b, 0x21, 0x03, 0x66, 0x1a, 0x0e, 0x9d, 0xfc, 0x02, 0x4e, 0x1f, 0xfa,
	0x10, 0x72, 0x8a, 0xf6, 0x46, 0xc8, 0x3c, 0x80, 0x08, 0xe2, 0xfd, 0xf4, 0x20, 0xe9, 0xd3, 0x12,
	0xa3, 0x9e, 0x85, 0xcc, 0xa9, 0x91, 0x24, 0x44, 0xe4, 0x9f, 0x9c, 0x35, 0x7a, 0x9d, 0x95, 0x3c,
	0x98, 0x44, 0x10, 0x3b, 0x74, 0xb4, 0x42, 0x0e, 0x71, 0x96, 0x8b, 0x82, 0xd7, 0x3a, 0xb0, 0x22,
	0x88, 0xe7, 0x74, 0x4b, 0xe4, 0x0a, 0x1d, 0xa6, 0xa4, 0xe6, 0x52, 0xd7, 0x81, 0x1d, 0x59, 0xb1,
	0x97, 0x1e, 0x8d, 0x2e, 0x48, 0xee, 0xb6, 0xce, 0x10, 0xdd, 0x6d, 0x5d, 0x3e, 0xe1, 0xe2, 0x9f,
	0x22, 0x3e, 0x5a, 0x1b, 0xde, 0x9a, 0x8e, 0x2e, 0xed, 0x47, 0x72, 0x8c, 0xd3, 0x8f, 0x75, 0xd9,
	0x0c, 0x65, 0xbc, 0xd4, 0x1b, 0xc5, 0xd2, 0xc1, 0xdc, 0x4c, 0xae, 0xe1, 0xfc, 0x0c, 0xdd, 0xdd,
	0x5b, 0xc8, 0x02, 0xdd, 0x7b, 0x51, 0x71, 0xa6, 0x55, 0xd5, 0xfa, 0x7b, 0xc4, 0x41, 0xfb, 0x51,
	0x94, 0xdc, 0x87, 0xdb, 0xf9, 0x57, 0x17, 0xc2, 0x77, 0x17, 0xc2, 0x4f, 0x17, 0x42, 0x36, 0x33,
	0x5f, 0x74, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xac, 0xa7, 0xd7, 0x66, 0x01, 0x00, 0x00,
}