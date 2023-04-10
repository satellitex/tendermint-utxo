package utxo_test

import (
	"testing"

	keepertest "utxo/testutil/keeper"
	"utxo/testutil/nullify"
	"utxo/x/utxo"
	"utxo/x/utxo/types"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func initKeys(num int) (pubKeys []string, privKeys []secp256k1.PrivKey) {
	for i := 0; i < num; i++ {
		privKey := secp256k1.GenPrivKey()
		privKeys = append(privKeys, privKey)
		pubKeys = append(pubKeys, privKey.PubKey().Address().String())
	}
	return
}

func TestGenesis(t *testing.T) {
	// Test addresses
	pubKeys, privKeys := initKeys(5)

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TransactionList: []types.Transaction{
			{
				Version: 1,
				TxIn:    []*types.TxIn{},
				TxOut: []*types.TxOut{
					{Value: 10000, PkScript: pubKeys[0]},
					{Value: 10000, PkScript: pubKeys[1]},
					{Value: 10000, PkScript: pubKeys[2]},
					{Value: 10000, PkScript: pubKeys[3]},
					{Value: 10000, PkScript: pubKeys[4]},
				},
				Locktime: 0,
			},
		},
		TransactionCount: 1,
	}

	t.Log("Generated private keys:")
	for i, privKey := range privKeys {
		t.Logf("Keys #%d: (%s, %s, %x)\n", i+1,
			genesisState.TransactionList[0].TxOut[i].PkScript, pubKeys[i], privKey.Bytes())
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
