package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ovrclk/akash/x/audit/keeper"
	"github.com/ovrclk/akash/x/audit/types"
)

// NewHandler returns a handler for "provider" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgSignProviderAttributes:
			return handleMsgSignProviderAttributes(ctx, keeper, msg)
		case *types.MsgDeleteProviderAttributes:
			return handleMsgDeleteProviderAttributes(ctx, keeper, msg)
		}

		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized bank message type: %T", msg)
	}
}

func handleMsgSignProviderAttributes(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSignProviderAttributes) (*sdk.Result, error) {
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	var owner sdk.AccAddress
	if owner, err = sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return nil, err
	}

	provId := types.ProviderID{
		Owner:     owner,
		Validator: validator,
	}

	if err = keeper.CreateOrUpdateProviderAttributes(ctx, provId, msg.Attributes); err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgDeleteProviderAttributes(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgDeleteProviderAttributes) (*sdk.Result, error) {
	validator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	var owner sdk.AccAddress
	if owner, err = sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return nil, err
	}

	provId := types.ProviderID{
		Owner:     owner,
		Validator: validator,
	}

	if err = keeper.DeleteProviderAttributes(ctx, provId, msg.Attributes); err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
