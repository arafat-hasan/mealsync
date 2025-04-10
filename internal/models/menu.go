package models

import (
	"time"

	"gorm.io/gorm"
)

type MenuItemType string

const (
	MenuItemTypeLunch MenuItemType = "lunch"
	MenuItemTypeSnack MenuItemType = "snack"
)

type MenuItem struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        MenuItemType `gorm:"type:varchar(20);not null" json:"type"`
	Date        time.Time    `json:"date"`
	IsActive    bool         `gorm:"default:true" json:"is_active"`
}

type MealRequest struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	MenuItemID  uint      `json:"menu_item_id"`
	MenuItem    MenuItem  `json:"menu_item"`
	Date        time.Time `json:"date"`
	Status      string    `gorm:"type:varchar(20);not null" json:"status"` // requested, cancelled
	RequestedAt time.Time `json:"requested_at"`
}
