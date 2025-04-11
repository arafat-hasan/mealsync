package model

// Notification represents a notification entity in the system
type Notification struct {
	Base
	UserID  uint             `json:"user_id" gorm:"not null"`
	Title   string           `json:"title" gorm:"not null"`
	Message string           `json:"message" gorm:"not null"`
	Type    NotificationType `json:"type" gorm:"not null"`
	Read    bool             `json:"read" gorm:"not null;default:false"`
	User    User             `json:"user" gorm:"foreignKey:UserID"`
}

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeMealRequest      NotificationType = "meal_request"
	NotificationTypeMealConfirmation NotificationType = "meal_confirmation"
	NotificationTypeMealReminder     NotificationType = "meal_reminder"
	NotificationTypeSystem           NotificationType = "system"
)
