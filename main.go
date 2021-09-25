package main

import (
	"github.com/rromodev/academy-go-q32021/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pokemon/:id", controllers.GetPokemon)

	router.Run("localhost:8085")
}
