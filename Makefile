DB_CONTAINER=pokemondb
DB_PASSWORD=1234
DB_USER=postgres
DB_NAME = postgres

GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable"
GOOSE_MIGRATION_DIR=./schema

# Initial call to migrate all tables
initdb: 
	@echo "starting docker container"
	docker-compose up -d
	@echo "initializing database schema"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } up
	@echo "initialization finished"
	
startPsql: 
	docker exec -ti ${DB_CONTAINER} psql -U ${DB_USER}

build:
	docker-compose build

run:
	docker-compose up -d

migrateup:
	@echo "migrating up"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } up
	@echo "migrating finished"

migratedown:
	@echo "migrating down"
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR } down
	@echo "migrating finished"