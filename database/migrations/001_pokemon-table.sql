-- +goose Up
CREATE TABLE pokemon (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR (100)[] ,
    hp INTEGER,
    attack INTEGER,
    defence INTEGER
);

-- +goose Down
DROP TABLE pokemon;