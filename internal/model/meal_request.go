package model

import "time"

// MealRequest represents a meal request entity in the system
type MealRequest struct {
	Base
	UserID       uint          `json:"user_id" gorm:"not null"`
	MenuItemID   uint          `json:"menu_item_id" gorm:"not null"`
	MenuID       uint          `json:"menu_id" gorm:"not null"`
	Quantity     int           `json:"quantity" gorm:"not null;default:1"`
	Status       RequestStatus `json:"status" gorm:"not null;default:'pending'"`
	RequestedFor time.Time     `json:"requested_for" gorm:"not null"`
	Notes        string        `json:"notes"`
	User         User          `json:"user" gorm:"foreignKey:UserID"`
	MenuItem     MenuItem      `json:"menu_item" gorm:"foreignKey:MenuItemID"`
	Menu         MealMenu      `json:"menu" gorm:"foreignKey:MenuID"`
}

// RequestStatus represents the status of a meal request
type RequestStatus string

const (
	RequestStatusPending   RequestStatus = "pending"
	RequestStatusApproved  RequestStatus = "approved"
	RequestStatusRejected  RequestStatus = "rejected"
	RequestStatusCompleted RequestStatus = "completed"
	RequestStatusCancelled RequestStatus = "cancelled"
)
