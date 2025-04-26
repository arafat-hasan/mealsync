package model

// MenuItem represents a menu item in the system
type MenuItem struct {
	Base
	Name             string            `json:"name" gorm:"not null;unique"`
	Description      string            `json:"description"`
	ImageURL         string            `json:"image_url"`
	IsActive         bool              `json:"is_active" gorm:"default:true"`
	CreatedBy        User              `json:"created_by" gorm:"foreignKey:CreatedBy"`
	UpdatedBy        User              `json:"updated_by" gorm:"foreignKey:UpdatedBy"`
	MenuSetItems     []MenuSetItem     `json:"menu_set_items" gorm:"foreignKey:MenuItemID"`
	MealRequestItems []MealRequestItem `json:"meal_request_items" gorm:"foreignKey:MenuItemID"`
	MenuItemComments []MenuItemComment `json:"menu_item_comments" gorm:"foreignKey:MenuItemID"`
	AverageRating    float64           `json:"average_rating" gorm:"type:numeric(3,2);default:0"`
}

// MealType represents the type of meal
type MealType string

const (
	MealTypeBreakfast MealType = "breakfast"
	MealTypeLunch     MealType = "lunch"
	MealTypeSnacks    MealType = "snacks"
)
