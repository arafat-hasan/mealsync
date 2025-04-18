package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// mealEventRepository implements MealEventRepository interface
type mealEventRepository struct {
	*baseRepository[model.MealEvent]
	db *gorm.DB
}

// NewMealEventRepository creates a new instance of MealEventRepository
func NewMealEventRepository(db *gorm.DB) MealEventRepository {
	return &mealEventRepository{
		baseRepository: NewBaseRepository[model.MealEvent](db),
		db:             db,
	}
}

// CreateRequest creates a new meal request
func (r *mealEventRepository) CreateRequest(ctx context.Context, request *model.MealRequest) error {
	return r.db.WithContext(ctx).Create(request).Error
}

// FindRequestByID finds a meal request by ID
func (r *mealEventRepository) FindRequestByID(ctx context.Context, id uint) (*model.MealRequest, error) {
	var request model.MealRequest
	err := r.db.WithContext(ctx).First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// FindAllRequests finds all meal requests
func (r *mealEventRepository) FindAllRequests(ctx context.Context) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindRequestsByUserID finds meal requests by user ID
func (r *mealEventRepository) FindRequestsByUserID(ctx context.Context, userID uint) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// DeleteRequest soft deletes a meal request
func (r *mealEventRepository) DeleteRequest(ctx context.Context, request *model.MealRequest) error {
	return r.db.WithContext(ctx).Delete(request).Error
}

// CreateComment creates a new meal comment
func (r *mealEventRepository) CreateComment(ctx context.Context, comment *model.MealComment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

// FindCommentsByMealEventID finds meal comments by meal event ID
func (r *mealEventRepository) FindCommentsByMealEventID(ctx context.Context, mealEventID uint) ([]model.MealComment, error) {
	var comments []model.MealComment
	err := r.db.WithContext(ctx).Where("meal_event_id = ?", mealEventID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Create creates a new meal event
func (r *mealEventRepository) Create(ctx context.Context, meal *model.MealEvent) error {
	return r.baseRepository.Create(ctx, meal)
}

// FindByID finds a meal event by ID
func (r *mealEventRepository) FindByID(ctx context.Context, id uint) (*model.MealEvent, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all meal events
func (r *mealEventRepository) FindAll(ctx context.Context) ([]model.MealEvent, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds all active meal events
func (r *mealEventRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MealEvent, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a meal event
func (r *mealEventRepository) Update(ctx context.Context, meal *model.MealEvent) error {
	return r.baseRepository.Update(ctx, meal)
}

// Delete deletes a meal event (soft delete)
func (r *mealEventRepository) Delete(ctx context.Context, meal *model.MealEvent) error {
	return r.baseRepository.Delete(ctx, meal)
}

// HardDelete permanently deletes a meal event
func (r *mealEventRepository) HardDelete(ctx context.Context, meal *model.MealEvent) error {
	return r.baseRepository.HardDelete(ctx, meal)
}

// FindByUserID finds meal events by user ID
func (r *mealEventRepository) FindByUserID(ctx context.Context, userID uint) ([]model.MealEvent, error) {
	var meals []model.MealEvent
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&meals).Error
	if err != nil {
		return nil, err
	}
	return meals, nil
}
