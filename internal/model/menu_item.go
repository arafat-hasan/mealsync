package model

// MenuItem represents a menu item in a restaurant
type MenuItem struct {
	Base
	RestaurantID uint          `json:"restaurant_id" gorm:"not null"`
	Name         string        `json:"name" gorm:"not null"`
	Description  string        `json:"description"`
	Price        float64       `json:"price" gorm:"not null"`
	Image        string        `json:"image"`
	Category     string        `json:"category" gorm:"not null"`
	MealType     MealType      `json:"meal_type" gorm:"not null"` // breakfast, lunch, snacks
	SetName      string        `json:"set_name"`                  // Set A, Set B, etc. (for lunch only)
	IsAvailable  bool          `json:"is_available" gorm:"not null;default:true"`
	Restaurant   Restaurant    `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	OrderItems   []OrderItem   `json:"order_items" gorm:"foreignKey:MenuItemID"`
	MealRequests []MealRequest `json:"meal_requests" gorm:"foreignKey:MenuItemID"`
}

// MealType represents the type of meal
type MealType string

const (
	MealTypeBreakfast MealType = "breakfast"
	MealTypeLunch     MealType = "lunch"
	MealTypeSnacks    MealType = "snacks"
)
