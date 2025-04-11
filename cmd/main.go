package main

import (
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/arafat-hasan/mealsync/docs"
	"github.com/arafat-hasan/mealsync/internal/api"
	"github.com/arafat-hasan/mealsync/internal/config"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           MealSync API
// @version         1.0
// @description     A meal management system API for employees
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @security BearerAuth
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize services
	authService := service.NewAuthService(db)
	mealService := service.NewMealService(db)

	// Initialize handlers
	authHandler := api.NewAuthHandler(authService)
	mealHandler := api.NewMealHandler(mealService)

	// Initialize router
	router := gin.Default()

	// Load HTML templates from docs directory
	router.LoadHTMLGlob(filepath.Join("docs", "*.html"))

	// API routes
	api.SetupRoutes(router, authHandler, mealHandler)

	// Documentation routes with custom configuration
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
		ginSwagger.DeepLinking(true),
		ginSwagger.DefaultModelsExpandDepth(-1), // Hide models section
		ginSwagger.DocExpansion("none"),         // Collapse all endpoints by default
	))

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
