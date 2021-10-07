package router

import (
	"github.com/gin-gonic/gin"
)

type PokemonRouter struct {
	pokemonController PokemonController
}

type PokemonController interface {
	GetPokemon(c *gin.Context)
}

func NewRouter(pokemonController PokemonController) PokemonRouter {
	return PokemonRouter{pokemonController}
}

func (pc PokemonRouter) SetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/pokemon/:id", pc.pokemonController.GetPokemon)
	return router
}
