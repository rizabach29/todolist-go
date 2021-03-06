package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/services"
)

type ITodolistController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type TodolistController struct {
	Services *services.Services
}

func NewTodolistController(services *services.Services) ITodolistController {
	return &TodolistController{services}
}

func (ctrl *TodolistController) Create(c *gin.Context) {
	var newTodolist models.CreateTodolistModel
	if err := helpers.PanicBindJSON(c, &newTodolist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	todolist := ctrl.Services.TodolistService.Create(newTodolist)
	c.JSON(http.StatusOK, gin.H{"message": "todolist created", "data": todolist})
}

func (ctrl *TodolistController) GetAll(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("todoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "todo id required"})
	}

	result, _ := ctrl.Services.TodolistService.GetByTodoId(id)
	c.JSON(http.StatusOK, result)
}

func (ctrl *TodolistController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id required"})
	}
	
	todolist, err := ctrl.Services.TodolistService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todolist)
}

func (ctrl *TodolistController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TodolistController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id required"})
	}
	
	err = ctrl.Services.TodolistService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "todolist deleted"})
}

