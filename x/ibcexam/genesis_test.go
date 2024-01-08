package ibcexam_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ibcExam/testutil/keeper"
	"ibcExam/testutil/nullify"
	"ibcExam/x/ibcexam"
	"ibcExam/x/ibcexam/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IbcexamKeeper(t)
	ibcexam.InitGenesis(ctx, *k, genesisState)
	got := ibcexam.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
