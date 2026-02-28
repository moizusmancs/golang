package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moizusmancs/students-api/services"
)

type NotesController struct {
	ns services.NoteService
}

func (n *NotesController) InitNotesControllerRoutes (router *gin.Engine, noteservice services.NoteService) {
	r := router.Group("/notes")

	r.GET("/", n.handleGetNotes())
	r.POST("/", n.handlePostNewNote())

	n.ns = noteservice
}

func (n *NotesController) handleGetNotes() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"notes" : n.ns.GetNotes(),
		})
	}
}

func (n *NotesController) handlePostNewNote() gin.HandlerFunc{
	return func(c *gin.Context){

		n.ns.CreateNote(1,"hello","pending")
		c.JSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "note created with body ",
			"note" : "i am new note",
		})
	}
}