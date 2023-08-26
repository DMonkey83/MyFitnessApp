package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomExercise(t *testing.T) Exercise {
	eq := CreateRandomEquipment(t)
	wk := CreateRandomWorkout(t)
	arg := CreateExerciseParams{
		ExerciseName:    util.GetRandomUsername(5),
		WorkoutID:       wk.WorkoutID,
		EquipmentName:   eq.EquipmentName,
		MuscleGroupName: MusclegroupenumAbs,
	}

	exerc, err := testStore.CreateExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exerc)

	require.Equal(t, arg.ExerciseName, exerc.ExerciseName)
	require.Equal(t, arg.WorkoutID, exerc.WorkoutID)

	require.NotZero(t, exerc.EquipmentName)
	return exerc
}

func TestCreateExercise(t *testing.T) {
	CreateRandomExercise(t)
}

func TestGetExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)
	ex2, err := testStore.GetExercise(context.Background(), ex1.ExerciseID)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, ex1.ExerciseID, ex2.ExerciseID)
	require.Equal(t, ex1.ExerciseName, ex2.ExerciseName)
	require.Equal(t, ex1.WorkoutID, ex2.WorkoutID)
}

func TestUpdateExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)

	arg := UpdateExerciseParams{
		EquipmentName:   ex1.EquipmentName,
		ExerciseID:      ex1.ExerciseID,
		ExerciseName:    util.GetRandomUsername(4),
		MuscleGroupName: MusclegroupenumBack,
	}

	eq2, err := testStore.UpdateExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, eq2)

	require.Equal(t, arg.EquipmentName, eq2.EquipmentName)
	require.Equal(t, arg.ExerciseName, eq2.ExerciseName)
}

func TestDeleteExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)
	err := testStore.DeleteExercise(context.Background(), ex1.ExerciseID)
	require.NoError(t, err)

	ex2, err := testStore.GetExercise(context.Background(), ex1.ExerciseID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, ex2)
}
