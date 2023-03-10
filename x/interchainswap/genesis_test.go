package interchainswap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tgntr/swapchain/testutil/keeper"
	"github.com/tgntr/swapchain/testutil/nullify"
	interchainswap "github.com/tgntr/swapchain/x/interchainswap"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InterchainswapKeeper(t)
	interchainswap.InitGenesis(ctx, *k, genesisState)
	got := interchainswap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
