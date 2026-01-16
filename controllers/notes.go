package controllers

import (
	"github.com/gin-gonic/gin"
	"main/services"
)
type NotesController struct {
	NotesService *services.Notesservice
}
func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine){
	notes:=router.Group("/notes")
	notes.GET("/",n.GetNotes())
	notes.POST("/",n.CreateNotes())

}
func(n *NotesController) GetNotes() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"notes":n.NotesService.GetNotes(),
		})
	}
}
func(n *NotesController) CreateNotes() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200,gin.H{
			"notes":n.NotesService.CreateNotes(),
		})
	}
}