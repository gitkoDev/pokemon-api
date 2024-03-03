-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pokemon (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR (100)[] ,
    hp INTEGER,
    attack INTEGER,
    defense INTEGER,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pokemon_trainers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR (255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS pokemon;
DROP TABLE IF EXISTS pokemon_trainers;