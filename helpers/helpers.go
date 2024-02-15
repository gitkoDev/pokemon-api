package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/models"
)

func DecodeJson(httpReq *http.Request) (models.Pokemon, error) {
	pokemon := models.Pokemon{}

	err := json.NewDecoder(httpReq.Body).Decode(&pokemon)
	if err != nil {
		log.Println("AddPokemon() error decoding from json", err)
		return models.Pokemon{}, err
	}
	return pokemon, nil
}

func EncodeJson(writer http.ResponseWriter, pokemon models.Pokemon) error {
	err := json.NewEncoder(writer).Encode(pokemon)
	if err != nil {
		return err
	}
	return nil
}
