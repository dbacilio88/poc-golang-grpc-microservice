package service

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/event/broker"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env"
	"go.uber.org/zap"
)

/**
 * messaging
 * <p>
 * This file contains core data structures and logic used throughout the application.
 *
 * <p><strong>Copyright © 2025 – All rights reserved.</strong></p>
 *
 * <p>This source code is distributed under a collaborative license.</p>
 *
 * <p>
 * Contributions, suggestions, and improvements are welcome!
 * You are free to fork, modify, and submit pull requests under the terms of the repository's license.
 * Please ensure proper attribution to the original author(s) and preserve this notice in derivative works.
 * </p>
 *
 * @author Christian Bacilio De La Cruz
 * @email dbacilio88@outlook.es
 * @since 5/9/2025
 */

type IMessaging interface {
	Receive(map[string]interface{})
	Send(payload []byte) error
}

type Messaging struct {
	env.Properties
	*zap.Logger
	broker.IBrokersFactory
}

func NewMessaging(log *zap.Logger, prop env.Properties) *Messaging {
	return &Messaging{
		Properties: prop,
		Logger:     log,
	}
}

func (p *Messaging) NewBrokerServer(instance int) *Messaging {
	factory, err := broker.NewBrokerFactory(p.Logger, instance)
	if err != nil {
		return nil
	}
	p.IBrokersFactory = factory
	return p
}

func (p *Messaging) Receive(params map[string]interface{}) {
	fmt.Println("Received message ", params)
	subscriber, err := p.Subscriber(params)
	if err != nil {
		return
	}

	msg := subscriber.(<-chan *message.Message)
	go func() {

		fmt.Println("Received message X", msg)
		for m := range msg {
			p.Info("Received message", zap.Any("message", m))
			m.Ack()
		}

	}()
}
func (p *Messaging) Send(payload []byte) error {
	params := map[string]interface{}{
		"binding": env.YAML.Rabbitmq.RoutingKey.Files,
		"data":    payload,
	}

	err := p.Publisher(params)
	if err != nil {
		return err
	}
	return nil
}
