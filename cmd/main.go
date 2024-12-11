package main

import (
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event/broker"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/server"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/server/router"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/task"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/utils"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/yaml"
	"go.uber.org/zap"
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

	yaml.LoadProperties()

	log, err := utils.LoggerConfiguration(yaml.YAML.Server.Environment)
	if err != nil {
		return
	}

	std := zap.RedirectStdLog(log)
	defer std()

	msg := event.NewBrokerConfig(log).
		NewBrokerServer(broker.RabbitMqInstance)

	//go routine subscriber instance rabbit or kafka
	go msg.Subscriber()

	//create instance task
	tsk := task.NewTask(log)
	scheduler := tsk.Create()
	tsk.Run(scheduler)

	app := server.NewHttpConfig(log).
		SetPort(yaml.YAML.Server.Port).
		SetName(router.NameRouterGorilla).
		NewHttpServer(router.InstanceRouterGorilla)

	// start instance server http
	app.Start()

	time.Sleep(1 * time.Second)
	//select {}

}
