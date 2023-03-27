package keeper

import (
	"context"

	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TransactionAll(goCtx context.Context, req *types.QueryAllTransactionRequest) (*types.QueryAllTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transactions []types.Transaction
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	transactionStore := prefix.NewStore(store, types.KeyPrefix(types.TransactionKey))

	pageRes, err := query.Paginate(transactionStore, req.Pagination, func(key []byte, value []byte) error {
		var transaction types.Transaction
		if err := k.cdc.Unmarshal(value, &transaction); err != nil {
			return err
		}

		transactions = append(transactions, transaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTransactionResponse{Transaction: transactions, Pagination: pageRes}, nil
}

func (k Keeper) Transaction(goCtx context.Context, req *types.QueryGetTransactionRequest) (*types.QueryGetTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	transaction, found := k.GetTransaction(ctx, []byte(req.Hash))
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTransactionResponse{Transaction: transaction}, nil
}
