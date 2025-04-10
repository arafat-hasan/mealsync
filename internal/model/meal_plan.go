package model

import "time"

// MealPlan represents a meal plan entity in the system
type MealPlan struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	StartDate time.Time      `json:"start_date" gorm:"not null"`
	EndDate   time.Time      `json:"end_date" gorm:"not null"`
	Status    string         `json:"status" gorm:"not null;default:'active'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	MealItems []MealPlanItem `json:"meal_items" gorm:"foreignKey:MealPlanID"`
}

// MealPlanItem represents a meal plan item entity in the system
type MealPlanItem struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MealPlanID uint      `json:"meal_plan_id" gorm:"not null"`
	MenuItemID uint      `json:"menu_item_id" gorm:"not null"`
	DayOfWeek  int       `json:"day_of_week" gorm:"not null"`
	MealType   string    `json:"meal_type" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	MealPlan   MealPlan  `json:"meal_plan" gorm:"foreignKey:MealPlanID"`
	MenuItem   MenuItem  `json:"menu_item" gorm:"foreignKey:MenuItemID"`
}
