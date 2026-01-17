package main

import (
	"main/controllers"
	"main/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	notesController := &controllers.NotesController{
		NotesService: services.NewNotesService(),
	}

	notesController.InitNotesControllerRoutes(router)
	router.Run(":8080")
}
