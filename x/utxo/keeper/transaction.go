package keeper

import (
	"crypto/sha256"
	"encoding/binary"

	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTransactionCount get the total number of transaction
func (k Keeper) GetTransactionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TransactionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTransactionCount set the total number of transaction
func (k Keeper) SetTransactionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TransactionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTransaction appends a transaction in the store with a new id and update the count
func (k Keeper) AppendTransaction(
	ctx sdk.Context,
	txHashStr string,
	transaction types.Transaction,
) uint64 {
	// Create the transaction
	count := k.GetTransactionCount(ctx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	appendedValue := k.cdc.MustMarshal(&transaction)
	store.Set([]byte(txHashStr), appendedValue)

	// Update transaction count
	k.SetTransactionCount(ctx, count+1)

	return count
}

// SetTransaction set a specific transaction in the store
func (k Keeper) SetTransaction(ctx sdk.Context, transaction types.Transaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	b := k.cdc.MustMarshal(&transaction)
	hash := sha256.Sum256(b)
	store.Set(hash[:], b)
}

// GetTransaction returns a transaction from its id
func (k Keeper) GetTransaction(ctx sdk.Context, hash []byte) (val types.Transaction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	b := store.Get(hash)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTransaction removes a transaction from the store
func (k Keeper) RemoveTransaction(ctx sdk.Context, hash []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	store.Delete(hash)
}

// GetAllTransaction returns all transaction
func (k Keeper) GetAllTransaction(ctx sdk.Context) (list []types.Transaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Transaction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTransactionIDBytes returns the byte representation of the ID
func GetTransactionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTransactionIDFromBytes returns ID in uint64 format from a byte array
func GetTransactionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
