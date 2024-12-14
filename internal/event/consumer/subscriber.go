package consumer

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
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

type BrokerSubscriber struct {
	*zap.Logger
}

type IBrokerSubscriber interface {
	SubscriberRabbitMq(cfg amqp.Config) (*amqp.Subscriber, error)
	SubscriberKafkaMq() error
}

func NewBrokerSubscriber(log *zap.Logger) IBrokerSubscriber {
	return &BrokerSubscriber{
		Logger: log,
	}
}

func (b *BrokerSubscriber) SubscriberRabbitMq(cfg amqp.Config) (*amqp.Subscriber, error) {
	sub, err := amqp.NewSubscriber(cfg, watermill.NewStdLogger(false, false))
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
