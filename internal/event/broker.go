package event

import (
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event/broker"
	"go.uber.org/zap"
)

/**
*
* instance
* <p>
* instance file
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
* @since 7/12/2024
*
 */

type BrokerConfig struct {
	broker.IBrokersFactory
	*zap.Logger
}

type IBrokerConfig interface {
	BrokerSubscriber()
	BrokerPublisher(data []byte) error
}

func NewBrokerConfig(log *zap.Logger) *BrokerConfig {
	return &BrokerConfig{
		Logger: log,
	}
}

func (b *BrokerConfig) NewBrokerServer(instance int) *BrokerConfig {
	factory, err := broker.NewBrokerFactory(b.Logger, instance)
	if err != nil {
		return nil
	}
	b.IBrokersFactory = factory
	return b
}

func (b *BrokerConfig) BrokerSubscriber() {
	b.Subscribe()
}

func (b *BrokerConfig) BrokerPublisher(data []byte) error {
	err := b.Publish(data)
	if err != nil {
		b.Error("Error publishing", zap.Error(err))
		return err
	}
	return nil
}
