package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/gitkoDev/pokemon-db/models"
	"github.com/sirupsen/logrus"
)

func DecodeJSON(httpReq *http.Request) (models.Pokemon, error) {
	pokemon := models.Pokemon{}

	err := json.NewDecoder(httpReq.Body).Decode(&pokemon)
	if err != nil {
		logrus.Println("AddPokemon() error decoding from json", err)
		return models.Pokemon{}, err
	}
	return pokemon, nil
}

func WriteJSON(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
