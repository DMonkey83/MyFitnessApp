// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: userProfile.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserProfile = `-- name: CreateUserProfile :one
INSERT INTO UserProfile (user_id, full_name, age, gender, height_cm, height_ft_in, preferred_unit)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING user_profile_id, user_id, full_name, age, gender, height_cm, height_ft_in, preferred_unit, created_at
`

type CreateUserProfileParams struct {
	UserID        int64       `json:"user_id"`
	FullName      string      `json:"full_name"`
	Age           int32       `json:"age"`
	Gender        string      `json:"gender"`
	HeightCm      float64     `json:"height_cm"`
	HeightFtIn    pgtype.Text `json:"height_ft_in"`
	PreferredUnit Weightunit  `json:"preferred_unit"`
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, createUserProfile,
		arg.UserID,
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
		&i.UserID,
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
WHERE user_id = $1
`

func (q *Queries) DeleteUserProfile(ctx context.Context, userID int64) error {
	_, err := q.db.Exec(ctx, deleteUserProfile, userID)
	return err
}

const getUserProfile = `-- name: GetUserProfile :one
SELECT user_profile_id, user_id, full_name, age, gender, height_cm, height_ft_in, preferred_unit
FROM UserProfile
WHERE user_id = $1
`

type GetUserProfileRow struct {
	UserProfileID int64       `json:"user_profile_id"`
	UserID        int64       `json:"user_id"`
	FullName      string      `json:"full_name"`
	Age           int32       `json:"age"`
	Gender        string      `json:"gender"`
	HeightCm      float64     `json:"height_cm"`
	HeightFtIn    pgtype.Text `json:"height_ft_in"`
	PreferredUnit Weightunit  `json:"preferred_unit"`
}

func (q *Queries) GetUserProfile(ctx context.Context, userID int64) (GetUserProfileRow, error) {
	row := q.db.QueryRow(ctx, getUserProfile, userID)
	var i GetUserProfileRow
	err := row.Scan(
		&i.UserProfileID,
		&i.UserID,
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
WHERE user_id = $1
RETURNING user_profile_id, user_id, full_name, age, gender, height_cm, height_ft_in, preferred_unit, created_at
`

type UpdateUserProfileParams struct {
	UserID        int64       `json:"user_id"`
	FullName      string      `json:"full_name"`
	Age           int32       `json:"age"`
	Gender        string      `json:"gender"`
	HeightCm      float64     `json:"height_cm"`
	HeightFtIn    pgtype.Text `json:"height_ft_in"`
	PreferredUnit Weightunit  `json:"preferred_unit"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (Userprofile, error) {
	row := q.db.QueryRow(ctx, updateUserProfile,
		arg.UserID,
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
		&i.UserID,
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
