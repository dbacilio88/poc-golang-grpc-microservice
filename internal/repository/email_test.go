package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/test"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/**
*
* email_test
* <p>
* email_test file
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

func createRandomEmail(t *testing.T) Email {
	params := CreateEmailParams{
		Title:  test.RandomTitle(),
		Body:   test.RandomBody(),
		Status: test.RandomStatus(),
	}
	email, err := testQueries.CreateEmail(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, email)

	require.Equal(t, params.Title, email.Title)
	require.Equal(t, params.Body, email.Body)
	require.Equal(t, params.Status, email.Status)

	require.NotZero(t, email.ID)
	require.NotZero(t, email.CreatedAt)
	return email
}

func TestQueries_CreateEmail(t *testing.T) {
	createRandomEmail(t)
}

func TestQueries_GetEmail(t *testing.T) {
	emailNew := createRandomEmail(t)
	email, err := testQueries.GetEmail(context.Background(), emailNew.ID)
	require.NoError(t, err)
	require.NotEmpty(t, email)
	require.Equal(t, emailNew.ID, email.ID)
	require.Equal(t, emailNew.Title, email.Title)
	require.Equal(t, emailNew.Body, email.Body)
	require.Equal(t, emailNew.Status, email.Status)
	require.Equal(t, emailNew.CreatedAt, email.CreatedAt)
	require.WithinDuration(t, emailNew.CreatedAt, email.CreatedAt, time.Second)
}

func TestQueries_UpdateEmail(t *testing.T) {
	emailNew := createRandomEmail(t)
	param := UpdateEmailParams{
		ID:     emailNew.ID,
		Title:  test.RandomTitle(),
		Body:   test.RandomBody(),
		Status: test.RandomStatus(),
	}
	email, err := testQueries.UpdateEmail(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, email)
	require.Equal(t, param.Title, email.Title)
	require.Equal(t, param.Body, email.Body)
	require.Equal(t, param.Status, email.Status)
	require.Equal(t, emailNew.ID, email.ID)
	require.Equal(t, emailNew.CreatedAt, email.CreatedAt)
	require.WithinDuration(t, emailNew.CreatedAt, email.CreatedAt, time.Second)

	require.NotZero(t, email.CreatedAt)
	require.NotZero(t, email.ID)
}

func TestQueries_DeleteBook(t *testing.T) {
	emailNew := createRandomEmail(t)

	err := testQueries.DeleteBook(context.Background(), emailNew.ID)
	require.NoError(t, err)

	email, err := testQueries.GetEmail(context.Background(), emailNew.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, email)
}

func TestQueries_ListEmails(t *testing.T) {
	var emailNew Email
	for i := 0; i < 10; i++ {
		emailNew = createRandomEmail(t)
	}

	param := ListEmailsParams{
		Status: emailNew.Status,
		Limit:  5,
		Offset: 1,
	}
	fmt.Println(param)

	emails, err := testQueries.ListEmails(context.Background(), param)
	fmt.Println(len(emails))
	require.NoError(t, err)
	require.Len(t, emails, 5)

	for _, email := range emails {
		require.NotEmpty(t, email)
		require.Equal(t, emailNew.Status, email.Status)
	}
}
