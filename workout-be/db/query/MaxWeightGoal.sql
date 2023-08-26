
-- name: CreateMaxWeightGoal :one
INSERT INTO MaxWeightGoal (username, exercise_id, goal_weight, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMaxWeightGoal :one
SELECT *
FROM MaxWeightGoal
WHERE goal_id = $1;

-- name: DeleteMaxWeightGoal :exec
DELETE FROM MaxWeightGoal
WHERE goal_id = $1;

-- name: UpdateMaxWeightGoal :one
UPDATE MaxWeightGoal
SET exercise_id = $2, goal_weight = $3, notes = $4
WHERE goal_id = $1
RETURNING *;

-- name: ListMaxWeightGoals :many
SELECT *
FROM MaxWeightGoal
WHERE username = $1 AND exercise_id = $2
ORDER BY goal_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $3
OFFSET $4;
