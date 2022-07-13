package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/sujine/test/testutil/keeper"
	"github.com/sujine/test/testutil/nullify"
	"github.com/sujine/test/x/test/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSymbolListQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSymbolList(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSymbolListRequest
		response *types.QueryGetSymbolListResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSymbolListRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSymbolListResponse{SymbolList: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSymbolListRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSymbolListResponse{SymbolList: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSymbolListRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SymbolList(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestSymbolListQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSymbolList(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSymbolListRequest {
		return &types.QueryAllSymbolListRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SymbolListAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SymbolList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SymbolList),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SymbolListAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SymbolList), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SymbolList),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SymbolListAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SymbolList),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SymbolListAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
