package model

import "time"

// User represents a user in the system
type User struct {
	Base
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"-" gorm:"not null"` // "-" means don't include in JSON
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	Role        UserRole  `json:"role" gorm:"not null;default:'employee'"`
	Department  string    `json:"department" gorm:"not null"`
	EmployeeID  string    `json:"employee_id" gorm:"unique;not null"`
	LastLoginAt time.Time `json:"last_login_at"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	Orders      []Order   `json:"orders" gorm:"foreignKey:UserID"`
}
