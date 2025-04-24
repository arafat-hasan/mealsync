package model

import "time"

// User represents a user in the system
type User struct {
	Base
	EmployeeID          string            `json:"employee_id" gorm:"unique;not null"`
	Username            string            `json:"username" gorm:"unique;not null"`
	PasswordHash        string            `json:"-" gorm:"column:password_hash;not null"` // Map to password_hash column
	Password            string            `json:"-" gorm:"-"`                             // Transient field for password input, not stored
	Name                string            `json:"name" gorm:"not null"`
	Email               string            `json:"email" gorm:"unique;not null"`
	Department          string            `json:"department" gorm:"not null"`
	Role                UserRole          `json:"role" gorm:"not null;default:'employee'"`
	IsActive            bool              `json:"is_active" gorm:"default:true"`
	NotificationEnabled bool              `json:"notification_enabled" gorm:"default:true"`
	LastLoginAt         time.Time         `json:"last_login_at"`
	CreatedBy           uint              `json:"created_by"`
	UpdatedBy           uint              `json:"updated_by"`
	CreatedByUser       *User             `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	UpdatedByUser       *User             `json:"updated_by_user" gorm:"foreignKey:UpdatedBy"`
	MealRequests        []MealRequest     `json:"meal_requests" gorm:"foreignKey:UserID"`
	MenuItemComments    []MenuItemComment `json:"menu_item_comments" gorm:"foreignKey:UserID"`
	Notifications       []Notification    `json:"notifications" gorm:"foreignKey:UserID"`
	CreatedMealEvents   []MealEvent       `json:"created_meal_events" gorm:"foreignKey:CreatedBy"`
	UpdatedMealEvents   []MealEvent       `json:"updated_meal_events" gorm:"foreignKey:UpdatedBy"`
}
