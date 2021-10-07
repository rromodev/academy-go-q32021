package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rromodev/academy-go-q32021/models"
)

type CSVGetter interface {
	GetPokemonById(id int) (*models.Pokemon, error)
}

type PokemonController struct {
	pokemonService CSVGetter
}

func NewPokemonController(pokemonService CSVGetter) PokemonController {
	return PokemonController{pokemonService}
}

func (ps PokemonController) GetPokemon(c *gin.Context) {
	idp := c.Param("id")
	id, err := strconv.Atoi(idp)

	if err != nil || id == 0 {
		c.IndentedJSON(http.StatusBadRequest, "{}")
		return
	} else {
		pokemon, err := ps.pokemonService.GetPokemonById(id)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusNotFound, "{}")
			return
		}
		c.IndentedJSON(http.StatusOK, pokemon)
	}
}
