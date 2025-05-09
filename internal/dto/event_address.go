package dto

// EventAddressResponse represents an event address in API responses
type EventAddressResponse struct {
	BaseAuditResponse
	Name        string `json:"name"`
	AddressLine string `json:"address_line"`
}

// EventAddressCreateRequest represents data for creating a new event address
type EventAddressCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	AddressLine string `json:"address_line" binding:"required"`
}

// EventAddressUpdateRequest represents data for updating an existing event address
type EventAddressUpdateRequest struct {
	Name        *string `json:"name"`
	AddressLine *string `json:"address_line"`
}
