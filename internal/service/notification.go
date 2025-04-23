package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/repository"
)

// notificationService handles business logic for notification-related operations
type notificationService struct {
	notificationRepo repository.NotificationRepository
	userRepo         repository.UserRepository
}

// NewNotificationService creates a new instance of NotificationService
func NewNotificationService(notificationRepo repository.NotificationRepository, userRepo repository.UserRepository) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		userRepo:         userRepo,
	}
}

// GetNotifications retrieves notifications for a user
func (s *notificationService) GetNotifications(ctx context.Context, userID uint) ([]model.Notification, error) {
	return s.notificationRepo.FindByUserID(ctx, userID)
}

// CreateNotification creates a new notification
func (s *notificationService) CreateNotification(ctx context.Context, notification *model.Notification, userID uint) error {
	if notification == nil {
		return errors.NewValidationError("notification cannot be nil", nil)
	}

	if notification.Type == "" {
		return errors.NewValidationError("notification type is required", nil)
	}

	// Set notification fields
	notification.UserID = userID
	notification.Read = false
	notification.CreatedBy = userID
	notification.UpdatedBy = userID

	return s.notificationRepo.Create(ctx, notification)
}

// MarkNotificationAsRead marks a notification as read
func (s *notificationService) MarkNotificationAsRead(ctx context.Context, notificationID uint, userID uint) error {
	notification, err := s.notificationRepo.FindByID(ctx, notificationID)
	if err != nil {
		return err
	}

	// Verify ownership
	if notification.UserID != userID {
		return errors.NewValidationError("unauthorized to mark this notification as read", nil)
	}

	return s.notificationRepo.MarkAsRead(ctx, notificationID)
}

// MarkNotificationAsDelivered marks a notification as delivered
func (s *notificationService) MarkNotificationAsDelivered(ctx context.Context, notificationID uint, userID uint) error {
	notification, err := s.notificationRepo.FindByID(ctx, notificationID)
	if err != nil {
		return err
	}

	// Verify ownership
	if notification.UserID != userID {
		return errors.NewValidationError("unauthorized to mark this notification as delivered", nil)
	}

	return s.notificationRepo.MarkAsDelivered(ctx, notificationID)
}

// DeleteNotification soft deletes a notification
func (s *notificationService) DeleteNotification(ctx context.Context, notificationID uint, userID uint) error {
	notification, err := s.notificationRepo.FindByID(ctx, notificationID)
	if err != nil {
		return err
	}

	// Verify ownership
	if notification.UserID != userID {
		return errors.NewValidationError("unauthorized to delete this notification", nil)
	}

	notification.UpdatedBy = userID
	return s.notificationRepo.Delete(ctx, notification)
}

// GetUnreadNotificationCount retrieves the count of unread notifications for a user
func (s *notificationService) GetUnreadNotificationCount(ctx context.Context, userID uint) (int64, error) {
	return s.notificationRepo.CountUnreadByUserID(ctx, userID)
}

// GetUndeliveredNotificationCount retrieves the count of undelivered notifications for a user
func (s *notificationService) GetUndeliveredNotificationCount(ctx context.Context, userID uint) (int64, error) {
	return s.notificationRepo.CountUndeliveredByUserID(ctx, userID)
}

// GetNotificationsByType retrieves notifications for a user by type
func (s *notificationService) GetNotificationsByType(ctx context.Context, userID uint, notificationType model.NotificationType) ([]model.Notification, error) {
	return s.notificationRepo.FindByType(ctx, userID, notificationType)
}

// GetUnreadNotifications retrieves unread notifications for a user
func (s *notificationService) GetUnreadNotifications(ctx context.Context, userID uint) ([]model.Notification, error) {
	return s.notificationRepo.FindUnreadByUserID(ctx, userID)
}

// CreateMealConfirmationNotification creates a notification for meal confirmation
func (s *notificationService) CreateMealConfirmationNotification(ctx context.Context, userID uint, mealEventID uint, message string) error {
	payload, err := json.Marshal(map[string]interface{}{
		"meal_event_id": mealEventID,
		"message":       message,
	})
	if err != nil {
		return err
	}

	notification := &model.Notification{
		UserID:    userID,
		Type:      model.NotificationTypeConfirmation,
		Payload:   payload,
		Message:   message,
		Read:      false,
		Delivered: false,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	return s.notificationRepo.Create(ctx, notification)
}

// CreateMealReminderNotification creates a notification for meal request reminder
func (s *notificationService) CreateMealReminderNotification(ctx context.Context, userID uint, mealEventID uint, message string, deadline time.Time) error {
	payload, err := json.Marshal(map[string]interface{}{
		"meal_event_id": mealEventID,
		"message":       message,
		"deadline":      deadline.Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	notification := &model.Notification{
		UserID:    userID,
		Type:      model.NotificationTypeReminder,
		Payload:   payload,
		Message:   message,
		Read:      false,
		Delivered: false,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	return s.notificationRepo.Create(ctx, notification)
}

// CreateMealCancellationNotification creates a notification for meal cancellation
func (s *notificationService) CreateMealCancellationNotification(ctx context.Context, userID uint, mealEventID uint, message string) error {
	payload, err := json.Marshal(map[string]interface{}{
		"meal_event_id": mealEventID,
		"message":       message,
	})
	if err != nil {
		return err
	}

	notification := &model.Notification{
		UserID:    userID,
		Type:      model.NotificationTypeConfirmation,
		Payload:   payload,
		Message:   message,
		Read:      false,
		Delivered: false,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	return s.notificationRepo.Create(ctx, notification)
}

// CreateAdminNotification creates an admin notification
func (s *notificationService) CreateAdminNotification(ctx context.Context, userID uint, message string, importance string) error {
	payload, err := json.Marshal(map[string]interface{}{
		"message":    message,
		"importance": importance,
	})
	if err != nil {
		return err
	}

	notification := &model.Notification{
		UserID:    userID,
		Type:      model.NotificationTypeAdminMessage,
		Payload:   payload,
		Message:   message,
		Read:      false,
		Delivered: false,
		CreatedBy: userID,
		UpdatedBy: userID,
	}

	return s.notificationRepo.Create(ctx, notification)
}
