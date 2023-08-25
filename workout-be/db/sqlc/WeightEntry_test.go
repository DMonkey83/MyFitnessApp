package db

import (
	"context"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomWeightEntry(t *testing.T) Weightentry {
	user := CreateRandomUser(t)

	arg := CreateWeightEntryParams{
		Username:  user.Username,
		Notes:     pgtype.Text{String: util.GetRandomUsername(79), Valid: true},
		WeightKg:  pgtype.Float8{Float64: float64(util.GetRandomAmount(1, 200)), Valid: true},
		EntryDate: pgtype.Date{Time: time.Now().UTC(), Valid: true, InfinityModifier: pgtype.Finite},
	}

	entry, err := testStore.CreateWeightEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Username, entry.Username)
	require.Equal(t, arg.Notes, entry.Notes)
	require.Equal(t, arg.WeightKg, entry.WeightKg)

	require.NotZero(t, entry.WeightEntryID)
	return entry
}

func TestCreateWeightEntry(t *testing.T) {
	CreateRandomWeightEntry(t)
}

func TestGetWeightEntry(t *testing.T) {
	entry1 := CreateRandomWeightEntry(t)
	entry2, err := testStore.GetWeightEntry(context.Background(), entry1.WeightEntryID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.Username, entry2.Username)
	require.Equal(t, entry1.Notes, entry2.Notes)
	require.Equal(t, entry1.WeightKg, entry2.WeightKg)
}

func TestUpdateWeightEntry(t *testing.T) {
	entry1 := CreateRandomWeightEntry(t)

	arg := UpdateWeightEntryParams{
		WeightEntryID: entry1.WeightEntryID,
		Notes:         pgtype.Text{String: util.GetRandomUsername(79), Valid: true},
		WeightKg:      pgtype.Float8{Float64: float64(util.GetRandomAmount(1, 200)), Valid: true},
		EntryDate:     pgtype.Date{Time: time.Now().UTC(), Valid: true},
	}

	wkout2, err := testStore.UpdateWeightEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, arg.Notes, wkout2.Notes)
	require.Equal(t, arg.WeightKg, wkout2.WeightKg)
	require.Equal(t, arg.EntryDate.Time.Day(), wkout2.EntryDate.Time.Day())
}

func TestDeleteWeightEntry(t *testing.T) {
	entry1 := CreateRandomWeightEntry(t)
	err := testStore.DeleteWeightEntry(context.Background(), entry1.WeightEntryID)
	require.NoError(t, err)

	entry2, err := testStore.GetWeightEntry(context.Background(), entry1.WeightEntryID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, entry2)
}
