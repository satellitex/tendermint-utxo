package cli

import (
	"encoding/json"
	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdSendTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-transaction [version] [tx-in] [tx-out] [locktime]",
		Short: "Send a new transaction",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVersion, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argTxIn := new(types.TxIn)
			err = json.Unmarshal([]byte(args[1]), argTxIn)
			if err != nil {
				return err
			}
			argTxOut := new(types.TxOut)
			err = json.Unmarshal([]byte(args[2]), argTxOut)
			if err != nil {
				return err
			}
			argLocktime, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendTransaction(
				&types.Transaction{
					Version:  argVersion,
					TxIn:     []*types.TxIn{argTxIn},
					TxOut:    []*types.TxOut{argTxOut},
					Locktime: argLocktime,
				})
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
