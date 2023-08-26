package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomMaxWeightGoal(t *testing.T) Maxweightgoal {
	ru := CreateRandomUser(t)
	ex := CreateRandomExercise(t)
	arg := CreateMaxWeightGoalParams{
		Username:   ru.Username,
		ExerciseID: ex.ExerciseID,
		GoalWeight: int32(util.GetRandomAmount(1, 100)),
		Notes:      util.GetRandomUsername(100),
	}

	maxWG, err := testStore.CreateMaxWeightGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, maxWG)

	require.Equal(t, arg.ExerciseID, maxWG.ExerciseID)
	require.Equal(t, arg.Notes, maxWG.Notes)
	require.Equal(t, arg.GoalWeight, maxWG.GoalWeight)
	require.Equal(t, arg.Username, maxWG.Username)

	require.NotZero(t, maxWG.GoalID)
	return maxWG
}

func TestCreateMaxWeightGoal(t *testing.T) {
	CreateRandomMaxWeightGoal(t)
}

func TestGetMaxWeightGoal(t *testing.T) {
	mwg1 := CreateRandomMaxWeightGoal(t)
	mwg2, err := testStore.GetMaxWeightGoal(context.Background(), mwg1.GoalID)
	require.NoError(t, err)
	require.NotEmpty(t, mwg1)

	require.Equal(t, mwg1.ExerciseID, mwg2.ExerciseID)
	require.Equal(t, mwg1.Notes, mwg2.Notes)
	require.Equal(t, mwg1.Username, mwg2.Username)
	require.Equal(t, mwg1.GoalWeight, mwg2.GoalWeight)
}

func TestUpdateMaxWeightGoal(t *testing.T) {
	mwg1 := CreateRandomMaxWeightGoal(t)

	arg := UpdateMaxWeightGoalParams{
		ExerciseID: mwg1.ExerciseID,
		GoalID:     mwg1.GoalID,
		GoalWeight: mwg1.GoalWeight,
		Notes:      mwg1.Notes,
	}

	mwg2, err := testStore.UpdateMaxWeightGoal(context.Background(), arg)
	require.Equal(t, mwg1.ExerciseID, mwg2.ExerciseID)
	require.Equal(t, mwg1.Notes, mwg2.Notes)
	require.Equal(t, mwg1.Username, mwg2.Username)
	require.Equal(t, mwg1.GoalWeight, mwg2.GoalWeight)
	require.NoError(t, err)
	require.NotEmpty(t, mwg2)

}

func TestDeleteMaxWeightGoal(t *testing.T) {
	mwg1 := CreateRandomMaxWeightGoal(t)
	err := testStore.DeleteMaxWeightGoal(context.Background(), mwg1.GoalID)
	require.NoError(t, err)

	mwg2, err := testStore.GetMaxWeightGoal(context.Background(), mwg1.GoalID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, mwg2)
}
