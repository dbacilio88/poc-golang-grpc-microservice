package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
*
* factory
* <p>
* factory file
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

type Port int64
type Name string

const InstanceRouterGin int = iota
const InstanceRouterGorilla = 1
const NameRouterGin string = "Gin"
const NameRouterGorilla string = "Gorilla Mux"

type ServerFactory interface {
	Run()
}

func NewRouterFactory(instance int, port Port, name Name) (ServerFactory, error) {
	switch instance {
	case InstanceRouterGin:
		return newGinFramework(port, name), nil
	case InstanceRouterGorilla:
		return newGorillaRouter(port, name), nil
	default:
		return nil, errors.New("invalid instance")
	}
}

func listenAndServe(port Port, name Name, middleware *negroni.Negroni) {
	srv := createHttpServer(port, middleware)

	stop := setupSignalHandler()

	go func() {
		fmt.Printf("start http server %s on %d\n", name, port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("http server listen error", err)
			return
		}
	}()

	// Esperar a recibir una señal
	<-stop

	fmt.Println("shutting down http server...")
	// Establece un tiempo límite para la parada del servidor.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	// Intenta cerrar el servidor de manera ordenada.
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	fmt.Println("http server shutdown successfully")
}

func createHttpServer(port Port, middleware *negroni.Negroni) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware,
	}
	return srv
}

// SetupSignalHandler configura el manejo de señales para una parada controlada.
func setupSignalHandler() (quitOs <-chan struct{}) {

	quit := make(chan struct{})
	// Canal para recibir señales del sistema
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Espera la primera señal y cierra el canal `stop`.
		next := <-s
		fmt.Println("caught signal next", next)
		close(quit)
		// Espera una segunda señal para terminar inmediatamente.
		next = <-s
		fmt.Println("caught signal next", next)
		os.Exit(1)
	}()
	return quit
}
