package repository

import (
	"context"
	"time"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// mealCommentRepository implements MealCommentRepository interface
type mealCommentRepository struct {
	*baseRepository[model.MealComment]
	db *gorm.DB
}

// NewMealCommentRepository creates a new instance of MealCommentRepository
func NewMealCommentRepository(db *gorm.DB) MealCommentRepository {
	return &mealCommentRepository{
		baseRepository: NewBaseRepository[model.MealComment](db),
		db:             db,
	}
}

// Create creates a new meal comment
func (r *mealCommentRepository) Create(ctx context.Context, comment *model.MealComment) error {
	return r.baseRepository.Create(ctx, comment)
}

// FindByID finds a meal comment by ID
func (r *mealCommentRepository) FindByID(ctx context.Context, id uint) (*model.MealComment, error) {
	return r.baseRepository.FindByID(ctx, id)
}

// FindAll finds all meal comments
func (r *mealCommentRepository) FindAll(ctx context.Context) ([]model.MealComment, error) {
	return r.baseRepository.FindAll(ctx)
}

// FindActive finds active meal comments based on conditions
func (r *mealCommentRepository) FindActive(ctx context.Context, conditions map[string]interface{}) ([]model.MealComment, error) {
	return r.baseRepository.FindActive(ctx, conditions)
}

// Update updates a meal comment
func (r *mealCommentRepository) Update(ctx context.Context, comment *model.MealComment) error {
	return r.baseRepository.Update(ctx, comment)
}

// Delete soft deletes a meal comment
func (r *mealCommentRepository) Delete(ctx context.Context, comment *model.MealComment) error {
	return r.baseRepository.Delete(ctx, comment)
}

// HardDelete permanently deletes a meal comment
func (r *mealCommentRepository) HardDelete(ctx context.Context, comment *model.MealComment) error {
	return r.baseRepository.HardDelete(ctx, comment)
}

// FindByMealEventID finds meal comments by meal event ID
func (r *mealCommentRepository) FindByMealEventID(ctx context.Context, mealEventID uint) ([]model.MealComment, error) {
	var comments []model.MealComment
	err := r.db.WithContext(ctx).Where("meal_event_id = ?", mealEventID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByUserID finds meal comments by user ID
func (r *mealCommentRepository) FindByUserID(ctx context.Context, userID uint) ([]model.MealComment, error) {
	var comments []model.MealComment
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// CountByMealEventID counts comments by meal event ID
func (r *mealCommentRepository) CountByMealEventID(ctx context.Context, mealEventID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.MealComment{}).Where("meal_event_id = ?", mealEventID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByDateRange finds comments within a date range
func (r *mealCommentRepository) FindByDateRange(ctx context.Context, startDate, endDate time.Time) ([]model.MealComment, error) {
	var comments []model.MealComment
	err := r.db.WithContext(ctx).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByMealRequestID finds comments by meal request ID
func (r *mealCommentRepository) FindByMealRequestID(ctx context.Context, mealRequestID uint) ([]model.MealComment, error) {
	var comments []model.MealComment
	err := r.db.WithContext(ctx).
		Where("meal_request_id = ?", mealRequestID).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByParentCommentID finds comments by parent comment ID
func (r *mealCommentRepository) FindByParentCommentID(ctx context.Context, parentCommentID uint) ([]model.MealComment, error) {
	var comments []model.MealComment
	err := r.db.WithContext(ctx).
		Where("parent_comment_id = ?", parentCommentID).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// FindRecentComments finds recent comments with a limit
func (r *mealCommentRepository) FindRecentComments(ctx context.Context, limit int) ([]model.MealComment, error) {
	var comments []model.MealComment
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
func (r *mealCommentRepository) FindWithUserDetails(ctx context.Context, commentID uint) (*model.MealComment, error) {
	var comment model.MealComment
	err := r.db.WithContext(ctx).
		Preload("User").
		First(&comment, commentID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
