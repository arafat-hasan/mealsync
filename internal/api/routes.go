package api

import (
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, authHandler *AuthHandler, mealHandler *MealHandler) {
	// Public routes (no auth required)
	public := r.Group("/api")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Meal routes
		protected.GET("/meals", mealHandler.GetMeals)
		protected.GET("/meals/:id", mealHandler.GetMealByID)
		protected.POST("/meals", mealHandler.CreateMeal)
		protected.PUT("/meals/:id", mealHandler.UpdateMeal)
		protected.DELETE("/meals/:id", mealHandler.DeleteMeal)
	}
}
