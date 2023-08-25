// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: Workout.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO Workout (username, workout_date, workout_duration, notes)
VALUES ($1, $2, $3, $4)
RETURNING workout_id, username, workout_date, workout_duration, notes
`

type CreateWorkoutParams struct {
	Username        string          `json:"username"`
	WorkoutDate     pgtype.Date     `json:"workout_date"`
	WorkoutDuration pgtype.Interval `json:"workout_duration"`
	Notes           pgtype.Text     `json:"notes"`
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error) {
	row := q.db.QueryRow(ctx, createWorkout,
		arg.Username,
		arg.WorkoutDate,
		arg.WorkoutDuration,
		arg.Notes,
	)
	var i Workout
	err := row.Scan(
		&i.WorkoutID,
		&i.Username,
		&i.WorkoutDate,
		&i.WorkoutDuration,
		&i.Notes,
	)
	return i, err
}

const deleteWorkout = `-- name: DeleteWorkout :exec
DELETE FROM Workout
WHERE workout_id = $1
`

func (q *Queries) DeleteWorkout(ctx context.Context, workoutID int64) error {
	_, err := q.db.Exec(ctx, deleteWorkout, workoutID)
	return err
}

const getWorkout = `-- name: GetWorkout :one
SELECT workout_id, username, workout_date, workout_duration, notes
FROM Workout
WHERE workout_id = $1
`

func (q *Queries) GetWorkout(ctx context.Context, workoutID int64) (Workout, error) {
	row := q.db.QueryRow(ctx, getWorkout, workoutID)
	var i Workout
	err := row.Scan(
		&i.WorkoutID,
		&i.Username,
		&i.WorkoutDate,
		&i.WorkoutDuration,
		&i.Notes,
	)
	return i, err
}

const listWorkouts = `-- name: ListWorkouts :many
SELECT workout_id, username, workout_date, workout_duration, notes
FROM Workout
ORDER BY workout_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2
`

type ListWorkoutsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListWorkouts(ctx context.Context, arg ListWorkoutsParams) ([]Workout, error) {
	rows, err := q.db.Query(ctx, listWorkouts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Workout{}
	for rows.Next() {
		var i Workout
		if err := rows.Scan(
			&i.WorkoutID,
			&i.Username,
			&i.WorkoutDate,
			&i.WorkoutDuration,
			&i.Notes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateWorkout = `-- name: UpdateWorkout :one
UPDATE Workout
SET username = $2, workout_date = $3, workout_duration = $4, notes = $5
WHERE workout_id = $1
RETURNING workout_id, username, workout_date, workout_duration, notes
`

type UpdateWorkoutParams struct {
	WorkoutID       int64           `json:"workout_id"`
	Username        string          `json:"username"`
	WorkoutDate     pgtype.Date     `json:"workout_date"`
	WorkoutDuration pgtype.Interval `json:"workout_duration"`
	Notes           pgtype.Text     `json:"notes"`
}

func (q *Queries) UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (Workout, error) {
	row := q.db.QueryRow(ctx, updateWorkout,
		arg.WorkoutID,
		arg.Username,
		arg.WorkoutDate,
		arg.WorkoutDuration,
		arg.Notes,
	)
	var i Workout
	err := row.Scan(
		&i.WorkoutID,
		&i.Username,
		&i.WorkoutDate,
		&i.WorkoutDuration,
		&i.Notes,
	)
	return i, err
}
