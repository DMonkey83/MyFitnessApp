-- name: CreateWeightEntry :one
INSERT INTO WeightEntry (user_id, entry_date, weight_kg, weight_lb, notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING weight_entry_id;

-- name: GetWeightEntry :one
SELECT weight_entry_id, user_id, entry_date, weight_kg, weight_lb, notes
FROM WeightEntry
WHERE weight_entry_id = $1;

-- name: DeleteWeightEntry :exec
DELETE FROM WeightEntry
WHERE weight_entry_id = $1;

-- name: UpdateWeightEntry :one
UPDATE WeightEntry
SET user_id = $2, entry_date = $3, weight_kg = $4, weight_lb = $5, notes = $6
WHERE weight_entry_id = $1
RETURNING weight_entry_id, user_id, entry_date, weight_kg, weight_lb, notes;

-- name: ListWeightEntries :many
SELECT entry_date, weight_kg, weight_lb, notes
FROM WeightEntry
ORDER BY weight_entry_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
