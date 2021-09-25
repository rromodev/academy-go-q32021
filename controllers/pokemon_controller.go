package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rromodev/academy-go-q32021/queries"

	"github.com/gin-gonic/gin"
)

func GetPokemon(c *gin.Context) {
	idp := c.Param("id")
	id, _ := strconv.Atoi(idp)
	fmt.Println(id)
	pokemon := queries.GetPokemonById(id)
	c.IndentedJSON(http.StatusOK, pokemon)
}
