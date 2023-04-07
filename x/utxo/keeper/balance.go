package keeper

import (
	"encoding/binary"

	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetOutpoint returns a outpoint from its pkScript
func (k Keeper) GetOutpoint(ctx sdk.Context, pkScript string) (val types.Outpoint, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutpointKey))
	b := store.Get([]byte(pkScript))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetOutpoint sets a specific outpoint in the store
func (k Keeper) SetOutpoint(ctx sdk.Context, pkScript string, outpoint types.Outpoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutpointKey))
	b := k.cdc.MustMarshal(&outpoint)
	store.Set([]byte(pkScript), b)
}

// RemoveOutpoint removes a outpoint from the store
func (k Keeper) RemoveOutpoint(ctx sdk.Context, pkScript string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutpointKey))
	store.Delete([]byte(pkScript))
}

// GetAllOutpoint returns all outpoint
func (k Keeper) GetAllOutpoint(ctx sdk.Context) (list []types.Outpoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutpointKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Outpoint
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOutpointIndexBytes returns the byte representation of the index
func GetOutpointIndexBytes(index uint32) []byte {
	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, index)
	return bz
}

// GetOutpointIndexFromBytes returns index in uint32 format from a byte array
func GetOutpointIndexFromBytes(bz []byte) uint32 {
	return binary.BigEndian.Uint32(bz)
}
