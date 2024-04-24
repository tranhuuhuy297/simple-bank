package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tranhuuhuy297/simple-bank-golang/util"
)

func createRandomAccount() (Account, error) {
	user, err := createRandomUser()

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	return account, err
}

func TestCreateAccount(t *testing.T) {
	account, err := createRandomAccount()

	require.NoError(t, err)
	require.NotEmpty(t, account)
}

func TestGetAccount(t *testing.T) {
	account1, _ := createRandomAccount()
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
}

func TestUpdateAccount(t *testing.T) {

}

func TestDeleteAccount(t *testing.T) {

}
