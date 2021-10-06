package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/controllers"
	"github.com/rizabach29/todolist-go/services"
)

func NewAuthRouter(router *gin.Engine, services *services.Services) {
	userCtrl := controllers.NewAuthController(services)

	userRoute := router.Group("/")
	{
		userRoute.POST("/login", userCtrl.Login)
		userRoute.POST("/register", userCtrl.Register)
	}
}