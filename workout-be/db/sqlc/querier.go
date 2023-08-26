// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreateEquipment(ctx context.Context, arg CreateEquipmentParams) (Equipment, error)
	CreateExercise(ctx context.Context, arg CreateExerciseParams) (Exercise, error)
	CreateMaxRepGoal(ctx context.Context, arg CreateMaxRepGoalParams) (Maxrepgoal, error)
	CreateMaxWeightGoal(ctx context.Context, arg CreateMaxWeightGoalParams) (Maxweightgoal, error)
	CreateMuscleGroup(ctx context.Context, arg CreateMuscleGroupParams) (Musclegroup, error)
	CreateRep(ctx context.Context, arg CreateRepParams) (Rep, error)
	CreateSet(ctx context.Context, arg CreateSetParams) (Set, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (Userprofile, error)
	CreateWeightEntry(ctx context.Context, arg CreateWeightEntryParams) (Weightentry, error)
	CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error)
	CreateWorkoutprogram(ctx context.Context, arg CreateWorkoutprogramParams) (Workoutprogram, error)
	DeleteEquipment(ctx context.Context, equipmentID int64) error
	DeleteExercise(ctx context.Context, exerciseID int64) error
	DeleteMaxRepGoal(ctx context.Context, goalID int64) error
	DeleteMaxWeightGoal(ctx context.Context, goalID int64) error
	DeleteMuscleGroup(ctx context.Context, muscleGroupID int64) error
	DeleteRep(ctx context.Context, repID int64) error
	DeleteSet(ctx context.Context, setID int64) error
	DeleteUser(ctx context.Context, username string) error
	DeleteUserProfile(ctx context.Context, username string) error
	DeleteWeightEntry(ctx context.Context, weightEntryID int64) error
	DeleteWorkout(ctx context.Context, workoutID int64) error
	DeleteWorkoutprogram(ctx context.Context, programID int64) error
	GetEquipment(ctx context.Context, equipmentID int64) (Equipment, error)
	GetExercise(ctx context.Context, exerciseID int64) (Exercise, error)
	GetMaxRepGoal(ctx context.Context, goalID int64) (Maxrepgoal, error)
	GetMaxWeightGoal(ctx context.Context, goalID int64) (Maxweightgoal, error)
	GetMuscleGroup(ctx context.Context, muscleGroupID int64) (Musclegroup, error)
	GetRep(ctx context.Context, repID int64) (Rep, error)
	GetSet(ctx context.Context, setID int64) (Set, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserProfile(ctx context.Context, username string) (Userprofile, error)
	GetWeightEntry(ctx context.Context, weightEntryID int64) (Weightentry, error)
	GetWorkout(ctx context.Context, workoutID int64) (Workout, error)
	GetWorkoutprogram(ctx context.Context, programID int64) (Workoutprogram, error)
	ListAllExercise(ctx context.Context, arg ListAllExerciseParams) ([]Exercise, error)
	ListAllWorkoutprograms(ctx context.Context, arg ListAllWorkoutprogramsParams) ([]Workoutprogram, error)
	ListEquipments(ctx context.Context, arg ListEquipmentsParams) ([]Equipment, error)
	ListMaxRepGoals(ctx context.Context, arg ListMaxRepGoalsParams) ([]Maxrepgoal, error)
	ListMaxWeightGoals(ctx context.Context, arg ListMaxWeightGoalsParams) ([]Maxweightgoal, error)
	ListMuscleGroups(ctx context.Context, arg ListMuscleGroupsParams) ([]Musclegroup, error)
	ListReps(ctx context.Context, arg ListRepsParams) ([]Rep, error)
	ListSets(ctx context.Context, arg ListSetsParams) ([]Set, error)
	ListWeightEntries(ctx context.Context, arg ListWeightEntriesParams) ([]Weightentry, error)
	ListWorkoutExercise(ctx context.Context, arg ListWorkoutExerciseParams) ([]Exercise, error)
	ListWorkoutprogramsForUser(ctx context.Context, arg ListWorkoutprogramsForUserParams) ([]Workoutprogram, error)
	ListWorkouts(ctx context.Context, arg ListWorkoutsParams) ([]Workout, error)
	UpdateEquipment(ctx context.Context, arg UpdateEquipmentParams) (Equipment, error)
	UpdateExercise(ctx context.Context, arg UpdateExerciseParams) (Exercise, error)
	UpdateMaxRepGoal(ctx context.Context, arg UpdateMaxRepGoalParams) (Maxrepgoal, error)
	UpdateMaxWeightGoal(ctx context.Context, arg UpdateMaxWeightGoalParams) (Maxweightgoal, error)
	UpdateMuscleGroup(ctx context.Context, arg UpdateMuscleGroupParams) (Musclegroup, error)
	UpdateRep(ctx context.Context, arg UpdateRepParams) (Rep, error)
	UpdateSet(ctx context.Context, arg UpdateSetParams) (Set, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (Userprofile, error)
	UpdateWeightEntry(ctx context.Context, arg UpdateWeightEntryParams) (Weightentry, error)
	UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (Workout, error)
	UpdateWorkoutprogram(ctx context.Context, arg UpdateWorkoutprogramParams) (Workoutprogram, error)
}

var _ Querier = (*Queries)(nil)
