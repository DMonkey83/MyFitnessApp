-- name: CreateUser :one
INSERT INTO users (
  username, 
  email, 
  password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2
WHERE username = $1
RETURNING *;


