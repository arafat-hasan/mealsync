package model

import "encoding/json"

// Notification represents a notification entity in the system
// @Description Notification entity containing user notifications and their details
type Notification struct {
	Base
	UserID        uint             `json:"user_id" gorm:"not null" example:"1"`
	Type          NotificationType `json:"type" gorm:"not null" example:"reminder" enums:"reminder,confirmation,admin-message"`
	Payload       json.RawMessage  `json:"payload" gorm:"type:jsonb" swaggertype:"string" example:"{\"message\":\"Your meal request has been confirmed\"}"`
	Read          bool             `json:"read" gorm:"not null;default:false" example:"false"`
	CreatedBy     uint             `json:"created_by" example:"1"`
	UpdatedBy     uint             `json:"updated_by" example:"1"`
	User          User             `json:"user" gorm:"foreignKey:UserID" swaggerignore:"true"`
	CreatedByUser User             `json:"created_by_user" gorm:"foreignKey:CreatedBy" swaggerignore:"true"`
	UpdatedByUser User             `json:"updated_by_user" gorm:"foreignKey:UpdatedBy" swaggerignore:"true"`
}

// NotificationType represents the type of notification
// @Description Type of notification (reminder, confirmation, or admin message)
type NotificationType string

const (
	NotificationTypeReminder     NotificationType = "reminder"
	NotificationTypeConfirmation NotificationType = "confirmation"
	NotificationTypeAdminMessage NotificationType = "admin-message"
)
