package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// mealRepository implements MealRepository interface
type mealRepository struct {
	*baseRepository[model.Meal]
}

// NewMealRepository creates a new meal repository
func NewMealRepository(db *gorm.DB) MealRepository {
	return &mealRepository{
		baseRepository: NewBaseRepository[model.Meal](db),
	}
}

// Create creates a new meal
func (r *mealRepository) Create(ctx context.Context, meal *model.Meal) error {
	return r.baseRepository.Create(ctx, meal)
}

// FindByID finds a meal by ID
func (r *mealRepository) FindByID(ctx context.Context, id uint) (*model.Meal, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all meals
func (r *mealRepository) FindAll(ctx context.Context) ([]model.Meal, error) {
	return r.baseRepository.FindAll(ctx)
}

// Update updates a meal
func (r *mealRepository) Update(ctx context.Context, meal *model.Meal) error {
	return r.baseRepository.Update(ctx, meal)
}

// Delete deletes a meal
func (r *mealRepository) Delete(ctx context.Context, id uint) error {
	return r.baseRepository.Delete(ctx, id)
}

// FindByUserID finds meals by user ID
func (r *mealRepository) FindByUserID(ctx context.Context, userID uint) ([]model.Meal, error) {
	var meals []model.Meal
	err := r.baseRepository.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&meals).Error
	if err != nil {
		return nil, err
	}
	return meals, nil
}

// FindByRestaurantID finds meals by restaurant ID
func (r *mealRepository) FindByRestaurantID(ctx context.Context, restaurantID uint) ([]model.Meal, error) {
	var meals []model.Meal
	err := r.baseRepository.DB.WithContext(ctx).Where("restaurant_id = ?", restaurantID).Find(&meals).Error
	if err != nil {
		return nil, err
	}
	return meals, nil
}
