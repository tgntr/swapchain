package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSendQueryOsmosisSpotPrice{}, "interchainswap/SendQueryOsmosisSpotPrice", nil)
	cdc.RegisterConcrete(&MsgRegisterIntechainAccount{}, "interchainswap/RegisterInterchainAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendQueryOsmosisSpotPrice{},
		&MsgRegisterIntechainAccount{},
	)
	// this line is used by starport scaffolding # 3

	// For printing purposes
	registry.RegisterImplementations((*proto.Message)(nil),
		&QueryOsmosisSpotPriceRequest{},
		&QueryOsmosisSpotPriceResponse{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
