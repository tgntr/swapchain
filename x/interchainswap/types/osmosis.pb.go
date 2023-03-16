// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainswap/osmosis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type OsmosisQuerySpotPriceRequest struct {
	PoolId          uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	BaseAssetDenom  string `protobuf:"bytes,2,opt,name=base_asset_denom,json=baseAssetDenom,proto3" json:"base_asset_denom,omitempty"`
	QuoteAssetDenom string `protobuf:"bytes,3,opt,name=quote_asset_denom,json=quoteAssetDenom,proto3" json:"quote_asset_denom,omitempty"`
}

func (m *OsmosisQuerySpotPriceRequest) Reset()         { *m = OsmosisQuerySpotPriceRequest{} }
func (m *OsmosisQuerySpotPriceRequest) String() string { return proto.CompactTextString(m) }
func (*OsmosisQuerySpotPriceRequest) ProtoMessage()    {}
func (*OsmosisQuerySpotPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8388c82863a51150, []int{0}
}
func (m *OsmosisQuerySpotPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmosisQuerySpotPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmosisQuerySpotPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmosisQuerySpotPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmosisQuerySpotPriceRequest.Merge(m, src)
}
func (m *OsmosisQuerySpotPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *OsmosisQuerySpotPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmosisQuerySpotPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OsmosisQuerySpotPriceRequest proto.InternalMessageInfo

func (m *OsmosisQuerySpotPriceRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *OsmosisQuerySpotPriceRequest) GetBaseAssetDenom() string {
	if m != nil {
		return m.BaseAssetDenom
	}
	return ""
}

func (m *OsmosisQuerySpotPriceRequest) GetQuoteAssetDenom() string {
	if m != nil {
		return m.QuoteAssetDenom
	}
	return ""
}

type OsmosisQuerySpotPriceResponse struct {
	SpotPrice string `protobuf:"bytes,1,opt,name=spot_price,json=spotPrice,proto3" json:"spot_price,omitempty"`
}

func (m *OsmosisQuerySpotPriceResponse) Reset()         { *m = OsmosisQuerySpotPriceResponse{} }
func (m *OsmosisQuerySpotPriceResponse) String() string { return proto.CompactTextString(m) }
func (*OsmosisQuerySpotPriceResponse) ProtoMessage()    {}
func (*OsmosisQuerySpotPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8388c82863a51150, []int{1}
}
func (m *OsmosisQuerySpotPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmosisQuerySpotPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmosisQuerySpotPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmosisQuerySpotPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmosisQuerySpotPriceResponse.Merge(m, src)
}
func (m *OsmosisQuerySpotPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *OsmosisQuerySpotPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmosisQuerySpotPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OsmosisQuerySpotPriceResponse proto.InternalMessageInfo

func (m *OsmosisQuerySpotPriceResponse) GetSpotPrice() string {
	if m != nil {
		return m.SpotPrice
	}
	return ""
}

type SwapAmountInRoute struct {
	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty"`
}

func (m *SwapAmountInRoute) Reset()         { *m = SwapAmountInRoute{} }
func (m *SwapAmountInRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountInRoute) ProtoMessage()    {}
func (*SwapAmountInRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_8388c82863a51150, []int{2}
}
func (m *SwapAmountInRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountInRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountInRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountInRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountInRoute.Merge(m, src)
}
func (m *SwapAmountInRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountInRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountInRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountInRoute proto.InternalMessageInfo

func (m *SwapAmountInRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *SwapAmountInRoute) GetTokenOutDenom() string {
	if m != nil {
		return m.TokenOutDenom
	}
	return ""
}

type OsmosisMsgSwapExactAmountIn struct {
	Sender            string                                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Routes            []SwapAmountInRoute                    `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes"`
	TokenIn           types.Coin                             `protobuf:"bytes,3,opt,name=token_in,json=tokenIn,proto3" json:"token_in"`
	TokenOutMinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=token_out_min_amount,json=tokenOutMinAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_out_min_amount"`
}

func (m *OsmosisMsgSwapExactAmountIn) Reset()         { *m = OsmosisMsgSwapExactAmountIn{} }
func (m *OsmosisMsgSwapExactAmountIn) String() string { return proto.CompactTextString(m) }
func (*OsmosisMsgSwapExactAmountIn) ProtoMessage()    {}
func (*OsmosisMsgSwapExactAmountIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8388c82863a51150, []int{3}
}
func (m *OsmosisMsgSwapExactAmountIn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmosisMsgSwapExactAmountIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmosisMsgSwapExactAmountIn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmosisMsgSwapExactAmountIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmosisMsgSwapExactAmountIn.Merge(m, src)
}
func (m *OsmosisMsgSwapExactAmountIn) XXX_Size() int {
	return m.Size()
}
func (m *OsmosisMsgSwapExactAmountIn) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmosisMsgSwapExactAmountIn.DiscardUnknown(m)
}

var xxx_messageInfo_OsmosisMsgSwapExactAmountIn proto.InternalMessageInfo

func (m *OsmosisMsgSwapExactAmountIn) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OsmosisMsgSwapExactAmountIn) GetRoutes() []SwapAmountInRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *OsmosisMsgSwapExactAmountIn) GetTokenIn() types.Coin {
	if m != nil {
		return m.TokenIn
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*OsmosisQuerySpotPriceRequest)(nil), "tgntr.swapchain.interchainswap.OsmosisQuerySpotPriceRequest")
	proto.RegisterType((*OsmosisQuerySpotPriceResponse)(nil), "tgntr.swapchain.interchainswap.OsmosisQuerySpotPriceResponse")
	proto.RegisterType((*SwapAmountInRoute)(nil), "tgntr.swapchain.interchainswap.SwapAmountInRoute")
	proto.RegisterType((*OsmosisMsgSwapExactAmountIn)(nil), "tgntr.swapchain.interchainswap.OsmosisMsgSwapExactAmountIn")
}

func init() { proto.RegisterFile("interchainswap/osmosis.proto", fileDescriptor_8388c82863a51150) }

var fileDescriptor_8388c82863a51150 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0x34, 0x4a, 0xc9, 0x56, 0x50, 0xb2, 0xaa, 0x20, 0x94, 0xd6, 0x8d, 0x72, 0xa8,
	0x22, 0x24, 0x76, 0x95, 0x72, 0xe3, 0x80, 0xd4, 0x00, 0x87, 0x08, 0x55, 0x01, 0x97, 0x13, 0x17,
	0xcb, 0x1f, 0x2b, 0x77, 0x55, 0xbc, 0xb3, 0xf5, 0x8e, 0x69, 0xfb, 0x10, 0x48, 0xbc, 0x03, 0x2f,
	0xd3, 0x63, 0x8f, 0x88, 0x43, 0x85, 0x92, 0x17, 0x41, 0xbb, 0x76, 0x28, 0x01, 0xa5, 0x27, 0x7b,
	0x67, 0xfe, 0x33, 0xf3, 0x9b, 0x0f, 0xb2, 0x23, 0x15, 0x8a, 0x22, 0x39, 0x89, 0xa4, 0x32, 0xe7,
	0x91, 0xe6, 0x60, 0x72, 0x30, 0xd2, 0x30, 0x5d, 0x00, 0x02, 0xf5, 0x31, 0x53, 0x58, 0x30, 0xeb,
	0x71, 0x12, 0xb6, 0xac, 0xde, 0xde, 0xca, 0x20, 0x03, 0x27, 0xe5, 0xf6, 0xaf, 0x8a, 0xda, 0xf6,
	0x13, 0x97, 0x85, 0xc7, 0x91, 0x11, 0xfc, 0xcb, 0x28, 0x16, 0x18, 0x8d, 0x78, 0x02, 0x52, 0x55,
	0xfe, 0xc1, 0x57, 0x8f, 0xec, 0x4c, 0xab, 0x3a, 0x1f, 0x4a, 0x51, 0x5c, 0x1e, 0x6b, 0xc0, 0xf7,
	0x85, 0x4c, 0x44, 0x20, 0xce, 0x4a, 0x61, 0x90, 0x3e, 0x26, 0xeb, 0x1a, 0xe0, 0x73, 0x28, 0xd3,
	0x9e, 0xd7, 0xf7, 0x86, 0xad, 0xa0, 0x6d, 0x9f, 0x93, 0x94, 0x0e, 0xc9, 0x43, 0x9b, 0x34, 0x8c,
	0x8c, 0x11, 0x18, 0xa6, 0x42, 0x41, 0xde, 0x6b, 0xf6, 0xbd, 0x61, 0x27, 0x78, 0x60, 0xed, 0x87,
	0xd6, 0xfc, 0xc6, 0x5a, 0xe9, 0x33, 0xd2, 0x3d, 0x2b, 0x01, 0x97, 0xa5, 0x6b, 0x4e, 0xba, 0xe9,
	0x1c, 0xb7, 0xda, 0xc1, 0x2b, 0xb2, 0xbb, 0x02, 0xc7, 0x68, 0x50, 0x46, 0xd0, 0x5d, 0x42, 0x8c,
	0x06, 0x0c, 0xb5, 0xb5, 0x3a, 0xa4, 0x4e, 0xd0, 0x31, 0x0b, 0xd9, 0xe0, 0x23, 0xe9, 0x1e, 0x9f,
	0x47, 0xfa, 0x30, 0x87, 0x52, 0xe1, 0x44, 0x05, 0x50, 0xa2, 0x58, 0xdd, 0xc3, 0x3e, 0xd9, 0x44,
	0x38, 0x15, 0x2a, 0x84, 0x72, 0xb9, 0x85, 0xfb, 0xce, 0x3c, 0x2d, 0x6b, 0xaa, 0xef, 0x4d, 0xf2,
	0xb4, 0xc6, 0x3a, 0x32, 0x99, 0x2d, 0xf0, 0xf6, 0x22, 0x4a, 0x70, 0x51, 0x85, 0x3e, 0x22, 0x6d,
	0x23, 0x54, 0x2a, 0x8a, 0x1a, 0xa8, 0x7e, 0xd1, 0x29, 0x69, 0x17, 0x96, 0xc0, 0xf4, 0x9a, 0xfd,
	0xb5, 0xe1, 0xc6, 0xc1, 0x88, 0xdd, 0xbd, 0x44, 0xf6, 0x1f, 0xfb, 0xb8, 0x75, 0x75, 0xb3, 0xd7,
	0x08, 0xea, 0x34, 0xf4, 0x25, 0xb9, 0x57, 0x01, 0x4b, 0xe5, 0x26, 0xb8, 0x71, 0xf0, 0x84, 0x55,
	0x1b, 0x66, 0x76, 0xe8, 0xac, 0xde, 0x30, 0x7b, 0x0d, 0x52, 0xd5, 0xa1, 0xeb, 0x2e, 0x60, 0xa2,
	0x68, 0x48, 0xb6, 0x6e, 0x9b, 0xcd, 0xa5, 0x0a, 0x23, 0x57, 0xa8, 0xd7, 0xb2, 0xc8, 0x63, 0x66,
	0xc5, 0x3f, 0x6f, 0xf6, 0xf6, 0x33, 0x89, 0x27, 0x65, 0xcc, 0x12, 0xc8, 0x79, 0x7d, 0x3b, 0xd5,
	0xe7, 0xb9, 0x49, 0x4f, 0x39, 0x5e, 0x6a, 0x61, 0xd8, 0x44, 0x61, 0xd0, 0x5d, 0x4c, 0xe8, 0x48,
	0xaa, 0x8a, 0x78, 0xfc, 0xee, 0x6a, 0xe6, 0x7b, 0xd7, 0x33, 0xdf, 0xfb, 0x35, 0xf3, 0xbd, 0x6f,
	0x73, 0xbf, 0x71, 0x3d, 0xf7, 0x1b, 0x3f, 0xe6, 0x7e, 0xe3, 0xd3, 0xe8, 0xaf, 0xa4, 0x6e, 0x02,
	0xfc, 0xcf, 0x04, 0xf8, 0x05, 0xff, 0xe7, 0xec, 0x5d, 0x8d, 0xb8, 0xed, 0xee, 0xf3, 0xc5, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xc8, 0x2c, 0xd7, 0x15, 0x03, 0x00, 0x00,
}

func (m *OsmosisQuerySpotPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmosisQuerySpotPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmosisQuerySpotPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QuoteAssetDenom) > 0 {
		i -= len(m.QuoteAssetDenom)
		copy(dAtA[i:], m.QuoteAssetDenom)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.QuoteAssetDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAssetDenom) > 0 {
		i -= len(m.BaseAssetDenom)
		copy(dAtA[i:], m.BaseAssetDenom)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.BaseAssetDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosis(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OsmosisQuerySpotPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmosisQuerySpotPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmosisQuerySpotPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpotPrice) > 0 {
		i -= len(m.SpotPrice)
		copy(dAtA[i:], m.SpotPrice)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.SpotPrice)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SwapAmountInRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountInRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountInRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOutDenom) > 0 {
		i -= len(m.TokenOutDenom)
		copy(dAtA[i:], m.TokenOutDenom)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.TokenOutDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosis(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OsmosisMsgSwapExactAmountIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmosisMsgSwapExactAmountIn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmosisMsgSwapExactAmountIn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenOutMinAmount.Size()
		i -= size
		if _, err := m.TokenOutMinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TokenIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Routes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOsmosis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOsmosis(dAtA []byte, offset int, v uint64) int {
	offset -= sovOsmosis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OsmosisQuerySpotPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosis(uint64(m.PoolId))
	}
	l = len(m.BaseAssetDenom)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	l = len(m.QuoteAssetDenom)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	return n
}

func (m *OsmosisQuerySpotPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpotPrice)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	return n
}

func (m *SwapAmountInRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosis(uint64(m.PoolId))
	}
	l = len(m.TokenOutDenom)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	return n
}

func (m *OsmosisMsgSwapExactAmountIn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovOsmosis(uint64(l))
		}
	}
	l = m.TokenIn.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.TokenOutMinAmount.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	return n
}

func sovOsmosis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOsmosis(x uint64) (n int) {
	return sovOsmosis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OsmosisQuerySpotPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: OsmosisQuerySpotPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmosisQuerySpotPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func (m *OsmosisQuerySpotPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: OsmosisQuerySpotPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmosisQuerySpotPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpotPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func (m *SwapAmountInRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: SwapAmountInRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountInRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOutDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func (m *OsmosisMsgSwapExactAmountIn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: OsmosisMsgSwapExactAmountIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmosisMsgSwapExactAmountIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, SwapAmountInRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutMinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOutMinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func skipOsmosis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOsmosis
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
					return 0, ErrIntOverflowOsmosis
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
					return 0, ErrIntOverflowOsmosis
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
				return 0, ErrInvalidLengthOsmosis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOsmosis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOsmosis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOsmosis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOsmosis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOsmosis = fmt.Errorf("proto: unexpected end of group")
)
