package cli

import (
	"context"

	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-transaction",
		Short: "list all transaction",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTransactionRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TransactionAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-transaction [id]",
		Short: "shows a transaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			hash := args[0]

			params := &types.QueryGetTransactionRequest{
				Hash: hash,
			}

			res, err := queryClient.Transaction(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
