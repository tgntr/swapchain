package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/tgntr/swapchain/x/interchainswap/types"
)

var _ = strconv.Itoa(0)

func CmdSendMsgOsmosisSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "send-msg-osmosis-swap [msg-swap-exact-amount-in-json] [connection-id]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
			cdc.InterfaceRegistry().RegisterImplementations((*sdk.Msg)(nil), &types.OsmosisMsgSwapExactAmountIn{})

			var swapMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(args[0]), &swapMsg); err != nil {
				return err
			}

			msg := types.NewMsgSendMsgOsmosisSwap(swapMsg.(*types.OsmosisMsgSwapExactAmountIn), args[1], clientCtx.GetFromAddress().String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
