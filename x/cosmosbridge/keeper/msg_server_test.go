package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tranquanghuy7198/CosmosBridge/testutil/keeper"
	"github.com/tranquanghuy7198/CosmosBridge/x/cosmosbridge/keeper"
	"github.com/tranquanghuy7198/CosmosBridge/x/cosmosbridge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosbridgeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
