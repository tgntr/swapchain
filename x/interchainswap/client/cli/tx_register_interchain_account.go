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

func CmdRegisterInterchainAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-interchain-account [connection-id]",
		Short: "Register ICA account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterIntechainAccount(
				clientCtx.GetFromAddress().String(),
				args[0], // connection id
				"",
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
