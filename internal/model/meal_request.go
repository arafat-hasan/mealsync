package model

import "time"

// MealRequest represents a meal request entity in the system
type MealRequest struct {
	Base
	UserID         uint              `json:"user_id" gorm:"not null"`
	MealEventID    uint              `json:"meal_event_id" gorm:"not null"`
	MenuSetID      uint              `json:"menu_set_id"`
	EventAddressID uint              `json:"event_address_id"`
	ConfirmedAt    *time.Time        `json:"confirmed_at"`
	User           User              `json:"user" gorm:"foreignKey:UserID"`
	MealEvent      MealEvent         `json:"meal_event" gorm:"foreignKey:MealEventID"`
	MenuSet        MenuSet           `json:"menu_set" gorm:"foreignKey:MenuSetID"`
	EventAddress   EventAddress      `json:"event_address" gorm:"foreignKey:EventAddressID"`
	RequestItems   []MealRequestItem `json:"request_items" gorm:"foreignKey:MealRequestID"`
}

// MealRequestItem represents an item in a meal request
type MealRequestItem struct {
	Base
	MealRequestID uint        `json:"meal_request_id" gorm:"not null"`
	MenuItemID    uint        `json:"menu_item_id" gorm:"not null"`
	MenuSetID     uint        `json:"menu_set_id" gorm:"not null"`
	IsSelected    bool        `json:"is_selected" gorm:"not null;default:true"`
	Quantity      int         `json:"quantity" gorm:"not null;default:1"`
	Notes         string      `json:"notes"`
	MealRequest   MealRequest `json:"meal_request" gorm:"foreignKey:MealRequestID"`
	MenuItem      MenuItem    `json:"menu_item" gorm:"foreignKey:MenuItemID"`
	MenuSet       MenuSet     `json:"menu_set" gorm:"foreignKey:MenuSetID"`
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
