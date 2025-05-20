package broker

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/event/consumer"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/event/producer"
	mq2 "github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env/mq"
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
	amqp.Config
	*zap.Logger
	mq2.Rabbitmq
	subscriber consumer.IBrokerSubscriber
	publisher  producer.IBrokerPublisher
}

func NewRabbitMq(log *zap.Logger, prop mq2.Rabbitmq) *RabbitMq {
	return &RabbitMq{
		Logger:     log,
		Rabbitmq:   prop,
		subscriber: consumer.NewBrokerSubscriber(log, prop),
		publisher:  producer.NewBrokerPublisher(log),
	}
}

func (r *RabbitMq) Publisher(params map[string]interface{}) error {

	r.Info("Publishing RabbitMq ...")

	routingKey, err := getParameter[string](params, "binding")
	fmt.Println("routingKey", routingKey)
	if err != nil {
		return err
	}

	data, err := getParameter[[]byte](params, "data")
	fmt.Println("data", data)
	if err != nil {
		return err
	}

	pub, err := r.publisher.PublisherRabbitMq(r.loadConfigurationPublisher())

	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewULID(), data)
	err = pub.Publish(routingKey, msg)
	if err != nil {
		return err
	}
	r.Info("Published RabbitMq")
	return nil
}

func (r *RabbitMq) Subscriber(params map[string]interface{}) (interface{}, error) {
	r.Info("Subscribing to RabbitMq")

	queueName, err := getParameter[string](params, "queueName")
	fmt.Println("queueName", queueName)
	if err != nil {
		return nil, err
	}

	routingKey, err := getParameter[string](params, "routingKey")
	fmt.Println("routingKey", routingKey)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	mq, err := r.subscriber.SubscriberRabbitMq(r.loadConfigurationSubscriber(queueName, routingKey))

	if err != nil {
		return nil, err
	}

	return mq.Subscribe(ctx, "LOAD.TRANSACTION.FILES")

	/*
		if err != nil {
			r.Error("Error subscribing to RabbitMq", zap.Error(err))
			return nil, err
		}

		service, err := r.SubscriberService(ctx, r.Rabbitmq.GetTopic().Email)
		if err != nil {
			r.Error("Error creating service: ", zap.Error(err))
			return
		}

		//chanel to receiving message
		msgChannel := make(chan *message.Message)

		go func() {
			for msg := range service {
				msgChannel <- msg
			}
		}()

		// process to messages
		for msg := range msgChannel {
			r.Info("Received message from RabbitMq", zap.String("message", string(msg.Payload)))
			msg.Ack()
		}

		r.Info("Process broker rabbitmq started")

	*/
}

func getParameter[T any](params map[string]interface{}, key string) (T, error) {
	val, ok := params[key]
	if !ok {
		var zero T
		return zero, fmt.Errorf("parameter '%s' is missing", key)
	}

	result, ok := val.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("parameter '%s' is not a string", key)
	}

	return result, nil
}

func (r *RabbitMq) loadConfigurationSubscriber(queueName, routingKey string) amqp.Config {
	fmt.Println("RabbitMq loadConfigurationSubscriber", queueName, routingKey)
	cfg := amqp.Config{
		Connection: amqp.ConnectionConfig{
			AmqpURI: r.Rabbitmq.GetUri(), // config ok
			AmqpConfig: &amqp091.Config{
				Vhost: r.Rabbitmq.Vhost, // config ok
			},
			Reconnect: amqp.DefaultReconnectConfig(),
		},

		Exchange: amqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				fmt.Println("topic exchange subscriber ", topic)
				//return r.Rabbitmq.Exchange.Name
				return "TOPIC.EXCHANGE.TRANSACTION"
			},
			//Binding * (star) can substitute for exactly one word.
			//Binding # (hash) can substitute for zero or more words.
			Type: "topic",
			//Type: r.Rabbitmq.Exchange.Types,

			Durable: r.Rabbitmq.Exchange.Durable,
			//AutoDeleted: false,
			//Internal:    false,
			//NoWait:      true,
			/*
				Arguments: map[string]interface{}{
					"alternative-exchange": "alt-exchange",
				},

			*/
		},

		Queue: amqp.QueueConfig{
			GenerateName: func(topic string) string {
				fmt.Println("topic queue subscriber ", topic)
				//return queueName // parameter ok
				return "QU-LOAD-FILES" // parameter ok
			},
			Durable:    true,
			Exclusive:  false,
			NoWait:     false,
			AutoDelete: false,
			Arguments:  nil,
			//		Arguments: map[string]interface{}{
			//			"x-message-ttl": 6000,
			//		"x-queue-type":  b.Rabbitmq.GetQueue().Email.Type,
			//},
		},

		QueueBind: amqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				// binding: example [*.transaction.*]
				// binding: example [*.data.#]
				// binding: example [*.*.validation]
				fmt.Println("topic queue binding subscriber ", topic)
				return "LOAD.TRANSACTION.FILES"
				//return routingKey
			},
		},
		Consume: amqp.ConsumeConfig{
			Qos: amqp.QosConfig{
				PrefetchCount: 1,
			},
		},
		Marshaler:       amqp.DefaultMarshaler{},
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
		//Publish:         b.loadPublishConfig(),
	}

	var ssl = false

	if r.Rabbitmq.TlsEnabled {
		ssl = true
	}

	cfg.Connection.TLSConfig = &tls.Config{
		InsecureSkipVerify: ssl,
	}

	r.Info("Configuration subscriber loaded")
	return cfg
}

func (r *RabbitMq) loadConfigurationPublisher() amqp.Config {
	cfg := amqp.Config{
		Connection: amqp.ConnectionConfig{
			AmqpURI: r.GetUri(), // config ok
			AmqpConfig: &amqp091.Config{
				Vhost: r.Vhost, // config ok
			},
			Reconnect: amqp.DefaultReconnectConfig(),
		},
		Exchange: amqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				fmt.Println("ExchangeConfig topic: ", topic)
				//return r.Rabbitmq.Exchange.Name // config ok
				return "TOPIC.EXCHANGE.TRANSACTION" // config ok
			},
			Type:    r.Rabbitmq.Exchange.Types, //config ok
			Durable: r.Rabbitmq.Exchange.Durable,
			//AutoDeleted: false,
			//Internal:    false,
			//NoWait:      true,
			/*
				Arguments: map[string]interface{}{
					"alternative-exchange": "alt-exchange",
				},

			*/
		},
		Publish: amqp.PublishConfig{
			GenerateRoutingKey: func(topic string) string {
				fmt.Println("PublishConfig topic: ", topic)
				return "LOAD.TRANSACTION.FILES"
			},
		},
		Marshaler: amqp.DefaultMarshaler{},
		//Queue:           b.loadQueueConfig(),
		//QueueBind:       b.loadQueueBindConfig(),
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
		//Consume:         b.loadConsumeConfig(),
	}

	var ssl = false

	if r.TlsEnabled {
		ssl = true
	}

	cfg.Connection.TLSConfig = &tls.Config{
		InsecureSkipVerify: ssl,
	}

	r.Info("Configuration publisher loaded")

	return cfg
}
