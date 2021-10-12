package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/controllers"
	"github.com/rizabach29/todolist-go/middleware"
	"github.com/rizabach29/todolist-go/services"
)

func NewUserRouter(router *gin.Engine, services *services.Services) {
	userCtrl := controllers.NewUserController(services)

	userRoute := router.Group("/")
	userRoute.Use(middleware.AuthorizeJWT())
	{
		userRoute.POST("/user/:id/asignrole/:userId", userCtrl.AsignRole)
		userRoute.GET("/user", userCtrl.GetAll)
	}
}