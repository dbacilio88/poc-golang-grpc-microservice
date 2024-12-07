package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	srv "github.com/dbacilio88/golang-grpc-email-microservice/internal/server"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/server/router"
	"github.com/rabbitmq/amqp091-go"
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

var routingKeySubscribe string = "service.*.srv.go.transaction.response"
var routingKeyPublish string = "service.app.go.transaction.request"
var queue string = "qu-topic-golang"
var exchange string = "exchange-topic-go"

func startPublisher(pub *amqp.Publisher, topic string, done chan bool) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			payload := []byte("message enviado a las # " + time.Now().Format(time.RFC3339))
			msg := message.NewMessage(watermill.NewUUID(), payload)
			err := pub.Publish(topic, msg)
			if err != nil {
				fmt.Println("Error publishing message")
				return
			} else {
				fmt.Printf("Published message %s\n", string(payload))
			}
		}
	}

}

func startSubscriber(sub *amqp.Subscriber, topic string) {
	messages, err := sub.Subscribe(context.Background(), topic)
	if err != nil {
		fmt.Println("Error subscribing to messages")
		return
	}

	for msg := range messages {
		fmt.Println("Received a message:", string(msg.Payload))
		msg.Ack()
	}
}

func main() {
	port := 3000

	sub := NewSubscriber()
	defer func(sub *amqp.Subscriber) {
		err := sub.Close()
		if err != nil {
			fmt.Println("Error closing subscriber")
		}
	}(sub)

	pub := NewPublisher()
	defer func(pub *amqp.Publisher) {
		err := pub.Close()
		if err != nil {
			fmt.Println("Error closing publisher")
		}
	}(pub)

	//
	done := make(chan bool)

	go startSubscriber(sub, routingKeyPublish)

	go startPublisher(pub, routingKeyPublish, done)
	//	<-done

	app := srv.NewHttpConfig().
		SetPort(port).
		SetName(router.NameRouterGin).
		NewHttpServer(router.InstanceRouterGin)

	app.Start()

	//time.Sleep(1 * time.Second)
	select {}
}

func CreateConnection() amqp.Config {

	cfg := amqp.Config{
		Connection: amqp.ConnectionConfig{
			AmqpURI: "amqps://kjsycjee:F6-m21vrQokbdCys_upLFsH2X2kfxl2J@chimpanzee.rmq.cloudamqp.com:5671/",
			AmqpConfig: &amqp091.Config{
				Vhost: "kjsycjee",
			},
			Reconnect: amqp.DefaultReconnectConfig(),
		},
		Exchange: amqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				fmt.Println("topic Exchange ", topic)
				return exchange
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
		},

		Queue: amqp.QueueConfig{
			GenerateName: func(topic string) string {
				fmt.Println("topic Queue ", topic)
				return queue
			},
			Durable:    true,
			AutoDelete: false,
			//Exclusive:  false,
			//NoWait:     true,
			Arguments: map[string]interface{}{
				"x-message-ttl": 6000,
				"x-queue-type":  "quorum",
			},
		},
		QueueBind: amqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				return routingKeyPublish
			},
			//NoWait: true,
		},
		Marshaler:       amqp.DefaultMarshaler{},
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
		Publish: amqp.PublishConfig{
			//ConfirmDelivery: true,
			GenerateRoutingKey: func(topic string) string {
				return topic
			},
		},
		Consume: amqp.ConsumeConfig{
			Qos: amqp.QosConfig{
				PrefetchCount: 10,
			},
			//NoWait: true,
		},
	}

	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	cfg.Connection.TLSConfig = tlsCfg

	return cfg
}

func NewPublisher() *amqp.Publisher {
	cfg := CreateConnection()
	publisher, err := amqp.NewPublisher(cfg, watermill.NewStdLogger(true, true))
	if err != nil {
		return nil
	}
	return publisher
}

func NewSubscriber() *amqp.Subscriber {
	cfg := CreateConnection()
	subscriber, err := amqp.NewSubscriber(cfg, watermill.NewStdLogger(true, true))
	if err != nil {
		return nil
	}
	return subscriber
}
