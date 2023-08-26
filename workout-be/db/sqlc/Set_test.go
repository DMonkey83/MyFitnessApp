package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomSet(t *testing.T) Set {
	ex := CreateRandomExercise(t)
	arg := CreateSetParams{
		ExerciseID:   ex.ExerciseID,
		SetNumber:    int32(util.GetRandomAmount(1, 10)),
		Weight:       int32(util.GetRandomAmount(1, 200)),
		Notes:        util.GetRandomUsername(100),
		RestDuration: "1h",
	}

	set, err := testStore.CreateSet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, set)

	require.Equal(t, arg.ExerciseID, set.ExerciseID)
	require.Equal(t, arg.Notes, set.Notes)
	require.Equal(t, arg.Weight, set.Weight)
	require.Equal(t, arg.Notes, set.Notes)

	require.NotZero(t, set.SetID)
	return set
}

func TestCreateSet(t *testing.T) {
	CreateRandomSet(t)
}

func TestGeSet(t *testing.T) {
	set1 := CreateRandomSet(t)
	set2, err := testStore.GetSet(context.Background(), set1.SetID)
	require.NoError(t, err)
	require.NotEmpty(t, set1)

	require.Equal(t, set1.ExerciseID, set2.ExerciseID)
	require.Equal(t, set1.Notes, set2.Notes)
	require.Equal(t, set1.Notes, set2.Notes)
	require.Equal(t, set1.Weight, set2.Weight)
}

func TestUpdateSet(t *testing.T) {
	set1 := CreateRandomSet(t)

	arg := UpdateSetParams{
		SetID:        set1.SetID,
		SetNumber:    int32(util.GetRandomAmount(1, 10)),
		Weight:       int32(util.GetRandomAmount(1, 200)),
		Notes:        util.GetRandomUsername(100),
		RestDuration: "1m",
	}

	set2, err := testStore.UpdateSet(context.Background(), arg)
	require.Equal(t, arg.Notes, set2.Notes)
	require.Equal(t, arg.Weight, set2.Weight)
	require.Equal(t, arg.SetNumber, set2.SetNumber)
	require.NoError(t, err)
	require.NotEmpty(t, set2)

}

func TestDeleteSet(t *testing.T) {
	set1 := CreateRandomSet(t)
	err := testStore.DeleteSet(context.Background(), set1.SetID)
	require.NoError(t, err)

	set2, err := testStore.GetMaxWeightGoal(context.Background(), set1.SetID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, set2)
}
