package model

import "time"

// Order represents an order in the system
type Order struct {
	Base
	UserID       uint        `json:"user_id" gorm:"not null"`
	RestaurantID uint        `json:"restaurant_id" gorm:"not null"`
	Status       OrderStatus `json:"status" gorm:"not null;default:'pending'"`
	TotalAmount  float64     `json:"total_amount" gorm:"not null"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	User         User        `json:"user" gorm:"foreignKey:UserID"`
	Restaurant   Restaurant  `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	OrderItems   []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	Base
	OrderID    uint     `json:"order_id" gorm:"not null"`
	MenuItemID uint     `json:"menu_item_id" gorm:"not null"`
	Quantity   int      `json:"quantity" gorm:"not null"`
	Price      float64  `json:"price" gorm:"not null"`
	Order      Order    `json:"order" gorm:"foreignKey:OrderID"`
	MenuItem   MenuItem `json:"menu_item" gorm:"foreignKey:MenuItemID"`
}
