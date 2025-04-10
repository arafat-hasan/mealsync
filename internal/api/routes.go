package api

import (
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize services
	authService := service.NewAuthService(db)
	menuService := service.NewMenuService(db)

	// Initialize handlers
	authHandler := NewAuthHandler(authService)
	menuHandler := NewMenuHandler(menuService)

	// Public routes
	public := r.Group("/api")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Menu routes for employees
		protected.GET("/menu", menuHandler.GetMenuItems)
		protected.POST("/meal-request", menuHandler.CreateMealRequest)
	}

	// Admin routes
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		// Menu management
		admin.POST("/menu", menuHandler.CreateMenuItem)
		admin.PUT("/menu/:id", menuHandler.UpdateMenuItem)
		admin.DELETE("/menu/:id", menuHandler.DeleteMenuItem)

		// Statistics
		admin.GET("/meal-requests/stats", menuHandler.GetMealRequestStats)
	}
}
