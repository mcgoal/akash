// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/resource.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Unit stores cpu, memory and storage metrics
type Unit struct {
	CPU     uint32 `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu" yaml:"cpu"`
	Memory  uint64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory" yaml:"memory"`
	Storage uint64 `protobuf:"varint,3,opt,name=storage,proto3" json:"storage" yaml:"storage"`
}

func (m *Unit) Reset()         { *m = Unit{} }
func (m *Unit) String() string { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()    {}
func (*Unit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e6470e3190728b6, []int{0}
}
func (m *Unit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unit.Merge(m, src)
}
func (m *Unit) XXX_Size() int {
	return m.Size()
}
func (m *Unit) XXX_DiscardUnknown() {
	xxx_messageInfo_Unit.DiscardUnknown(m)
}

var xxx_messageInfo_Unit proto.InternalMessageInfo

func (m *Unit) GetCPU() uint32 {
	if m != nil {
		return m.CPU
	}
	return 0
}

func (m *Unit) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Unit) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func init() {
	proto.RegisterType((*Unit)(nil), "akash.Unit")
}

func init() { proto.RegisterFile("akash/resource.proto", fileDescriptor_1e6470e3190728b6) }

var fileDescriptor_1e6470e3190728b6 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x8b, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x45, 0xf4, 0x41, 0x2c, 0x88,
	0xa4, 0xd2, 0x2a, 0x46, 0x2e, 0x96, 0xd0, 0xbc, 0xcc, 0x12, 0x21, 0x3d, 0x2e, 0xe6, 0xe4, 0x82,
	0x52, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x5e, 0x27, 0x99, 0x47, 0xf7, 0xe4, 0x99, 0x9d, 0x03, 0x42,
	0x5f, 0xdd, 0x93, 0x07, 0x89, 0x7e, 0xba, 0x27, 0xcf, 0x55, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x94,
	0x5c, 0x50, 0xaa, 0x14, 0x04, 0x12, 0x12, 0x32, 0xe6, 0x62, 0xcb, 0x4d, 0xcd, 0xcd, 0x2f, 0xaa,
	0x94, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x71, 0x92, 0x7e, 0x75, 0x4f, 0x1e, 0x2a, 0xf2, 0xe9, 0x9e,
	0x3c, 0x2f, 0x44, 0x39, 0x84, 0xaf, 0x14, 0x04, 0x95, 0x10, 0x32, 0xe7, 0x62, 0x2f, 0x2e, 0xc9,
	0x2f, 0x4a, 0x4c, 0x4f, 0x95, 0x60, 0x06, 0xeb, 0x92, 0x7d, 0x75, 0x4f, 0x1e, 0x26, 0xf4, 0xe9,
	0x9e, 0x3c, 0x1f, 0x44, 0x1b, 0x54, 0x40, 0x29, 0x08, 0x26, 0x65, 0xc5, 0xf2, 0x62, 0x81, 0x3c,
	0xa3, 0x93, 0xf9, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xc9, 0xa6, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x97, 0x15, 0x25, 0xe7, 0x64, 0xeb,
	0x43, 0xc2, 0xa2, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x59, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0xe7, 0xbb, 0x58, 0x21, 0x01, 0x00, 0x00,
}

func (this *Unit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Unit)
	if !ok {
		that2, ok := that.(Unit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CPU != that1.CPU {
		return false
	}
	if this.Memory != that1.Memory {
		return false
	}
	if this.Storage != that1.Storage {
		return false
	}
	return true
}
func (m *Unit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Unit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Storage != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.Storage))
		i--
		dAtA[i] = 0x18
	}
	if m.Memory != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.Memory))
		i--
		dAtA[i] = 0x10
	}
	if m.CPU != 0 {
		i = encodeVarintResource(dAtA, i, uint64(m.CPU))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Unit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CPU != 0 {
		n += 1 + sovResource(uint64(m.CPU))
	}
	if m.Memory != 0 {
		n += 1 + sovResource(uint64(m.Memory))
	}
	if m.Storage != 0 {
		n += 1 + sovResource(uint64(m.Storage))
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Unit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPU", wireType)
			}
			m.CPU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CPU |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memory", wireType)
			}
			m.Memory = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Memory |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			m.Storage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Storage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResource
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
			if length < 0 {
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)
