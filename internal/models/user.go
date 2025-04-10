package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleEmployee Role = "employee"
)

type User struct {
	gorm.Model
	Email       string    `gorm:"uniqueIndex;not null" json:"email"`
	Password    string    `json:"-"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        Role      `gorm:"type:varchar(20);not null" json:"role"`
	Department  string    `json:"department"`
	EmployeeID  string    `gorm:"uniqueIndex" json:"employee_id"`
	LastLoginAt time.Time `json:"last_login_at"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
}
