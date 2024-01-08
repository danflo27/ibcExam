package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ibcExam/testutil/keeper"
	"ibcExam/x/exam/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ExamKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
