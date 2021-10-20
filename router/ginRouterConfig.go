package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	pokemonController PokemonController
	userController    UserController
	workerController  WorkerController
}

type PokemonController interface {
	GetPokemon(c *gin.Context)
}

type UserController interface {
	GetUser(c *gin.Context)
	StoreNewInfo(c *gin.Context)
}

type WorkerController interface {
	Reader(c *gin.Context)
}

func NewRouter(pokemonController PokemonController,
	userController UserController,
	workerController WorkerController) Router {
	return Router{
		pokemonController,
		userController,
		workerController,
	}
}

func (r Router) SetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/pokemon/:id", r.pokemonController.GetPokemon)
	router.GET("/user/:id", r.userController.GetUser)
	router.GET("/user/storeNewInfo", r.userController.StoreNewInfo)
	router.GET("/worker/reader", r.workerController.Reader)
	return router
}
