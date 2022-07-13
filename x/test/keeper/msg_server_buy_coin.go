package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sujine/test/x/test/types"
)

func (k msgServer) BuyCoin(goCtx context.Context, msg *types.MsgBuyCoin) (*types.MsgBuyCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	getCoin, isFound := k.GetSellCoin(ctx, msg.Symbol)
	if isFound {
		if getCoin.Amount < msg.Amount {
			return nil, sdkerrors.Wrap(types.ErrInsufficientAmount, "Too many coins")
		}
		//base, _ := sdk.GetBaseDenom()
		//buyAmount := sdk.Coins{sdk.NewInt64Coin(base, int64(msg.Amount)*int64(getCoin.Price))} 

		addressBuyer, _ := sdk.AccAddressFromBech32(msg.Creator) 
		addressSeller, _ := sdk.AccAddressFromBech32(getCoin.Creator) 
		buyAmount := sdk.Coins{k.bankKeeper.GetBalance(ctx, addressBuyer, "token")}
		if buyAmount[0].Amount.Uint64() < msg.Amount*getCoin.Price {
			return nil, sdkerrors.Wrap(types.ErrInsufficientAmount, "account balance")
		}
		//코인 가격 지불  
		buyAmount[0].Amount = sdk.NewIntFromUint64(msg.Amount * getCoin.Price)
		k.bankKeeper.SendCoins(ctx, addressBuyer, addressSeller, buyAmount)

		sendAmount := sdk.Coins{k.bankKeeper.GetBalance(ctx, addressSeller, msg.Symbol)}
		sendAmount[0].Amount = sdk.NewIntFromUint64(msg.Amount)
		k.bankKeeper.SendCoins(ctx, addressSeller, addressBuyer, sendAmount)
		resultAmount := getCoin.Amount - msg.Amount

		newSell := types.SellCoin{
			Index:   getCoin.Symbole,
			Creator: getCoin.Creator,
			Symbole: getCoin.Symbole,
			Price:   getCoin.Price,
			Amount:  resultAmount,
		}
		k.SetSellCoin(ctx, newSell)
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Coin does not exist")
	}
	_ = ctx

	return &types.MsgBuyCoinResponse{}, nil
}
