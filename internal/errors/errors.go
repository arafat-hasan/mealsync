package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Common errors
var (
	// ErrNotFound represents a not found error
	ErrNotFound = errors.New("record not found")
)

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "VALIDATION_ERROR"
	// ErrorTypeNotFound represents not found errors
	ErrorTypeNotFound ErrorType = "NOT_FOUND"
	// ErrorTypeUnauthorized represents unauthorized errors
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	// ErrorTypeForbidden represents forbidden errors
	ErrorTypeForbidden ErrorType = "FORBIDDEN"
	// ErrorTypeInternal represents internal server errors
	ErrorTypeInternal ErrorType = "INTERNAL_ERROR"
	// ErrorTypeConflict represents conflict errors
	ErrorTypeConflict ErrorType = "CONFLICT"
)

// ErrorResponse represents the error response for the Swagger documentation
// This is used to document the error responses in the Swagger documentation
type ErrorResponse struct {
	Type      ErrorType `json:"type"`
	Message   string    `json:"message"`
	Code      int       `json:"code"`
	Details   string    `json:"details,omitempty"`
	RequestID string    `json:"request_id,omitempty"`
}

// AppError represents an application error
type AppError struct {
	Type      ErrorType `json:"type"`
	Message   string    `json:"message"`
	Code      int       `json:"code"`
	Details   string    `json:"details,omitempty"`
	RequestID string    `json:"request_id,omitempty"`
	Err       error     `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(errType ErrorType, message string, code int, err error) *AppError {
	return &AppError{
		Type:    errType,
		Message: message,
		Code:    code,
		Err:     err,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string, err error) *AppError {
	return New(ErrorTypeValidation, message, http.StatusBadRequest, err)
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string, err error) *AppError {
	return New(ErrorTypeNotFound, message, http.StatusNotFound, err)
}

// NewUnauthorizedError creates a new unauthorized error
func NewUnauthorizedError(message string, err error) *AppError {
	return New(ErrorTypeUnauthorized, message, http.StatusUnauthorized, err)
}

// NewForbiddenError creates a new forbidden error
func NewForbiddenError(message string, err error) *AppError {
	return New(ErrorTypeForbidden, message, http.StatusForbidden, err)
}

// NewInternalError creates a new internal server error
func NewInternalError(message string, err error) *AppError {
	return New(ErrorTypeInternal, message, http.StatusInternalServerError, err)
}

// NewConflictError creates a new conflict error
func NewConflictError(message string, err error) *AppError {
	return New(ErrorTypeConflict, message, http.StatusConflict, err)
}

// WithRequestID adds a request ID to the error
func (e *AppError) WithRequestID(requestID string) *AppError {
	e.RequestID = requestID
	return e
}

// WithDetails adds details to the error
func (e *AppError) WithDetails(details string) *AppError {
	e.Details = details
	return e
}
