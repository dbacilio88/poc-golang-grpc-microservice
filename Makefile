update:
	go get -u all

postgres:
	docker run --name go-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.15-alpine3.21

docker_stop:
	docker container stop go-postgres

docker_delete:
	docker container rm go-postgres

create_db:
	docker exec -it go-postgres createdb --username=root --owner=root postgres

drop_db:
	docker exec -it go-postgres dropdb postgres

migrate_up:
	 migrate -path workspace/database/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose up

migrate_down:
	 migrate -path workspace/database/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose down -all

sqlc:
	sqlc generate -f workspace/database/sqlc.yaml

test:
	go test -v -cover ./...

.PHONY: update postgres create_db drop_db migrate_up migrate_down test
