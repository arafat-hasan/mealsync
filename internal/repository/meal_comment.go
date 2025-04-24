package repository

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// menuItemCommentRepository implements MenuItemCommentRepository interface
type menuItemCommentRepository struct {
	*baseRepository[model.MenuItemComment]
	db *gorm.DB
}

// NewMenuItemCommentRepository creates a new instance of MenuItemCommentRepository
func NewMenuItemCommentRepository(db *gorm.DB) MenuItemCommentRepository {
	return &menuItemCommentRepository{
		baseRepository: NewBaseRepository[model.MenuItemComment](db),
		db:             db,
	}
}

// Create creates a new menu item comment
func (r *menuItemCommentRepository) Create(ctx context.Context, comment *model.MenuItemComment) error {
	return r.baseRepository.Create(ctx, comment)
}

// FindByID finds a menu item comment by ID
func (r *menuItemCommentRepository) FindByID(ctx context.Context, id uint) (*model.MenuItemComment, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all menu item comments
func (r *menuItemCommentRepository) FindAll(ctx context.Context) ([]model.MenuItemComment, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds active menu item comments based on conditions
func (r *menuItemCommentRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MenuItemComment, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a menu item comment
func (r *menuItemCommentRepository) Update(ctx context.Context, comment *model.MenuItemComment) error {
	return r.baseRepository.Update(ctx, comment)
}

// Delete soft deletes a menu item comment
func (r *menuItemCommentRepository) Delete(ctx context.Context, comment *model.MenuItemComment) error {
	return r.baseRepository.Delete(ctx, comment)
}

// HardDelete permanently deletes a menu item comment
func (r *menuItemCommentRepository) HardDelete(ctx context.Context, comment *model.MenuItemComment) error {
	return r.baseRepository.HardDelete(ctx, comment)
}

// FindByMealEventID finds menu item comments by meal event ID
func (r *menuItemCommentRepository) FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MenuItemComment, error) {
	var comments []model.MenuItemComment
	err := r.db.WithContext(ctx).Where("meal_event_id = ?", mealEventID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByUserID finds menu item comments by user ID
func (r *menuItemCommentRepository) FindByUserID(ctx context.Context, userID uint) ([]model.MenuItemComment, error) {
	var comments []model.MenuItemComment
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByMenuItemID finds comments by menu item ID
func (r *menuItemCommentRepository) FindByMenuItemID(ctx context.Context, menuItemID uint) ([]model.MenuItemComment, error) {
	var comments []model.MenuItemComment
	err := r.db.WithContext(ctx).Where("menu_item_id = ?", menuItemID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// CountByMealEventID counts comments by meal event ID
func (r *menuItemCommentRepository) CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.MenuItemComment{}).Where("meal_event_id = ?", mealEventID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByDateRange finds comments within a date range
func (r *menuItemCommentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MenuItemComment, error) {
	var comments []model.MenuItemComment
	err := r.db.WithContext(ctx).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindRecentComments finds recent comments with a limit
func (r *menuItemCommentRepository) FindRecentComments(ctx context.Context, limit int) ([]model.MenuItemComment, error) {
	var comments []model.MenuItemComment
	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindWithUserDetails finds a comment with user details by ID
func (r *menuItemCommentRepository) FindWithUserDetails(ctx context.Context, commentID uint) (*model.MenuItemComment, error) {
	var comment model.MenuItemComment
	err := r.db.WithContext(ctx).
		Preload("User").
		First(&comment, commentID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// FindReplies finds all replies to a parent comment
func (r *menuItemCommentRepository) FindReplies(ctx context.Context, parentID uint) ([]model.MenuItemComment, error) {
	var replies []model.MenuItemComment
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("parent_id = ?", parentID).
		Find(&replies).Error
	if err != nil {
		return nil, err
	}
	return replies, nil
}
