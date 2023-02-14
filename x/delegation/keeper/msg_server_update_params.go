package keeper

import (
	"context"
	"encoding/json"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	// Delegation
	"github.com/KYVENetwork/chain/x/delegation/types"
	// Gov
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// UpdateParams is a governance message to update module-wide parameters.
// req.payload is a valid json string
// This is already checked and validated by the `types/params.go`
// Only the provided properties will be updated, the rest remains the same.
func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.authority != req.Authority {
		return nil, errors.Wrapf(govTypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	payload := params
	_ = json.Unmarshal([]byte(req.Payload), &payload)
	k.SetParams(ctx, payload)

	return &types.MsgUpdateParamsResponse{}, nil
}