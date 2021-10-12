package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/controllers"
	"github.com/rizabach29/todolist-go/middleware"
	"github.com/rizabach29/todolist-go/services"
)

func NewUserRouter(router *gin.Engine, services *services.Services) {
	userCtrl := controllers.NewUserController(services)

	userRoute := router.Group("/user")
	userRoute.Use(middleware.AuthorizeJWT())
	{
		userRoute.POST("/:userId/assignrole/:roleId", userCtrl.AsignRole)
		userRoute.GET("/", userCtrl.GetAll)
	}
}