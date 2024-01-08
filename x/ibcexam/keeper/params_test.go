package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ibcExam/testutil/keeper"
	"ibcExam/x/ibcexam/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IbcexamKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
