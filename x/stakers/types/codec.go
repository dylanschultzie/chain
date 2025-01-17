package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateStaker{}, "kyve/stakers/MsgCreateStaker", nil)
	cdc.RegisterConcrete(&MsgUpdateCommission{}, "kyve/stakers/MsgUpdateCommission", nil)
	cdc.RegisterConcrete(&MsgUpdateMetadata{}, "kyve/stakers/MsgUpdateMetadata", nil)
	cdc.RegisterConcrete(&MsgJoinPool{}, "kyve/stakers/MsgJoinPool", nil)
	cdc.RegisterConcrete(&MsgLeavePool{}, "kyve/stakers/MsgLeavePool", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "kyve/stakers/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCreateStaker{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateCommission{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateMetadata{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgJoinPool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgLeavePool{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateParams{})
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptoCodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
