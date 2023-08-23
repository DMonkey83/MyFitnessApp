// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: Exercise.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExercise = `-- name: CreateExercise :one
INSERT INTO Exercise (exercise_name,muscle_group, description, equipment_id)
VALUES ($1, $2, $3, $4)
RETURNING exercise_id
`

type CreateExerciseParams struct {
	ExerciseName string          `json:"exercise_name"`
	MuscleGroup  MuscleGroupEnum `json:"muscle_group"`
	Description  pgtype.Text     `json:"description"`
	EquipmentID  pgtype.Int8     `json:"equipment_id"`
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) (int64, error) {
	row := q.db.QueryRow(ctx, createExercise,
		arg.ExerciseName,
		arg.MuscleGroup,
		arg.Description,
		arg.EquipmentID,
	)
	var exercise_id int64
	err := row.Scan(&exercise_id)
	return exercise_id, err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM Exercise
WHERE exercise_id = $1
`

func (q *Queries) DeleteExercise(ctx context.Context, exerciseID int64) error {
	_, err := q.db.Exec(ctx, deleteExercise, exerciseID)
	return err
}

const getExercise = `-- name: GetExercise :one
SELECT exercise_id, exercise_name, muscle_group, description, equipment_id
FROM Exercise
WHERE exercise_id = $1
`

func (q *Queries) GetExercise(ctx context.Context, exerciseID int64) (Exercise, error) {
	row := q.db.QueryRow(ctx, getExercise, exerciseID)
	var i Exercise
	err := row.Scan(
		&i.ExerciseID,
		&i.ExerciseName,
		&i.MuscleGroup,
		&i.Description,
		&i.EquipmentID,
	)
	return i, err
}

const listExercise = `-- name: ListExercise :many
SELECT exercise_id, exercise_name, description
FROM Exercise
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2
`

type ListExerciseParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListExerciseRow struct {
	ExerciseID   int64       `json:"exercise_id"`
	ExerciseName string      `json:"exercise_name"`
	Description  pgtype.Text `json:"description"`
}

func (q *Queries) ListExercise(ctx context.Context, arg ListExerciseParams) ([]ListExerciseRow, error) {
	rows, err := q.db.Query(ctx, listExercise, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListExerciseRow{}
	for rows.Next() {
		var i ListExerciseRow
		if err := rows.Scan(&i.ExerciseID, &i.ExerciseName, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExercise = `-- name: UpdateExercise :one
UPDATE Exercise
SET exercise_name = $2, muscle_group = $3, description = $4, equipment_id = $5
WHERE exercise_id = $1
RETURNING exercise_id, exercise_name, muscle_group, description, equipment_id
`

type UpdateExerciseParams struct {
	ExerciseID   int64           `json:"exercise_id"`
	ExerciseName string          `json:"exercise_name"`
	MuscleGroup  MuscleGroupEnum `json:"muscle_group"`
	Description  pgtype.Text     `json:"description"`
	EquipmentID  pgtype.Int8     `json:"equipment_id"`
}

func (q *Queries) UpdateExercise(ctx context.Context, arg UpdateExerciseParams) (Exercise, error) {
	row := q.db.QueryRow(ctx, updateExercise,
		arg.ExerciseID,
		arg.ExerciseName,
		arg.MuscleGroup,
		arg.Description,
		arg.EquipmentID,
	)
	var i Exercise
	err := row.Scan(
		&i.ExerciseID,
		&i.ExerciseName,
		&i.MuscleGroup,
		&i.Description,
		&i.EquipmentID,
	)
	return i, err
}
