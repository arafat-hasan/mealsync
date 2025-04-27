package model

// MenuSet represents a set of menu items
type MenuSet struct {
	Base
	MenuSetName        string         `json:"menu_set_name" gorm:"not null"`
	MenuSetDescription string         `json:"menu_set_description"`
	MenuSetItems       []MenuSetItem  `json:"menu_set_items" gorm:"foreignKey:MenuSetID"`
	MealEventSets      []MealEventSet `json:"meal_event_sets" gorm:"foreignKey:MenuSetID"`
}

// MenuSetItem represents a menu item in a menu set
type MenuSetItem struct {
	Base
	MenuSetID  uint     `json:"menu_set_id" gorm:"not null"`
	MenuItemID uint     `json:"menu_item_id" gorm:"not null"`
	MenuSet    MenuSet  `json:"menu_set" gorm:"foreignKey:MenuSetID"`
	MenuItem   MenuItem `json:"menu_item" gorm:"foreignKey:MenuItemID"`
}
