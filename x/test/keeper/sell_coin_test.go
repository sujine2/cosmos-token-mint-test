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

func createNSellCoin(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SellCoin {
	items := make([]types.SellCoin, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSellCoin(ctx, items[i])
	}
	return items
}

func TestSellCoinGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSellCoin(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSellCoin(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSellCoinRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSellCoin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSellCoin(ctx,
			item.Index,
		)
		_, found := keeper.GetSellCoin(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSellCoinGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSellCoin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSellCoin(ctx)),
	)
}
