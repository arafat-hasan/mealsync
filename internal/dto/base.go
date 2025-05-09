package dto

import "time"

// BaseResponse contains common fields for all response DTOs
type BaseResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// BaseAuditResponse contains common audit fields for response DTOs
type BaseAuditResponse struct {
	BaseResponse
	CreatedByID uint          `json:"created_by_id,omitempty"`
	CreatedBy   *UserResponse `json:"created_by,omitempty"`
	UpdatedByID uint          `json:"updated_by_id,omitempty"`
	UpdatedBy   *UserResponse `json:"updated_by,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error     string `json:"error"`                // User-friendly error message
	Details   string `json:"details,omitempty"`    // Additional details for debugging
	Code      string `json:"code,omitempty"`       // Error code for client-side error handling
	RequestID string `json:"request_id,omitempty"` // Request ID for tracing
}

// PaginationRequest represents pagination parameters for requests
type PaginationRequest struct {
	Page     int `json:"page" form:"page" binding:"omitempty,min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"omitempty,min=1,max=100"`
}

// PaginationResponse represents pagination metadata in responses
type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

// PaginatedResponse wraps any response with pagination metadata
type PaginatedResponse struct {
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string `json:"message"`
}
