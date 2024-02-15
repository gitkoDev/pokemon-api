-- +goose Up

INSERT INTO pokemon (name, type, hp, attack, defense)
VALUES
('Bulbasaur', ARRAY ['Grass', 'Poison'], 45, 49, 49),
('Charmander', ARRAY ['Fire'], 39, 52, 43),
('Squirtle', ARRAY ['Water'], 44, 48, 65),
('Caterpie', ARRAY ['Bug'], 45, 30, 35),
('Metapod', ARRAY ['Bug'], 50, 20, 55),
('Weedle', ARRAY ['Bug', 'Poison'], 40, 35, 30),
('Kakuna', ARRAY ['Grass', 'Poison'], 45, 25, 50),
('Beedrill', ARRAY ['Normal', 'Flying'], 65, 150, 40),
('Pidgey', ARRAY ['Grass', 'Poison'], 40, 45, 40),
('Rattata', ARRAY ['Normal'], 30, 56, 35),
('Spearow', ARRAY ['Normal', 'Flying'], 40, 60, 30),
('Ekans', ARRAY ['Poison'], 35, 60, 44),
('Pikachu', ARRAY ['Electric'], 35, 55, 40),
('Sandshrew', ARRAY ['Ground'], 50, 75, 85)
;

-- +goose Down
DELETE FROM pokemon 
WHERE id >= 1 AND id <= 14;