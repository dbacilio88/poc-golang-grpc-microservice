package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
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

type ServerFactory struct {
	log *zap.Logger
}

type IServerFactory interface {
	Run()
}

func NewRouterFactory(instance int, port Port, name Name, log *zap.Logger) (IServerFactory, error) {
	switch instance {
	case InstanceRouterGin:
		return newGinFramework(port, name, log), nil
	case InstanceRouterGorilla:
		return newGorillaRouter(port, name, log), nil
	default:
		return nil, errors.New("invalid instance")
	}
}

func listenAndServe(port Port, name Name, middleware *negroni.Negroni, log *zap.Logger) {
	srv := createHttpServer(port, middleware)

	go func() {
		log.Info(fmt.Sprintf("Starting http server on port %d [%s]...", port, name))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error starting http server", zap.Error(err))
			return
		}
	}()

	stop := setupSignalHandler(log)
	// stop o shutdown server
	<-stop
	log.Info("shutting down http server", zap.Int("port", int(port)))
	// Set a time limit for the server to stop.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Try to shut down the server in an orderly manner.
	if err := srv.Shutdown(ctx); err != nil {
		//error shutdown
		log.Error("error shutting down http server", zap.Error(err))
	}
	os.Exit(0)

}

func createHttpServer(port Port, middleware *negroni.Negroni) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware,
	}
	return srv
}

// SetupSignalHandler configures signal handling for a controlled stop.
func setupSignalHandler(log *zap.Logger) (quitOs <-chan struct{}) {

	quit := make(chan struct{})
	// Channel to receive signals from the system
	s := make(chan os.Signal, 2)
	//	Interrupt Signal = syscall.SIGINT
	//	Kill      Signal = syscall.SIGKILL

	//signal.Notify(s, os.Interrupt, os.Kill, syscall.SIGTERM)
	//signal.Notify(s, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(s, os.Interrupt, syscall.SIGKILL)

	go func() {
		// Wait for the first signal and close the `stop` channel.
		next := <-s
		log.Info("Caught signal; shutting down...", zap.Any("signal", next))
		close(quit)
		// Wait for a second signal to finish immediately.
		next = <-s
		log.Info("Caught signal next; shutting down...", zap.Any("signal", next))
	}()
	return quit
}
