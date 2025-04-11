package model

import "time"

// MealComment represents a comment on a meal for a specific date
type MealComment struct {
	Base
	UserID  uint      `json:"user_id" gorm:"not null"`
	Date    time.Time `json:"date" gorm:"not null"`
	Comment string    `json:"comment" gorm:"not null"`
	User    User      `json:"user" gorm:"foreignKey:UserID"`
}
