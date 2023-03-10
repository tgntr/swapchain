package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tgntr/swapchain/testutil/keeper"
	"github.com/tgntr/swapchain/x/interchainswap/keeper"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterchainswapKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
