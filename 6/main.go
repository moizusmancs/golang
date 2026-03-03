package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moizusmancs/internal/auth"
	internal "github.com/moizusmancs/internal/db"
	"github.com/moizusmancs/internal/middleware"
)


func main(){

	// db initialization 
	database := internal.InitializePostgre()

	// migrations
	if err := auth.AutoMigrate(database); err != nil {
		log.Fatalf("❌ Failed to migrate auth module: %v", err)
	}

	// gin initialization
	r := gin.Default()

	// dependancy injection
	repo := auth.NewRepository(database)


	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	api := r.Group("/api/v1")

	// public routes
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	//protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			c.JSON(200, gin.H{
				"message": "Protected route",
				"user_id": userID,
			})
		})
	}

	// server startup
	log.Println("✅ Starting server on port :8080")
	if err := r.Run(":8080"); err != nil {
    	log.Fatalf("❌ Failed to run server: %v", err)
	}
}