-- name: CreateEquipment :one
INSERT INTO Equipment (equipment_name, description, equipment_type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEquipment :one
SELECT *
FROM Equipment
WHERE equipment_name = $1;

-- name: DeleteEquipment :exec
DELETE FROM Equipment
WHERE equipment_name = $1;

-- name: UpdateEquipment :one
UPDATE Equipment
SET description = $2, equipment_type = $3
WHERE equipment_name = $1
RETURNING *;

-- name: ListExpuipments :many
SELECT *
FROM Equipment
ORDER BY equipment_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;

