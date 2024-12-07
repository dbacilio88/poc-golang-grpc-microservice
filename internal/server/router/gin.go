package router

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
	"net/http"
	"time"
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
	router       *gin.Engine
	middleware   *negroni.Negroni
	port         Port
	name         Name
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
}

func newGinFramework(port Port, name Name) *GinFramework {
	return &GinFramework{
		router:     gin.Default(),
		middleware: negroni.New(),
		port:       port,
		name:       name,
	}
}

func (f *GinFramework) Run() {
	f.router.GET("/health", healthCheckHandlerGin)
	f.middleware.UseHandler(f.router)
	listenAndServe(f.port, f.name, f.middleware)
}

func healthCheckHandlerGin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"alive": true,
	})
}
