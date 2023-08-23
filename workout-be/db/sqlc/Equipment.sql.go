// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: Equipment.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEquipment = `-- name: CreateEquipment :one
INSERT INTO Equipment (equipment_name, description, equipment_type)
VALUES ($1, $2, $3)
RETURNING equipment_id
`

type CreateEquipmentParams struct {
	EquipmentName string        `json:"equipment_name"`
	Description   pgtype.Text   `json:"description"`
	EquipmentType Equipmenttype `json:"equipment_type"`
}

func (q *Queries) CreateEquipment(ctx context.Context, arg CreateEquipmentParams) (int64, error) {
	row := q.db.QueryRow(ctx, createEquipment, arg.EquipmentName, arg.Description, arg.EquipmentType)
	var equipment_id int64
	err := row.Scan(&equipment_id)
	return equipment_id, err
}

const deleteEquipment = `-- name: DeleteEquipment :exec
DELETE FROM Equipment
WHERE equipment_id = $1
`

func (q *Queries) DeleteEquipment(ctx context.Context, equipmentID int64) error {
	_, err := q.db.Exec(ctx, deleteEquipment, equipmentID)
	return err
}

const getEquipment = `-- name: GetEquipment :one
SELECT equipment_id, equipment_name, description, equipment_type
FROM Equipment
WHERE equipment_id = $1
`

func (q *Queries) GetEquipment(ctx context.Context, equipmentID int64) (Equipment, error) {
	row := q.db.QueryRow(ctx, getEquipment, equipmentID)
	var i Equipment
	err := row.Scan(
		&i.EquipmentID,
		&i.EquipmentName,
		&i.Description,
		&i.EquipmentType,
	)
	return i, err
}

const listEquipments = `-- name: ListEquipments :many
SELECT equipment_id, equipment_name, equipment_type, description
FROM Equipment
ORDER BY equipment_type -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2
`

type ListEquipmentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListEquipmentsRow struct {
	EquipmentID   int64         `json:"equipment_id"`
	EquipmentName string        `json:"equipment_name"`
	EquipmentType Equipmenttype `json:"equipment_type"`
	Description   pgtype.Text   `json:"description"`
}

func (q *Queries) ListEquipments(ctx context.Context, arg ListEquipmentsParams) ([]ListEquipmentsRow, error) {
	rows, err := q.db.Query(ctx, listEquipments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListEquipmentsRow{}
	for rows.Next() {
		var i ListEquipmentsRow
		if err := rows.Scan(
			&i.EquipmentID,
			&i.EquipmentName,
			&i.EquipmentType,
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

const updateEquipment = `-- name: UpdateEquipment :one
UPDATE Equipment
SET equipment_name = $2, description = $3, equipment_type = $4
WHERE equipment_id = $1
RETURNING equipment_id, equipment_name, description, equipment_type
`

type UpdateEquipmentParams struct {
	EquipmentID   int64         `json:"equipment_id"`
	EquipmentName string        `json:"equipment_name"`
	Description   pgtype.Text   `json:"description"`
	EquipmentType Equipmenttype `json:"equipment_type"`
}

func (q *Queries) UpdateEquipment(ctx context.Context, arg UpdateEquipmentParams) (Equipment, error) {
	row := q.db.QueryRow(ctx, updateEquipment,
		arg.EquipmentID,
		arg.EquipmentName,
		arg.Description,
		arg.EquipmentType,
	)
	var i Equipment
	err := row.Scan(
		&i.EquipmentID,
		&i.EquipmentName,
		&i.Description,
		&i.EquipmentType,
	)
	return i, err
}
