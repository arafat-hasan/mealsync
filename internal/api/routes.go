package api

import (
	"github.com/arafat-hasan/mealsync/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine, authHandler *AuthHandler, mealHandler *MealEventHandler, menuSetHandler *MenuSetHandler, mealCommentHandler *MealCommentHandler, menuItemHandler *MenuItemHandler, mealRequestHandler *MealRequestHandler, notificationHandler *NotificationHandler) {
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
		// Meal event routes
		meals := protected.Group("/meals")
		{
			// List and create routes (no parameters)
			meals.GET("", mealHandler.GetMealEvents)
			meals.POST("", mealHandler.CreateMealEvent)

			// Routes with meal event ID parameter
			meal := meals.Group("/:meal_id")
			{
				// Meal event operations
				meal.GET("", mealHandler.GetMealEventByID)
				meal.PUT("", mealHandler.UpdateMealEvent)
				meal.DELETE("", mealHandler.DeleteMealEvent)

				// Comment routes under meal event
				comments := meal.Group("/comments")
				{
					comments.GET("", mealCommentHandler.GetComments)
					comments.POST("", mealCommentHandler.CreateComment)
				}
			}
		}

		// Menu set routes
		menus := protected.Group("/menus")
		{
			menus.GET("", menuSetHandler.GetMenuSets)
			menus.GET("/:id", menuSetHandler.GetMenuSetByID)
			menus.POST("", menuSetHandler.CreateMenuSet)
			menus.PUT("/:id", menuSetHandler.UpdateMenuSet)
			menus.DELETE("/:id", menuSetHandler.DeleteMenuSet)
			menus.GET("/:id/items", menuSetHandler.GetMenuSetItems)
			menus.POST("/:id/items", menuSetHandler.AddMenuItemToMenuSet)
			menus.DELETE("/:id/items/:item_id", menuSetHandler.RemoveMenuItemFromMenuSet)
		}

		// Menu item routes
		menuItems := protected.Group("/menu-items")
		{
			menuItems.GET("", menuItemHandler.GetMenuItems)
			menuItems.GET("/:id", menuItemHandler.GetMenuItemByID)
			menuItems.POST("", menuItemHandler.CreateMenuItem)
			menuItems.PUT("/:id", menuItemHandler.UpdateMenuItem)
			menuItems.DELETE("/:id", menuItemHandler.DeleteMenuItem)
			menuItems.GET("/category/:category", menuItemHandler.GetMenuItemsByCategory)
			menuItems.GET("/menu-set/:menu_set_id", menuItemHandler.GetMenuItemsByMenuSet)
		}

		// Meal request routes
		mealRequests := protected.Group("/meal-requests")
		{
			mealRequests.GET("", mealRequestHandler.GetMealRequests)
			mealRequests.GET("/:id", mealRequestHandler.GetMealRequestByID)
			mealRequests.POST("", mealRequestHandler.CreateMealRequest)
			mealRequests.PUT("/:id", mealRequestHandler.UpdateMealRequest)
			mealRequests.DELETE("/:id", mealRequestHandler.DeleteMealRequest)
			mealRequests.PUT("/:id/status", mealRequestHandler.UpdateRequestStatus)

			// Meal request items
			mealRequests.GET("/:id/items", mealRequestHandler.GetRequestItems)
			mealRequests.POST("/:id/items", mealRequestHandler.AddRequestItem)
			mealRequests.DELETE("/:id/items/:item_id", mealRequestHandler.RemoveRequestItem)
		}

		// Comment routes
		comments := protected.Group("/comments")
		{
			comments.GET("/:id", mealCommentHandler.GetCommentByID)
			comments.PUT("/:id", mealCommentHandler.UpdateComment)
			comments.DELETE("/:id", mealCommentHandler.DeleteComment)
			comments.GET("/:id/replies", mealCommentHandler.GetReplies)
		}

		// User comment routes
		users := protected.Group("/users")
		{
			users.GET("/:user_id/comments", mealCommentHandler.GetUserComments)
		}

		// Notification routes
		notifications := protected.Group("/notifications")
		{
			notifications.GET("", notificationHandler.GetNotifications)
			notifications.GET("/unread", notificationHandler.GetUnreadNotifications)
			notifications.GET("/unread/count", notificationHandler.GetUnreadNotificationCount)
			notifications.GET("/type/:type", notificationHandler.GetNotificationsByType)
			notifications.PUT("/:notification_id/read", notificationHandler.MarkNotificationAsRead)
			notifications.PUT("/:notification_id/delivered", notificationHandler.MarkNotificationAsDelivered)
			notifications.DELETE("/:notification_id", notificationHandler.DeleteNotification)
		}
	}
}
