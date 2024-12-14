package database

import (
	"database/sql"
	"github.com/dbacilio88/golang-grpc-email-microservice/internal/repository"
	"github.com/dbacilio88/golang-grpc-email-microservice/pkg/yaml"
	_ "github.com/lib/pq"
	"log"
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
}

type IConnection interface {
}

func NewConnection() *Connection {
	open, err := sql.Open(yaml.YAML.Database.Driver, yaml.GetUriDatasource())
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return nil
	}
	return &Connection{
		repository.NewStore(open),
	}
}
