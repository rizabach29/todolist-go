package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/services"
)

// initialize all router
func InitRouter(router *gin.Engine, services *services.Services) {
	NewTodolistRouter(router,services)
	NewAuthRouter(router, services)
	NewTodoRouter(router, services)
	NewUserRouter(router, services)
}