package main

import (
	srv "github.com/dbacilio88/golang-grpc-email-microservice/internal/server"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/server/router"
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

	app := srv.NewHttpConfig().
		SetPort(port).
		SetName(router.NameRouterGin).
		NewHttpServer(router.InstanceRouterGin)

	app.Start()

	//usar gin
	/*
		gr.GET("/health", HealthCheckHandlerGin)

		http.Handle("/", gmr)

		server := http.Server{
			Handler: gr,
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

	*/
}
