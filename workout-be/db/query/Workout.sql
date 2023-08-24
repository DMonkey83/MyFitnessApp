
-- name: CreateWorkout :one
INSERT INTO Workout (username, workout_date, workout_duration, notes)
VALUES ($1, $2, $3, $4)
RETURNING workout_id;

-- name: GetWorkout :one
SELECT workout_id, username, workout_date, workout_duration, notes
FROM Workout
WHERE workout_id = $1;

-- name: DeleteWorkout :exec
DELETE FROM Workout
WHERE workout_id = $1;

-- name: UpdateWorkout :one
UPDATE Workout
SET username = $2, workout_date = $3, workout_duration = $4, notes = $5
WHERE workout_id = $1
RETURNING workout_id, username, workout_date, workout_duration, notes;

-- name: ListWorkouts :many
SELECT workout_id, workout_date, workout_duration, notes
FROM Workout
ORDER BY workout_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
