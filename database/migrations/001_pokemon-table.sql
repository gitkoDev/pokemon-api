-- +goose Up
CREATE TABLE IF EXISTS pokemon (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR (100)[] ,
    hp INTEGER,
    attack INTEGER,
    defence INTEGER
);

-- +goose Down
DROP TABLE IF EXISTS pokemon;