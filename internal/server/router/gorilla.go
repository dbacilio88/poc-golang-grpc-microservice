package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"io"
	"net/http"
	"time"
)

/**
*
* gorilla
* <p>
* gorilla file
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

type GorillaRouter struct {
	router       *mux.Router
	middleware   *negroni.Negroni
	port         Port
	name         Name
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
}

func newGorillaRouter(port Port, name Name) *GorillaRouter {
	return &GorillaRouter{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		port:       port,
		name:       name,
	}
}

func (r *GorillaRouter) Run() {
	r.router.HandleFunc("/health", healthCheckHandlerMux)
	r.middleware.UseHandler(r.router)
	listenAndServe(r.port, r.name, r.middleware)
}

func healthCheckHandlerMux(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request from ", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		return
	}
}
