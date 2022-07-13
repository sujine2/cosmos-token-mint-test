package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/sujine/test/x/test/types"
)

var _ = strconv.Itoa(0)

func CmdRegisterCoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-coin [symbol] [price] [amount]",
		Short: "Broadcast message register-coin",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSymbol := args[0]
			argPrice, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argAmount, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterCoin(
				clientCtx.GetFromAddress().String(),
				argSymbol,
				argPrice,
				argAmount,
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
