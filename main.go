package main

import (
	"main/controllers"
	"main/database"
	"main/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()

	notesController := &controllers.NotesController{
		NotesService: services.NewNotesService(),
	}

	notesController.InitNotesControllerRoutes(router)
	router.Run(":8081")
}
