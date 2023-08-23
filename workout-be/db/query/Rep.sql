
-- name: CreateRep :one
INSERT INTO Rep (set_id, rep_number, completed, notes)
VALUES ($1, $2, $3, $4)
RETURNING rep_id;

-- name: GetRep :one
SELECT rep_id, set_id, rep_number, completed, notes
FROM Rep
WHERE rep_id = $1;

-- name: DeleteRep :exec
DELETE FROM Rep
WHERE rep_id = $1;

-- name: UpdateRep :one
UPDATE Rep
SET set_id = $2, rep_number = $3, completed = $4, notes = $5
WHERE rep_id = $1
RETURNING rep_id, set_id, rep_number, completed, notes;

-- name: ListReps :many
SELECT rep_number, completed, notes
FROM Rep
ORDER BY rep_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
