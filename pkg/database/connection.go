package database

import (
	"database/sql"
	"github.com/dbacilio88/poc-golang-grpc-microservice/internal/repository"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

/**
*
* connection
* <p>
* connection file
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
* @since 14/12/2024
*
 */

type Connection struct {
	*repository.Store
	*zap.Logger
}

type IConnection interface {
}

func NewConnection(log *zap.Logger) *Connection {
	open, err := sql.Open(env.YAML.Database.Driver, env.YAML.Database.GetUrl())
	if err != nil {
		log.Error("Error connecting to database ", zap.Error(err))
		return nil
	}
	return &Connection{
		Store:  repository.NewStore(open),
		Logger: log,
	}
}
