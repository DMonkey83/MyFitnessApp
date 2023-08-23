-- name: CreateExercise :one
INSERT INTO Exercise (exercise_name,muscle_group, description, equipment_id)
VALUES ($1, $2, $3, $4)
RETURNING exercise_id;

-- name: GetExercise :one
SELECT exercise_id, exercise_name, muscle_group, description, equipment_id
FROM Exercise
WHERE exercise_id = $1;

-- name: DeleteExercise :exec
DELETE FROM Exercise
WHERE exercise_id = $1;

-- name: UpdateExercise :one
UPDATE Exercise
SET exercise_name = $2, muscle_group = $3, description = $4, equipment_id = $5
WHERE exercise_id = $1
RETURNING exercise_id, exercise_name, muscle_group, description, equipment_id;

-- name: ListExercise :many
SELECT exercise_id, exercise_name, description
FROM Exercise
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;

