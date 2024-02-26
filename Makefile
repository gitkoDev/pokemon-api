DOCKER_CONTAINER=pokemon_db
BINARY_NAME=pokemonapi
DB_PASSWORD=1234
DB_USER=pokemon_db

GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_USER}?sslmode=disable"
GOOSE_MIGRATION_DIR=./db/migrations

# Initial call to migrate all tables
initdb: 
	@echo "starting docker container"
	docker-compose up -d
	@echo "initializing database schema"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } up
	@echo "initialization finished"


reload: 
	docker-compose down
	./script.sh
	@echo "container stopped"
	docker-compose up -d
	go run cmd/main.go
	
startPsql: 
	docker exec -ti ${DOCKER_CONTAINER} psql -U ${DOCKER_CONTAINER}

run:
	docker-compose up -d
	go run cmd/main.go


migrate-up:
	@echo "migrating up"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } up
	@echo "migrating finished"

migrate-down:
	@echo "migrating down"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } down
	@echo "migrating finished"

