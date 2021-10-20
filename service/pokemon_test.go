package service

import (
	"errors"
	"testing"

	"github.com/rromodev/academy-go-q32021/model"
)

func TestGetPokemonByIdTableDriven(t *testing.T) {
	bulbasaur := model.Pokemon{ID: 1, Name: "bulbasaur"}

	type test struct {
		name    string
		id      int
		pokemon model.Pokemon
		err     error // bool
		path    string
	}

	tests := []test{
		{
			name:    "test de pokemon 1",
			id:      1,
			pokemon: bulbasaur,
			err:     errors.New("error"),
			path:    "./data.csv",
		},
		{
			name:    "test de pokemon 2",
			id:      2,
			pokemon: bulbasaur,
			err:     errors.New("error"),
			path:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := NewPokemonService(tt.path)
			ans, err := ps.GetPokemonById(tt.id)

			if err != nil {
				t.Error(err)
			}
			if ans.Name != tt.pokemon.Name {
				t.Errorf("got %q, want %q", ans.Name, tt.pokemon.Name)
			}
		})
	}
}
