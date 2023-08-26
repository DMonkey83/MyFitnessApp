package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomReps(t *testing.T) Rep {
	set := CreateRandomSet(t)
	arg := CreateRepParams{
		SetID:            set.SetID,
		RepNumber:        int32(util.GetRandomAmount(1, 20)),
		CompletionStatus: CompletionenumIncomplete,
		Notes:            util.GetRandomUsername(100),
	}

	rep, err := testStore.CreateRep(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rep)

	require.Equal(t, arg.CompletionStatus, rep.CompletionStatus)
	require.Equal(t, arg.Notes, rep.Notes)
	require.Equal(t, arg.RepNumber, rep.RepNumber)
	require.Equal(t, arg.SetID, rep.SetID)

	require.NotZero(t, rep.RepID)
	return rep
}

func TestCreateRep(t *testing.T) {
	CreateRandomReps(t)
}

func TestGetReps(t *testing.T) {
	set1 := CreateRandomSet(t)
	set2, err := testStore.GetSet(context.Background(), set1.SetID)
	require.NoError(t, err)
	require.NotEmpty(t, set1)

	require.Equal(t, set1.ExerciseID, set2.ExerciseID)
	require.Equal(t, set1.Notes, set2.Notes)
	require.Equal(t, set1.Notes, set2.Notes)
	require.Equal(t, set1.Weight, set2.Weight)
}

func TestUpdateReps(t *testing.T) {
	rep1 := CreateRandomReps(t)

	arg := UpdateRepParams{
		RepID:            rep1.RepID,
		RepNumber:        int32(util.GetRandomAmount(1, 20)),
		CompletionStatus: CompletionenumCompleted,
		Notes:            util.GetRandomUsername(100),
	}

	rep2, err := testStore.UpdateRep(context.Background(), arg)
	require.Equal(t, arg.CompletionStatus, rep2.CompletionStatus)
	require.Equal(t, arg.RepNumber, rep2.RepNumber)
	require.Equal(t, arg.Notes, rep2.Notes)
	require.NoError(t, err)
	require.NotEmpty(t, rep2)

}

func TestDeleteReps(t *testing.T) {
	rep1 := CreateRandomReps(t)
	err := testStore.DeleteRep(context.Background(), rep1.RepID)
	require.NoError(t, err)

	rep2, err := testStore.GetRep(context.Background(), rep1.RepID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, rep2)
}
