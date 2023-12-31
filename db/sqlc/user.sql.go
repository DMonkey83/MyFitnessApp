// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, 
  email, 
  password_hash)
VALUES ($1, $2, $3)
RETURNING username, email, password_hash, password_changed_at, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT username, email, password_hash, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
email = COALESCE($1,email),
password_hash = COALESCE($2,password_hash),
password_changed_at = COALESCE($3,password_changed_at)
WHERE 
username = $4
RETURNING username, email, password_hash, password_changed_at, created_at
`

type UpdateUserParams struct {
	Email             pgtype.Text        `json:"email"`
	PasswordHash      pgtype.Text        `json:"password_hash"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	Username          string             `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Email,
		arg.PasswordHash,
		arg.PasswordChangedAt,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
