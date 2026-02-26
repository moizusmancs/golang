package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesController struct {

}

func (n *NotesController) InitNotesControllerRoutes (router *gin.Engine) {
	r := router.Group("/notes")

	r.GET("/", n.handleGetNotes())
	r.POST("/", n.handlePostNewNote())
}

func (n *NotesController) handleGetNotes() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"notes" : "i am all notes",
		})
	}
}

func (n *NotesController) handlePostNewNote() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"note" : "i am new note",
		})
	}
}