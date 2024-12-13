package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

/**
*
* store
* <p>
* store file
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

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {

	//s.db.BeginTx(ctx,sql.TxOptions{})
	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		log.Fatal("Error starting transaction: ", err)
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {

		if txErr := tx.Rollback(); txErr != nil {
			return fmt.Errorf("error tx: %v, rolling back transaction: %v", err, txErr)
		}

		return err
	}

	return tx.Commit()
}

type TransactionTxParams struct {
	Email       Email       `json:"email"`
	Transaction Transaction `json:"transaction"`
}

type TransactionTxResult struct {
	Transaction Transaction `json:"transaction"`
}

func (s *Store) TransactionTx(ctx context.Context, arg TransactionTxParams) (TransactionTxResult, error) {
	var result TransactionTxResult
	err := s.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Transaction, err = queries.CreateTransaction(ctx, CreateTransactionParams{
			EmailID: arg.Email.ID,
			Status:  arg.Email.Status,
		})
		if err != nil {
			log.Fatal("Error creating transaction:")
			return err
		}
		return nil
	})

	return result, err
}
