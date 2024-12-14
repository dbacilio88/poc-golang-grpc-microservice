package broker

import (
	"errors"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/yaml"
	"go.uber.org/zap"
)

/**
*
* factory
* <p>
* factory file
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

const RabbitMqInstance int = iota
const KafkaMqInstance = 1

type BrokersFactory struct{}

type IBrokersFactory interface {
	Subscribe()
	Publish(data []byte) error
	LoadConfiguration()
}

func NewBrokerFactory(log *zap.Logger, instance int) (IBrokersFactory, error) {
	switch instance {
	case RabbitMqInstance:
		return NewRabbitMq(log, &yaml.YAML.Rabbitmq), nil
	case KafkaMqInstance:
		return nil, errors.New("kafka no implementado aún")
	default:
		return nil, errors.New("invalid instance")
	}
}
