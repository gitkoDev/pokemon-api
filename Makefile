DOCKER_CONTAINER=pokemon_db
BINARY_NAME=pokemonapi
DB_PASSWORD=1234
DB_USER=pokemon_db
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_USER}?sslmode=disable"
GOOSE_MIGRATION_DIR=./db/migrations

postgres:
	docker run --name ${DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_PASSWORD=1234 -e POSTGRES_USER=${DOCKER_CONTAINER} -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres

initdb: 
	@echo "initializing database schema"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } up
	@echo "initialization finished"

startdb: start
	@echo "starting database"
	docker exec -ti ${DOCKER_CONTAINER} psql -U ${DOCKER_CONTAINER}

start: 
	@echo "starting container"
	docker start ${DOCKER_CONTAINER}

stop:
	docker stop ${DOCKER_CONTAINER}
	kill -9 [$(lsof -t -i tcp:8080)]
	@echo "container stopped"

run: 
	@echo "starting database"
	docker start ${DOCKER_CONTAINER}
	go run main.go
	docker exec -ti ${DOCKER_CONTAINER} psql -U ${DOCKER_CONTAINER}


migrate-up:
	@echo "migrating up"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } up
	@echo "migrating finished"

migrate-down:
	@echo "migrating down"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } down
	@echo "migrating finished"

stop_delete:
	docker stop pokemon_db
	docker rm pokemon_db

delete:
	docker rm pokemon_db
