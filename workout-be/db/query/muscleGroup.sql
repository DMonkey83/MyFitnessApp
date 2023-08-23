-- name: CreateMuscleGroup :one
INSERT INTO MuscleGroup (muscle_group_name)
VALUES ($1)
RETURNING muscle_group_id;

-- name: GetMuscleGroup :one
SELECT muscle_group_id, muscle_group_name
FROM MuscleGroup
WHERE muscle_group_id = $1;

-- name: DeleteMuscleGroup :exec
DELETE FROM MuscleGroup
WHERE muscle_group_id = $1;

-- name: UpdateMuscleGroup :one
UPDATE MuscleGroup
SET muscle_group_name = $2
WHERE muscle_group_id = $1
RETURNING muscle_group_id, muscle_group_name;

-- name: ListMuscleGroups :many
SELECT muscle_group_id, muscle_group_name
FROM MuscleGroup
ORDER BY muscle_group_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;

