// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/v3/x/fixationstore/types"
	types1 "github.com/lavanet/lava/v3/x/timerstore/types"
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

// GenesisState defines the subscription module's genesis state.
type GenesisState struct {
	Params      Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SubsFS      types.GenesisState  `protobuf:"bytes,2,opt,name=subsFS,proto3" json:"subsFS"`
	SubsTS      types1.GenesisState `protobuf:"bytes,3,opt,name=subsTS,proto3" json:"subsTS"`
	CuTrackerFS types.GenesisState  `protobuf:"bytes,4,opt,name=cuTrackerFS,proto3" json:"cuTrackerFS"`
	CuTrackerTS types1.GenesisState `protobuf:"bytes,5,opt,name=cuTrackerTS,proto3" json:"cuTrackerTS"`
	Adjustments []Adjustment        `protobuf:"bytes,6,rep,name=adjustments,proto3" json:"adjustments"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc6c60f9c112fe52, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSubsFS() types.GenesisState {
	if m != nil {
		return m.SubsFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetSubsTS() types1.GenesisState {
	if m != nil {
		return m.SubsTS
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetCuTrackerFS() types.GenesisState {
	if m != nil {
		return m.CuTrackerFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetCuTrackerTS() types1.GenesisState {
	if m != nil {
		return m.CuTrackerTS
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetAdjustments() []Adjustment {
	if m != nil {
		return m.Adjustments
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.subscription.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/genesis.proto", fileDescriptor_dc6c60f9c112fe52)
}

var fileDescriptor_dc6c60f9c112fe52 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0x93, 0xbf, 0xfd, 0x33, 0xa4, 0x4c, 0x11, 0x43, 0xe8, 0x10, 0x0a, 0x08, 0x28, 0x0c,
	0xb6, 0x54, 0x0e, 0x80, 0x28, 0x52, 0x11, 0x03, 0x52, 0x45, 0x32, 0xb1, 0xb9, 0xc1, 0x84, 0x00,
	0x89, 0x23, 0xdb, 0xa9, 0xca, 0x2d, 0x38, 0x56, 0x07, 0x86, 0x8e, 0x4c, 0x08, 0x35, 0x17, 0x41,
	0xb6, 0xd3, 0xd6, 0x16, 0xca, 0x00, 0x93, 0x1d, 0xe9, 0x79, 0x9f, 0xbc, 0x5f, 0xf2, 0xb9, 0xc7,
	0x2f, 0x68, 0x8a, 0x72, 0xcc, 0xa1, 0x38, 0x21, 0x2b, 0x27, 0x2c, 0xa6, 0x69, 0xc1, 0x53, 0x92,
	0xc3, 0x04, 0xe7, 0x98, 0xa5, 0x0c, 0x14, 0x94, 0x70, 0xe2, 0xed, 0xd4, 0x20, 0x10, 0x27, 0xd0,
	0xc1, 0xee, 0x76, 0x42, 0x12, 0x22, 0x29, 0x28, 0x6e, 0x2a, 0xd0, 0x3d, 0x6a, 0x36, 0x17, 0x88,
	0xa2, 0xac, 0x16, 0x77, 0x4f, 0x9b, 0x39, 0x74, 0xff, 0x54, 0x32, 0x9e, 0xe1, 0x9c, 0xd7, 0xec,
	0x89, 0xc1, 0x3e, 0xa4, 0x33, 0x24, 0x38, 0xc6, 0x09, 0xc5, 0xeb, 0xa7, 0x1a, 0x3d, 0x30, 0x50,
	0x9e, 0x66, 0x98, 0x2a, 0x4e, 0x5e, 0x15, 0xb4, 0xff, 0xde, 0x72, 0xb7, 0xae, 0xd4, 0x98, 0x21,
	0x47, 0x1c, 0x7b, 0xe7, 0xae, 0xa3, 0xca, 0xf9, 0x76, 0xcf, 0xee, 0x77, 0x06, 0x7b, 0xa0, 0x71,
	0x6c, 0x30, 0x96, 0xe0, 0xb0, 0x3d, 0xff, 0xdc, 0xb5, 0x6e, 0xeb, 0x98, 0x37, 0x72, 0x1d, 0x01,
	0x8d, 0x42, 0xff, 0x9f, 0x14, 0xf4, 0x4d, 0x81, 0x51, 0x19, 0xe8, 0xaf, 0x5e, 0x79, 0x54, 0xda,
	0xbb, 0x54, 0x9e, 0x28, 0xf4, 0x5b, 0xd2, 0x73, 0x68, 0x7a, 0x36, 0xf3, 0x34, 0x4a, 0xa2, 0xd0,
	0x1b, 0xbb, 0x9d, 0xb8, 0x8c, 0x28, 0x8a, 0x9f, 0x31, 0x1d, 0x85, 0x7e, 0xfb, 0x4f, 0x8d, 0x74,
	0x85, 0x77, 0xa3, 0x19, 0xa3, 0xd0, 0xff, 0xff, 0xfb, 0x6e, 0x7a, 0x5e, 0xe8, 0x36, 0xff, 0x98,
	0xf9, 0x4e, 0xaf, 0xf5, 0x53, 0x67, 0x7c, 0xf3, 0x8b, 0x35, 0xbd, 0xd2, 0x69, 0xf9, 0xe1, 0xf5,
	0x7c, 0x19, 0xd8, 0x8b, 0x65, 0x60, 0x7f, 0x2d, 0x03, 0xfb, 0xad, 0x0a, 0xac, 0x45, 0x15, 0x58,
	0x1f, 0x55, 0x60, 0xdd, 0xc1, 0x24, 0xe5, 0x8f, 0xe5, 0x04, 0xc4, 0x24, 0x83, 0xc6, 0x62, 0x4c,
	0x07, 0x70, 0x66, 0x2e, 0x1d, 0x7f, 0x2d, 0x30, 0x9b, 0x38, 0x72, 0x41, 0xce, 0xbe, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0xd1, 0xb2, 0xd4, 0x20, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Adjustments) > 0 {
		for iNdEx := len(m.Adjustments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Adjustments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.CuTrackerTS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.CuTrackerFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.SubsTS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SubsFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SubsFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SubsTS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.CuTrackerFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.CuTrackerTS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Adjustments) > 0 {
		for _, e := range m.Adjustments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubsFS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubsFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubsTS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubsTS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuTrackerFS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CuTrackerFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuTrackerTS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CuTrackerTS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Adjustments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Adjustments = append(m.Adjustments, Adjustment{})
			if err := m.Adjustments[len(m.Adjustments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
