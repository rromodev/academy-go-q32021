package main

import (
	"github.com/rromodev/academy-go-q32021/controller"
	"github.com/rromodev/academy-go-q32021/csvservice"
	"github.com/rromodev/academy-go-q32021/router"
)

const URL = "localhost:8085"
const filePath = "./data.csv"

func main() {

	pokemonService := csvservice.NewPokemonService(filePath)
	pokemonController := controller.NewPokemonController(pokemonService)

	ginServer := router.NewRouter(pokemonController).SetRoutes()

	ginServer.Run(URL)
}
