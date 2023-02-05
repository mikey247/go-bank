package bankdb

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store  {
	return &Store{
		db: 	db,
	Queries :   New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
	tx,err:= store.db.BeginTx(ctx, nil)

	if err != nil{
		return err
	}

	q:= New(tx)
	err= fn(q)
	
	if err != nil{
		if rBErr:= tx.Rollback(); rBErr !=nil{
			return fmt.Errorf("tx err: %v, rberr: %w", err,rBErr)
		}
		return err 
	}

	return tx.Commit()
}

type  TransferTxParams struct{
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

type TransferTxResult struct{
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

var txKey = struct{}{}

// TransferTx performs a money transfer from one account to the other.
// It creates a transfer record, add account entries, and updates accounts/ balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult,error){
 var result TransferTxResult

 
 err:= store.execTx(ctx, func(q *Queries) error {
	var txName = ctx.Value(txKey)
	var err error

	fmt.Println(txName,"create Transfer")
	// create the transfer record
	result.Transfer, err = q.CreateTransfer(ctx,CreateTransferParams{
		FromAccountID:	arg.FromAccountID,
		ToAccountID: 	arg.ToAccountID,
		Amount: 		arg.Amount,
	})

	if err!= nil{
		return err
	} 


	// create from entry
	fmt.Println(txName, "create from entrys")
	result.FromEntry, err = q.CreateEntry(ctx,CreateEntryParams{
		AccountID:	 		arg.FromAccountID,
		Amount: 			-arg.Amount,
	})

	if err!= nil{
		return err
	} 

	// create to entry
	fmt.Println(txName, "create to entry")
	result.ToEntry, err = q.CreateEntry(ctx,CreateEntryParams{
		AccountID:	 		arg.ToAccountID,
		Amount: 			arg.Amount,
	})

	if err!= nil{
		return err
	} 

	// TODO: update account's Balance
	fmt.Println(txName,"get account 1" )
	account1,err:= q.GetAccountForUpdate(ctx,arg.FromAccountID)

	if err!= nil{
		return err
	} 

	fmt.Println(txName, "update account 1")
	result.FromAccount,err=q.UpdateAccount(ctx, UpdateAccountParams{
		ID: arg.FromAccountID,
		Balance: account1.Balance - arg.Amount,
	})

	if err!= nil{
		return err
	} 

	fmt.Println(txName,"get account 2" )
	account2,err:= q.GetAccountForUpdate(ctx,arg.ToAccountID)
	
	if err!= nil{
		return err
	} 

	fmt.Println(txName, "update account 2")
	result.ToAccount,err=q.UpdateAccount(ctx, UpdateAccountParams{
		ID: arg.ToAccountID,
		Balance: account2.Balance + arg.Amount,
	})

	if err!= nil{
		return err
	} 

	return nil
	})

 	return result,err
}