package server

import (
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/server/router"
	"time"
)

/**
*
* http
* <p>
* http file
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

type HttpConfig struct {
	factory router.ServerFactory
	port    router.Port
	name    router.Name
	timeout time.Duration
}

func NewHttpConfig() *HttpConfig {
	return &HttpConfig{
		timeout: time.Second * 10,
	}
}

func (c *HttpConfig) NewHttpServer(instance int) *HttpConfig {
	factory, err := router.NewRouterFactory(
		instance, c.port, c.name)

	if err != nil {
		return nil
	}

	c.factory = factory
	return c
}

func (c *HttpConfig) SetPort(port int) *HttpConfig {
	c.port = router.Port(port)
	return c
}

func (c *HttpConfig) SetName(name string) *HttpConfig {
	c.name = router.Name(name)
	return c
}

func (c *HttpConfig) Start() {
	c.factory.Run()
}
