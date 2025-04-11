package api

import (
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, authHandler *AuthHandler, mealHandler *MealHandler, menuHandler *MenuHandler) {
	// Public routes (no auth required)
	public := r.Group("/api")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.POST("/refresh", authHandler.RefreshToken)
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

		// Menu routes
		protected.GET("/menus", menuHandler.GetMenus)
		protected.GET("/menus/:id", menuHandler.GetMenuByID)
		protected.POST("/menus", menuHandler.CreateMenu)
		protected.PUT("/menus/:id", menuHandler.UpdateMenu)
		protected.DELETE("/menus/:id", menuHandler.DeleteMenu)
		protected.GET("/menus/:id/items", menuHandler.GetMenuItems)
		protected.POST("/menus/:id/items", menuHandler.AddMenuItem)
		protected.DELETE("/menus/:id/items/:item_id", menuHandler.RemoveMenuItem)
	}
}
