
-- name: CreateMaxRepGoal :one
INSERT INTO MaxRepGoal (username, exercise_id, goal_reps, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMaxRepGoal :one
SELECT *
FROM MaxRepGoal
WHERE goal_id = $1;

-- name: DeleteMaxRepGoal :exec
DELETE FROM MaxRepGoal
WHERE goal_id = $1;

-- name: UpdateMaxRepGoal :one
UPDATE MaxRepGoal
SET exercise_id = $2, goal_reps = $3, notes = $4
WHERE goal_id = $1
RETURNING *;

-- name: ListMaxRepGoals :many
SELECT *
FROM MaxRepGoal
WHERE username = $1 AND exercise_id = $2
ORDER BY goal_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $3
OFFSET $4;
