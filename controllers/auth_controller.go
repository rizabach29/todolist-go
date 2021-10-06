package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/services"
)


type IAuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}


type AuthController struct {
	services *services.Services
}


func NewAuthController(services *services.Services) IAuthController {
	return &AuthController{
		services: services,
	}
}


func (ctrl *AuthController) Register(c *gin.Context) {
	var newUser models.RegisterModel

	if err := helpers.PanicBindJSON(c, &newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// registering new account
	if err := ctrl.services.AuthService.Register(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}


func (ctrl *AuthController) Login(c *gin.Context) {
	var loginUser models.LoginModel

	if err := helpers.PanicBindJSON(c, &loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := ctrl.services.AuthService.Login(loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type": "Bearer",
		"token": token,
	})
}