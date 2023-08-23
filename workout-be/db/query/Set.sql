
-- name: CreateSet :one
INSERT INTO Set (workout_id, exercise_id, set_number, weight, rest_duration, notes)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING set_id;

-- name: GetSet :one
SELECT set_id, workout_id, exercise_id, set_number, weight, rest_duration, notes
FROM Set
WHERE set_id = $1;

-- name: DeleteSet :exec
DELETE FROM Set
WHERE set_id = $1;

-- name: UpdateSet :one
UPDATE Set
SET workout_id = $2, exercise_id = $3, set_number = $4, weight = $5, rest_duration = $6, notes = $7
WHERE set_id = $1
RETURNING set_id, workout_id, exercise_id, set_number, weight, rest_duration, notes;

-- name: ListSets :many
SELECT set_number, weight, rest_duration, notes
FROM Set
ORDER BY set_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
