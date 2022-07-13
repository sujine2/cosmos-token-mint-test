package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sujine/test/x/test/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SymbolListAll(c context.Context, req *types.QueryAllSymbolListRequest) (*types.QueryAllSymbolListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var symbolLists []types.SymbolList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	symbolListStore := prefix.NewStore(store, types.KeyPrefix(types.SymbolListKeyPrefix))

	pageRes, err := query.Paginate(symbolListStore, req.Pagination, func(key []byte, value []byte) error {
		var symbolList types.SymbolList
		if err := k.cdc.Unmarshal(value, &symbolList); err != nil {
			return err
		}

		symbolLists = append(symbolLists, symbolList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSymbolListResponse{SymbolList: symbolLists, Pagination: pageRes}, nil
}

func (k Keeper) SymbolList(c context.Context, req *types.QueryGetSymbolListRequest) (*types.QueryGetSymbolListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSymbolList(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSymbolListResponse{SymbolList: val}, nil
}
