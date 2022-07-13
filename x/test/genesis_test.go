package test_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/sujine/test/testutil/keeper"
	"github.com/sujine/test/testutil/nullify"
	"github.com/sujine/test/x/test"
	"github.com/sujine/test/x/test/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SellCoinList: []types.SellCoin{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		SymbolListList: []types.SymbolList{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestKeeper(t)
	test.InitGenesis(ctx, *k, genesisState)
	got := test.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SellCoinList, got.SellCoinList)
	require.ElementsMatch(t, genesisState.SymbolListList, got.SymbolListList)
	// this line is used by starport scaffolding # genesis/test/assert
}
