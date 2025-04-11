package service

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
	"gorm.io/gorm"
)

// MealService defines the interface for meal-related operations
type MealService interface {
	CreateMeal(ctx context.Context, meal *model.Meal) error
	GetMeals(ctx context.Context) ([]model.Meal, error)
	GetMealByID(ctx context.Context, id uint) (*model.Meal, error)
	UpdateMeal(ctx context.Context, meal *model.Meal) error
	DeleteMeal(ctx context.Context, id uint) error
	GetMealsByUserID(ctx context.Context, userID uint) ([]model.Meal, error)
	GetMealsByRestaurantID(ctx context.Context, restaurantID uint) ([]model.Meal, error)
}

// mealService implements the MealService interface
type mealService struct {
	mealRepo repository.MealRepository
}

// NewMealService creates a new instance of MealService
func NewMealService(db *gorm.DB) MealService {
	return &mealService{
		mealRepo: repository.NewMealRepository(db),
	}
}

// CreateMeal creates a new meal
func (s *mealService) CreateMeal(ctx context.Context, meal *model.Meal) error {
	return s.mealRepo.Create(ctx, meal)
}

// GetMeals retrieves all meals
func (s *mealService) GetMeals(ctx context.Context) ([]model.Meal, error) {
	return s.mealRepo.FindAll(ctx)
}

// GetMealByID retrieves a meal by its ID
func (s *mealService) GetMealByID(ctx context.Context, id uint) (*model.Meal, error) {
	return s.mealRepo.FindByID(ctx, id)
}

// UpdateMeal updates an existing meal
func (s *mealService) UpdateMeal(ctx context.Context, meal *model.Meal) error {
	return s.mealRepo.Update(ctx, meal)
}

// DeleteMeal deletes a meal by its ID
func (s *mealService) DeleteMeal(ctx context.Context, id uint) error {
	return s.mealRepo.Delete(ctx, id)
}

// GetMealsByUserID retrieves meals by user ID
func (s *mealService) GetMealsByUserID(ctx context.Context, userID uint) ([]model.Meal, error) {
	return s.mealRepo.FindByUserID(ctx, userID)
}

// GetMealsByRestaurantID retrieves meals by restaurant ID
func (s *mealService) GetMealsByRestaurantID(ctx context.Context, restaurantID uint) ([]model.Meal, error) {
	return s.mealRepo.FindByRestaurantID(ctx, restaurantID)
}
