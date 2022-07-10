package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sujine/test/x/test/types"
)

func (k msgServer) MintCoin(goCtx context.Context, msg *types.MsgMintCoin) (*types.MsgMintCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newCoin := sdk.NewCoin(msg.Denom, sdk.NewInt(int64(msg.Amount)))
	newCoins := sdk.NewCoins(newCoin)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(msg.Creator), newCoins)

	return &types.MsgMintCoinResponse{}, nil
}
