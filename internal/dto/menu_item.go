package dto

// MenuItemResponse represents a menu item in API responses
type MenuItemResponse struct {
	BaseAuditResponse
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// MenuItemCreateRequest represents data for creating a new menu item
type MenuItemCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=0"`
}

// MenuItemUpdateRequest represents data for updating an existing menu item
type MenuItemUpdateRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,min=0"`
}
