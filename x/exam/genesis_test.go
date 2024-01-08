package exam_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ibcExam/testutil/keeper"
	"ibcExam/testutil/nullify"
	"ibcExam/x/exam"
	"ibcExam/x/exam/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExamKeeper(t)
	exam.InitGenesis(ctx, *k, genesisState)
	got := exam.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
