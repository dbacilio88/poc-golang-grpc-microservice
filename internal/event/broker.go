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
	factory broker.BrokersFactory
	log     *zap.Logger
}

func NewBrokerConfig(log *zap.Logger) *BrokerConfig {
	return &BrokerConfig{
		log: log,
	}
}

func (b *BrokerConfig) NewBrokerServer(instance int) *BrokerConfig {
	factory, err := broker.NewBrokerFactory(b.log, instance)
	if err != nil {
		return nil
	}
	b.factory = factory
	return b
}

func (b *BrokerConfig) Subscriber() {
	b.factory.Subscribe()
}

func (b *BrokerConfig) Publisher(data []byte) error {
	err := b.factory.Publish(data)
	if err != nil {
		b.log.Error("Error publishing", zap.Error(err))
		return err
	}
	return nil
}
