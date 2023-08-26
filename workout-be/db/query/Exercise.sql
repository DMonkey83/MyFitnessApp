-- name: CreateExercise :one
INSERT INTO Exercise (exercise_name,workout_id, description, equipment_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetExercise :one
SELECT *
FROM Exercise
WHERE exercise_id = $1;

-- name: DeleteExercise :exec
DELETE FROM Exercise
WHERE exercise_id = $1;

-- name: UpdateExercise :one
UPDATE Exercise
SET exercise_name = $2, description = $3, equipment_id = $4
WHERE exercise_id = $1
RETURNING *;

-- name: ListWorkoutExercise :many
SELECT *
FROM Exercise
WHERE workout_id = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;

-- name: ListAllExercise :many
SELECT *
FROM Exercise
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
