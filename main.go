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
	db := app.InitDB()
	app.CreateSchema(db)

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	
	gin.ForceConsoleColor()
	server := gin.Default()
	
	routes.InitRouter(server, service)
	
	log.Fatal(server.Run())
}