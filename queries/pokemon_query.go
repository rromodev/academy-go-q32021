package queries

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/rromodev/academy-go-q32021/models"
)

func GetPokemonById(id int) models.Pokemon {
	pokemon := models.Pokemon{}

	file, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Println(records)
	var pokemones []models.Pokemon
	for _, record := range records {
		idPokemon, _ := strconv.Atoi(record[0])
		data := models.Pokemon{
			ID:   idPokemon,
			Name: record[1],
		}
		pokemones = append(pokemones, data)
	}

	for _, a := range pokemones {
		if a.ID == id {
			pokemon = a
		}
	}

	return pokemon
}
