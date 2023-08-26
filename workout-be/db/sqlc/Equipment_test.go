package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEquipment(t *testing.T) Equipment {
	arg := CreateEquipmentParams{
		EquipmentName: util.GetRandomUsername(8),
		Description:   util.GetRandomUsername(40),
		EquipmentType: EquipmenttypeBarbell,
	}

	eq, err := testStore.CreateEquipment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, eq)

	require.Equal(t, arg.EquipmentName, eq.EquipmentName)
	require.Equal(t, arg.Description, eq.Description)
	require.Equal(t, arg.EquipmentType, eq.EquipmentType)
	return eq
}

func TestCreateEqupment(t *testing.T) {
	CreateRandomEquipment(t)
}

func TestGetEquipment(t *testing.T) {
	eq1 := CreateRandomEquipment(t)
	eq2, err := testStore.GetEquipment(context.Background(), eq1.EquipmentName)
	require.NoError(t, err)
	require.NotEmpty(t, eq2)

	require.Equal(t, eq1.EquipmentName, eq2.EquipmentName)
	require.Equal(t, eq1.EquipmentType, eq2.EquipmentType)
	require.Equal(t, eq1.Description, eq2.Description)
}

func TestUpdateEquipment(t *testing.T) {
	ex1 := CreateRandomEquipment(t)

	arg := UpdateEquipmentParams{
		EquipmentName: ex1.EquipmentName,
		Description:   util.GetRandomUsername(40),
		EquipmentType: EquipmenttypeMachine,
	}

	eq2, err := testStore.UpdateEquipment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, eq2)

	require.Equal(t, arg.EquipmentType, eq2.EquipmentType)
	require.Equal(t, arg.Description, eq2.Description)
}

func TestDeleteEquipment(t *testing.T) {
	eq1 := CreateRandomEquipment(t)
	err := testStore.DeleteEquipment(context.Background(), eq1.EquipmentName)
	require.NoError(t, err)

	eq2, err := testStore.GetEquipment(context.Background(), eq1.EquipmentName)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, eq2)
}
