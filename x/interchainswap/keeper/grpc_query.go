package keeper

import (
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

var _ types.QueryServer = Keeper{}
