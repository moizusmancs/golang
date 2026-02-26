package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moizusmancs/students-api/controllers"
  "github.com/moizusmancs/students-api/internal/database"

)

func main() {
	r := gin.Default()
  startDb := internal.InitPostgre()

  if startDb == nil{
    // handle the error
  }

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/me/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"Id": id,
		})
	})

	r.POST("/me", func(c *gin.Context) {

		type User struct {
			Fullname string `json:"fullname" binding:"required,min=4,max=20"` 
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8,max=24"`
		}

		var givenUser User

		if err:=c.BindJSON(&givenUser); err != nil{
      c.JSON(http.StatusBadRequest, gin.H{
        "success" : false,
        "error" : err.Error(),
		})
    return
    }

		c.JSON(http.StatusOK, gin.H{
			"full-name": givenUser.Fullname,
      "email" : givenUser.Email,
      "password" : givenUser.Password,
		})

	})

  notesController := controllers.NotesController{}
  notesController.InitNotesControllerRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	} else {
		fmt.Println("Server started on port 8080")
	}
}
