package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

/**
*
* main_test
* <p>
* main_test file
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
* @since 12/12/2024
*
 */

const (
	driverName      = "postgres"
	datasSourceName = "postgresql://postgres:Hos5XoBdI2ps9V1alhCd@db-golang-postgresql.clck2ieqa8u0.us-east-1.rds.amazonaws.com:5432/postgres"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(driverName, datasSourceName)
	if err != nil {
		log.Fatal("Error connecting to database:")
		return
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
