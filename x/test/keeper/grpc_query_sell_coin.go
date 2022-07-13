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

func (k Keeper) SellCoinAll(c context.Context, req *types.QueryAllSellCoinRequest) (*types.QueryAllSellCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sellCoins []types.SellCoin
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sellCoinStore := prefix.NewStore(store, types.KeyPrefix(types.SellCoinKeyPrefix))

	pageRes, err := query.Paginate(sellCoinStore, req.Pagination, func(key []byte, value []byte) error {
		var sellCoin types.SellCoin
		if err := k.cdc.Unmarshal(value, &sellCoin); err != nil {
			return err
		}

		sellCoins = append(sellCoins, sellCoin)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSellCoinResponse{SellCoin: sellCoins, Pagination: pageRes}, nil
}

func (k Keeper) SellCoin(c context.Context, req *types.QueryGetSellCoinRequest) (*types.QueryGetSellCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSellCoin(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSellCoinResponse{SellCoin: val}, nil
}
