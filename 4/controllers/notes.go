package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moizusmancs/students-api/services"
)

type NotesController struct {
	ns services.NoteService
}

func (n *NotesController) InitNotesControllerRoutes (router *gin.Engine) {
	r := router.Group("/notes")

	r.GET("/", n.handleGetNotes())
	r.POST("/", n.handlePostNewNote())
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

		c.JSON(http.StatusOK, gin.H{
			"success" : true,
			"message" : "note created with body " + strconv.Itoa(n.ns.CreateNote(1,"hello")),
			"note" : "i am new note",
		})
	}
}