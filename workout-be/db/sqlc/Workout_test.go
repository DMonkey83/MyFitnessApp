package db

import (
	"context"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomWorkout(t *testing.T) Workout {
	user := CreateRandomUser(t)

	arg := CreateWorkoutParams{
		Username:        user.Username,
		Notes:           util.GetRandomUsername(100),
		WorkoutDuration: "0h10m",
		WorkoutDate:     time.Now(),
	}

	workout, err := testStore.CreateWorkout(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, workout)

	require.Equal(t, arg.Username, workout.Username)
	require.Equal(t, arg.Notes, workout.Notes)
	require.Equal(t, arg.WorkoutDuration, workout.WorkoutDuration)

	require.NotZero(t, workout.WorkoutID)
	return workout
}

func TestCreateWorkout(t *testing.T) {
	CreateRandomWorkout(t)
}

func TestGetWorkout(t *testing.T) {
	wkout1 := CreateRandomWorkout(t)
	wkout2, err := testStore.GetWorkout(context.Background(), wkout1.WorkoutID)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, wkout1.Username, wkout2.Username)
	require.Equal(t, wkout1.Notes, wkout2.Notes)
	require.Equal(t, wkout1.WorkoutDuration, wkout2.WorkoutDuration)
}

func TestUpdateWorkout(t *testing.T) {
	wkout1 := CreateRandomWorkout(t)

	arg := UpdateWorkoutParams{
		WorkoutID:       wkout1.WorkoutID,
		Username:        wkout1.Username,
		Notes:           wkout1.Notes,
		WorkoutDuration: "1h",
		WorkoutDate:     time.Now(),
	}

	wkout2, err := testStore.UpdateWorkout(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, arg.WorkoutID, wkout2.WorkoutID)
	require.Equal(t, arg.Username, wkout2.Username)
	require.Equal(t, arg.WorkoutDuration, wkout2.WorkoutDuration)
}

func TestDeleteWorkout(t *testing.T) {
	wkout1 := CreateRandomWorkout(t)
	err := testStore.DeleteWorkout(context.Background(), wkout1.WorkoutID)
	require.NoError(t, err)

	wkout2, err := testStore.GetWorkout(context.Background(), wkout1.WorkoutID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, wkout2)
}
