package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sujine/test/x/test/types"
)

func (k msgServer) MintCoin(goCtx context.Context, msg *types.MsgMintCoin) (*types.MsgMintCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, isFound := k.GetSymbolList(ctx, msg.Denom)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Symbol already in used")
	}
	newCoin := sdk.NewInt64Coin(msg.Denom, int64(msg.Amount))
	newCoins := sdk.NewCoins(newCoin)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
	address, _ := sdk.AccAddressFromBech32(msg.Creator)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, newCoins)

	newSymbol := types.SymbolList{
		Index:   msg.Denom,
		Creator: msg.Creator,
		Symbol: msg.Denom,
	}
	k.SetSymbolList(ctx, newSymbol)

	return &types.MsgMintCoinResponse{}, nil
}
