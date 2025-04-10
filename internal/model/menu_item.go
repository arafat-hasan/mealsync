package model

// MenuItem represents a menu item in a restaurant
type MenuItem struct {
	Base
	RestaurantID uint        `json:"restaurant_id" gorm:"not null"`
	Name         string      `json:"name" gorm:"not null"`
	Description  string      `json:"description"`
	Price        float64     `json:"price" gorm:"not null"`
	Image        string      `json:"image"`
	Category     string      `json:"category" gorm:"not null"`
	IsAvailable  bool        `json:"is_available" gorm:"not null;default:true"`
	Restaurant   Restaurant  `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	OrderItems   []OrderItem `json:"order_items" gorm:"foreignKey:MenuItemID"`
}
