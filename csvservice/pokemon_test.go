package csvservice_test

import (
	"errors"
	"testing"

	"github.com/rromodev/academy-go-q32021/csvservice"
	"github.com/rromodev/academy-go-q32021/models"
)

func TestGetPokemonByIdTableDriven(t *testing.T) {
	bulbasaur := models.Pokemon{ID: 1, Name: "bulbasaur"}

	type test struct {
		name    string
		id      int
		pokemon models.Pokemon
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
			service := csvservice.NewPokemonService(tt.path)
			ans, err := service.GetPokemonById(tt.id)

			if err != nil {
				t.Error(err)
			}
			if ans.Name != tt.pokemon.Name {
				t.Errorf("got %q, want %q", ans.Name, tt.pokemon.Name)
			}
		})
	}
}
