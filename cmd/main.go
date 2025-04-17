package main

import (
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/arafat-hasan/mealsync/docs"
	"github.com/arafat-hasan/mealsync/internal/api"
	"github.com/arafat-hasan/mealsync/internal/config"
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/arafat-hasan/mealsync/internal/repository"
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

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	mealEventRepo := repository.NewMealEventRepository(db)
	menuSetRepo := repository.NewMenuSetRepository(db)
	menuItemRepo := repository.NewMenuItemRepository(db)
	mealRequestRepo := repository.NewMealRequestRepository(db)
	mealCommentRepo := repository.NewMealCommentRepository(db)
	eventAddressRepo := repository.NewEventAddressRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)

	// Initialize services
	authService := service.NewAuthService(db)
	notificationService := service.NewNotificationService(notificationRepo, userRepo)
	mealEventService := service.NewMealEventService(
		mealEventRepo,
		userRepo,
		menuSetRepo,
		eventAddressRepo,
		mealRequestRepo,
		mealCommentRepo,
		notificationService,
	)
	menuSetService := service.NewMenuSetService(
		menuSetRepo,
		menuItemRepo,
		userRepo,
	)
	menuItemService := service.NewMenuItemService(
		menuItemRepo,
		userRepo,
	)
	mealRequestService := service.NewMealRequestService(
		mealRequestRepo,
		mealEventRepo,
		userRepo,
	)
	mealCommentService := service.NewMealCommentService(
		mealCommentRepo,
		mealEventRepo,
		userRepo,
	)

	// Initialize handlers
	authHandler := api.NewAuthHandler(authService)
	mealEventHandler := api.NewMealEventHandler(mealEventService)
	menuSetHandler := api.NewMenuSetHandler(menuSetService)
	menuItemHandler := api.NewMenuItemHandler(menuItemService)
	mealRequestHandler := api.NewMealRequestHandler(mealRequestService)
	mealCommentHandler := api.NewMealCommentHandler(mealCommentService)

	// Initialize router with custom middleware
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())              // Add logging
	router.Use(middleware.Recovery())     // Add custom recovery middleware
	router.Use(middleware.ErrorHandler()) // Add custom error handling middleware
	router.Use(gin.Recovery())            // Add gin's recovery as a fallback

	// Load HTML templates from docs directory
	router.LoadHTMLGlob(filepath.Join("docs", "*.html"))

	// API routes
	api.SetupRoutes(router, authHandler, mealEventHandler, menuSetHandler, mealCommentHandler, menuItemHandler, mealRequestHandler)

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
