package model

// EventAddress represents a location where meal events can be held
type EventAddress struct {
	Base
	Address   string `json:"address" gorm:"not null"`
	Label     string `json:"label" gorm:"not null"`
}
