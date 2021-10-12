package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rizabach29/todolist-go/app"
	"github.com/rizabach29/todolist-go/repositories"
	"github.com/rizabach29/todolist-go/routes"
	"github.com/rizabach29/todolist-go/services"
)

func main() {
	app.GetDatabase()
	app.CreateSchema()

	repository := repositories.NewRepository()
	service := services.NewService(repository)
	
	gin.ForceConsoleColor()
	server := gin.Default()
	
	routes.InitRouter(server, service)
	
	log.Fatal(server.Run())
}