
-- name: CreateSet :one
INSERT INTO Set (exercise_id, set_number, weight, rest_duration, notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSet :one
SELECT *
FROM Set
WHERE set_id = $1;

-- name: DeleteSet :exec
DELETE FROM Set
WHERE set_id = $1;

-- name: UpdateSet :one
UPDATE Set
SET set_number = $2, weight = $3, rest_duration = $4, notes = $5
WHERE set_id = $1
RETURNING *;

-- name: ListSets :many
SELECT *
FROM Set
WHERE exercise_id = $1
ORDER BY set_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
