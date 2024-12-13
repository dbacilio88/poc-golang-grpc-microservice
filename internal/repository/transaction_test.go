package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/**
*
* transaction_test
* <p>
* transaction_test file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 12/12/2024
*
 */

func createTransactionRandom(t *testing.T, email Email) Transaction {
	param := CreateTransactionParams{
		Status:  email.Status,
		EmailID: email.ID,
	}
	transaction, err := testQueries.CreateTransaction(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, param.Status, transaction.Status)
	require.Equal(t, param.EmailID, transaction.EmailID)
	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)
	return transaction
}

func TestQueries_CreateTransaction(t *testing.T) {
	email := createRandomEmail(t)
	createTransactionRandom(t, email)
}

func TestQueries_GetTransaction(t *testing.T) {
	email := createRandomEmail(t)
	transactionNew := createTransactionRandom(t, email)
	transaction, err := testQueries.GetTransaction(context.Background(), transactionNew.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, transactionNew.ID, transaction.ID)
	require.Equal(t, transactionNew.Status, transaction.Status)
	require.Equal(t, transactionNew.EmailID, transaction.EmailID)
	require.WithinDuration(t, transactionNew.CreatedAt, transaction.CreatedAt, time.Second)

}

func TestQueries_ListTransactions(t *testing.T) {
	email := createRandomEmail(t)
	for i := 0; i < 5; i++ {
		createTransactionRandom(t, email)
	}
	param := ListTransactionsParams{
		EmailID: email.ID,
		Limit:   2,
		Offset:  2,
	}

	transactions, err := testQueries.ListTransactions(context.Background(), param)

	require.NoError(t, err)
	require.Len(t, transactions, 2)

	for _, tx := range transactions {
		require.NotEmpty(t, tx)
		require.Equal(t, param.EmailID, tx.EmailID)
	}
}
