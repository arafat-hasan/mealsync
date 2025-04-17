package model

// EventAddress represents a location where meal events can be held
type EventAddress struct {
	Base
	Address       string `json:"address" gorm:"not null"`
	IsActive      bool   `json:"is_active" gorm:"default:true"`
	CreatedBy     uint   `json:"created_by"`
	UpdatedBy     uint   `json:"updated_by"`
	CreatedByUser User   `json:"created_by_user" gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User   `json:"updated_by_user" gorm:"foreignKey:UpdatedBy"`
}
