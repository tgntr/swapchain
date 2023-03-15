package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendQueryOsmosisSpotPrice = "send_query_osmosis_spot_price"

var _ sdk.Msg = &MsgSendQueryOsmosisSpotPrice{}

func NewMsgSendQueryOsmosisSpotPrice(creator, channelId string, poolId uint64, baseAssetDenom, quoteAssetDenom string) *MsgSendQueryOsmosisSpotPrice {
	return &MsgSendQueryOsmosisSpotPrice{
		Creator:         creator,
		ChannelId:       channelId,
		PoolId:          poolId,
		BaseAssetDenom:  baseAssetDenom,
		QuoteAssetDenom: quoteAssetDenom,
	}
}

func (msg *MsgSendQueryOsmosisSpotPrice) Route() string {
	return RouterKey
}

func (msg *MsgSendQueryOsmosisSpotPrice) Type() string {
	return TypeMsgSendQueryOsmosisSpotPrice
}

func (msg *MsgSendQueryOsmosisSpotPrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendQueryOsmosisSpotPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendQueryOsmosisSpotPrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.PoolId < 1 {
		return sdkerrors.Wrapf(ErrInvalidPoolId, "invalid base asset denom (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.BaseAssetDenom); err != nil {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid base asset denom (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.QuoteAssetDenom); err != nil {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid quote asset denom (%s)", err)
	}

	return nil
}
