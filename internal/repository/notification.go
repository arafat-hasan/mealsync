package repository

import (
	"context"

	"github.com/arafat-hasan/mealsync/internal/model"
	"gorm.io/gorm"
)

// notificationRepository implements NotificationRepository interface
type notificationRepository struct {
	*baseRepository[model.Notification]
	db *gorm.DB
}

// NewNotificationRepository creates a new instance of NotificationRepository
func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{
		baseRepository: NewBaseRepository[model.Notification](db),
		db:             db,
	}
}

// FindByUserID finds notifications by user ID
func (r *notificationRepository) FindByUserID(ctx context.Context, userID uint) ([]model.Notification, error) {
	var notifications []model.Notification
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// CountUnreadByUserID counts unread notifications by user ID
func (r *notificationRepository) CountUnreadByUserID(ctx context.Context, userID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
