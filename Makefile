DOCKER_CONTAINER=pokemon_db
BINARY_NAME=pokemonapi
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres@localhost:5432/pokemon_db?sslmode=disable
GOOSE_MIGRATION_DIR=./db/migrations

postgres:
	docker run --name ${DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_PASSWORD=1234 -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres

createdb:
	docker exec -ti ${DOCKER_CONTAINER} psql -U postgres 

run:
	go run cmd/main.go

migrate-up:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR }up 

migrate-down:
	goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} ${GOOSE_DBSTRING} ${GOOSE_MIGRATION_DIR }down