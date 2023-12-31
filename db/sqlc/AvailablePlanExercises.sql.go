// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: AvailablePlanExercises.sql

package db

import (
	"context"
)

const createAvailablePlanExercise = `-- name: CreateAvailablePlanExercise :one
INSERT INTO AvailablePlanExercises (
  exercise_name, 
  plan_id, 
  sets, 
  rest_duration,
  notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, plan_id, exercise_name, sets, rest_duration, notes
`

type CreateAvailablePlanExerciseParams struct {
	ExerciseName string `json:"exercise_name"`
	PlanID       int64  `json:"plan_id"`
	Sets         int32  `json:"sets"`
	RestDuration string `json:"rest_duration"`
	Notes        string `json:"notes"`
}

func (q *Queries) CreateAvailablePlanExercise(ctx context.Context, arg CreateAvailablePlanExerciseParams) (Availableplanexercise, error) {
	row := q.db.QueryRow(ctx, createAvailablePlanExercise,
		arg.ExerciseName,
		arg.PlanID,
		arg.Sets,
		arg.RestDuration,
		arg.Notes,
	)
	var i Availableplanexercise
	err := row.Scan(
		&i.ID,
		&i.PlanID,
		&i.ExerciseName,
		&i.Sets,
		&i.RestDuration,
		&i.Notes,
	)
	return i, err
}

const deleteAvailablePlanExercise = `-- name: DeleteAvailablePlanExercise :exec
DELETE FROM AvailablePlanExercises
WHERE id = $1
`

func (q *Queries) DeleteAvailablePlanExercise(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAvailablePlanExercise, id)
	return err
}

const getAvailablePlanExercise = `-- name: GetAvailablePlanExercise :one
SELECT id, plan_id, exercise_name, sets, rest_duration, notes
FROM AvailablePlanExercises
WHERE id = $1
`

func (q *Queries) GetAvailablePlanExercise(ctx context.Context, id int64) (Availableplanexercise, error) {
	row := q.db.QueryRow(ctx, getAvailablePlanExercise, id)
	var i Availableplanexercise
	err := row.Scan(
		&i.ID,
		&i.PlanID,
		&i.ExerciseName,
		&i.Sets,
		&i.RestDuration,
		&i.Notes,
	)
	return i, err
}

const listAllAvailablePlanExercises = `-- name: ListAllAvailablePlanExercises :many
SELECT id, plan_id, exercise_name, sets, rest_duration, notes
FROM AvailablePlanExercises
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2
`

type ListAllAvailablePlanExercisesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllAvailablePlanExercises(ctx context.Context, arg ListAllAvailablePlanExercisesParams) ([]Availableplanexercise, error) {
	rows, err := q.db.Query(ctx, listAllAvailablePlanExercises, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Availableplanexercise{}
	for rows.Next() {
		var i Availableplanexercise
		if err := rows.Scan(
			&i.ID,
			&i.PlanID,
			&i.ExerciseName,
			&i.Sets,
			&i.RestDuration,
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

const updateAvailablePlanExercise = `-- name: UpdateAvailablePlanExercise :one
UPDATE AvailablePlanExercises
SET notes = $2, sets = $3, rest_duration = $4
WHERE id = $1
RETURNING id, plan_id, exercise_name, sets, rest_duration, notes
`

type UpdateAvailablePlanExerciseParams struct {
	ID           int64  `json:"id"`
	Notes        string `json:"notes"`
	Sets         int32  `json:"sets"`
	RestDuration string `json:"rest_duration"`
}

func (q *Queries) UpdateAvailablePlanExercise(ctx context.Context, arg UpdateAvailablePlanExerciseParams) (Availableplanexercise, error) {
	row := q.db.QueryRow(ctx, updateAvailablePlanExercise,
		arg.ID,
		arg.Notes,
		arg.Sets,
		arg.RestDuration,
	)
	var i Availableplanexercise
	err := row.Scan(
		&i.ID,
		&i.PlanID,
		&i.ExerciseName,
		&i.Sets,
		&i.RestDuration,
		&i.Notes,
	)
	return i, err
}
