package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/arafat-hasan/mealsync/internal/api"
	"github.com/arafat-hasan/mealsync/internal/config"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// Initialize router
	r := gin.Default()

	// Setup routes
	api.SetupRoutes(r, db)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
