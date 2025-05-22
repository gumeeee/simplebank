package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gumeeee/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	params := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountSaved := createRandomAccount(t)
	accountSearched, err := testQueries.GetAccount(context.Background(), accountSaved.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accountSearched)

	require.Equal(t, accountSaved.ID, accountSearched.ID)
	require.Equal(t, accountSaved.Owner, accountSearched.Owner)
	require.Equal(t, accountSaved.Balance, accountSearched.Balance)
	require.Equal(t, accountSaved.Currency, accountSearched.Currency)
	require.WithinDuration(t, accountSaved.CreatedAt, accountSearched.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accountSaved := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      accountSaved.ID,
		Balance: util.RandomMoney(),
	}

	accountUpdated, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountUpdated)

	require.Equal(t, accountSaved.ID, accountUpdated.ID)
	require.Equal(t, accountSaved.Owner, accountUpdated.Owner)
	require.Equal(t, arg.Balance, accountUpdated.Balance)
	require.Equal(t, accountSaved.Currency, accountUpdated.Currency)
	require.WithinDuration(t, accountSaved.CreatedAt, accountUpdated.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	accountSaved := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accountSaved.ID)
	require.NoError(t, err)

	accountDeleted, err := testQueries.GetAccount(context.Background(), accountSaved.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountDeleted)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
