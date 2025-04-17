package repository

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// mealRequestRepository implements MealRequestRepository interface
type mealRequestRepository struct {
	*baseRepository[model.MealRequest]
	db *gorm.DB
}

// NewMealRequestRepository creates a new instance of MealRequestRepository
func NewMealRequestRepository(db *gorm.DB) MealRequestRepository {
	return &mealRequestRepository{
		baseRepository: NewBaseRepository[model.MealRequest](db),
		db:             db,
	}
}

// Create creates a new meal request
func (r *mealRequestRepository) Create(ctx context.Context, request *model.MealRequest) error {
	return r.baseRepository.Create(ctx, request)
}

// FindByID finds a meal request by ID
func (r *mealRequestRepository) FindByID(ctx context.Context, id uint) (*model.MealRequest, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all meal requests
func (r *mealRequestRepository) FindAll(ctx context.Context) ([]model.MealRequest, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds active meal requests based on conditions
func (r *mealRequestRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MealRequest, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a meal request
func (r *mealRequestRepository) Update(ctx context.Context, request *model.MealRequest) error {
	return r.baseRepository.Update(ctx, request)
}

// Delete soft deletes a meal request
func (r *mealRequestRepository) Delete(ctx context.Context, request *model.MealRequest) error {
	return r.baseRepository.Delete(ctx, request)
}

// HardDelete permanently deletes a meal request
func (r *mealRequestRepository) HardDelete(ctx context.Context, request *model.MealRequest) error {
	return r.baseRepository.HardDelete(ctx, request)
}

// FindByUserID finds meal requests by user ID
func (r *mealRequestRepository) FindByUserID(ctx context.Context, userID uint) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindByMealEventID finds meal requests by meal event ID
func (r *mealRequestRepository) FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).Where("meal_event_id = ?", mealEventID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// AddRequestItem adds a meal request item
func (r *mealRequestRepository) AddRequestItem(ctx context.Context, item *model.MealRequestItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// RemoveRequestItem removes a meal request item
func (r *mealRequestRepository) RemoveRequestItem(ctx context.Context, item *model.MealRequestItem) error {
	return r.db.WithContext(ctx).Delete(item).Error
}

// FindRequestItems finds all items for a meal request
func (r *mealRequestRepository) FindRequestItems(ctx context.Context, requestID uint) ([]model.MealRequestItem, error) {
	var items []model.MealRequestItem
	err := r.db.WithContext(ctx).Where("meal_request_id = ?", requestID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// CountByMealEventID counts requests by meal event ID
func (r *mealRequestRepository) CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.MealRequest{}).Where("meal_event_id = ?", mealEventID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindApprovedRequests finds all approved meal requests
func (r *mealRequestRepository) FindApprovedRequests(ctx context.Context) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).Where("status = ?", "approved").Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindByDateRange finds meal requests within a date range
func (r *mealRequestRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).
		Where("requested_for BETWEEN ? AND ?", startDate, endDate).
		Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindByMenuSetID finds meal requests by menu set ID
func (r *mealRequestRepository) FindByMenuSetID(ctx context.Context, menuSetID uint) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).
		Where("menu_set_id = ?", menuSetID).
		Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindPendingRequests finds all pending meal requests
func (r *mealRequestRepository) FindPendingRequests(ctx context.Context) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).
		Where("status = ?", model.RequestStatusPending).
		Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindRejectedRequests finds all rejected meal requests
func (r *mealRequestRepository) FindRejectedRequests(ctx context.Context) ([]model.MealRequest, error) {
	var requests []model.MealRequest
	err := r.db.WithContext(ctx).
		Where("status = ?", model.RequestStatusRejected).
		Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindWithDetails finds a meal request with all related details
func (r *mealRequestRepository) FindWithDetails(ctx context.Context, id uint) (*model.MealRequest, error) {
	var request model.MealRequest
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("MealEvent").
		Preload("MealEvent.Menu").
		Preload("MealEvent.Menu.CreatedBy").
		First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// UpdateRequestStatus updates the status of a meal request
func (r *mealRequestRepository) UpdateRequestStatus(ctx context.Context, id uint, status model.RequestStatus) error {
	return r.db.WithContext(ctx).
		Model(&model.MealRequest{}).
		Where("id = ?", id).
		Update("status", status).Error
}
