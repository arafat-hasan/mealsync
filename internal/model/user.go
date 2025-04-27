package model

import "time"

// User represents a user in the system
type User struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	EmployeeID          string    `json:"employee_id" gorm:"unique;not null"`
	Username            string    `json:"username" gorm:"unique;not null"`
	PasswordHash        string    `json:"-" gorm:"column:password_hash;not null"` // Map to password_hash column
	Password            string    `json:"-" gorm:"-"`                             // Transient field for password input, not stored
	Name                string    `json:"name" gorm:"not null"`
	Email               string    `json:"email" gorm:"unique;not null"`
	Department          string    `json:"department" gorm:"not null"`
	Role                UserRole  `json:"role" gorm:"not null;default:'employee'"`
	NotificationEnabled bool      `json:"notification_enabled" gorm:"default:true"`
	LastLoginAt         time.Time `json:"last_login_at"`
	// MealRequests        []MealRequest     `json:"meal_requests" gorm:"foreignKey:UserID"`
	// MenuItemComments    []MenuItemComment `json:"menu_item_comments" gorm:"foreignKey:UserID"`
	// Notifications       []Notification    `json:"notifications" gorm:"foreignKey:UserID"`
	// CreatedMealEvents   []MealEvent       `json:"created_meal_events" gorm:"foreignKey:CreatedBy"`
	// UpdatedMealEvents   []MealEvent       `json:"updated_meal_events" gorm:"foreignKey:UpdatedBy"`

	IsActive    bool       `json:"is_active" gorm:"default:true"`
	CreatedByID uint       `json:"created_by_id" gorm:"column:created_by"`
	UpdatedByID uint       `json:"updated_by_id" gorm:"column:updated_by"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
