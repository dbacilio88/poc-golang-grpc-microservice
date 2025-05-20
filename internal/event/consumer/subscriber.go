package consumer

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env/mq"
	"go.uber.org/zap"
)

/**
*
* consumer
* <p>
* consumer file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 8/12/2024
*
 */

type IBrokerSubscriber interface {
	SubscriberRabbitMq(config amqp.Config) (*amqp.Subscriber, error)
	SubscriberKafkaMq() error
}

type BrokerSubscriber struct {
	*zap.Logger
	mq.Rabbitmq
}

func NewBrokerSubscriber(log *zap.Logger, prop mq.Rabbitmq) IBrokerSubscriber {
	return &BrokerSubscriber{
		Logger:   log,
		Rabbitmq: prop,
	}
}

func (b *BrokerSubscriber) SubscriberRabbitMq(config amqp.Config) (*amqp.Subscriber, error) {
	sub, err := amqp.NewSubscriber(config, watermill.NewStdLogger(false, false))
	if err != nil {
		b.Error("Error creating subscriber", zap.Error(err))
		return nil, err
	}
	return sub, nil
}

func (b *BrokerSubscriber) SubscriberKafkaMq() error {
	b.Info("Kafka subscription is not implemented yet.")
	return fmt.Errorf("kafka subscription is not implemented yet")
}
