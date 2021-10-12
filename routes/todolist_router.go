package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/controllers"
	"github.com/rizabach29/todolist-go/middleware"
	"github.com/rizabach29/todolist-go/services"
)

func NewTodolistRouter(router *gin.Engine, services *services.Services) {
	todolistCtrl := controllers.NewTodolistController(services)

	todolistRouter := router.Group("/todolist")
	todolistRouter.Use(middleware.AuthorizeJWT())
	{
		todolistRouter.POST("", todolistCtrl.Create)
		todolistRouter.GET("all/:todoId", todolistCtrl.GetAll)
		todolistRouter.GET("/:id", todolistCtrl.GetById)
		todolistRouter.PUT("/:id", todolistCtrl.Update)
		todolistRouter.DELETE("/:id", todolistCtrl.Delete)
	}
}