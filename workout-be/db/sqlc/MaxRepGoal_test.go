package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomMaxRepGoal(t *testing.T) Maxrepgoal {
	ru := CreateRandomUser(t)
	ex := CreateRandomExercise(t)
	arg := CreateMaxRepGoalParams{
		Username:   ru.Username,
		ExerciseID: ex.ExerciseID,
		GoalReps:   int32(util.GetRandomAmount(1, 100)),
		Notes:      util.GetRandomUsername(100),
	}

	maxRepG, err := testStore.CreateMaxRepGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, maxRepG)

	require.Equal(t, arg.ExerciseID, maxRepG.ExerciseID)
	require.Equal(t, arg.Notes, maxRepG.Notes)
	require.Equal(t, arg.GoalReps, maxRepG.GoalReps)
	require.Equal(t, arg.Username, maxRepG.Username)

	require.NotZero(t, maxRepG.GoalID)
	return maxRepG
}

func TestCreateMaxRepGoal(t *testing.T) {
	CreateRandomMaxRepGoal(t)
}

func TestGetMaxRepGoal(t *testing.T) {
	mrg1 := CreateRandomMaxRepGoal(t)
	mrg2, err := testStore.GetMaxRepGoal(context.Background(), mrg1.GoalID)
	require.NoError(t, err)
	require.NotEmpty(t, mrg1)

	require.Equal(t, mrg1.ExerciseID, mrg2.ExerciseID)
	require.Equal(t, mrg1.Notes, mrg2.Notes)
	require.Equal(t, mrg1.Username, mrg2.Username)
	require.Equal(t, mrg1.GoalReps, mrg2.GoalReps)
}

func TestUpdateMaxRepGoal(t *testing.T) {
	mrg1 := CreateRandomMaxRepGoal(t)

	arg := UpdateMaxRepGoalParams{
		ExerciseID: mrg1.ExerciseID,
		GoalID:     mrg1.GoalID,
		GoalReps:   mrg1.GoalReps,
		Notes:      mrg1.Notes,
	}

	mrg2, err := testStore.UpdateMaxRepGoal(context.Background(), arg)
	require.Equal(t, mrg1.ExerciseID, mrg2.ExerciseID)
	require.Equal(t, mrg1.Notes, mrg2.Notes)
	require.Equal(t, mrg1.Username, mrg2.Username)
	require.Equal(t, mrg1.GoalReps, mrg2.GoalReps)
	require.NoError(t, err)
	require.NotEmpty(t, mrg2)

}

func TestDeleteMaxRepGoal(t *testing.T) {
	mrg1 := CreateRandomMaxRepGoal(t)
	err := testStore.DeleteMaxRepGoal(context.Background(), mrg1.GoalID)
	require.NoError(t, err)

	mrg2, err := testStore.GetMaxRepGoal(context.Background(), mrg1.GoalID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, mrg2)
}
