
-- name: CreateMaxRepGoal :one
INSERT INTO MaxRepGoal (user_id, exercise_id, goal_reps, notes)
VALUES ($1, $2, $3, $4)
RETURNING goal_id;

-- name: GetMaxRepGoal :one
SELECT goal_id, user_id, exercise_id, goal_reps, notes
FROM MaxRepGoal
WHERE goal_id = $1;

-- name: DeleteMaxRepGoal :exec
DELETE FROM MaxRepGoal
WHERE goal_id = $1;

-- name: UpdateMaxRepGoal :one
UPDATE MaxRepGoal
SET user_id = $2, exercise_id = $3, goal_reps = $4, notes = $5
WHERE goal_id = $1
RETURNING goal_id, user_id, exercise_id, goal_reps, notes;

-- name: ListMaxRepGoals :many
SELECT goal_reps, notes
FROM MaxRepGoal
ORDER BY goal_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
