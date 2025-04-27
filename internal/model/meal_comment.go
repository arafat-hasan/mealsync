package model

// MenuItemComment represents a comment on a menu item for a specific event
type MenuItemComment struct {
	Base
	UserID      uint              `json:"user_id" gorm:"not null"`
	MealEventID uint              `json:"meal_event_id" gorm:"not null"`
	MenuItemID  uint              `json:"menu_item_id" gorm:"not null"`
	Comment     string            `json:"comment" gorm:"not null"`
	Rating      int               `json:"rating" gorm:"check:rating >= 1 AND rating <= 5;not null"`
	ParentID    *uint             `json:"parent_id" gorm:"default:null"` // Added to support replies
	MealEvent   MealEvent         `json:"meal_event" gorm:"foreignKey:MealEventID"`
	MenuItem    MenuItem          `json:"menu_item" gorm:"foreignKey:MenuItemID"`
	Parent      *MenuItemComment  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies     []MenuItemComment `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
}
