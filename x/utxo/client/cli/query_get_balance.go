package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"utxo/x/utxo/types"
)

var _ = strconv.Itoa(0)

func CmdGetBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-balance [pubkey]",
		Short: "Query get-balance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPubkey := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetBalanceRequest{

				Pubkey: reqPubkey,
			}

			res, err := queryClient.GetBalance(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
