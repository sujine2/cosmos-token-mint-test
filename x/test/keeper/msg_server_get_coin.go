package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sujine/test/x/test/types"
)

func (k msgServer) GetCoin(goCtx context.Context, msg *types.MsgGetCoin) (*types.MsgGetCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGetCoinResponse{}, nil
}
