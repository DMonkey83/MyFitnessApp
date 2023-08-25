package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomWorkoutProgram(t *testing.T) Workoutprogram {
	user := CreateRandomUser(t)

	arg := CreateWorkoutprogramParams{
		Username:    user.Username,
		ProgramName: util.GetRandomUsername(10),
		Description: pgtype.Text{String: util.GetRandomUsername(100), Valid: true},
	}

	workoutP, err := testStore.CreateWorkoutprogram(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, workoutP)

	require.Equal(t, arg.Username, workoutP.Username)
	require.Equal(t, arg.Description, workoutP.Description)
	require.Equal(t, arg.ProgramName, workoutP.ProgramName)

	require.NotZero(t, workoutP.ProgramID)
	return workoutP
}

func TestCreateWorkoutProgram(t *testing.T) {
	CreateRandomWorkoutProgram(t)
}

func TestGetWorkoutProgram(t *testing.T) {
	wkoutp1 := CreateRandomWorkoutProgram(t)
	wkoutp2, err := testStore.GetWorkoutprogram(context.Background(), wkoutp1.ProgramID)
	require.NoError(t, err)
	require.NotEmpty(t, wkoutp2)

	require.Equal(t, wkoutp1.Username, wkoutp2.Username)
	require.Equal(t, wkoutp1.Description, wkoutp2.Description)
	require.Equal(t, wkoutp1.ProgramName, wkoutp2.ProgramName)
}

func TestUpdateWorkoutProgram(t *testing.T) {
	wkout1 := CreateRandomWorkoutProgram(t)

	arg := UpdateWorkoutprogramParams{
		ProgramID:   wkout1.ProgramID,
		ProgramName: util.GetRandomUsername(10),
		Description: pgtype.Text{String: util.GetRandomUsername(100), Valid: true},
	}

	wkout2, err := testStore.UpdateWorkoutprogram(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, arg.Description, wkout2.Description)
	require.Equal(t, arg.ProgramName, wkout2.ProgramName)
}

func TestDeleteWorkoutProgram(t *testing.T) {
	wkoutp1 := CreateRandomWorkoutProgram(t)
	err := testStore.DeleteWorkoutprogram(context.Background(), wkoutp1.ProgramID)
	require.NoError(t, err)

	wkoutp2, err := testStore.GetWorkoutprogram(context.Background(), wkoutp1.ProgramID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, wkoutp2)
}
