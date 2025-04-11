package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
)

// BaseRepository defines common CRUD operations
type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id uint) (*T, error)
	FindAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uint) error
}

// UserRepository defines user-specific operations
type UserRepository interface {
	BaseRepository[model.User]
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
}

// MealRepository defines meal-specific operations
type MealRepository interface {
	BaseRepository[model.Meal]
	FindByUserID(ctx context.Context, userID uint) ([]model.Meal, error)
	FindByRestaurantID(ctx context.Context, restaurantID uint) ([]model.Meal, error)
}

// RestaurantRepository defines restaurant-specific operations
type RestaurantRepository interface {
	BaseRepository[model.Restaurant]
	FindByUserID(ctx context.Context, userID uint) ([]model.Restaurant, error)
}

// MenuRepository defines menu-specific operations
type MenuRepository interface {
	BaseRepository[model.MenuItem]
	FindByRestaurantID(ctx context.Context, restaurantID uint) ([]model.MenuItem, error)
}

// OrderRepository defines order-specific operations
type OrderRepository interface {
	BaseRepository[model.Order]
	FindByUserID(ctx context.Context, userID uint) ([]model.Order, error)
	FindByRestaurantID(ctx context.Context, restaurantID uint) ([]model.Order, error)
}

// MealPlanRepository defines meal plan-specific operations
type MealPlanRepository interface {
	BaseRepository[model.MealPlan]
	FindByUserID(ctx context.Context, userID uint) ([]model.MealPlan, error)
}

// MealCommentRepository defines meal comment-specific operations
type MealCommentRepository interface {
	BaseRepository[model.MealComment]
	FindByMealID(ctx context.Context, mealID uint) ([]model.MealComment, error)
}

// NotificationRepository defines notification-specific operations
type NotificationRepository interface {
	BaseRepository[model.Notification]
	FindByUserID(ctx context.Context, userID uint) ([]model.Notification, error)
	MarkAsRead(ctx context.Context, id uint) error
}
