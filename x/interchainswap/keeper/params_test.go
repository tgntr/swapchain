package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tgntr/swapchain/testutil/keeper"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterchainswapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
