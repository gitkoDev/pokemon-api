package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/gitkoDev/pokemon-db/models"
	log "github.com/sirupsen/logrus"
)

type ErrResponse struct {
	Message string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrResponse{Message: err})
}

func DecodeTrainerJSON(httpReq *http.Request) (models.Trainer, error) {
	trainer := models.Trainer{}

	decoder := json.NewDecoder(httpReq.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&trainer); err != nil {
		log.Println(err)
		return models.Trainer{}, err
	}

	return trainer, nil
}

func DecodePokemonJSON(httpReq *http.Request) (models.Pokemon, error) {
	pokemon := models.Pokemon{}

	err := json.NewDecoder(httpReq.Body).Decode(&pokemon)
	if err != nil {
		log.Println(err)
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
