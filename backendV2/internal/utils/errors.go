package utils

import "fmt"

// AppError represents a custom application error with HTTP status code
type AppError struct {
	Code    int
	Message string
	Err     error
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewValidationError creates a new validation error (400 Bad Request)
func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Code:    400,
		Message: message,
		Err:     err,
	}
}

// NewNotFoundError creates a new not found error (404 Not Found)
func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Code:    404,
		Message: fmt.Sprintf("%s not found", resource),
		Err:     nil,
	}
}

// NewConflictError creates a new conflict error (409 Conflict)
func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    409,
		Message: message,
		Err:     nil,
	}
}

// NewInternalError creates a new internal server error (500 Internal Server Error)
func NewInternalError(err error) *AppError {
	return &AppError{
		Code:    500,
		Message: "Internal server error",
		Err:     err,
	}
}
