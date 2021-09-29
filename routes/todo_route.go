package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/controllers"
	"github.com/rizabach29/todolist-go/middleware"

	"github.com/rizabach29/todolist-go/services"
)

func NewTodoRouter(router *gin.Engine, services *services.Services) {
	todoController := controllers.NewTodoController(services)

	todoRouter := router.Group("/todo")
	todoRouter.Use(middleware.AuthorizeJWT())
	{
		todoRouter.GET("/", todoController.GetAll)
		todoRouter.GET("/:id", todoController.GetById)
		todoRouter.POST("", todoController.Create)
		todoRouter.PUT("/:id", todoController.Update)
		todoRouter.DELETE("/:id", todoController.Delete)
	}
}