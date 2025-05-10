package repository

import (
	"context"
	"time"

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
func (r *mealEventRepository) CreateComment(ctx context.Context, comment *model.MenuItemComment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

// FindCommentsByMealEventID finds meal comments by meal event ID
func (r *mealEventRepository) FindCommentsByMealEventID(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error) {
	var comments []model.MenuItemComment
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

// FindByID finds a meal event by ID with preloaded relationships
func (r *mealEventRepository) FindByID(ctx context.Context, id uint) (*model.MealEvent, error) {
	var meal model.MealEvent
	err := r.db.WithContext(ctx).
		Preload("MenuSets").
		Preload("MenuSets.MenuSet").
		Preload("Addresses").
		Preload("Addresses.Address").
		Preload("MealRequests").
		Preload("MenuItemComments").
		First(&meal, id).Error
	if err != nil {
		return nil, err
	}
	return &meal, nil
}

// FindAll finds all meal events with preloaded relationships
func (r *mealEventRepository) FindAll(ctx context.Context) ([]model.MealEvent, error) {
	var meals []model.MealEvent
	err := r.db.WithContext(ctx).
		Preload("MenuSets").
		Preload("MenuSets.MenuSet").
		Preload("Addresses").
		Preload("Addresses.Address").
		Preload("MealRequests").
		Preload("MenuItemComments").
		Find(&meals).Error
	if err != nil {
		return nil, err
	}
	return meals, nil
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

// FindUpcomingAndActive finds upcoming and active meal events
func (r *mealEventRepository) FindUpcomingAndActive(ctx context.Context) ([]model.MealEvent, error) {
	var meals []model.MealEvent
	err := r.db.WithContext(ctx).
		Preload("MenuSets").
		Preload("MenuSets.MenuSet").
		Preload("Addresses").
		Preload("Addresses.Address").
		Preload("MealRequests").
		Preload("MenuItemComments").
		Where("is_active = ?", true).
		Where("event_date >= ?", time.Now()).
		Order("event_date ASC").
		Find(&meals).Error
	if err != nil {
		return nil, err
	}
	return meals, nil
}

// AddMenuSetToEvent associates a menu set with a meal event
func (r *mealEventRepository) AddMenuSetToEvent(ctx context.Context, MealEventSet *model.MealEventSet) error {
	return r.db.WithContext(ctx).Create(MealEventSet).Error
}

// UpdateMenuSetInEvent updates the menu set association details in a meal event
func (r *mealEventRepository) UpdateMenuSetInEvent(ctx context.Context, MealEventSet *model.MealEventSet) error {
	return r.db.WithContext(ctx).
		Model(&model.MealEventSet{}).
		Where("meal_event_id = ? AND menu_set_id = ?", MealEventSet.MealEventID, MealEventSet.MenuSetID).
		Updates(map[string]interface{}{
			"label":      MealEventSet.Label,
			"note":       MealEventSet.Note,
			"updated_by": MealEventSet.UpdatedByID,
			"updated_at": time.Now(),
		}).Error
}

// RemoveMenuSetFromEvent removes a menu set association from a meal event
func (r *mealEventRepository) RemoveMenuSetFromEvent(ctx context.Context, mealEventID uint, menuSetID uint) error {
	return r.db.WithContext(ctx).
		Where("meal_event_id = ? AND menu_set_id = ?", mealEventID, menuSetID).
		Delete(&model.MealEventSet{}).Error
}

// FindMenuSetsByEventID finds all menu sets associated with a meal event
func (r *mealEventRepository) FindMenuSetsByEventID(ctx context.Context, mealEventID uint) ([]model.MealEventSet, error) {
	var menuSets []model.MealEventSet
	err := r.db.WithContext(ctx).
		Preload("MenuSet").
		Where("meal_event_id = ?", mealEventID).
		Find(&menuSets).Error
	if err != nil {
		return nil, err
	}
	return menuSets, nil
}

// FindByDateRange finds meal events within a specified date range
func (r *mealEventRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealEvent, error) {
	var meals []model.MealEvent
	err := r.db.WithContext(ctx).
		Preload("MenuSets").
		Preload("MenuSets.MenuSet").
		Preload("MenuSets.MenuSet.MenuSetItems").
		Preload("Addresses").
		Preload("Addresses.Address").
		Where("event_date BETWEEN ? AND ?", startDate, endDate).
		Order("event_date ASC").
		Find(&meals).Error
	if err != nil {
		return nil, err
	}
	return meals, nil
}
