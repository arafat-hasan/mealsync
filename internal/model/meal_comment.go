package model

// MealComment represents a comment on a meal for a specific event
type MealComment struct {
	Base
	UserID         uint             `json:"user_id" gorm:"not null"`
	MealEventID    uint             `json:"meal_event_id" gorm:"not null"`
	EventMenuSetID uint             `json:"event_menu_set_id"`
	MenuItemID     uint             `json:"menu_item_id"`
	Comment        string           `json:"comment" gorm:"not null"`
	Rating         int              `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	CreatedBy      uint             `json:"created_by"`
	UpdatedBy      uint             `json:"updated_by"`
	User           User             `json:"user" gorm:"foreignKey:UserID"`
	MealEvent      MealEvent        `json:"meal_event" gorm:"foreignKey:MealEventID"`
	EventMenuSet   MealEventMenuSet `json:"event_menu_set" gorm:"foreignKey:EventMenuSetID"`
	MenuItem       MenuItem         `json:"menu_item" gorm:"foreignKey:MenuItemID"`
	CreatedByUser  User             `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	UpdatedByUser  User             `json:"updated_by_user" gorm:"foreignKey:UpdatedBy"`
}
