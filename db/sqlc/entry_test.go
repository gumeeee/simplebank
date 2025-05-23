package db

import (
	"context"
	"testing"

	"github.com/gumeeee/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entrySaved := createRandomEntry(t, account)

	entrySearched, err := testQueries.GetEntry(context.Background(), entrySaved.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entrySearched)

	require.Equal(t, entrySaved.ID, entrySearched.ID)
	require.Equal(t, entrySaved.Amount, entrySearched.Amount)
	require.Equal(t, entrySaved.AccountID, entrySearched.AccountID)

	require.NotZero(t, entrySearched.ID)
	require.NotZero(t, entrySearched.CreatedAt)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
