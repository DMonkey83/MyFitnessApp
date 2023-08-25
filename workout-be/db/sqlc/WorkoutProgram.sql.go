// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: WorkoutProgram.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createWorkoutprogram = `-- name: CreateWorkoutprogram :one
INSERT INTO WorkoutProgram (username, program_name, description)
VALUES ($1, $2, $3)
RETURNING program_id, username, program_name, description
`

type CreateWorkoutprogramParams struct {
	Username    string      `json:"username"`
	ProgramName string      `json:"program_name"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) CreateWorkoutprogram(ctx context.Context, arg CreateWorkoutprogramParams) (Workoutprogram, error) {
	row := q.db.QueryRow(ctx, createWorkoutprogram, arg.Username, arg.ProgramName, arg.Description)
	var i Workoutprogram
	err := row.Scan(
		&i.ProgramID,
		&i.Username,
		&i.ProgramName,
		&i.Description,
	)
	return i, err
}

const deleteWorkoutprogram = `-- name: DeleteWorkoutprogram :exec
DELETE FROM WorkoutProgram
WHERE program_id = $1
`

func (q *Queries) DeleteWorkoutprogram(ctx context.Context, programID int64) error {
	_, err := q.db.Exec(ctx, deleteWorkoutprogram, programID)
	return err
}

const getWorkoutprogram = `-- name: GetWorkoutprogram :one
SELECT program_id, username, program_name, description
FROM WorkoutProgram
WHERE program_id = $1
`

func (q *Queries) GetWorkoutprogram(ctx context.Context, programID int64) (Workoutprogram, error) {
	row := q.db.QueryRow(ctx, getWorkoutprogram, programID)
	var i Workoutprogram
	err := row.Scan(
		&i.ProgramID,
		&i.Username,
		&i.ProgramName,
		&i.Description,
	)
	return i, err
}

const listWorkoutprograms = `-- name: ListWorkoutprograms :many
SELECT program_id, username, program_name, description
FROM WorkoutProgram
ORDER BY program_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2
`

type ListWorkoutprogramsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListWorkoutprograms(ctx context.Context, arg ListWorkoutprogramsParams) ([]Workoutprogram, error) {
	rows, err := q.db.Query(ctx, listWorkoutprograms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Workoutprogram{}
	for rows.Next() {
		var i Workoutprogram
		if err := rows.Scan(
			&i.ProgramID,
			&i.Username,
			&i.ProgramName,
			&i.Description,
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

const updateWorkoutprogram = `-- name: UpdateWorkoutprogram :one
UPDATE WorkoutProgram
SET program_name = $2, description = $3
WHERE program_id = $1
RETURNING program_id, username, program_name, description
`

type UpdateWorkoutprogramParams struct {
	ProgramID   int64       `json:"program_id"`
	ProgramName string      `json:"program_name"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) UpdateWorkoutprogram(ctx context.Context, arg UpdateWorkoutprogramParams) (Workoutprogram, error) {
	row := q.db.QueryRow(ctx, updateWorkoutprogram, arg.ProgramID, arg.ProgramName, arg.Description)
	var i Workoutprogram
	err := row.Scan(
		&i.ProgramID,
		&i.Username,
		&i.ProgramName,
		&i.Description,
	)
	return i, err
}
