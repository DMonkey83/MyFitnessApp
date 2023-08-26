-- name: CreateWeightEntry :one
INSERT INTO WeightEntry (username, entry_date, weight_kg, weight_lb, notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetWeightEntry :one
SELECT *
FROM WeightEntry
WHERE weight_entry_id = $1;

-- name: DeleteWeightEntry :exec
DELETE FROM WeightEntry
WHERE weight_entry_id = $1;

-- name: UpdateWeightEntry :one
UPDATE WeightEntry
SET entry_date = $2, weight_kg = $3, weight_lb = $4, notes = $5
WHERE weight_entry_id = $1
RETURNING *;

-- name: ListWeightEntries :many
SELECT *
FROM WeightEntry
WHERE username = $1
ORDER BY weight_entry_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
