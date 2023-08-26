// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: userProfile.sql

package db

import (
	"context"
)

const createUserProfile = `-- name: CreateUserProfile :one
INSERT INTO UserProfile (username, full_name, age, gender, height_cm, height_ft_in, preferred_unit)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING user_profile_id, username, full_name, age, gender, height_cm, height_ft_in, preferred_unit, created_at
`

type CreateUserProfileParams struct {
	Username      string     `json:"username"`
	FullName      string     `json:"full_name"`
	Age           int32      `json:"age"`
	Gender        string     `json:"gender"`
	HeightCm      int32      `json:"height_cm"`
	HeightFtIn    string     `json:"height_ft_in"`
	PreferredUnit Weightunit `json:"preferred_unit"`
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, createUserProfile,
		arg.Username,
		arg.FullName,
		arg.Age,
		arg.Gender,
		arg.HeightCm,
		arg.HeightFtIn,
		arg.PreferredUnit,
	)
	var i Userprofile
	err := row.Scan(
		&i.UserProfileID,
		&i.Username,
		&i.FullName,
		&i.Age,
		&i.Gender,
		&i.HeightCm,
		&i.HeightFtIn,
		&i.PreferredUnit,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUserProfile = `-- name: DeleteUserProfile :exec
DELETE FROM UserProfile
WHERE username = $1
`

func (q *Queries) DeleteUserProfile(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUserProfile, username)
	return err
}

const getUserProfile = `-- name: GetUserProfile :one
SELECT user_profile_id, username, full_name, age, gender, height_cm, height_ft_in, preferred_unit
FROM UserProfile
WHERE username = $1
`

type GetUserProfileRow struct {
	UserProfileID int64      `json:"user_profile_id"`
	Username      string     `json:"username"`
	FullName      string     `json:"full_name"`
	Age           int32      `json:"age"`
	Gender        string     `json:"gender"`
	HeightCm      int32      `json:"height_cm"`
	HeightFtIn    string     `json:"height_ft_in"`
	PreferredUnit Weightunit `json:"preferred_unit"`
}

func (q *Queries) GetUserProfile(ctx context.Context, username string) (GetUserProfileRow, error) {
	row := q.db.QueryRow(ctx, getUserProfile, username)
	var i GetUserProfileRow
	err := row.Scan(
		&i.UserProfileID,
		&i.Username,
		&i.FullName,
		&i.Age,
		&i.Gender,
		&i.HeightCm,
		&i.HeightFtIn,
		&i.PreferredUnit,
	)
	return i, err
}

const updateUserProfile = `-- name: UpdateUserProfile :one
UPDATE UserProfile
SET full_name = $2, age = $3, gender = $4, height_cm = $5, height_ft_in = $6, preferred_unit = $7
WHERE username = $1
RETURNING user_profile_id, username, full_name, age, gender, height_cm, height_ft_in, preferred_unit, created_at
`

type UpdateUserProfileParams struct {
	Username      string     `json:"username"`
	FullName      string     `json:"full_name"`
	Age           int32      `json:"age"`
	Gender        string     `json:"gender"`
	HeightCm      int32      `json:"height_cm"`
	HeightFtIn    string     `json:"height_ft_in"`
	PreferredUnit Weightunit `json:"preferred_unit"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, updateUserProfile,
		arg.Username,
		arg.FullName,
		arg.Age,
		arg.Gender,
		arg.HeightCm,
		arg.HeightFtIn,
		arg.PreferredUnit,
	)
	var i Userprofile
	err := row.Scan(
		&i.UserProfileID,
		&i.Username,
		&i.FullName,
		&i.Age,
		&i.Gender,
		&i.HeightCm,
		&i.HeightFtIn,
		&i.PreferredUnit,
		&i.CreatedAt,
	)
	return i, err
}
