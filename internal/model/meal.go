package model

import "time"

// Meal represents a meal entity in the system
type Meal struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        string    `json:"date" gorm:"not null"`
	MenuItems   []string  `json:"menu_items" gorm:"type:jsonb"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
