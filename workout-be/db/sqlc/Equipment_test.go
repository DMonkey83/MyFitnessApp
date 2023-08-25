package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomEquipment(t *testing.T) Equipment {

	arg := CreateEquipmentParams{
		EquipmentName: util.GetRandomUsername(7),
		Description:   pgtype.Text{String: util.GetRandomUsername(40), Valid: true},
		EquipmentType: EquipmenttypeOther,
	}

	equipment, err := testStore.CreateEquipment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, equipment)

	require.Equal(t, arg.EquipmentName, equipment.EquipmentName)
	require.Equal(t, arg.EquipmentType, equipment.EquipmentType)
	require.Equal(t, arg.Description, equipment.Description)

	require.NotZero(t, equipment.EquipmentID)
	return equipment
}

func TestCreateEquipment(t *testing.T) {
	CreateRandomEquipment(t)
}

func TestGetEquipment(t *testing.T) {
	eq1 := CreateRandomEquipment(t)
	eq2, err := testStore.GetEquipment(context.Background(), eq1.EquipmentID)
	require.NoError(t, err)
	require.NotEmpty(t, eq2)

	require.Equal(t, eq1.EquipmentName, eq2.EquipmentName)
	require.Equal(t, eq1.Description, eq2.Description)
	require.Equal(t, eq1.EquipmentType, eq2.EquipmentType)
}

func TestUpdateEquipment(t *testing.T) {
	eq1 := CreateRandomEquipment(t)

	arg := UpdateEquipmentParams{
		EquipmentID:   eq1.EquipmentID,
		EquipmentName: util.GetRandomUsername(4),
		EquipmentType: EquipmenttypeBarbell,
		Description:   pgtype.Text{String: util.GetRandomUsername(40), Valid: true},
	}

	eq2, err := testStore.UpdateEquipment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, eq2)

	require.Equal(t, arg.EquipmentID, eq2.EquipmentID)
	require.Equal(t, arg.EquipmentName, eq2.EquipmentName)
	require.Equal(t, arg.EquipmentType, eq2.EquipmentType)
	require.Equal(t, arg.Description, eq2.Description)
}

func TestDeleteEquipment(t *testing.T) {
	eq1 := CreateRandomEquipment(t)
	err := testStore.DeleteEquipment(context.Background(), eq1.EquipmentID)
	require.NoError(t, err)

	eq2, err := testStore.GetEquipment(context.Background(), eq1.EquipmentID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, eq2)
}
