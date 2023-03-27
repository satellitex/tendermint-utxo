package utxo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "utxo/testutil/keeper"
	"utxo/testutil/nullify"
	"utxo/x/utxo"
	"utxo/x/utxo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TransactionList: []types.Transaction{
		{
			Id: 0,
		},
		{
			Id: 1,
		},
	},
	TransactionCount: 2,
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UtxoKeeper(t)
	utxo.InitGenesis(ctx, *k, genesisState)
	got := utxo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TransactionList, got.TransactionList)
require.Equal(t, genesisState.TransactionCount, got.TransactionCount)
// this line is used by starport scaffolding # genesis/test/assert
}
