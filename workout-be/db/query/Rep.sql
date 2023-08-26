
-- name: CreateRep :one
INSERT INTO Rep (set_id, rep_number, completion_status, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetRep :one
SELECT *
FROM Rep
WHERE rep_id = $1;

-- name: DeleteRep :exec
DELETE FROM Rep
WHERE rep_id = $1;

-- name: UpdateRep :one
UPDATE Rep
SET rep_number = $2, completion_status = $3, notes = $4
WHERE rep_id = $1
RETURNING *;

-- name: ListReps :many
SELECT *
FROM Rep
WHERE set_id = $1
ORDER BY rep_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
