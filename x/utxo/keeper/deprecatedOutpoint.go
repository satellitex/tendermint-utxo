package keeper

import (
	"encoding/binary"
	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDeprecatedOutpoint returns a deprecated status from its PreviousOutpoint
func (k Keeper) GetDeprecatedOutpoint(ctx sdk.Context, previousOutpoint types.Outpoint) (deprecated bool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeprecatedOutpointKey))
	outpointKey := outpointToByte(previousOutpoint)
	b := store.Get(outpointKey)
	if b == nil {
		return deprecated, false
	}
	deprecated = decodeBool(b)
	return deprecated, true
}

// SetDeprecatedOutpoint sets a specific deprecated status in the store
func (k Keeper) SetDeprecatedOutpoint(ctx sdk.Context, previousOutpoint types.Outpoint, deprecated bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeprecatedOutpointKey))
	outpointKey := outpointToByte(previousOutpoint)
	b := encodeBool(deprecated)
	store.Set(outpointKey, b)
}

// RemoveDeprecatedOutpoint removes a deprecated status from the store
func (k Keeper) RemoveDeprecatedOutpoint(ctx sdk.Context, previousOutpoint types.Outpoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeprecatedOutpointKey))
	outpointKey := outpointToByte(previousOutpoint)
	store.Delete(outpointKey)
}

// GetAllDeprecatedOutpoint returns all deprecated statuses
func (k Keeper) GetAllDeprecatedOutpoint(ctx sdk.Context) (outpoints map[types.Outpoint]bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeprecatedOutpointKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	outpoints = make(map[types.Outpoint]bool)

	for ; iterator.Valid(); iterator.Next() {
		outpoint := byteToOutpoint(iterator.Key())
		outpoints[outpoint] = decodeBool(iterator.Value())
	}

	return
}

func outpointToByte(outpoint types.Outpoint) []byte {
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, outpoint.Index)
	return append([]byte(outpoint.Hash), indexBytes...)
}

func byteToOutpoint(bz []byte) types.Outpoint {
	hash := string(bz[:len(bz)-4])
	index := binary.BigEndian.Uint32(bz[len(bz)-4:])
	return types.Outpoint{Hash: hash, Index: index}
}

func decodeBool(bz []byte) bool {
	return bz[0] == 1
}

func encodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}
