// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: WeightEntry.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createWeightEntry = `-- name: CreateWeightEntry :one
INSERT INTO WeightEntry (username, entry_date, weight_kg, weight_lb, notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING weight_entry_id, username, entry_date, weight_kg, weight_lb, notes
`

type CreateWeightEntryParams struct {
	Username  string        `json:"username"`
	EntryDate pgtype.Date   `json:"entry_date"`
	WeightKg  pgtype.Float8 `json:"weight_kg"`
	WeightLb  pgtype.Float8 `json:"weight_lb"`
	Notes     pgtype.Text   `json:"notes"`
}

func (q *Queries) CreateWeightEntry(ctx context.Context, arg CreateWeightEntryParams) (Weightentry, error) {
	row := q.db.QueryRow(ctx, createWeightEntry,
		arg.Username,
		arg.EntryDate,
		arg.WeightKg,
		arg.WeightLb,
		arg.Notes,
	)
	var i Weightentry
	err := row.Scan(
		&i.WeightEntryID,
		&i.Username,
		&i.EntryDate,
		&i.WeightKg,
		&i.WeightLb,
		&i.Notes,
	)
	return i, err
}

const deleteWeightEntry = `-- name: DeleteWeightEntry :exec
DELETE FROM WeightEntry
WHERE weight_entry_id = $1
`

func (q *Queries) DeleteWeightEntry(ctx context.Context, weightEntryID int64) error {
	_, err := q.db.Exec(ctx, deleteWeightEntry, weightEntryID)
	return err
}

const getWeightEntry = `-- name: GetWeightEntry :one
SELECT weight_entry_id, username, entry_date, weight_kg, weight_lb, notes
FROM WeightEntry
WHERE weight_entry_id = $1
`

func (q *Queries) GetWeightEntry(ctx context.Context, weightEntryID int64) (Weightentry, error) {
	row := q.db.QueryRow(ctx, getWeightEntry, weightEntryID)
	var i Weightentry
	err := row.Scan(
		&i.WeightEntryID,
		&i.Username,
		&i.EntryDate,
		&i.WeightKg,
		&i.WeightLb,
		&i.Notes,
	)
	return i, err
}

const listWeightEntries = `-- name: ListWeightEntries :many
SELECT weight_entry_id, username, entry_date, weight_kg, weight_lb, notes
FROM WeightEntry
WHERE username = $1
ORDER BY weight_entry_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListWeightEntriesParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListWeightEntries(ctx context.Context, arg ListWeightEntriesParams) ([]Weightentry, error) {
	rows, err := q.db.Query(ctx, listWeightEntries, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Weightentry{}
	for rows.Next() {
		var i Weightentry
		if err := rows.Scan(
			&i.WeightEntryID,
			&i.Username,
			&i.EntryDate,
			&i.WeightKg,
			&i.WeightLb,
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

const updateWeightEntry = `-- name: UpdateWeightEntry :one
UPDATE WeightEntry
SET entry_date = $2, weight_kg = $3, weight_lb = $4, notes = $5
WHERE weight_entry_id = $1
RETURNING weight_entry_id, username, entry_date, weight_kg, weight_lb, notes
`

type UpdateWeightEntryParams struct {
	WeightEntryID int64         `json:"weight_entry_id"`
	EntryDate     pgtype.Date   `json:"entry_date"`
	WeightKg      pgtype.Float8 `json:"weight_kg"`
	WeightLb      pgtype.Float8 `json:"weight_lb"`
	Notes         pgtype.Text   `json:"notes"`
}

func (q *Queries) UpdateWeightEntry(ctx context.Context, arg UpdateWeightEntryParams) (Weightentry, error) {
	row := q.db.QueryRow(ctx, updateWeightEntry,
		arg.WeightEntryID,
		arg.EntryDate,
		arg.WeightKg,
		arg.WeightLb,
		arg.Notes,
	)
	var i Weightentry
	err := row.Scan(
		&i.WeightEntryID,
		&i.Username,
		&i.EntryDate,
		&i.WeightKg,
		&i.WeightLb,
		&i.Notes,
	)
	return i, err
}
