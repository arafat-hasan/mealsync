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
		Where("user_id = ? AND read = ?", userID, false).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountUndeliveredByUserID counts undelivered notifications by user ID
func (r *notificationRepository) CountUndeliveredByUserID(ctx context.Context, userID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? AND delivered = ?", userID, false).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// MarkAsRead marks a notification as read
func (r *notificationRepository) MarkAsRead(ctx context.Context, id uint) error {
	now := gorm.Expr("NOW()")
	return r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"read":    true,
			"read_at": now,
		}).Error
}

// MarkAsDelivered marks a notification as delivered
func (r *notificationRepository) MarkAsDelivered(ctx context.Context, id uint) error {
	now := gorm.Expr("NOW()")
	return r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"delivered":    true,
			"delivered_at": now,
		}).Error
}

// FindUnreadByUserID finds unread notifications by user ID
func (r *notificationRepository) FindUnreadByUserID(ctx context.Context, userID uint) ([]model.Notification, error) {
	var notifications []model.Notification
	err := r.db.WithContext(ctx).Where("user_id = ? AND read = ?", userID, false).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// FindByType finds notifications by type
func (r *notificationRepository) FindByType(ctx context.Context, userID uint, notificationType model.NotificationType) ([]model.Notification, error) {
	var notifications []model.Notification
	err := r.db.WithContext(ctx).Where("user_id = ? AND type = ?", userID, notificationType).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
