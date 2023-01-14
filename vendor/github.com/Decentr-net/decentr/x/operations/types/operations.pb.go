// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operations/operations.proto

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

type Params struct {
	Supervisors []string       `protobuf:"bytes,1,rep,name=supervisors,proto3" json:"supervisors,omitempty"`
	FixedGas    FixedGasParams `protobuf:"bytes,2,opt,name=fixed_gas,json=fixedGas,proto3" json:"fixed_gas"`
	MinGasPrice types.FurCoin  `protobuf:"bytes,3,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_07148eb1bb237ba9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSupervisors() []string {
	if m != nil {
		return m.Supervisors
	}
	return nil
}

func (m *Params) GetFixedGas() FixedGasParams {
	if m != nil {
		return m.FixedGas
	}
	return FixedGasParams{}
}

func (m *Params) GetMinGasPrice() types.FurCoin {
	if m != nil {
		return m.MinGasPrice
	}
	return types.FurCoin{}
}

type FixedGasParams struct {
	ResetAccount      uint64 `protobuf:"varint,1,opt,name=reset_account,json=resetAccount,proto3" json:"reset_account,omitempty"`
	DistributeRewards uint64 `protobuf:"varint,2,opt,name=distribute_rewards,json=distributeRewards,proto3" json:"distribute_rewards,omitempty"`
}

func (m *FixedGasParams) Reset()         { *m = FixedGasParams{} }
func (m *FixedGasParams) String() string { return proto.CompactTextString(m) }
func (*FixedGasParams) ProtoMessage()    {}
func (*FixedGasParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_07148eb1bb237ba9, []int{1}
}
func (m *FixedGasParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FixedGasParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FixedGasParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FixedGasParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedGasParams.Merge(m, src)
}
func (m *FixedGasParams) XXX_Size() int {
	return m.Size()
}
func (m *FixedGasParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedGasParams.DiscardUnknown(m)
}

var xxx_messageInfo_FixedGasParams proto.InternalMessageInfo

func (m *FixedGasParams) GetResetAccount() uint64 {
	if m != nil {
		return m.ResetAccount
	}
	return 0
}

func (m *FixedGasParams) GetDistributeRewards() uint64 {
	if m != nil {
		return m.DistributeRewards
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "operations.Params")
	proto.RegisterType((*FixedGasParams)(nil), "operations.FixedGasParams")
}

func init() { proto.RegisterFile("operations/operations.proto", fileDescriptor_07148eb1bb237ba9) }

var fileDescriptor_07148eb1bb237ba9 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xbf, 0x55, 0xf5, 0xd7, 0xa5, 0x48, 0x44, 0x0c, 0x51, 0x41, 0x21, 0x2a, 0x4b,
	0x97, 0xda, 0x2a, 0xcc, 0x0c, 0x94, 0xaa, 0x0c, 0x2c, 0x28, 0x23, 0x4b, 0xe4, 0x38, 0xb7, 0xc1,
	0x43, 0xec, 0xc8, 0x76, 0x4a, 0x79, 0x0b, 0x5e, 0x85, 0xb7, 0xe8, 0xd8, 0x91, 0x09, 0xa1, 0xf6,
	0x45, 0x50, 0xe2, 0x4a, 0x2d, 0xdb, 0xd1, 0x77, 0x7c, 0x8f, 0x8f, 0x7d, 0xf1, 0x85, 0x2a, 0x41,
	0x33, 0x2b, 0x94, 0x34, 0xf4, 0x20, 0x49, 0xa9, 0x95, 0x55, 0x3e, 0x3e, 0x90, 0xc1, 0x79, 0xae,
	0x72, 0xd5, 0x60, 0x5a, 0x2b, 0x77, 0x62, 0x10, 0x72, 0x65, 0x0a, 0x65, 0x68, 0xca, 0x0c, 0xd0,
	0xe5, 0x24, 0x05, 0xcb, 0x26, 0x94, 0x2b, 0x21, 0x9d, 0x3f, 0xfc, 0x44, 0xb8, 0xf3, 0xcc, 0x34,
	0x2b, 0x8c, 0x1f, 0xe1, 0x9e, 0xa9, 0x4a, 0xd0, 0x4b, 0x61, 0x94, 0x36, 0x01, 0x8a, 0x5a, 0xa3,
	0x6e, 0x7c, 0x8c, 0xfc, 0x3b, 0xdc, 0x5d, 0x88, 0x15, 0x64, 0x49, 0xce, 0x4c, 0xf0, 0x2f, 0x42,
	0xa3, 0xde, 0xcd, 0x80, 0x1c, 0x95, 0x9a, 0xd7, 0xe6, 0x23, 0x33, 0x2e, 0x70, 0xda, 0x5e, 0x7f,
	0x5f, 0x79, 0xf1, 0xff, 0xc5, 0x9e, 0xfa, 0x73, 0xdc, 0x2f, 0x84, 0xac, 0x87, 0x93, 0x52, 0x0b,
	0x0e, 0x41, 0xab, 0x89, 0xb8, 0x24, 0xae, 0x23, 0xa9, 0x3b, 0x92, 0x7d, 0x47, 0x32, 0x03, 0xfe,
	0xa0, 0x84, 0xdc, 0x87, 0xf4, 0x0a, 0x21, 0xeb, 0xe0, 0x7a, 0x6c, 0x98, 0xe1, 0xd3, 0xbf, 0x37,
	0xf9, 0xd7, 0xb8, 0xaf, 0xc1, 0x80, 0x4d, 0x18, 0xe7, 0xaa, 0x92, 0x36, 0x40, 0x11, 0x1a, 0xb5,
	0xe3, 0x93, 0x06, 0xde, 0x3b, 0xe6, 0x8f, 0xb1, 0x9f, 0x09, 0x63, 0xb5, 0x48, 0x2b, 0x0b, 0x89,
	0x86, 0x37, 0xa6, 0x33, 0xf7, 0x8c, 0x76, 0x7c, 0x76, 0x70, 0x62, 0x67, 0x4c, 0x9f, 0xd6, 0xdb,
	0x10, 0x6d, 0xb6, 0x21, 0xfa, 0xd9, 0x86, 0xe8, 0x63, 0x17, 0x7a, 0x9b, 0x5d, 0xe8, 0x7d, 0xed,
	0x42, 0xef, 0x65, 0x92, 0x0b, 0xfb, 0x5a, 0xa5, 0x84, 0xab, 0x82, 0xce, 0x80, 0x83, 0xb4, 0x7a,
	0x2c, 0xc1, 0xd2, 0xcc, 0x69, 0xba, 0x3a, 0x5a, 0x14, 0xb5, 0xef, 0x25, 0x98, 0xb4, 0xd3, 0xfc,
	0xf6, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x99, 0x7b, 0x35, 0xce, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOperations(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.FixedGas.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOperations(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Supervisors) > 0 {
		for iNdEx := len(m.Supervisors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Supervisors[iNdEx])
			copy(dAtA[i:], m.Supervisors[iNdEx])
			i = encodeVarintOperations(dAtA, i, uint64(len(m.Supervisors[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FixedGasParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FixedGasParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FixedGasParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DistributeRewards != 0 {
		i = encodeVarintOperations(dAtA, i, uint64(m.DistributeRewards))
		i--
		dAtA[i] = 0x10
	}
	if m.ResetAccount != 0 {
		i = encodeVarintOperations(dAtA, i, uint64(m.ResetAccount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOperations(dAtA []byte, offset int, v uint64) int {
	offset -= sovOperations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Supervisors) > 0 {
		for _, s := range m.Supervisors {
			l = len(s)
			n += 1 + l + sovOperations(uint64(l))
		}
	}
	l = m.FixedGas.Size()
	n += 1 + l + sovOperations(uint64(l))
	l = m.MinGasPrice.Size()
	n += 1 + l + sovOperations(uint64(l))
	return n
}

func (m *FixedGasParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResetAccount != 0 {
		n += 1 + sovOperations(uint64(m.ResetAccount))
	}
	if m.DistributeRewards != 0 {
		n += 1 + sovOperations(uint64(m.DistributeRewards))
	}
	return n
}

func sovOperations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOperations(x uint64) (n int) {
	return sovOperations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supervisors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supervisors = append(m.Supervisors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedGas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FixedGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOperations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOperations
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
func (m *FixedGasParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: FixedGasParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixedGasParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResetAccount", wireType)
			}
			m.ResetAccount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResetAccount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeRewards", wireType)
			}
			m.DistributeRewards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributeRewards |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOperations
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
func skipOperations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
				return 0, ErrInvalidLengthOperations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOperations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOperations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOperations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOperations = fmt.Errorf("proto: unexpected end of group")
)
