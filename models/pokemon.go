package models

type Pokemon struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	PokemonType []string `json:"type"`
	Hp          uint     `json:"hp"`
	Attack      uint     `json:"attack"`
	Defense     uint     `json:"defense"`
}
