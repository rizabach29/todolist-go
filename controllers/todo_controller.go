package controllers

import (
	"github.com/gin-gonic/gin"
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
	
}

func (ctrl *TodoController) GetAll(c *gin.Context) {

}

func (ctrl *TodoController) GetById(c *gin.Context) {

}

func (ctrl *TodoController) Update(c *gin.Context) {

}

func (ctrl *TodoController) Delete(c *gin.Context) {

}