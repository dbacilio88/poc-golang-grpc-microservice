package broker

import (
	"context"
	"crypto/tls"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event/consumer"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/event/producer"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/yaml"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

/**
*
* rabbitmq
* <p>
* rabbitmq file
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

type RabbitMq struct {
	config     amqp.Config
	subscriber *consumer.BrokerSubscriber
	publisher  *producer.BrokerPublisher
	log        *zap.Logger
	parameters yaml.IParameterBroker
}

type IRabbitMq interface {
	SubscriberService() (<-chan *message.Message, error)
	PublisherService(topic string, data []byte) error
	Close() error
}

func NewRabbitMq(log *zap.Logger, parameters yaml.IParameterBroker) *RabbitMq {
	return &RabbitMq{
		log:        log,
		subscriber: consumer.NewBrokerSubscriber(),
		publisher:  producer.NewBrokerPublisher(),
		parameters: parameters,
	}
}

func (r *RabbitMq) SubscriberService(ctx context.Context, topic string) (<-chan *message.Message, error) {
	r.log.Info("Subscribing RabbitMq to topic...", zap.String("topic", topic))
	sub, err := r.subscriber.SubscriberRabbitMq(r.config)
	if err != nil {
		return nil, err
	}
	subscribe, err := sub.Subscribe(ctx, topic)
	if err != nil {
		r.log.Error("Error subscribing to RabbitMq", zap.Error(err))
		return nil, err
	}
	r.log.Info("RabbitMq subscribed", zap.String("topic", topic))
	return subscribe, nil
}

func (r *RabbitMq) PublisherService(topic string, data []byte) error {
	r.log.Info("Publishing RabbitMq to topic...", zap.String("topic", topic))
	pub, err := r.publisher.PublisherRabbitMq(r.config)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewULID(), data)
	err = pub.Publish(topic, msg)
	if err != nil {
		r.log.Error("Error publishing message in RabbitMq", zap.Error(err))
		return err
	}
	r.log.Info("RabbitMq published message", zap.String("topic", topic))
	return nil
}

func (r *RabbitMq) Subscribe() {
	r.LoadConfiguration()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service, err := r.SubscriberService(ctx, "service.app.go.transaction.request")
	if err != nil {
		r.log.Error("Error creating service: ", zap.Error(err))
		return
	}

	// Canal para recibir mensajes
	msgChannel := make(chan *message.Message)

	go func() {
		for msg := range service {
			msgChannel <- msg
		}
	}()

	// Procesamiento de los mensajes
	for msg := range msgChannel {
		r.log.Info("Received message from RabbitMq", zap.String("message", string(msg.Payload)))
		msg.Ack()
	}

	r.log.Info("Process broker rabbitmq started")
}

func (r *RabbitMq) Publish(data []byte) error {
	r.LoadConfiguration()
	err := r.PublisherService("service.app.go.transaction.request", data)
	if err != nil {
		r.log.Error("Error publishing message", zap.Error(err))
		return err
	}
	return nil
}

func (r *RabbitMq) LoadConfiguration() {
	r.log.Info("Loading configuration for RabbitMq...")
	cfg := amqp.Config{
		Connection:      r.loadConnectionConfig(),
		Exchange:        r.loadExchangeConfig(),
		Queue:           r.loadQueueConfig(),
		QueueBind:       r.loadQueueBindConfig(),
		Marshaler:       amqp.DefaultMarshaler{},
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
		Publish:         r.loadPublishConfig(),
		Consume:         r.loadConsumeConfig(),
	}

	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	cfg.Connection.TLSConfig = tlsCfg

	r.config = cfg
	r.log.Info("Configuration Loaded")
}

func (r *RabbitMq) loadConnectionConfig() amqp.ConnectionConfig {
	return amqp.ConnectionConfig{
		AmqpURI: r.parameters.GetUri(),
		AmqpConfig: &amqp091.Config{
			Vhost: r.parameters.GetVhost(),
		},
		Reconnect: amqp.DefaultReconnectConfig(),
	}
}

func (r *RabbitMq) loadExchangeConfig() amqp.ExchangeConfig {
	return amqp.ExchangeConfig{
		GenerateName: func(topic string) string {
			return r.parameters.GetExchange()
		},
		Type:    "topic",
		Durable: true,
		//AutoDeleted: false,
		//Internal:    false,
		//NoWait:      true,
		/*
			Arguments: map[string]interface{}{
				"alternative-exchange": "alt-exchange",
			},

		*/
	}
}

func (r *RabbitMq) loadQueueConfig() amqp.QueueConfig {
	return amqp.QueueConfig{
		GenerateName: func(topic string) string {
			return r.parameters.GetQueueName()
		},
		Durable: true,
		//Exclusive:  false,
		//NoWait:     true,
		AutoDelete: false,
		Arguments: map[string]interface{}{
			"x-message-ttl": 6000,
			"x-queue-type":  "quorum",
		},
	}
}

func (r *RabbitMq) loadQueueBindConfig() amqp.QueueBindConfig {
	return amqp.QueueBindConfig{
		GenerateRoutingKey: func(topic string) string {
			return r.parameters.GetRoutingKey()
		},
	}
}

func (r *RabbitMq) loadPublishConfig() amqp.PublishConfig {
	return amqp.PublishConfig{
		GenerateRoutingKey: func(topic string) string {
			return topic
		},
	}
}

func (r *RabbitMq) loadConsumeConfig() amqp.ConsumeConfig {
	return amqp.ConsumeConfig{
		Qos: amqp.QosConfig{
			PrefetchCount: 10,
		},
	}
}
