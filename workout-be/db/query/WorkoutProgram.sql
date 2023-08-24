
-- name: CreateWorkoutprogram :one
INSERT INTO WorkoutProgram (username, program_name, description)
VALUES ($1, $2, $3)
RETURNING program_id;

-- name: GetWorkoutprogram :one
SELECT program_id, username, program_name, description
FROM WorkoutProgram
WHERE program_id = $1;

-- name: DeleteWorkoutprogram :exec
DELETE FROM WorkoutProgram
WHERE program_id = $1;

-- name: UpdateWorkoutprogram :one
UPDATE WorkoutProgram
SET username = $2, program_name = $3, description = $4
WHERE program_id = $1
RETURNING program_id, username, program_name, description;

-- name: ListWorkoutprograms :many
SELECT program_name, description
FROM WorkoutProgram
ORDER BY program_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
