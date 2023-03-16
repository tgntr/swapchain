package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterIntechainAccount = "register_interchain_account"

var _ sdk.Msg = &MsgRegisterIntechainAccount{}

// NewMsgRegisterIntechainAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterIntechainAccount(owner, connectionID, version string) *MsgRegisterIntechainAccount {
	return &MsgRegisterIntechainAccount{
		Owner:        owner,
		ConnectionId: connectionID,
		Version:      version,
	}
}

func (msg *MsgRegisterIntechainAccount) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIntechainAccount) Type() string {
	return TypeMsgRegisterIntechainAccount
}

func (msg *MsgRegisterIntechainAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIntechainAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterIntechainAccount) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "failed to parse address: %s", msg.Owner)
	}

	if strings.TrimSpace(msg.ConnectionId) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing connection id")
	}

	return nil
}
