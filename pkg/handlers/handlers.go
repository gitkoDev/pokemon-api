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
		responseString := fmt.Sprintf("%d pokemon in database", len(selectedPokemon))
		w.Write([]byte(responseString))
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
				fmt.Printf("No pokemon with name %s found\n", name)
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

		query := `DELETE FROM pokemon WHERE name = $1`

		_, err := db.Exec(query, name)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("No pokemon with name %s found\n", name)
				return
			}
			fmt.Println(err)
			return
		}

		fmt.Printf("Pokemon %s deleted\n", name)
	}
}
