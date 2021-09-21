package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pokemon struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var pokemones = []pokemon{
	{ID: "1", Name: "bulbasaur"},
	{ID: "2", Name: "ivysaur"},
}

func main() {
	router := gin.Default()
	// router.GET("/hello", hello)
	router.GET("/pokemon/:id", getPokemon)

	router.Run("localhost:8085")
}

func getPokemon(c *gin.Context) {
	id := c.Param("id")

	for _, a := range pokemones {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pokemon not found"})
}
