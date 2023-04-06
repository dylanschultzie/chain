package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	println("----->", ModuleName)
	//cdc.RegisterConcrete(&MsgFundPool{}, ModuleName+"/FundPool", nil)
	legacy.RegisterAminoMsg(cdc, &MsgFundPool{}, "cosmos-sdk/MsgFundPool")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgFundPool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgDefundPool{})

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCreatePool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdatePool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgDisablePool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgEnablePool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgScheduleRuntimeUpgrade{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCancelRuntimeUpgrade{})
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	println("init call")
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	//Amino.Seal()
}
