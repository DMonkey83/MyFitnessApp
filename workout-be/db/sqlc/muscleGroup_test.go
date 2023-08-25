package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomMuscleGroup(t *testing.T) Musclegroup {

	mgroup, err := testStore.CreateMuscleGroup(context.Background(), string(MuscleGroupEnumLegs))
	require.NoError(t, err)
	require.NotEmpty(t, mgroup)

	require.Equal(t, string("Legs"), mgroup.MuscleGroupName)

	require.NotZero(t, mgroup.MuscleGroupID)
	return mgroup
}

func TestCreateMuscleGroup(t *testing.T) {
	CreateRandomMuscleGroup(t)
}

func TestGetMuscleGroup(t *testing.T) {
	mg1 := CreateRandomMuscleGroup(t)
	mg2, err := testStore.GetMuscleGroup(context.Background(), mg1.MuscleGroupID)
	require.NoError(t, err)
	require.NotEmpty(t, mg2)

	require.Equal(t, mg1.MuscleGroupName, mg2.MuscleGroupName)
}

func TestUpdateMuscleGroup(t *testing.T) {
	mg1 := CreateRandomMuscleGroup(t)

	arg := UpdateMuscleGroupParams{
		MuscleGroupID:   mg1.MuscleGroupID,
		MuscleGroupName: string(MuscleGroupEnumBack),
	}

	mg2, err := testStore.UpdateMuscleGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, mg2)

	require.Equal(t, arg.MuscleGroupName, mg2.MuscleGroupName)
}

func TestDeleteMuscleGroup(t *testing.T) {
	mg1 := CreateRandomMuscleGroup(t)
	err := testStore.DeleteMuscleGroup(context.Background(), mg1.MuscleGroupID)
	require.NoError(t, err)

	mg2, err := testStore.GetMuscleGroup(context.Background(), mg1.MuscleGroupID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, mg2)
}
