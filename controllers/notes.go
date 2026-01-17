package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	NotesService *services.Notesservice
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	notes.DELETE("/:id", n.DeleteNotes())

}
func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.NotesService.GetNotes(),
		})
	}
}
func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var note services.Note
		err := c.BindJSON(&note)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Something wrong",
			})
			return
		}
		created := n.NotesService.CreateNotes(note)

		c.JSON(http.StatusOK, gin.H{
			"note": created,
		})
	}
}
func (n *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		deletedNote := n.NotesService.DeleteNotes(id)
		if deletedNote.ID.IsZero() {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Note not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Note deleted successfully",
		})
	}
}
