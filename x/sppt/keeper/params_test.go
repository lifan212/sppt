package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "sppt/testutil/keeper"
	"sppt/x/sppt/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SpptKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
