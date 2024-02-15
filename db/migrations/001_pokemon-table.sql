-- +goose Up
CREATE TABLE IF NOT EXISTS pokemon (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR (100)[] ,
    hp INTEGER,
    attack INTEGER,
    defense INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS pokemon;