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
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user exists"})
		return
	}

	// registering new account
	ctrl.services.UserService.Register(newUser)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var userLogin models.LoginModel

	if isErr := helpers.PanicBindJSON(c, &userLogin); isErr {
		return
	}	
}