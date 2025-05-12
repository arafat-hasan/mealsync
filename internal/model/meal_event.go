package model

import "time"

// MealEvent represents a meal event in the system
type MealEvent struct {
	Base
	Name             string             `json:"name" gorm:"not null"`
	Description      string             `json:"description"`
	EventDate        time.Time          `json:"event_date" gorm:"not null"`
	EventDuration    int                `json:"event_duration" gorm:"not null"` // in minutes
	CutoffTime       time.Time          `json:"cutoff_time" gorm:"not null"`
	IsActive         bool               `json:"is_active" gorm:"default:true"`
	ConfirmedAt      *time.Time         `json:"confirmed_at"`
	MenuSets         []MealEventSet     `json:"menu_sets" gorm:"foreignKey:MealEventID"`
	Addresses        []MealEventAddress `json:"addresses" gorm:"foreignKey:MealEventID"`
	MealRequests     []MealRequest      `json:"meal_requests" gorm:"foreignKey:MealEventID"`
	MenuItemComments []MenuItemComment  `json:"menu_item_comments" gorm:"foreignKey:MealEventID"`
}

// MealEventSet represents a junction table between meal events and menu sets
type MealEventSet struct {
	Base
	MealEventID uint       `json:"meal_event_id" gorm:"primaryKey;not null"`
	MenuSetID   uint       `json:"menu_set_id" gorm:"primaryKey;not null"`
	Label       string     `json:"label"`
	Note        string     `json:"note"`
	MealEvent   *MealEvent `json:"meal_event" gorm:"foreignKey:MealEventID"`
	MenuSet     *MenuSet   `json:"menu_set" gorm:"foreignKey:MenuSetID"`
}

// MealEventAddress represents an address for a meal event
type MealEventAddress struct {
	Base
	MealEventID uint         `json:"meal_event_id" gorm:"not null"`
	AddressID   uint         `json:"address_id" gorm:"not null"`
	MealEvent   MealEvent    `json:"meal_event" gorm:"foreignKey:MealEventID"`
	Address     EventAddress `json:"address" gorm:"foreignKey:AddressID"`
}
