package controllers

import (
	"net/http"
	"strconv"

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
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id required"})
		return
	}

	roleId, err := strconv.Atoi(c.Param("RoleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "role id required"})
		return
	}
	
	res := ctrl.Services.UserService.UpdateRole(id, roleId)
	if res {
		c.JSON(http.StatusOK, gin.H{"message": "update role success"})
		return
	}
	
	c.JSON(http.StatusBadRequest, gin.H{"message": "update role failed"})
}

