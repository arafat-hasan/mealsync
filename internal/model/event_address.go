package model

// EventAddress represents a location where meal events can be held
type EventAddress struct {
	Base
	Address   string `json:"address" gorm:"not null"`
	Label     string `json:"label" gorm:"not null"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
	CreatedBy User   `json:"created_by" gorm:"foreignKey:CreatedBy"`
	UpdatedBy User   `json:"updated_by" gorm:"foreignKey:UpdatedBy"`
}
