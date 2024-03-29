package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/gitkoDev/pokemon-api/models"
	log "github.com/sirupsen/logrus"
)

type Error struct {
	Msg    string
	Status int
}

type ErrResponseJSON struct {
	ErrMsg string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, receivedError error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(receivedError.Error())
	if err != nil {
		log.Println("error encoding json:", err)
	}
}

func DecodeAuthJSON(httpReq *http.Request) (models.SingInInput, error) {
	var input models.SingInInput

	decoder := json.NewDecoder(httpReq.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, err
	}

	return input, nil
}

func DecodeTrainerJSON(httpReq *http.Request) (models.Trainer, error) {
	trainer := models.Trainer{}

	decoder := json.NewDecoder(httpReq.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&trainer); err != nil {
		return trainer, err
	}

	return trainer, nil
}

func DecodePokemonJSON(httpReq *http.Request) (models.Pokemon, error) {
	pokemon := models.Pokemon{}

	err := json.NewDecoder(httpReq.Body).Decode(&pokemon)
	if err != nil {
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
