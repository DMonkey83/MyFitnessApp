
-- name: CreateMaxWeightGoal :one
INSERT INTO MaxWeightGoal (username, exercise_id, goal_weight, notes)
VALUES ($1, $2, $3, $4)
RETURNING goal_id;

-- name: GetMaxWeightGoal :one
SELECT goal_id, username, exercise_id, goal_weight, notes
FROM MaxWeightGoal
WHERE goal_id = $1;

-- name: DeleteWeightRepGoal :exec
DELETE FROM MaxWeightGoal
WHERE goal_id = $1;

-- name: UpdateMaxWeightGoal :one
UPDATE MaxWeightGoal
SET exercise_id = $2, goal_weight = $3, notes = $4
WHERE goal_id = $1
RETURNING goal_id, username, exercise_id, goal_weight, notes;

-- name: ListMaxWeightGoals :many
SELECT goal_weight, notes
FROM MaxWeightGoal
ORDER BY goal_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
