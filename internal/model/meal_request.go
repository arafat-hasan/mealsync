package model

import "time"

// MealRequest represents a meal request entity in the system
type MealRequest struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	MenuItemID   uint      `json:"menu_item_id" gorm:"not null"`
	Quantity     int       `json:"quantity" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null;default:'pending'"`
	RequestedFor time.Time `json:"requested_for" gorm:"not null"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
	MenuItem     MenuItem  `json:"menu_item" gorm:"foreignKey:MenuItemID"`
}
