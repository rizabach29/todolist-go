package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/services"
)

type ITodoController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type TodoController struct {
	services *services.Services
}

func NewTodoController(services *services.Services) ITodoController {
	return &TodoController{services}
}

func (ctrl *TodoController) Create(c *gin.Context) {
	var newTodo models.CreateTodoModel

	if err := helpers.PanicBindJSON(c, &newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	todo := ctrl.services.TodoService.Create(newTodo)

	c.JSON(http.StatusOK, gin.H{"message": "create todo success", "data": todo})
}

func (ctrl *TodoController) GetAll(c *gin.Context) {
	todos := ctrl.services.TodoService.GetAll()
	c.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := ctrl.services.TodoService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) Update(c *gin.Context) {
	var updatedTodo models.UpdateTodoModel

	if err := helpers.PanicBindJSON(c, &updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "todo updated"})
}

func (ctrl *TodoController) Delete(c *gin.Context) {
	
	c.JSON(http.StatusBadRequest, gin.H{"message": "todo deleted"})
}