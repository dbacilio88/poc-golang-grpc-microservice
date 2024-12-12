update:
	go get -u all

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

create_bd:
	docker exec -it postgres12 createdb --username=root --owner=root name_database

drop_bd:
	docker exec -it postgres12 dropdb name_database

migrate_up:
	 migrate -path workspace/database/migration -database "postgresql://postgres:Hos5XoBdI2ps9V1alhCd@db-golang-postgresql.clck2ieqa8u0.us-east-1.rds.amazonaws.com:5432/postgres" -verbose up

migrate_down:
	 migrate -path workspace/database/migration -database "postgresql://postgres:Hos5XoBdI2ps9V1alhCd@db-golang-postgresql.clck2ieqa8u0.us-east-1.rds.amazonaws.com:5432/postgres" -verbose down -all

sqlc:
	sqlc generate -f workspace/database/sqlc.yaml

test:
	go test -v -cover ./...
.PHONY: update postgres create_bd drop_bd migrate_up migrate_down test
