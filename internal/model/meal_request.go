package model

import "time"

// MealRequest represents a meal request entity in the system
type MealRequest struct {
	Base
	UserID             uint              `json:"user_id" gorm:"not null"`
	MealEventID        uint              `json:"meal_event_id" gorm:"not null"`
	EventMenuSetID     uint              `json:"event_menu_set_id"`
	MealEventAddressID uint              `json:"meal_event_address_id"`
	ConfirmedAt        *time.Time        `json:"confirmed_at"`
	CreatedBy          uint              `json:"created_by"`
	UpdatedBy          uint              `json:"updated_by"`
	User               User              `json:"user" gorm:"foreignKey:UserID"`
	MealEvent          MealEvent         `json:"meal_event" gorm:"foreignKey:MealEventID"`
	EventMenuSet       MealEventMenuSet  `json:"event_menu_set" gorm:"foreignKey:EventMenuSetID"`
	EventAddress       MealEventAddress  `json:"event_address" gorm:"foreignKey:MealEventAddressID"`
	CreatedByUser      User              `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	UpdatedByUser      User              `json:"updated_by_user" gorm:"foreignKey:UpdatedBy"`
	RequestItems       []MealRequestItem `json:"request_items" gorm:"foreignKey:MealRequestID"`
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
	CreatedBy     uint        `json:"created_by"`
	UpdatedBy     uint        `json:"updated_by"`
	CreatedByUser User        `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User        `json:"updated_by_user" gorm:"foreignKey:UpdatedBy"`
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
