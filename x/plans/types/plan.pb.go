// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plans/plan.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Plan struct {
	Index                    string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Duration                 uint64     `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Block                    uint64     `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	Price                    types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	ComputeUnits             uint64     `protobuf:"varint,5,opt,name=compute_units,json=computeUnits,proto3" json:"compute_units,omitempty"`
	ComputeUnitsPerEpoch     uint64     `protobuf:"varint,6,opt,name=compute_units_per_epoch,json=computeUnitsPerEpoch,proto3" json:"compute_units_per_epoch,omitempty"`
	ServicersToPair          uint64     `protobuf:"varint,7,opt,name=servicers_to_pair,json=servicersToPair,proto3" json:"servicers_to_pair,omitempty"`
	AllowOveruse             bool       `protobuf:"varint,8,opt,name=allow_overuse,json=allowOveruse,proto3" json:"allow_overuse,omitempty"`
	OveruseRate              uint64     `protobuf:"varint,9,opt,name=overuse_rate,json=overuseRate,proto3" json:"overuse_rate,omitempty"`
	Name                     string     `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description              string     `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Type                     string     `protobuf:"bytes,12,opt,name=type,proto3" json:"type,omitempty"`
	AnnualDiscountPercentage uint64     `protobuf:"varint,13,opt,name=annual_discount_percentage,json=annualDiscountPercentage,proto3" json:"annual_discount_percentage,omitempty"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5909a10cd0e3497, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Plan) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Plan) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Plan) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func (m *Plan) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

func (m *Plan) GetComputeUnitsPerEpoch() uint64 {
	if m != nil {
		return m.ComputeUnitsPerEpoch
	}
	return 0
}

func (m *Plan) GetServicersToPair() uint64 {
	if m != nil {
		return m.ServicersToPair
	}
	return 0
}

func (m *Plan) GetAllowOveruse() bool {
	if m != nil {
		return m.AllowOveruse
	}
	return false
}

func (m *Plan) GetOveruseRate() uint64 {
	if m != nil {
		return m.OveruseRate
	}
	return 0
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Plan) GetAnnualDiscountPercentage() uint64 {
	if m != nil {
		return m.AnnualDiscountPercentage
	}
	return 0
}

func init() {
	proto.RegisterType((*Plan)(nil), "lavanet.lava.plans.Plan")
}

func init() { proto.RegisterFile("plans/plan.proto", fileDescriptor_e5909a10cd0e3497) }

var fileDescriptor_e5909a10cd0e3497 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x8e, 0x69, 0x52, 0x52, 0x27, 0x15, 0x60, 0x45, 0xc2, 0xe4, 0xb0, 0x2c, 0x20, 0xa4, 0x88,
	0xc3, 0xae, 0x0a, 0xea, 0x8d, 0x53, 0x0b, 0x67, 0xa2, 0x15, 0x5c, 0xb8, 0xac, 0xbc, 0xce, 0x28,
	0xb5, 0xd8, 0x78, 0x2c, 0xdb, 0x1b, 0xca, 0x5b, 0xf0, 0x18, 0x3c, 0x4a, 0x8f, 0x3d, 0x72, 0x01,
	0xa1, 0xe4, 0x45, 0x90, 0xed, 0x55, 0xd4, 0x5e, 0xd6, 0x33, 0xdf, 0xcf, 0xec, 0x67, 0x6b, 0xe8,
	0x63, 0xd3, 0x0a, 0xed, 0xca, 0xf0, 0x2d, 0x8c, 0x45, 0x8f, 0x8c, 0xb5, 0x62, 0x2b, 0x34, 0xf8,
	0x22, 0x9c, 0x45, 0xa4, 0xe7, 0xb3, 0x35, 0xae, 0x31, 0xd2, 0x65, 0xa8, 0x92, 0x72, 0x9e, 0x49,
	0x74, 0x1b, 0x74, 0x65, 0x23, 0x1c, 0x94, 0xdb, 0xb3, 0x06, 0xbc, 0x38, 0x2b, 0x25, 0xaa, 0x7e,
	0xd2, 0xcb, 0x3f, 0x47, 0x74, 0xb8, 0x6c, 0x85, 0x66, 0x33, 0x3a, 0x52, 0x7a, 0x05, 0xd7, 0x9c,
	0xe4, 0x64, 0x71, 0x52, 0xa5, 0x86, 0xcd, 0xe9, 0x78, 0xd5, 0x59, 0xe1, 0x15, 0x6a, 0xfe, 0x20,
	0x27, 0x8b, 0x61, 0x75, 0xe8, 0x83, 0xa3, 0x69, 0x51, 0x7e, 0xe3, 0x47, 0x91, 0x48, 0x0d, 0x3b,
	0xa7, 0x23, 0x63, 0x95, 0x04, 0x3e, 0xcc, 0xc9, 0x62, 0xf2, 0xf6, 0x59, 0x91, 0x02, 0x14, 0x21,
	0x40, 0xd1, 0x07, 0x28, 0x2e, 0x51, 0xe9, 0x8b, 0xe1, 0xcd, 0xdf, 0xe7, 0x83, 0x2a, 0xa9, 0xd9,
	0x2b, 0x7a, 0x2a, 0x71, 0x63, 0x3a, 0x0f, 0x75, 0xa7, 0x95, 0x77, 0x7c, 0x14, 0x87, 0x4e, 0x7b,
	0xf0, 0x4b, 0xc0, 0xd8, 0x39, 0x7d, 0x7a, 0x4f, 0x54, 0x1b, 0xb0, 0x35, 0x18, 0x94, 0x57, 0xfc,
	0x38, 0xca, 0x67, 0x77, 0xe5, 0x4b, 0xb0, 0x1f, 0x03, 0xc7, 0xde, 0xd0, 0x27, 0x0e, 0xec, 0x56,
	0x49, 0xb0, 0xae, 0xf6, 0x58, 0x1b, 0xa1, 0x2c, 0x7f, 0x18, 0x0d, 0x8f, 0x0e, 0xc4, 0x67, 0x5c,
	0x0a, 0x65, 0x43, 0x0e, 0xd1, 0xb6, 0xf8, 0xbd, 0xc6, 0x2d, 0xd8, 0xce, 0x01, 0x1f, 0xe7, 0x64,
	0x31, 0xae, 0xa6, 0x11, 0xfc, 0x94, 0x30, 0xf6, 0x82, 0x4e, 0x7b, 0xba, 0xb6, 0xc2, 0x03, 0x3f,
	0x89, 0xb3, 0x26, 0x3d, 0x56, 0x09, 0x0f, 0x8c, 0xd1, 0xa1, 0x16, 0x1b, 0xe0, 0x34, 0xbe, 0x66,
	0xac, 0x59, 0x4e, 0x27, 0x2b, 0x70, 0xd2, 0x2a, 0x13, 0xdf, 0x73, 0x12, 0xa9, 0xbb, 0x50, 0x70,
	0xf9, 0x1f, 0x06, 0xf8, 0x34, 0xb9, 0x42, 0xcd, 0xde, 0xd3, 0xb9, 0xd0, 0xba, 0x13, 0x6d, 0xbd,
	0x52, 0x4e, 0x62, 0xa7, 0x7d, 0xb8, 0xb6, 0x04, 0xed, 0xc5, 0x1a, 0xf8, 0x69, 0xfc, 0x35, 0x4f,
	0x8a, 0x0f, 0xbd, 0x60, 0x79, 0xe0, 0x2f, 0x2e, 0x7f, 0xed, 0x32, 0x72, 0xb3, 0xcb, 0xc8, 0xed,
	0x2e, 0x23, 0xff, 0x76, 0x19, 0xf9, 0xb9, 0xcf, 0x06, 0xb7, 0xfb, 0x6c, 0xf0, 0x7b, 0x9f, 0x0d,
	0xbe, 0xbe, 0x5e, 0x2b, 0x7f, 0xd5, 0x35, 0x85, 0xc4, 0x4d, 0xd9, 0xaf, 0x54, 0x3c, 0xcb, 0xeb,
	0x32, 0xed, 0x5c, 0x48, 0xe0, 0x9a, 0xe3, 0xb8, 0x2b, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x26, 0xc4, 0x7c, 0x6c, 0x89, 0x02, 0x00, 0x00,
}

func (this *Plan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plan)
	if !ok {
		that2, ok := that.(Plan)
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
	if this.Index != that1.Index {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.Block != that1.Block {
		return false
	}
	if !this.Price.Equal(&that1.Price) {
		return false
	}
	if this.ComputeUnits != that1.ComputeUnits {
		return false
	}
	if this.ComputeUnitsPerEpoch != that1.ComputeUnitsPerEpoch {
		return false
	}
	if this.ServicersToPair != that1.ServicersToPair {
		return false
	}
	if this.AllowOveruse != that1.AllowOveruse {
		return false
	}
	if this.OveruseRate != that1.OveruseRate {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.AnnualDiscountPercentage != that1.AnnualDiscountPercentage {
		return false
	}
	return true
}
func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AnnualDiscountPercentage != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.AnnualDiscountPercentage))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x52
	}
	if m.OveruseRate != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.OveruseRate))
		i--
		dAtA[i] = 0x48
	}
	if m.AllowOveruse {
		i--
		if m.AllowOveruse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.ServicersToPair != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.ServicersToPair))
		i--
		dAtA[i] = 0x38
	}
	if m.ComputeUnitsPerEpoch != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.ComputeUnitsPerEpoch))
		i--
		dAtA[i] = 0x30
	}
	if m.ComputeUnits != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Block != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if m.Duration != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovPlan(uint64(m.Duration))
	}
	if m.Block != 0 {
		n += 1 + sovPlan(uint64(m.Block))
	}
	l = m.Price.Size()
	n += 1 + l + sovPlan(uint64(l))
	if m.ComputeUnits != 0 {
		n += 1 + sovPlan(uint64(m.ComputeUnits))
	}
	if m.ComputeUnitsPerEpoch != 0 {
		n += 1 + sovPlan(uint64(m.ComputeUnitsPerEpoch))
	}
	if m.ServicersToPair != 0 {
		n += 1 + sovPlan(uint64(m.ServicersToPair))
	}
	if m.AllowOveruse {
		n += 2
	}
	if m.OveruseRate != 0 {
		n += 1 + sovPlan(uint64(m.OveruseRate))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.AnnualDiscountPercentage != 0 {
		n += 1 + sovPlan(uint64(m.AnnualDiscountPercentage))
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnitsPerEpoch", wireType)
			}
			m.ComputeUnitsPerEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnitsPerEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServicersToPair", wireType)
			}
			m.ServicersToPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServicersToPair |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowOveruse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowOveruse = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OveruseRate", wireType)
			}
			m.OveruseRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OveruseRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualDiscountPercentage", wireType)
			}
			m.AnnualDiscountPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnnualDiscountPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)