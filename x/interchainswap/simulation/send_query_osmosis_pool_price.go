package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/tgntr/swapchain/x/interchainswap/keeper"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

func SimulateMsgSendQueryOsmosisSpotPrice(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSendQueryOsmosisSpotPrice{
			Creator:         simAccount.Address.String(),
			BaseAssetDenom:  "test",
			QuoteAssetDenom: "testt",
		}

		// TODO: Handling the SendQueryOsmosisSpotPrice simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendQueryOsmosisSpotPrice simulation not implemented"), nil, nil
	}
}
