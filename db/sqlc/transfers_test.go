package bankdb

import (
	"context"
	"testing"

	"github.com/mikey247/go-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account.ID,
		ToAccountID:   account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer,err:= testQueries.CreateTransfer(context.Background(),arg)

	require.NoError(t,err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T){
	transfer:=createRandomTransfer(t)
	transfer2,err := testQueries.GetTransfer(context.Background(), transfer.ID)
	
	require.NoError(t,err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
}

func TestListTransfer(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg:= ListTransfersParams{
      Limit: 5,
	  Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t,err)
	require.Len(t,transfers,5)

	for _,transfer := range transfers{
		require.NotEmpty(t,transfer)
	}
}