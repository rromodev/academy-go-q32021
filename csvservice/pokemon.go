package csvservice

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/rromodev/academy-go-q32021/models"
)

type PokemonService struct {
	filePath string
}

func NewPokemonService(filepath string) PokemonService {
	return PokemonService{filePath: filepath}
}

func (ps PokemonService) GetPokemonById(id int) (*models.Pokemon, error) {
	pokemon := models.Pokemon{}

	file, err := os.Open(ps.filePath)
	if err != nil {
		return nil, errors.New("pokemon csv not found")
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var pokemones []models.Pokemon
	for _, record := range records {
		idPokemon, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, errors.New("pokemon not found")
		}
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

	return &pokemon, nil
}
