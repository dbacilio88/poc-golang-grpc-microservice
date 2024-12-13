package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
*
* store_test
* <p>
* store_test file
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

func TestStore_TransactionTx(t *testing.T) {
	store := NewStore(testDb)

	email := createRandomEmail(t)
	transaction := createTransactionRandom(t, email)

	n := 5
	errs := make(chan error)
	results := make(chan TransactionTxResult)
	for i := 0; i < n; i++ {
		go func() {
			tx, err := store.TransactionTx(context.Background(), TransactionTxParams{
				Email:       email,
				Transaction: transaction,
			})
			errs <- err
			results <- tx
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		tx := result.Transaction

		require.NotEmpty(t, tx)
		require.Equal(t, email.ID, tx.EmailID)
		require.Equal(t, email.Status, tx.Status)
		require.NotZero(t, tx.ID)
		require.NotZero(t, tx.CreatedAt)

		_, err = store.GetTransaction(context.Background(), tx.ID)
		require.NoError(t, err)

	}
}
