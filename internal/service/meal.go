package service

import (
	"context"
	"errors"

	apperrors "github.com/arafat-hasan/mealsync/internal/errors"
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
	if meal == nil {
		return apperrors.NewValidationError("Meal is required", nil)
	}

	if meal.Date == "" {
		return apperrors.NewValidationError("Date is required", nil)
	}

	if len(meal.MenuItems) == 0 {
		return apperrors.NewValidationError("Menu items are required", nil)
	}

	if err := s.mealRepo.Create(ctx, meal); err != nil {
		return apperrors.NewInternalError("Failed to create meal", err)
	}

	return nil
}

// GetMeals retrieves all meals
func (s *mealService) GetMeals(ctx context.Context) ([]model.Meal, error) {
	meals, err := s.mealRepo.FindAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch meals", err)
	}
	return meals, nil
}

// GetMealByID retrieves a meal by its ID
func (s *mealService) GetMealByID(ctx context.Context, id uint) (*model.Meal, error) {
	if id == 0 {
		return nil, apperrors.NewValidationError("Meal ID is required", nil)
	}

	meal, err := s.mealRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Meal not found", err)
		}
		return nil, apperrors.NewInternalError("Failed to fetch meal", err)
	}
	return meal, nil
}

// UpdateMeal updates an existing meal
func (s *mealService) UpdateMeal(ctx context.Context, meal *model.Meal) error {
	if meal == nil {
		return apperrors.NewValidationError("Meal is required", nil)
	}

	if meal.ID == 0 {
		return apperrors.NewValidationError("Meal ID is required", nil)
	}

	// Check if meal exists
	if _, err := s.mealRepo.FindByID(ctx, meal.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Meal not found", err)
		}
		return apperrors.NewInternalError("Failed to fetch meal", err)
	}

	// Update meal
	if err := s.mealRepo.Update(ctx, meal); err != nil {
		return apperrors.NewInternalError("Failed to update meal", err)
	}

	return nil
}

// DeleteMeal deletes a meal by its ID
func (s *mealService) DeleteMeal(ctx context.Context, id uint) error {
	if id == 0 {
		return apperrors.NewValidationError("Meal ID is required", nil)
	}

	// Check if meal exists
	_, err := s.mealRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Meal not found", err)
		}
		return apperrors.NewInternalError("Failed to fetch meal", err)
	}

	// Delete meal
	if err := s.mealRepo.Delete(ctx, id); err != nil {
		return apperrors.NewInternalError("Failed to delete meal", err)
	}

	return nil
}

// GetMealsByUserID retrieves meals by user ID
func (s *mealService) GetMealsByUserID(ctx context.Context, userID uint) ([]model.Meal, error) {
	if userID == 0 {
		return nil, apperrors.NewValidationError("User ID is required", nil)
	}

	meals, err := s.mealRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch meals", err)
	}
	return meals, nil
}

// GetMealsByRestaurantID retrieves meals by restaurant ID
func (s *mealService) GetMealsByRestaurantID(ctx context.Context, restaurantID uint) ([]model.Meal, error) {
	if restaurantID == 0 {
		return nil, apperrors.NewValidationError("Restaurant ID is required", nil)
	}

	meals, err := s.mealRepo.FindByRestaurantID(ctx, restaurantID)
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to fetch meals", err)
	}
	return meals, nil
}
