package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterCoin = "register_coin"

var _ sdk.Msg = &MsgRegisterCoin{}

func NewMsgRegisterCoin(creator string, symbol string, price uint64, amount uint64) *MsgRegisterCoin {
	return &MsgRegisterCoin{
		Creator: creator,
		Symbol:  symbol,
		Price:   price,
		Amount:  amount,
	}
}

func (msg *MsgRegisterCoin) Route() string {
	return RouterKey
}

func (msg *MsgRegisterCoin) Type() string {
	return TypeMsgRegisterCoin
}

func (msg *MsgRegisterCoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
