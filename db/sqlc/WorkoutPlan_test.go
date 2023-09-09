package db

import (
	"context"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomWorkoutPlan(t *testing.T) Workoutplan {
	user := CreateRandomUser(t)

	arg := CreatePlanParams{
		Username:    user.Username,
		PlanName:    util.GetRandomUsername(8),
		Description: util.GetRandomUsername(60),
		StartDate:   time.Now(),
		Goal:        WorkoutgoalenumLoseWeight,
		IsPublic:    VisibilityPrivate,
		Difficulty:  DifficultyIntermediate,
	}

	plan, err := testStore.CreatePlan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plan)

	require.Equal(t, arg.Username, plan.Username)
	require.Equal(t, arg.PlanName, plan.PlanName)
	require.Equal(t, arg.Description, plan.Description)
	require.Equal(t, arg.StartDate.Minute(), plan.StartDate.Minute())
	require.Equal(t, arg.Goal, plan.Goal)
	require.Equal(t, arg.Difficulty, plan.Difficulty)

	require.NotZero(t, plan.PlanID)
	return plan
}

func TestWorkoutPlan(t *testing.T) {
	CreateRandomWorkoutPlan(t)
}

func TestGetWorkoutPlan(t *testing.T) {
	plan1 := CreateRandomWorkoutPlan(t)
	arg := GetPlanParams{
		PlanID:   plan1.PlanID,
		Username: plan1.Username,
	}
	plan2, err := testStore.GetPlan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plan2)

	require.Equal(t, plan1.Username, plan2.Username)
	require.Equal(t, plan1.Description, plan2.Description)
	require.Equal(t, plan1.StartDate, plan2.StartDate)
	require.Equal(t, plan1.Goal, plan2.Goal)
	require.Equal(t, plan1.Difficulty, plan2.Difficulty)
	require.Equal(t, plan1.IsPublic, plan2.IsPublic)
}

func TestUpdateWorkoutPlan(t *testing.T) {
	plan1 := CreateRandomWorkoutPlan(t)

	arg := UpdatePlanParams{
		PlanID:      plan1.PlanID,
		Username:    plan1.Username,
		PlanName:    util.GetRandomUsername(6),
		Description: util.GetRandomUsername(49),
		Difficulty:  DifficultyBeginner,
		Goal:        WorkoutgoalenumBuildMuscle,
		StartDate:   plan1.StartDate,
		EndDate:     time.Now().Add(time.Duration(time.Minute)),
		IsPublic:    VisibilityPublic,
	}

	plan2, err := testStore.UpdatePlan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plan2)

	require.Equal(t, plan1.Username, plan2.Username)
	require.Equal(t, plan1.StartDate.Minute(), plan2.StartDate.Minute())
	require.Equal(t, arg.EndDate.Minute(), plan2.EndDate.Minute())
	require.Equal(t, arg.PlanID, plan2.PlanID)
	require.Equal(t, arg.Description, plan2.Description)
	require.Equal(t, arg.Goal, plan2.Goal)
	require.Equal(t, arg.Difficulty, plan2.Difficulty)
	require.Equal(t, arg.IsPublic, plan2.IsPublic)
}

func TestDeleteWorkoutPlan(t *testing.T) {
	plan1 := CreateRandomWorkoutPlan(t)
	arg1 := DeletePlanParams{
		PlanID:   plan1.PlanID,
		Username: plan1.Username,
	}
	err := testStore.DeletePlan(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := GetPlanParams{
		PlanID:   plan1.PlanID,
		Username: plan1.Username,
	}

	plan2, err := testStore.GetPlan(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, plan2)
}
