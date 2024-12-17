package service

import (
	"context"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/models"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/repository"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/database"
	"log"
)

/**
*
* email
* <p>
* email file
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
* @since 13/12/2024
*
 */

type EmailService struct {
	*database.Connection
}

type IEmailService interface {
	CreateEmailService(ctx context.Context, req models.CreateEmailRequest) (models.CreateEmailResponse, error)
	GetEmailService(ctx context.Context, id int64) (models.CreateEmailResponse, error)
	GetEmailsService(ctx context.Context) ([]models.CreateEmailResponse, error)
}

func NewEmailService() IEmailService {
	return &EmailService{
		database.NewConnection(),
	}
}
func (e *EmailService) GetEmailsService(ctx context.Context) ([]models.CreateEmailResponse, error) {

	param := repository.ListEmailsParams{
		Offset: 1,
		Limit:  10,
		Status: "SEND",
	}
	emails, err := e.Store.ListEmails(ctx, param)
	if err != nil {
		log.Fatal("Error connecting to database")
		return nil, err
	}
	var response []models.CreateEmailResponse
	for _, email := range emails {
		response = append(response, models.CreateEmailResponse{
			Title:  email.Title,
			Body:   email.Body,
			Status: email.Status,
		})
	}
	return response, nil
}
func (e *EmailService) GetEmailService(ctx context.Context, id int64) (models.CreateEmailResponse, error) {
	email, err := e.Store.GetEmail(ctx, id)
	if err != nil {
		return models.CreateEmailResponse{}, err
	}
	response := models.CreateEmailResponse{
		Title:  email.Title,
		Body:   email.Body,
		Status: email.Status,
	}
	return response, nil
}

func (e *EmailService) CreateEmailService(ctx context.Context, req models.CreateEmailRequest) (models.CreateEmailResponse, error) {
	param := repository.CreateEmailParams{
		Status: req.Status,
		Body:   req.Body,
		Title:  req.Title,
	}
	email, err := e.Store.CreateEmail(ctx, param)
	if err != nil {
		return models.CreateEmailResponse{}, err
	}
	response := models.CreateEmailResponse{
		Title:  email.Title,
		Body:   email.Body,
		Status: email.Status,
	}
	return response, nil
}
