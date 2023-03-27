package types

import (
	"crypto/sha256"
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TransactionList: []Transaction{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in transaction
	transactionIdMap := make(map[string]bool)
	transactionCount := gs.GetTransactionCount()
	if len(gs.TransactionList) >= int(transactionCount) {
		return fmt.Errorf("transaction id should be lower or equal than the last id")
	}
	for _, elem := range gs.TransactionList {
		data, _ := elem.Marshal()
		elmHash := sha256.Sum256(data)
		elmHashStr := string(elmHash[:])
		if _, ok := transactionIdMap[elmHashStr]; ok {
			return fmt.Errorf("duplicated id for transaction")
		}
		transactionIdMap[elmHashStr] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
