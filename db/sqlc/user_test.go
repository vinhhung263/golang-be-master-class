package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vinhhung263/simplebank/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	usr1 := createRandomUser(t)
	urs2, err := testQueries.GetUser(context.Background(), usr1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, urs2)

	require.Equal(t, usr1.Username, urs2.Username)
	require.Equal(t, usr1.HashedPassword, urs2.HashedPassword)
	require.Equal(t, usr1.FullName, urs2.FullName)
	require.Equal(t, usr1.Email, urs2.Email)
	require.WithinDuration(t, usr1.CreatedAt, urs2.CreatedAt, time.Second)
	require.WithinDuration(t, usr1.PasswordChangedAt, urs2.PasswordChangedAt, time.Second)
}
