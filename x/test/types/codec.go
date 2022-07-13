package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintCoin{}, "test/MintCoin", nil)
	cdc.RegisterConcrete(&MsgRegisterCoin{}, "test/RegisterCoin", nil)
	cdc.RegisterConcrete(&MsgBuyCoin{}, "test/BuyCoin", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintCoin{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterCoin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCoin{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
