// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainswap/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgSendQueryAllBalances struct {
	Creator    string             `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChannelId  string             `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Address    string             `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *MsgSendQueryAllBalances) Reset()         { *m = MsgSendQueryAllBalances{} }
func (m *MsgSendQueryAllBalances) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryAllBalances) ProtoMessage()    {}
func (*MsgSendQueryAllBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cb2d9d84eeca56b, []int{0}
}
func (m *MsgSendQueryAllBalances) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryAllBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryAllBalances.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryAllBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryAllBalances.Merge(m, src)
}
func (m *MsgSendQueryAllBalances) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryAllBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryAllBalances.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryAllBalances proto.InternalMessageInfo

func (m *MsgSendQueryAllBalances) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendQueryAllBalances) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *MsgSendQueryAllBalances) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgSendQueryAllBalances) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type MsgSendQueryAllBalancesResponse struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgSendQueryAllBalancesResponse) Reset()         { *m = MsgSendQueryAllBalancesResponse{} }
func (m *MsgSendQueryAllBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryAllBalancesResponse) ProtoMessage()    {}
func (*MsgSendQueryAllBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cb2d9d84eeca56b, []int{1}
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryAllBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryAllBalancesResponse.Merge(m, src)
}
func (m *MsgSendQueryAllBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryAllBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryAllBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryAllBalancesResponse proto.InternalMessageInfo

func (m *MsgSendQueryAllBalancesResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSendQueryAllBalances)(nil), "tgntr.swapchain.interchainswap.MsgSendQueryAllBalances")
	proto.RegisterType((*MsgSendQueryAllBalancesResponse)(nil), "tgntr.swapchain.interchainswap.MsgSendQueryAllBalancesResponse")
}

func init() { proto.RegisterFile("interchainswap/tx.proto", fileDescriptor_6cb2d9d84eeca56b) }

var fileDescriptor_6cb2d9d84eeca56b = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0x2b, 0x31,
	0x14, 0xc6, 0x9b, 0xdb, 0x72, 0xef, 0x6d, 0xdc, 0x0d, 0x42, 0x87, 0x82, 0xb1, 0x74, 0x21, 0xc5,
	0x45, 0x42, 0xeb, 0xc2, 0x95, 0x88, 0x5d, 0x08, 0x22, 0x05, 0x1d, 0x77, 0x6e, 0x24, 0x93, 0x39,
	0x4c, 0x07, 0xa6, 0xc9, 0x34, 0x27, 0xd5, 0xf6, 0x2d, 0xba, 0xf4, 0x41, 0x7c, 0x08, 0x97, 0x5d,
	0xba, 0x94, 0xf6, 0x45, 0x64, 0xa6, 0x7f, 0xfc, 0x83, 0x75, 0xe1, 0xf2, 0x3b, 0xe7, 0x7c, 0xf9,
	0x7e, 0x49, 0x0e, 0xad, 0x25, 0xda, 0x81, 0x55, 0x7d, 0x99, 0x68, 0x7c, 0x90, 0x99, 0x70, 0x63,
	0x9e, 0x59, 0xe3, 0x8c, 0xc7, 0x5c, 0xac, 0x9d, 0xe5, 0x79, 0xb1, 0xe8, 0xf2, 0xcf, 0x83, 0xf5,
	0x43, 0x65, 0x70, 0x60, 0x50, 0x84, 0x12, 0x41, 0x0c, 0x47, 0x60, 0x27, 0xe2, 0xbe, 0x1d, 0x82,
	0x93, 0x6d, 0x91, 0xc9, 0x38, 0xd1, 0xd2, 0x25, 0x46, 0x2f, 0xcf, 0x6a, 0x3e, 0x11, 0x5a, 0xeb,
	0x61, 0x7c, 0x03, 0x3a, 0xba, 0xce, 0x27, 0xcf, 0xd2, 0xb4, 0x2b, 0x53, 0xa9, 0x15, 0xa0, 0xe7,
	0xd3, 0x7f, 0xca, 0x82, 0x74, 0xc6, 0xfa, 0xa4, 0x41, 0x5a, 0xd5, 0x60, 0x2d, 0xbd, 0x3d, 0x4a,
	0x55, 0x5f, 0x6a, 0x0d, 0xe9, 0x5d, 0x12, 0xf9, 0x7f, 0x8a, 0x66, 0x75, 0x55, 0xb9, 0x88, 0x72,
	0xa3, 0x8c, 0x22, 0x0b, 0x88, 0x7e, 0x79, 0x69, 0x5c, 0x49, 0xef, 0x9c, 0xd2, 0x77, 0x04, 0xbf,
	0xd2, 0x20, 0xad, 0x9d, 0xce, 0x01, 0x5f, 0xf2, 0xf2, 0x9c, 0x97, 0x17, 0xbc, 0x7c, 0xc5, 0xcb,
	0xaf, 0x64, 0x0c, 0x01, 0x0c, 0x47, 0x80, 0x2e, 0xf8, 0xe0, 0x6c, 0x9e, 0xd0, 0xfd, 0x2d, 0xd4,
	0x01, 0x60, 0x66, 0x34, 0x82, 0x57, 0xa7, 0xff, 0x31, 0x77, 0x6a, 0x05, 0x05, 0x7e, 0x25, 0xd8,
	0xe8, 0xce, 0x23, 0xa1, 0xe5, 0x1e, 0xc6, 0xde, 0x94, 0xd0, 0xdd, 0x6f, 0xaf, 0x7e, 0xcc, 0x7f,
	0x7e, 0x63, 0xbe, 0x25, 0xbd, 0x7e, 0xfa, 0x4b, 0xe3, 0x1a, 0xbb, 0x7b, 0xf9, 0x3c, 0x67, 0x64,
	0x36, 0x67, 0xe4, 0x75, 0xce, 0xc8, 0x74, 0xc1, 0x4a, 0xb3, 0x05, 0x2b, 0xbd, 0x2c, 0x58, 0xe9,
	0xb6, 0x1d, 0x27, 0xae, 0x3f, 0x0a, 0xb9, 0x32, 0x03, 0x51, 0x84, 0x88, 0x4d, 0x88, 0x18, 0x8b,
	0xaf, 0xcb, 0x32, 0xc9, 0x00, 0xc3, 0xbf, 0xc5, 0x27, 0x1f, 0xbd, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0x63, 0xf6, 0x3a, 0x4b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendQueryAllBalances(ctx context.Context, in *MsgSendQueryAllBalances, opts ...grpc.CallOption) (*MsgSendQueryAllBalancesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendQueryAllBalances(ctx context.Context, in *MsgSendQueryAllBalances, opts ...grpc.CallOption) (*MsgSendQueryAllBalancesResponse, error) {
	out := new(MsgSendQueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, "/tgntr.swapchain.interchainswap.Msg/SendQueryAllBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendQueryAllBalances(context.Context, *MsgSendQueryAllBalances) (*MsgSendQueryAllBalancesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendQueryAllBalances(ctx context.Context, req *MsgSendQueryAllBalances) (*MsgSendQueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQueryAllBalances not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendQueryAllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendQueryAllBalances)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendQueryAllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tgntr.swapchain.interchainswap.Msg/SendQueryAllBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendQueryAllBalances(ctx, req.(*MsgSendQueryAllBalances))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tgntr.swapchain.interchainswap.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendQueryAllBalances",
			Handler:    _Msg_SendQueryAllBalances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchainswap/tx.proto",
}

func (m *MsgSendQueryAllBalances) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryAllBalances) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryAllBalances) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryAllBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryAllBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryAllBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendQueryAllBalances) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendQueryAllBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendQueryAllBalances) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendQueryAllBalances: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryAllBalances: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSendQueryAllBalancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendQueryAllBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryAllBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)