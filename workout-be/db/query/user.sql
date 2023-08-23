-- name: CreateUser :one
INSERT INTO "User" (username, email, password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT user_id, username, email
FROM "User"
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE user_id = $1;

-- name: UpdateUser :one
UPDATE "User"
SET username = $2, email = $3
WHERE user_id = $1
RETURNING *;

-- name: ListUsers :many
SELECT user_id, username, email
FROM "User"
ORDER BY user_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;

