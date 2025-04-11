package model

import "time"

// MealMenu represents a menu for a specific date and meal type
type MealMenu struct {
	Base
	Date          time.Time      `json:"date" gorm:"not null"`
	MealType      MealType       `json:"meal_type" gorm:"not null"`
	CutoffTime    time.Time      `json:"cutoff_time" gorm:"not null"`
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	CreatedBy     uint           `json:"created_by" gorm:"not null"`
	CreatedByUser User           `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	MenuItems     []MealMenuItem `json:"menu_items" gorm:"foreignKey:MealMenuID"`
	MealRequests  []MealRequest  `json:"meal_requests" gorm:"foreignKey:MenuID"`
}

// MealMenuItem represents a menu item in a meal menu
type MealMenuItem struct {
	Base
	MealMenuID uint     `json:"meal_menu_id" gorm:"not null"`
	MenuItemID uint     `json:"menu_item_id" gorm:"not null"`
	SetName    string   `json:"set_name"` // Set A, Set B, etc. (for lunch only)
	MealMenu   MealMenu `json:"meal_menu" gorm:"foreignKey:MealMenuID"`
	MenuItem   MenuItem `json:"menu_item" gorm:"foreignKey:MenuItemID"`
}
