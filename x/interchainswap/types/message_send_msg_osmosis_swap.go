package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendMsgOsmosisSwap = "send_msg_osmosis_swap"

var (
	_ sdk.Msg = &MsgSendMsgOsmosisSwap{}
)

func NewMsgSendMsgOsmosisSwap(swapMsg *OsmosisMsgSwapExactAmountIn, connectionID string, sender string) *MsgSendMsgOsmosisSwap {
	return &MsgSendMsgOsmosisSwap{
		Owner:        sender,
		ConnectionId: connectionID,
		Msg:          swapMsg,
	}
}

func (msg *MsgSendMsgOsmosisSwap) Route() string {
	return RouterKey
}

func (msg *MsgSendMsgOsmosisSwap) Type() string {
	return TypeMsgSendMsgOsmosisSwap
}

func (msg *MsgSendMsgOsmosisSwap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendMsgOsmosisSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendMsgOsmosisSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if strings.TrimSpace(msg.ConnectionId) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing connection id")
	}

	return msg.Msg.ValidateBasic()
}
