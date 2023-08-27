-- name: CreateExercise :one
INSERT INTO Exercise (exercise_name, description, equipment_required, muscle_group_name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetExercise :one
SELECT *
FROM Exercise
WHERE exercise_name = $1;

-- name: DeleteExercise :exec
DELETE FROM Exercise
WHERE exercise_name = $1;

-- name: UpdateExercise :one
UPDATE Exercise
SET description = $2, equipment_required = $3, muscle_group_name = $4
WHERE exercise_name = $1
RETURNING *;

-- name: ListEquipmentExercise :many
SELECT *
FROM Exercise
WHERE equipment_required = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;

-- name: ListMuscleGroupExercise :many
SELECT *
FROM Exercise
WHERE muscle_group_name = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;

-- name: ListAllExercise :many
SELECT *
FROM Exercise
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
