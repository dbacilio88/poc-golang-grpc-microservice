package handler

import (
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/models"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
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

type EmailHandler struct {
	es service.IEmailService
	*zap.Logger
}

type IEmailHandler interface {
	CreateEmailHandler(ctx *gin.Context)
	GetEmailHandler(ctx *gin.Context)
	GetEmailsHandler(ctx *gin.Context)
}

func NewEmailHandler(log *zap.Logger) IEmailHandler {
	srv := service.NewEmailService(log)
	return &EmailHandler{
		Logger: log,
		es:     srv,
	}
}

func (e *EmailHandler) GetEmailsHandler(ctx *gin.Context) {
	emailsService, err := e.es.GetEmailsService(ctx)
	if err != nil {
		log.Fatal("Error getting emails:")
		return
	}
	ctx.JSON(http.StatusOK, emailsService)
}

func (e *EmailHandler) CreateEmailHandler(ctx *gin.Context) {
	var request models.CreateEmailRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	emailService, err := e.es.CreateEmailService(ctx, request)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, emailService)
}

func (e *EmailHandler) GetEmailHandler(ctx *gin.Context) {
	var request models.CreateEmailRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
