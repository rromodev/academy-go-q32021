package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rromodev/academy-go-q32021/model"
)

type UserServiceGetter interface {
	GetUserById(id int) (*model.User, error)
	StoreNewInfo() (string, error)
}

type UserController struct {
	userService UserServiceGetter
}

func NewUserController(userService UserServiceGetter) UserController {
	return UserController{userService}
}

func (uc UserController) GetUser(c *gin.Context) {
	idp := c.Param("id")
	id, err := strconv.Atoi(idp)

	if err != nil || id == 0 {
		c.IndentedJSON(http.StatusBadRequest, "{}")
		return
	} else {
		user, err := uc.userService.GetUserById(id)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusNotFound, "{}")
			return
		}
		c.IndentedJSON(http.StatusOK, user)
	}
}

func (uc UserController) StoreNewInfo(c *gin.Context) {
	status, err := uc.userService.StoreNewInfo()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "{}")
		return
	}
	c.IndentedJSON(http.StatusOK, status)
}
