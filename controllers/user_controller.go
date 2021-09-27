package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/services"
)


type IUserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type UserController struct {
	services *services.Services
}

func NewUserController(services *services.Services) IUserController {
	return &UserController{
		services: services,
	}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var newUser models.RegisterModel

	if isErr := helpers.PanicBindJSON(c, &newUser); isErr {
		return
	}

	// registering new account
	if isErr := ctrl.services.AuthService.Register(newUser); isErr {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl *UserController) Login(c *gin.Context) {
	
}