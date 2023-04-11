package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"cosmossdk.io/errors"

	"github.com/KYVENetwork/chain/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgUpdatePool{}
	_ sdk.Msg = &MsgDisablePool{}
	_ sdk.Msg = &MsgEnablePool{}
	_ sdk.Msg = &MsgScheduleRuntimeUpgrade{}
	_ sdk.Msg = &MsgCancelRuntimeUpgrade{}

	_ legacytx.LegacyMsg = &MsgFundPool{}
)

func (msg MsgFundPool) Route() string { return ModuleName }

func (msg MsgFundPool) Type() string {
	return sdk.MsgTypeURL(&msg)
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgFundPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgCreatePool message.
func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgCreatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := util.ValidatePositiveNumber(msg.UploadInterval); err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, "invalid upload interval")
	}

	if err := util.ValidatePositiveNumber(msg.OperatingCost); err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, "invalid operating cost")
	}

	if err := util.ValidatePositiveNumber(msg.MinDelegation); err != nil {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, "invalid minimum delegation")
	}

	return nil
}

// GetSigners returns the expected signers for a MsgUpdatePool message.
func (msg *MsgUpdatePool) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// PoolUpdate ...
type PoolUpdate struct {
	Name              *string
	Runtime           *string
	Logo              *string
	Config            *string
	UploadInterval    *uint64
	OperatingCost     *uint64
	MinDelegation     *uint64
	MaxBundleSize     *uint64
	StorageProviderId *uint32
	CompressionId     *uint32
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgUpdatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	var payload PoolUpdate
	if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
		return err
	}

	if payload.UploadInterval != nil {
		if err := util.ValidatePositiveNumber(*payload.UploadInterval); err != nil {
			return errors.Wrapf(errorsTypes.ErrInvalidRequest, "invalid upload interval")
		}
	}

	if payload.OperatingCost != nil {
		if err := util.ValidatePositiveNumber(*payload.OperatingCost); err != nil {
			return errors.Wrapf(errorsTypes.ErrInvalidRequest, "invalid operating cost")
		}
	}

	if payload.MinDelegation != nil {
		if err := util.ValidatePositiveNumber(*payload.MinDelegation); err != nil {
			return errors.Wrapf(errorsTypes.ErrInvalidRequest, "invalid minimum delegation")
		}
	}

	return nil
}

// GetSigners returns the expected signers for a MsgDisablePool message.
func (msg *MsgDisablePool) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgDisablePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	return nil
}

// GetSigners returns the expected signers for a MsgEnablePool message.
func (msg *MsgEnablePool) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgEnablePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	return nil
}

// GetSigners returns the expected signers for a MsgScheduleRuntimeUpgrade message.
func (msg *MsgScheduleRuntimeUpgrade) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgScheduleRuntimeUpgrade) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	return nil
}

// GetSigners returns the expected signers for a MsgCancelRuntimeUpgrade message.
func (msg *MsgCancelRuntimeUpgrade) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgCancelRuntimeUpgrade) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	return nil
}
