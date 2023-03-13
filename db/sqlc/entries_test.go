package bankdb

import (
	"context"
	"testing"

	"github.com/mikey247/go-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry  {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry,err:= testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t,err)
	require.NotEmpty(t,entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
   createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry:=createRandomEntry(t)
	entry2,err := testQueries.GetEntry(context.Background(), entry.ID)
	
	require.NoError(t,err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg:= ListEntriesParams{
      Limit: 5,
	  Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t,err)
	require.Len(t,entries,5)

	for _,account := range entries{
		require.NotEmpty(t,account)
	}
}