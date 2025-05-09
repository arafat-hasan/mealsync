package dto

import "time"

// MealEventResponse represents a meal event in API responses
type MealEventResponse struct {
	BaseAuditResponse
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	EventDate     time.Time                  `json:"event_date"`
	EventDuration int                        `json:"event_duration"` // in minutes
	CutoffTime    time.Time                  `json:"cutoff_time"`
	ConfirmedAt   *time.Time                 `json:"confirmed_at,omitempty"`
	MenuSets      []MealEventSetResponse     `json:"menu_sets,omitempty"`
	Addresses     []MealEventAddressResponse `json:"addresses,omitempty"`
}

// MealEventCreateRequest represents data for creating a new meal event
type MealEventCreateRequest struct {
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description"`
	EventDate     time.Time `json:"event_date" binding:"required"`
	EventDuration int       `json:"event_duration" binding:"required,min=1"` // in minutes
	CutoffTime    time.Time `json:"cutoff_time" binding:"required"`
	MenuSetIDs    []uint    `json:"menu_set_ids,omitempty"`
	AddressIDs    []uint    `json:"address_ids,omitempty"`
}

// MealEventUpdateRequest represents data for updating an existing meal event
type MealEventUpdateRequest struct {
	Name          *string    `json:"name"`
	Description   *string    `json:"description"`
	EventDate     *time.Time `json:"event_date"`
	EventDuration *int       `json:"event_duration" binding:"omitempty,min=1"` // in minutes
	CutoffTime    *time.Time `json:"cutoff_time"`
	MenuSetIDs    []uint     `json:"menu_set_ids,omitempty"`
	AddressIDs    []uint     `json:"address_ids,omitempty"`
}

// MealEventSetResponse represents a junction between meal events and menu sets
type MealEventSetResponse struct {
	BaseAuditResponse
	MealEventID uint             `json:"meal_event_id"`
	MenuSetID   uint             `json:"menu_set_id"`
	Label       string           `json:"label"`
	Note        string           `json:"note"`
	MenuSet     *MenuSetResponse `json:"menu_set,omitempty"`
}

// MealEventSetCreateRequest represents data for adding a menu set to a meal event
type MealEventSetCreateRequest struct {
	MenuSetID uint   `json:"menu_set_id" binding:"required"`
	Label     string `json:"label"`
	Note      string `json:"note"`
}

// MealEventAddressResponse represents an address for a meal event
type MealEventAddressResponse struct {
	BaseAuditResponse
	MealEventID uint                  `json:"meal_event_id"`
	AddressID   uint                  `json:"address_id"`
	Address     *EventAddressResponse `json:"address,omitempty"`
}
