package keeper

import (
	"utxo/x/utxo/types"
)

var _ types.QueryServer = Keeper{}
