package sppt_test

import (
	"testing"

	keepertest "sppt/testutil/keeper"
	"sppt/testutil/nullify"
	sppt "sppt/x/sppt/module"
	"sppt/x/sppt/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SpptKeeper(t)
	sppt.InitGenesis(ctx, k, genesisState)
	got := sppt.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
