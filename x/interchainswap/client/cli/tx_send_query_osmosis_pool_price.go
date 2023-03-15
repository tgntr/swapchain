package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tgntr/swapchain/x/interchainswap/types"
)

var _ = strconv.Itoa(0)

// TODO test
func CmdSendQueryOsmosisSpotPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-query-osmosis-spot-price [channel-id] [pool-id] [base-asset-denom] [quote-asset-denom]",
		Short: "Query osmosis pool spot price via ICQ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendQueryOsmosisSpotPrice(
				clientCtx.GetFromAddress().String(),
				args[0], // channel id
				poolId,
				args[2], // base asset denom
				args[3], // quote asset denom
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
