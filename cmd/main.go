package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

/**
*
* main
* <p>
* main file
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

func main() {
	port := 3000

	//usar gin

	router := gin.Default()

	//usar gorilla mux
	//router := mux.NewRouter()

	fmt.Println("grpc server start...", port)

	//router.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)
	router.GET("/health", gin.WrapF(HealthCheckHandler))
	//http.Handle("/", router)

	server := http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		return
	}
	/*
		fmt.Println("helloWorld")
		err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		if err != nil {
			return
		}
	*/
}
