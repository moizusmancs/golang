package main

import (
	"log"

	"github.com/gin-gonic/gin"
	internal "github.com/moizusmancs/internal/db"
)


func main(){

	internal.InitializePostgre()

	gin := gin.Default()

	log.Println("✅ Starting server on port :8080")
	if err := gin.Run(":8080"); err != nil {
    	log.Fatalf("❌ Failed to run server: %v", err)
	}
}