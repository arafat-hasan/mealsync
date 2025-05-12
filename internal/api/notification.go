package api

import (
	"net/http"
	"strconv"

	"github.com/arafat-hasan/mealsync/internal/errors"
	"github.com/arafat-hasan/mealsync/internal/model"
	"github.com/arafat-hasan/mealsync/internal/service"
	"github.com/arafat-hasan/mealsync/internal/utils"
	"github.com/gin-gonic/gin"
)

// NotificationHandler handles notification-related API requests
type NotificationHandler struct {
	notificationService service.NotificationService
}

// NewNotificationHandler creates a new instance of NotificationHandler
func NewNotificationHandler(notificationService service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

// handleError properly formats and returns API errors
func handleError(c *gin.Context, err error) {
	if appErr, ok := err.(*errors.AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
}

// GetNotifications godoc
//	@Summary		Get notifications for current user
//	@Description	Retrieves all notifications for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{array}		model.Notification
//	@Failure		401	{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		500	{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications [get]
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	notifications, err := h.notificationService.GetNotifications(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetNotificationsByType godoc
//	@Summary		Get notifications by type
//	@Description	Retrieves notifications by type for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			type	path		string	true	"Notification type"	Enums(reminder, confirmation, admin-message, event-info)
//	@Success		200		{array}		model.Notification
//	@Failure		400		{object}	errors.ErrorResponse	"Bad Request"
//	@Failure		401		{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		500		{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications/type/{type} [get]
func (h *NotificationHandler) GetNotificationsByType(c *gin.Context) {
	notificationType := model.NotificationType(c.Param("type"))
	if !isValidNotificationType(notificationType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification type"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	notifications, err := h.notificationService.GetNotificationsByType(c.Request.Context(), userID, notificationType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetUnreadNotifications godoc
//	@Summary		Get unread notifications
//	@Description	Retrieves unread notifications for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{array}		model.Notification
//	@Failure		401	{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		500	{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications/unread [get]
func (h *NotificationHandler) GetUnreadNotifications(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	notifications, err := h.notificationService.GetUnreadNotifications(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetUnreadNotificationCount godoc
//	@Summary		Get unread notification count
//	@Description	Retrieves the count of unread notifications for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	map[string]int64
//	@Failure		401	{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		500	{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications/unread/count [get]
func (h *NotificationHandler) GetUnreadNotificationCount(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	count, err := h.notificationService.GetUnreadNotificationCount(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// MarkNotificationAsRead godoc
//	@Summary		Mark notification as read
//	@Description	Marks a notification as read for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			notification_id	path		int	true	"Notification ID"
//	@Success		200				{object}	map[string]string
//	@Failure		400				{object}	errors.ErrorResponse	"Bad Request"
//	@Failure		401				{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		404				{object}	errors.ErrorResponse	"Not Found"
//	@Failure		500				{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications/{notification_id}/read [put]
func (h *NotificationHandler) MarkNotificationAsRead(c *gin.Context) {
	notificationIDStr := c.Param("notification_id")
	notificationID, err := strconv.ParseUint(notificationIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.notificationService.MarkNotificationAsRead(c.Request.Context(), uint(notificationID), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}

// MarkNotificationAsDelivered godoc
//	@Summary		Mark notification as delivered
//	@Description	Marks a notification as delivered for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			notification_id	path		int	true	"Notification ID"
//	@Success		200				{object}	map[string]string
//	@Failure		400				{object}	errors.ErrorResponse	"Bad Request"
//	@Failure		401				{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		404				{object}	errors.ErrorResponse	"Not Found"
//	@Failure		500				{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications/{notification_id}/delivered [put]
func (h *NotificationHandler) MarkNotificationAsDelivered(c *gin.Context) {
	notificationIDStr := c.Param("notification_id")
	notificationID, err := strconv.ParseUint(notificationIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.notificationService.MarkNotificationAsDelivered(c.Request.Context(), uint(notificationID), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification marked as delivered"})
}

// DeleteNotification godoc
//	@Summary		Delete notification
//	@Description	Deletes a notification for the authenticated user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			notification_id	path		int	true	"Notification ID"
//	@Success		200				{object}	map[string]string
//	@Failure		400				{object}	errors.ErrorResponse	"Bad Request"
//	@Failure		401				{object}	errors.ErrorResponse	"Unauthorized"
//	@Failure		404				{object}	errors.ErrorResponse	"Not Found"
//	@Failure		500				{object}	errors.ErrorResponse	"Internal Server Error"
//	@Router			/notifications/{notification_id} [delete]
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	notificationIDStr := c.Param("notification_id")
	notificationID, err := strconv.ParseUint(notificationIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.notificationService.DeleteNotification(c.Request.Context(), uint(notificationID), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}

// Helper function to check if a notification type is valid
func isValidNotificationType(notificationType model.NotificationType) bool {
	switch notificationType {
	case model.NotificationTypeReminder, model.NotificationTypeConfirmation,
		model.NotificationTypeAdminMessage, model.NotificationTypeEventInfo:
		return true
	default:
		return false
	}
}
