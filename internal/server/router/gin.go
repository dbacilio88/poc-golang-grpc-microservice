package router

import (
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
	"net/http"
)

/**
*
* gin
* <p>
* gin file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 6/12/2024
*
 */

type GinFramework struct {
	log        *zap.Logger
	router     *gin.Engine
	middleware *negroni.Negroni
	port       Port
	name       Name
	h          handler.IEmailHandler
}

func newGinFramework(port Port, name Name, log *zap.Logger) *GinFramework {
	han := handler.NewEmailHandler()
	return &GinFramework{
		log:        log,
		router:     gin.Default(),
		middleware: negroni.New(),
		port:       port,
		name:       name,
		h:          han,
	}
}

func (f *GinFramework) Run() {

	f.router.GET("/health", healthCheckHandlerGin)
	f.router.GET("/emails", f.h.GetEmailsHandler)
	f.middleware.UseHandler(f.router)
	listenAndServe(f.port, f.name, f.middleware, f.log)
}

func healthCheckHandlerGin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"alive": true,
	})
}
