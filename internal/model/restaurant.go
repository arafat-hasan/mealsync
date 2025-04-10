package model

import "time"

// Restaurant represents a restaurant in the system
type Restaurant struct {
	Base
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	Address     string     `json:"address" gorm:"not null"`
	Phone       string     `json:"phone" gorm:"not null"`
	Email       string     `json:"email" gorm:"not null"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	MenuItems   []MenuItem `json:"menu_items" gorm:"foreignKey:RestaurantID"`
	Orders      []Order    `json:"orders" gorm:"foreignKey:RestaurantID"`
}
