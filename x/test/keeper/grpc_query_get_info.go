package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sujine/test/x/test/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetInfo(goCtx context.Context, req *types.QueryGetInfoRequest) (*types.QueryGetInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	address, _ := sdk.AccAddressFromBech32(req.Creator)
	conInfo := k.bankKeeper.GetBalance(ctx, address, req.Denom)
	// TODO: Process the query
	_ = ctx

	return &types.QueryGetInfoResponse{Denom: conInfo.GetDenom(), Info: conInfo.String()}, nil
}
