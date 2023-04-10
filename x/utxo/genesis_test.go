package utxo_test

import (
	"testing"

	keepertest "utxo/testutil/keeper"
	"utxo/testutil/nullify"
	"utxo/x/utxo"
	"utxo/x/utxo/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	// Test addresses
	addresses := []string{
		"address1",
		"address2",
		"address3",
		"address4",
		"address5",
	}

	// Create initial UTXOs
	var utxoList []*types.TxOut
	for _, addr := range addresses {
		utxoList = append(utxoList, &types.TxOut{
			PkScript: addr,
			Value:    10000,
		})
	}

	// Create initial transaction with all UTXOs in a single txOut
	transaction := types.Transaction{
		Version:  1,
		TxOut:    utxoList,
		Locktime: 0,
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TransactionList:  []types.Transaction{transaction},
		TransactionCount: 1,
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
