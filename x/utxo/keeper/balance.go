package keeper

import (
	"encoding/binary"

	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxOut returns a outpoint from its pkScript
func (k Keeper) GetTxOut(ctx sdk.Context, pkScript string) (val types.Outpoint, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxOutKey))
	b := store.Get([]byte(pkScript))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetTxOut sets a specific outpoint in the store
func (k Keeper) SetTxOut(ctx sdk.Context, pkScript string, outpoint types.Outpoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxOutKey))
	b := k.cdc.MustMarshal(&outpoint)
	store.Set([]byte(pkScript), b)
}

// RemoveTxOut removes a outpoint from the store
func (k Keeper) RemoveTxOut(ctx sdk.Context, pkScript string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxOutKey))
	store.Delete([]byte(pkScript))
}

// GetAllTxOut returns all outpoint
func (k Keeper) GetAllTxOut(ctx sdk.Context) (list []types.Outpoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxOutKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Outpoint
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTxOutIndexBytes returns the byte representation of the index
func GetTxOutIndexBytes(index uint32) []byte {
	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, index)
	return bz
}

// GetTxOutIndexFromBytes returns index in uint32 format from a byte array
func GetTxOutIndexFromBytes(bz []byte) uint32 {
	return binary.BigEndian.Uint32(bz)
}
