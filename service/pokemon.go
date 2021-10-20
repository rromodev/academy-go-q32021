package service

import (
	"github.com/rromodev/academy-go-q32021/model"
)

type CSVPokemonGetter interface {
	GetRecordById(id int, rowType string) ([]string, error)
}

type PokemonService struct {
	pokemonData CSVPokemonGetter
}

func NewPokemonService(pokemonData CSVPokemonGetter) PokemonService {
	return PokemonService{pokemonData: pokemonData}
}

const POKEMON string = "pokemon"

func (ps PokemonService) GetPokemonById(id int) (*model.Pokemon, error) {

	record, err := ps.pokemonData.GetRecordById(id, POKEMON)

	if err != nil {
		return nil, err
	}
	pokemon := model.Pokemon{
		ID:   id,
		Name: record[2],
	}

	return &pokemon, nil
}
