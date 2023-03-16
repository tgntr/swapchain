package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

func (k msgServer) RegisterInterchainAccount(goCtx context.Context, msg *types.MsgRegisterIntechainAccount) (*types.MsgRegisterIntechainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, msg.Owner, msg.Version); err != nil {
		return nil, err
	}

	return &types.MsgRegisterIntechainAccountResponse{}, nil
}
