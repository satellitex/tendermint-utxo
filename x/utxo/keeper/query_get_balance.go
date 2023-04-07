package keeper

import (
	"context"

	"utxo/x/utxo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBalance(goCtx context.Context, req *types.QueryGetBalanceRequest) (*types.QueryGetBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the outpint from pubkey by keepr.
	outpoint, found := k.GetOutpoint(ctx, *&req.Pubkey)
	if !found {
		return nil, status.Error(codes.NotFound, "outpoint not found")
	}
	// Get the transaction from outpoint's txHash by keeper.
	tx, found := k.GetTransaction(ctx, []byte(outpoint.Hash))
	if !found {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}
	// Get the txOut from outpoint's index by keeper.
	txOut := tx.TxOut[outpoint.Index]

	return &types.QueryGetBalanceResponse{
		Amount:   txOut.Value,
		Outpoint: &outpoint,
	}, nil
}
