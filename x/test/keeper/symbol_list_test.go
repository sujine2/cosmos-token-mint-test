package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/sujine/test/testutil/keeper"
	"github.com/sujine/test/testutil/nullify"
	"github.com/sujine/test/x/test/keeper"
	"github.com/sujine/test/x/test/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSymbolList(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SymbolList {
	items := make([]types.SymbolList, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSymbolList(ctx, items[i])
	}
	return items
}

func TestSymbolListGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSymbolList(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSymbolList(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSymbolListRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSymbolList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSymbolList(ctx,
			item.Index,
		)
		_, found := keeper.GetSymbolList(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSymbolListGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSymbolList(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSymbolList(ctx)),
	)
}
