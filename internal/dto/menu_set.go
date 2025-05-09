package dto

// MenuSetResponse represents a menu set in API responses
type MenuSetResponse struct {
	BaseAuditResponse
	Name        string             `json:"name"`
	Description string             `json:"description"`
	MenuItems   []MenuItemResponse `json:"menu_items,omitempty"`
}

// MenuSetCreateRequest represents data for creating a new menu set
type MenuSetCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	MenuItemIDs []uint `json:"menu_item_ids,omitempty"`
}

// MenuSetUpdateRequest represents data for updating an existing menu set
type MenuSetUpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	MenuItemIDs []uint  `json:"menu_item_ids,omitempty"`
}
