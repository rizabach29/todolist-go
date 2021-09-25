package routes

import "github.com/gin-gonic/gin"

func NewTodoRouter(router *gin.Engine) {
	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("/")
		todoRouter.GET("/:id")
		todoRouter.POST("")
		todoRouter.PUT("/:id")
		todoRouter.DELETE("/:id")
	}
}