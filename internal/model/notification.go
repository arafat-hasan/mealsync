package model

import (
	"encoding/json"
	"time"
)

// Notification represents a notification entity in the system
// @Description Notification entity containing user notifications and their details
type Notification struct {
	Base
	UserID      uint             `json:"user_id" gorm:"not null" example:"1"`
	Type        NotificationType `json:"type" gorm:"not null" example:"reminder" enums:"reminder,confirmation,admin-message,event-info"`
	Payload     json.RawMessage  `json:"payload" gorm:"type:jsonb" swaggertype:"string" example:"{\"message\":\"Your meal request has been confirmed\"}"`
	Message     string           `json:"message" gorm:"not null" example:"Your meal request has been confirmed."`
	Read        bool             `json:"read" gorm:"not null;default:false" example:"false"`
	Delivered   bool             `json:"delivered" gorm:"not null;default:false" example:"false"`
	ReadAt      *time.Time       `json:"read_at" gorm:"default:null" example:"2025-04-24T10:15:00Z"`
	DeliveredAt *time.Time       `json:"delivered_at" gorm:"default:null" example:"2025-04-24T10:00:00Z"`
	User        User             `json:"user" gorm:"foreignKey:UserID" swaggerignore:"true"`
}

// NotificationType represents the type of notification
// @Description Type of notification (reminder, confirmation, admin message, or event info)
type NotificationType string

const (
	NotificationTypeReminder     NotificationType = "reminder"
	NotificationTypeConfirmation NotificationType = "confirmation"
	NotificationTypeAdminMessage NotificationType = "admin-message"
	NotificationTypeEventInfo    NotificationType = "event-info"
)
