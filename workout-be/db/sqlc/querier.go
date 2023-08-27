// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreateAvailablePlan(ctx context.Context, arg CreateAvailablePlanParams) (Availableworkoutplan, error)
	CreateAvailablePlanExercise(ctx context.Context, arg CreateAvailablePlanExerciseParams) (Availableplanexercise, error)
	CreateExercise(ctx context.Context, arg CreateExerciseParams) (Exercise, error)
	CreateExerciseLog(ctx context.Context, arg CreateExerciseLogParams) (Exerciselog, error)
	CreateExerciseSet(ctx context.Context, arg CreateExerciseSetParams) (Exerciseset, error)
	CreateMaxRepGoal(ctx context.Context, arg CreateMaxRepGoalParams) (Maxrepgoal, error)
	CreateMaxWeightGoal(ctx context.Context, arg CreateMaxWeightGoalParams) (Maxweightgoal, error)
	CreateOneOffWorkoutExercise(ctx context.Context, arg CreateOneOffWorkoutExerciseParams) (Oneoffworkoutexercise, error)
	CreatePlan(ctx context.Context, arg CreatePlanParams) (Workoutplan, error)
	CreateSet(ctx context.Context, arg CreateSetParams) (Set, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (Userprofile, error)
	CreateWeightEntry(ctx context.Context, arg CreateWeightEntryParams) (Weightentry, error)
	CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error)
	CreateWorkoutLog(ctx context.Context, arg CreateWorkoutLogParams) (Workoutlog, error)
	DeleteAvailablePlan(ctx context.Context, planID int64) error
	DeleteAvailablePlanExercise(ctx context.Context, id int64) error
	DeleteExercise(ctx context.Context, exerciseName string) error
	DeleteExerciseLog(ctx context.Context, exerciseLogID int64) error
	DeleteExerciseSet(ctx context.Context, setID int64) error
	DeleteMaxRepGoal(ctx context.Context, arg DeleteMaxRepGoalParams) error
	DeleteMaxWeightGoal(ctx context.Context, arg DeleteMaxWeightGoalParams) error
	DeleteOneOffWorkoutExercise(ctx context.Context, id int32) error
	DeletePlan(ctx context.Context, arg DeletePlanParams) error
	DeleteSet(ctx context.Context, setID int64) error
	DeleteUser(ctx context.Context, username string) error
	DeleteUserProfile(ctx context.Context, username string) error
	DeleteWeightEntry(ctx context.Context, weightEntryID int64) error
	DeleteWorkout(ctx context.Context, workoutID int64) error
	DeleteWorkoutLog(ctx context.Context, logID int64) error
	GetAvailablePlan(ctx context.Context, planID int64) (Availableworkoutplan, error)
	GetAvailablePlanExercise(ctx context.Context, id int64) (Availableplanexercise, error)
	GetExercise(ctx context.Context, exerciseName string) (Exercise, error)
	GetExerciseLog(ctx context.Context, exerciseLogID int64) (Exerciselog, error)
	GetExerciseSet(ctx context.Context, setID int64) (Exerciseset, error)
	GetMaxRepGoal(ctx context.Context, arg GetMaxRepGoalParams) (Maxrepgoal, error)
	GetMaxWeightGoal(ctx context.Context, arg GetMaxWeightGoalParams) (Maxweightgoal, error)
	GetOneOffWorkoutExercise(ctx context.Context, id int32) (Oneoffworkoutexercise, error)
	GetPlan(ctx context.Context, arg GetPlanParams) (Workoutplan, error)
	GetSet(ctx context.Context, setID int64) (Set, error)
	GetUser(ctx context.Context, username string) (GetUserRow, error)
	GetUserProfile(ctx context.Context, username string) (GetUserProfileRow, error)
	GetWeightEntry(ctx context.Context, weightEntryID int64) (Weightentry, error)
	GetWorkout(ctx context.Context, workoutID int64) (Workout, error)
	GetWorkoutLog(ctx context.Context, logID int64) (Workoutlog, error)
	ListAllAvailablePlanExercises(ctx context.Context, arg ListAllAvailablePlanExercisesParams) ([]Availableplanexercise, error)
	ListAllAvailablePlans(ctx context.Context, arg ListAllAvailablePlansParams) ([]Availableworkoutplan, error)
	ListAllExercise(ctx context.Context, arg ListAllExerciseParams) ([]Exercise, error)
	ListAllOneOffWorkoutExercises(ctx context.Context, arg ListAllOneOffWorkoutExercisesParams) ([]Availableplanexercise, error)
	ListAvailablePlansByCreator(ctx context.Context, arg ListAvailablePlansByCreatorParams) ([]Availableworkoutplan, error)
	ListEquipmentExercise(ctx context.Context, arg ListEquipmentExerciseParams) ([]Exercise, error)
	ListExerciseLog(ctx context.Context, arg ListExerciseLogParams) ([]Exerciselog, error)
	ListExerciseSets(ctx context.Context, arg ListExerciseSetsParams) ([]Exerciseset, error)
	ListMaxRepGoals(ctx context.Context, arg ListMaxRepGoalsParams) ([]Maxrepgoal, error)
	ListMaxWeightGoals(ctx context.Context, arg ListMaxWeightGoalsParams) ([]Maxweightgoal, error)
	ListMuscleGroupExercise(ctx context.Context, arg ListMuscleGroupExerciseParams) ([]Exercise, error)
	ListSets(ctx context.Context, arg ListSetsParams) ([]Set, error)
	ListWeightEntries(ctx context.Context, arg ListWeightEntriesParams) ([]Weightentry, error)
	ListWorkoutLogs(ctx context.Context, arg ListWorkoutLogsParams) ([]Workoutlog, error)
	ListWorkouts(ctx context.Context, arg ListWorkoutsParams) ([]Workout, error)
	UpdateAvailablePlan(ctx context.Context, arg UpdateAvailablePlanParams) (Availableworkoutplan, error)
	UpdateAvailablePlanExercise(ctx context.Context, arg UpdateAvailablePlanExerciseParams) (Availableplanexercise, error)
	UpdateExercise(ctx context.Context, arg UpdateExerciseParams) (Exercise, error)
	UpdateExerciseLog(ctx context.Context, arg UpdateExerciseLogParams) (Exerciselog, error)
	UpdateExerciseSet(ctx context.Context, arg UpdateExerciseSetParams) (Exerciseset, error)
	UpdateMaxRepGoal(ctx context.Context, arg UpdateMaxRepGoalParams) (Maxrepgoal, error)
	UpdateMaxWeightGoal(ctx context.Context, arg UpdateMaxWeightGoalParams) (Maxweightgoal, error)
	UpdateOneOffWorkoutExercise(ctx context.Context, arg UpdateOneOffWorkoutExerciseParams) (Oneoffworkoutexercise, error)
	UpdatePlan(ctx context.Context, arg UpdatePlanParams) (Workoutplan, error)
	UpdateSet(ctx context.Context, arg UpdateSetParams) (Set, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (Userprofile, error)
	UpdateWeightEntry(ctx context.Context, arg UpdateWeightEntryParams) (Weightentry, error)
	UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (Workout, error)
	UpdateWorkoutLog(ctx context.Context, arg UpdateWorkoutLogParams) (Workoutlog, error)
}

var _ Querier = (*Queries)(nil)
