package cosmosbridge_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tranquanghuy7198/CosmosBridge/testutil/keeper"
	"github.com/tranquanghuy7198/CosmosBridge/x/cosmosbridge"
	"github.com/tranquanghuy7198/CosmosBridge/x/cosmosbridge/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosbridgeKeeper(t)
	cosmosbridge.InitGenesis(ctx, *k, genesisState)
	got := cosmosbridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
