package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/services"
)

type IUserController interface {
	GetAll(c *gin.Context)
	AsignRole(c *gin.Context)
}

type UserController struct {
	Services *services.Services
}

func NewUserController(services *services.Services) IUserController {
	return &UserController{services}
}

func (ctrl *UserController) GetAll(c *gin.Context) {
	if !helpers.IsRoleAllowed(c, "admin") {return}
	users := ctrl.Services.UserService.GetAll()
	
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) AsignRole(c *gin.Context) {
	if !helpers.IsRoleAllowed(c, "admin") {return}
}

