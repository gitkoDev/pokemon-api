package models

type Pokemon struct {
	Name        string   `json:"name"`
	PokemonType []string `json:"type"`
	Hp          uint     `json:"hp"`
	Attack      uint     `json:"attack"`
	Defence     uint     `json:"defense"`
}
