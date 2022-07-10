package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGetCoin = "get_coin"

var _ sdk.Msg = &MsgGetCoin{}

func NewMsgGetCoin(creator string) *MsgGetCoin {
	return &MsgGetCoin{
		Creator: creator,
	}
}

func (msg *MsgGetCoin) Route() string {
	return RouterKey
}

func (msg *MsgGetCoin) Type() string {
	return TypeMsgGetCoin
}

func (msg *MsgGetCoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGetCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGetCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
