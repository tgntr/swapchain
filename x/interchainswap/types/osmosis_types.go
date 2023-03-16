package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &OsmosisMsgSwapExactAmountIn{}

func (msg *OsmosisMsgSwapExactAmountIn) XXX_MessageName() string {
	return "osmosis.gamm.v1beta1.MsgSwapExactAmountIn"
}

func (msg *OsmosisMsgSwapExactAmountIn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.GetFromBech32(string(msg.Sender), "osmo")
	if err != nil {
		panic(err)
	}

	if err := sdk.VerifyAddressFormat(addr); err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(addr)}
}

func (msg *OsmosisMsgSwapExactAmountIn) ValidateBasic() error {
	addr, err := sdk.GetFromBech32(string(msg.Sender), "osmo")
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if err := sdk.VerifyAddressFormat(addr); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	for _, route := range msg.Routes {
		err := sdk.ValidateDenom(route.TokenOutDenom)
		if err != nil {
			return err
		}
	}

	if !msg.TokenIn.IsValid() || !msg.TokenIn.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.TokenIn.String())
	}

	if !msg.TokenOutMinAmount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.TokenOutMinAmount.String())
	}

	return nil
}

func (q *OsmosisQuerySpotPriceRequest) XXX_MessageName() string {
	return "osmosis.gamm.v2.QuerySpotPriceRequest"
}

func (q *OsmosisQuerySpotPriceResponse) XXX_MessageName() string {
	return "osmosis.gamm.v2.QuerySpotPriceResponse"
}
