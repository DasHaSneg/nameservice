package nameservice

import (
	"fmt"
	"nameservice/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case types.MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case types.MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSetName(ctx sdk.Context, k Keeper, msg types.MsgSetName) sdk.Result {
	if !msg.Owner.Equals(k.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	k.SetName(ctx, msg.Name, msg.Value)
	return sdk.Result{}
}

func handleMsgBuyName(ctx sdk.Context, k Keeper, msg types.MsgBuyName) sdk.Result {
	if k.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result()
	}
	if k.HasOwner(ctx, msg.Name) {
		err := k.coinKeeper.SendCoins(ctx, msg.Buyer, k.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := k.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	k.SetOwner(ctx, msg.Name, msg.Buyer)
	k.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}

func handleMsgDeleteName(ctx sdk.Context, k Keeper, msg types.MsgDeleteName) sdk.Result {
	if !k.IsNamePresent(ctx, msg.Name) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(k.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	k.DeleteWhois(ctx, msg.Name)
	return sdk.Result{}
}
