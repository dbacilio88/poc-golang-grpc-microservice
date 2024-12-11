package producer

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

/**
*
* producer
* <p>
* producer file
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

type BrokerPublisher struct {
}

type IBrokerPublisher interface {
	PublisherRabbitMq() (*amqp.Publisher, error)
	PublisherKafkaMq() error
}

func NewBrokerPublisher() *BrokerPublisher {
	return &BrokerPublisher{}
}

func (b *BrokerPublisher) PublisherRabbitMq(cfg amqp.Config) (*amqp.Publisher, error) {
	pub, err := amqp.NewPublisher(cfg, watermill.NewStdLogger(false, false))
	if err != nil {
		fmt.Println("Error creating publisher", err)
		return nil, err
	}
	return pub, nil
}

func (b *BrokerPublisher) PublisherKafkaMq() error {
	fmt.Println("Kafka publishing is not implemented yet.")
	return fmt.Errorf("kafka publishing not implemented yet")
}
