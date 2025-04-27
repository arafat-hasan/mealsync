package model

import "time"

// Base contains common fields for all models
type Base struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	IsActive    bool       `json:"is_active" gorm:"default:true"`
	CreatedByID uint       `json:"created_by_id" gorm:"column:created_by"`
	CreatedBy   *User      `gorm:"foreignKey:CreatedByID" json:"created_by,omitempty"`
	UpdatedByID uint       `json:"updated_by_id" gorm:"column:updated_by"`
	UpdatedBy   *User      `gorm:"foreignKey:UpdatedByID" json:"updated_by,omitempty"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusAccepted  OrderStatus = "accepted"
	OrderStatusRejected  OrderStatus = "rejected"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// UserRole represents the role of a user
type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleEmployee UserRole = "employee"
	UserRoleManager  UserRole = "manager"
)
