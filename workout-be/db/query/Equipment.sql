-- name: CreateEquipment :one
INSERT INTO Equipment (equipment_name, description, equipment_type)
VALUES ($1, $2, $3)
RETURNING equipment_id;

-- name: GetEquipment :one
SELECT equipment_id, equipment_name, description, equipment_type
FROM Equipment
WHERE equipment_id = $1;

-- name: DeleteEquipment :exec
DELETE FROM Equipment
WHERE equipment_id = $1;

-- name: UpdateEquipment :one
UPDATE Equipment
SET equipment_name = $2, description = $3, equipment_type = $4
WHERE equipment_id = $1
RETURNING equipment_id, equipment_name, description, equipment_type;

-- name: ListEquipments :many
SELECT equipment_id, equipment_name, equipment_type, description
FROM Equipment
ORDER BY equipment_type -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;

