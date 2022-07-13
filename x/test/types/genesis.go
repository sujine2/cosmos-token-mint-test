package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SellCoinList:   []SellCoin{},
		SymbolListList: []SymbolList{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in sellCoin
	sellCoinIndexMap := make(map[string]struct{})

	for _, elem := range gs.SellCoinList {
		index := string(SellCoinKey(elem.Index))
		if _, ok := sellCoinIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sellCoin")
		}
		sellCoinIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in symbolList
	symbolListIndexMap := make(map[string]struct{})

	for _, elem := range gs.SymbolListList {
		index := string(SymbolListKey(elem.Index))
		if _, ok := symbolListIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for symbolList")
		}
		symbolListIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
