package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tranhuuhuy297/simple-bank-golang/util"
)

func createRandomUser() (User, error) {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	return user, err
}

func TestCreateUser(t *testing.T) {
	user, err := createRandomUser()

	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestGetUser(t *testing.T) {
	user1, _ := createRandomUser()
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
}
