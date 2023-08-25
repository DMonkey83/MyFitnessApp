package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashPassword(util.GetRandomUsername(8))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:     util.GetRandomUsername(7),
		Email:        util.GetRandomEmail(8),
		PasswordHash: hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.Username)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	arg := UpdateUserParams{
		Username: user1.Username,
		Email:    util.GetRandomEmail(8),
	}

	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Email, user2.Email)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user2)
}
