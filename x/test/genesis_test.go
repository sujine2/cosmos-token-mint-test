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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestKeeper(t)
	test.InitGenesis(ctx, *k, genesisState)
	got := test.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
