package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/sujine/test/x/test/types"
)

var _ = strconv.Itoa(0)

func CmdGetInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-info [creator] [denom]",
		Short: "Query get-info",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCreator := args[0]
			reqDenom := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetInfoRequest{

				Creator: reqCreator,
				Denom:   reqDenom,
			}

			res, err := queryClient.GetInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
