package main

import (
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/event/broker"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/server"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/server/router"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/service"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/task"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/utils"
	"go.uber.org/zap"
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

	message := env.LoadProperties()

	log, err := utils.LoggerConfiguration(env.YAML.Server.Environment)

	if err != nil {
		return
	}

	log.Info(message)

	std := zap.RedirectStdLog(log)
	defer std()

	//go routine subscriber [producer] instance rabbit or kafka
	msg := service.NewMessaging(log, env.YAML).
		NewBrokerServer(broker.RabbitMqInstance)

	params := map[string]interface{}{
		"queueName":  env.YAML.Rabbitmq.Queue.Files.Name,
		"routingKey": env.YAML.Rabbitmq.RoutingKey.Files,
	}

	go msg.Receive(params)

	//msg := event.NewBrokerConfig(log).
	//	NewBrokerServer(broker.RabbitMqInstance)
	//go msg.BrokerSubscriber()

	// instance new task
	tsk := task.NewTask(log)
	// create instance task pending read file data
	scheduler := tsk.Create()
	if env.YAML.Scheduler.Enable {
		tsk.Run(scheduler, env.YAML.Workspace.Files.Path, msg)
	}

	app := server.NewHttpConfig(log).
		SetPort(env.YAML.Server.Port).
		SetName(router.NameRouterGorilla).
		NewHttpServer(router.InstanceRouterGorilla)

	// start instance server http
	app.Start()

	//time.Sleep(1 * time.Second)
	select {}

}
