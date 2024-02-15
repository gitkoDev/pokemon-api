package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/helpers"
	"github.com/gitkoDev/pokemon-db/models"
	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

func AddPokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode pokemon data from json

		pokemon, err := helpers.DecodeJson(r)
		if err != nil {
			log.Println("decodeJson() error:", err)
			return
		}

		isExisting := checkForExistence(db, pokemon.Name)
		if isExisting {
			responseString := fmt.Sprintf("%s already exists", pokemon.Name)
			w.Write([]byte(responseString))
			return
		}

		// Insert pokemon into database
		query := `INSERT INTO pokemon (name, type, hp, attack, defence) VALUES($1, $2, $3, $4, $5)`
		_, err = db.Exec(query, pokemon.Name, pokemon.PokemonType, pokemon.Hp, pokemon.Attack, pokemon.Defence)
		if err != nil {
			log.Println("InsertPokemon() error", err)
			return
		}

		responseString := fmt.Sprintf("%s added", pokemon.Name)
		w.Write([]byte(responseString))

	}

}

func GetAll(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get pokemon from database
		query := `SELECT name, type, hp, attack, defence FROM pokemon`

		res, err := db.Query(query)
		if err != nil {
			log.Println("GetAll() error:", err)
			return
		}

		defer res.Close()

		selectedPokemon := []models.Pokemon{}

		for res.Next() {
			pokemon := models.Pokemon{}

			err := res.Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defence)
			if err != nil {
				log.Println("GetAll() error scanning row:", err)
				return
			}

			selectedPokemon = append(selectedPokemon, pokemon)
		}

		// Respond with json
		helpers.EncodeJson(w, selectedPokemon)
	}
}

func GetByName(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pokemon := models.Pokemon{}

		name := chi.URLParam(r, "name")

		query := `SELECT name, type, hp, attack, defence FROM pokemon WHERE name = $1`

		err := db.QueryRow(query, name).Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defence)
		if err != nil {
			if err == sql.ErrNoRows {
				responseString := fmt.Sprintf("%s not found", name)
				w.Write([]byte(responseString))
				return
			}
			fmt.Println("GetByName() error:", err)
			return
		}

		// Print pokemon data if found
		err = helpers.EncodeJson(w, pokemon)
		if err != nil {
			log.Println("encodeJson() error:", err)
			return
		}

		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Type:", pokemon.PokemonType)
		fmt.Println("Hp:", pokemon.Hp)
		fmt.Println("Attack:", pokemon.Attack)
		fmt.Println("Defence:", pokemon.Defence)

	}
}

func DeletePokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		name := chi.URLParam(r, "name")

		isExisting := checkForExistence(db, name)
		if !isExisting {
			responseString := fmt.Sprintf("%s not found", name)
			w.Write([]byte(responseString))
			return
		}

		query := `DELETE FROM pokemon WHERE name = $1`

		_, err := db.Exec(query, name)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("%s not found\n", name)
				return
			}
			fmt.Println(err)
			return
		}

		responseString := fmt.Sprintf("%s deleted", name)
		w.Write([]byte(responseString))
	}
}

func checkForExistence(db *sql.DB, name string) bool {
	query := `SELECT id FROM pokemon WHERE name = $1`
	var id uint

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println("checkForExistence() error:", err)
	}

	return true
}
