package main

import (
	"github.com/rromodev/academy-go-q32021/controller"
	"github.com/rromodev/academy-go-q32021/data"
	"github.com/rromodev/academy-go-q32021/router"
	"github.com/rromodev/academy-go-q32021/service"
)

const URL = "localhost:8085"
const filePath = "./data/data.csv"
const externalApi = "https://reqres.in/api/users/2"

func main() {

	csvData := data.NewCSVData(filePath)
	externalData := data.NewExternalData(externalApi)

	pokemonService := service.NewPokemonService(csvData)
	pokemonController := controller.NewPokemonController(pokemonService)

	userService := service.NewUserService(csvData, externalData)
	userController := controller.NewUserController(userService)

	workerService := service.NewWorkerService()
	workerController := controller.NewWorkerController(workerService)

	ginServer := router.NewRouter(
		pokemonController, 
		userController,
		workerController
	).SetRoutes()

	ginServer.Run(URL)
}
