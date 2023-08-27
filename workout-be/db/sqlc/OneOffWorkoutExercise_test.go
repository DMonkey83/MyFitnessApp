package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomOneOffWorkoutExercise(t *testing.T) Oneoffworkoutexercise {

	exercise := CreateRandomExercise(t)
	workout := CreateRandomWorkout(t)
	arg := CreateOneOffWorkoutExerciseParams{
		ExerciseName:    exercise.ExerciseName,
		WorkoutID:       workout.WorkoutID,
		Description:     util.GetRandomUsername(100),
		MuscleGroupName: MusclegroupenumAbs,
	}

	ex, err := testStore.CreateOneOffWorkoutExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex)

	require.Equal(t, arg.ExerciseName, ex.ExerciseName)
	require.Equal(t, arg.Description, ex.Description)
	require.Equal(t, arg.WorkoutID, ex.WorkoutID)
	require.Equal(t, arg.MuscleGroupName, ex.MuscleGroupName)

	require.NotEmpty(t, ex.ID)

	return ex
}

func TestCreateOneOffWorkoutExercise(t *testing.T) {
	CreateRandomOneOffWorkoutExercise(t)
}

func TestGetOneOffWorkoutExercise(t *testing.T) {
	ex1 := CreateRandomOneOffWorkoutExercise(t)
	ex2, err := testStore.GetOneOffWorkoutExercise(context.Background(), ex1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, ex1.ExerciseName, ex2.ExerciseName)
	require.Equal(t, ex1.Description, ex2.Description)
	require.Equal(t, ex1.WorkoutID, ex2.WorkoutID)
	require.Equal(t, ex1.MuscleGroupName, ex2.MuscleGroupName)
}

func TestUpdateOffWorkoutExercise(t *testing.T) {
	ex1 := CreateRandomOneOffWorkoutExercise(t)

	arg := UpdateOneOffWorkoutExerciseParams{
		Description:     util.GetRandomUsername(100),
		MuscleGroupName: MusclegroupenumAbs,
		ID:              ex1.ID,
		WorkoutID:       ex1.WorkoutID,
	}

	ex2, err := testStore.UpdateOneOffWorkoutExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, ex1.ExerciseName, ex2.ExerciseName)
	require.Equal(t, arg.ID, ex2.ID)
	require.Equal(t, arg.Description, ex2.Description)
	require.Equal(t, arg.MuscleGroupName, ex2.MuscleGroupName)
	require.Equal(t, arg.WorkoutID, ex2.WorkoutID)

}

func TestDeleteOffWorkoutExercise(t *testing.T) {
	ex1 := CreateRandomOneOffWorkoutExercise(t)
	err := testStore.DeleteOneOffWorkoutExercise(context.Background(), ex1.ID)
	require.NoError(t, err)

	ex2, err := testStore.GetOneOffWorkoutExercise(context.Background(), ex1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, ex2)
}