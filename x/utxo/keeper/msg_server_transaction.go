package keeper

import (
	"context"
	"crypto/sha256"

	"utxo/x/utxo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Change
func (k msgServer) SendTransaction(goCtx context.Context, msg *types.MsgSendTransaction) (*types.MsgSendTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var transaction = msg.Tx
	data, _ := transaction.Marshal()
	txHash := sha256.Sum256(data)
	txHashStr := string(txHash[:])

	k.AppendTransaction(
		ctx,
		txHashStr,
		*transaction,
	)
	return &types.MsgSendTransactionResponse{
		Hash: txHashStr,
	}, nil
}
