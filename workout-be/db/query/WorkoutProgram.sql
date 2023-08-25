
-- name: CreateWorkoutprogram :one
INSERT INTO WorkoutProgram (username, program_name, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetWorkoutprogram :one
SELECT *
FROM WorkoutProgram
WHERE program_id = $1;

-- name: DeleteWorkoutprogram :exec
DELETE FROM WorkoutProgram
WHERE program_id = $1;

-- name: UpdateWorkoutprogram :one
UPDATE WorkoutProgram
SET program_name = $2, description = $3
WHERE program_id = $1
RETURNING *;

-- name: ListWorkoutprograms :many
SELECT *
FROM WorkoutProgram
ORDER BY program_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
