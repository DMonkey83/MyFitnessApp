-- name: CreateMuscleGroup :one
INSERT INTO MuscleGroup (muscle_group_name)
VALUES ($1)
RETURNING *;

-- name: GetMuscleGroup :one
SELECT *
FROM MuscleGroup
WHERE muscle_group_id = $1;

-- name: DeleteMuscleGroup :exec
DELETE FROM MuscleGroup
WHERE muscle_group_id = $1;

-- name: UpdateMuscleGroup :one
UPDATE MuscleGroup
SET muscle_group_name = $2
WHERE muscle_group_id = $1
RETURNING *;

-- name: ListMuscleGroups :many
SELECT *
FROM MuscleGroup
ORDER BY muscle_group_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;

