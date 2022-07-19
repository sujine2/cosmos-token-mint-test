package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sujine/test/x/test/types"
)

func (k msgServer) RegisterCoin(goCtx context.Context, msg *types.MsgRegisterCoin) (*types.MsgRegisterCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Handling the message
	sell, isFound := k.GetSymbolList(ctx, msg.Symbol)
	if isFound {
		address, _ := sdk.AccAddressFromBech32(msg.Creator)
		conInfo := k.bankKeeper.GetBalance(ctx, address, msg.Symbol)
		if conInfo.Amount.Int64() < int64(msg.Amount) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Not enough tokens")
		}

		_, isSell := k.GetSellCoin(ctx, msg.Symbol)
		if isSell {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "This coin is already on sale")
		}

		newSell := types.SellCoin{
			Index:   sell.Symbol,
			Creator: sell.Creator,
			Symbol: sell.Symbol,
			Price:   msg.Price,
			Amount:  msg.Amount,
		}
		k.SetSellCoin(ctx, newSell)

	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Coin does not exist")
	}

	_ = ctx

	return &types.MsgRegisterCoinResponse{}, nil
}
