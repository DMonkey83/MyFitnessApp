package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomExercise(t *testing.T) Exercise {
	eq := CreateRandomEquipment(t)
	arg := CreateExerciseParams{
		ExerciseName: util.GetRandomUsername(5),
		MuscleGroup:  MuscleGroupEnumAbs,
		EquipmentID:  pgtype.Int8{Int64: int64(eq.EquipmentID), Valid: true},
	}

	exerc, err := testStore.CreateExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exerc)

	require.Equal(t, arg.MuscleGroup, exerc.MuscleGroup)
	require.Equal(t, arg.ExerciseName, exerc.ExerciseName)
	require.Equal(t, arg.MuscleGroup, exerc.MuscleGroup)

	require.NotZero(t, exerc.EquipmentID)
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
	require.Equal(t, ex1.MuscleGroup, ex2.MuscleGroup)
}

func TestUpdateExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)

	arg := UpdateExerciseParams{
		EquipmentID:  ex1.EquipmentID,
		ExerciseID:   ex1.ExerciseID,
		ExerciseName: util.GetRandomUsername(4),
		MuscleGroup:  MuscleGroupEnumArms,
	}

	eq2, err := testStore.UpdateExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, eq2)

	require.Equal(t, arg.EquipmentID, eq2.EquipmentID)
	require.Equal(t, arg.ExerciseName, eq2.ExerciseName)
	require.Equal(t, arg.MuscleGroup, eq2.MuscleGroup)
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
