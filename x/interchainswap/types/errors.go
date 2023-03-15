package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/interchainswap module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrInvalidDenom         = sdkerrors.Register(ModuleName, 1502, "invalid denom")
	ErrInvalidPoolId        = sdkerrors.Register(ModuleName, 1503, "non-positive pool id")
)
