-- name: CreateUser :one
INSERT INTO "User" (
  username, 
  email, 
  password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT username, email, password_hash, password_changed_at
FROM "User"
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE username = $1;

-- name: UpdateUser :one
UPDATE "User"
SET email = $2
WHERE username = $1
RETURNING *;


