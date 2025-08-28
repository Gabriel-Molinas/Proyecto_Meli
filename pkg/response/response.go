package response

import (
	"encoding/json"
	"net/http"

	"meli-products-api/domain"
)

// APIResponse represents a standard API response structure
type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Request completed successfully"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// ErrorInfo represents error information in the response
type ErrorInfo struct {
	Code    string `json:"code" example:"PRODUCT_NOT_FOUND"`
	Message string `json:"message" example:"Product with ID 'INVALID_ID' not found"`
	Details string `json:"details,omitempty" example:"Please check the product ID and try again"`
}

// Meta represents metadata information in the response
type Meta struct {
	Timestamp  string `json:"timestamp" example:"2024-01-15T10:30:00Z"`
	RequestID  string `json:"request_id,omitempty" example:"req-12345-abcde"`
	Version    string `json:"version" example:"v1"`
	TotalCount int    `json:"total_count,omitempty" example:"150"`
	Page       int    `json:"page,omitempty" example:"1"`
	PageSize   int    `json:"page_size,omitempty" example:"20"`
	TotalPages int    `json:"total_pages,omitempty" example:"8"`
}

// JSON sends a JSON response with the given status code
func JSON(w http.ResponseWriter, statusCode int, response *APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Success sends a successful JSON response
func Success(w http.ResponseWriter, data interface{}, message string) {
	JSON(w, http.StatusOK, &APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessWithMeta sends a successful JSON response with metadata
func SuccessWithMeta(w http.ResponseWriter, data interface{}, message string, meta *Meta) {
	JSON(w, http.StatusOK, &APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Created sends a 201 Created response
func Created(w http.ResponseWriter, data interface{}, message string) {
	JSON(w, http.StatusCreated, &APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(w http.ResponseWriter, code, message, details string) {
	JSON(w, http.StatusBadRequest, &APIResponse{
		Success: false,
		Message: "Bad Request",
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// NotFound sends a 404 Not Found response
func NotFound(w http.ResponseWriter, code, message, details string) {
	JSON(w, http.StatusNotFound, &APIResponse{
		Success: false,
		Message: "Resource Not Found",
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(w http.ResponseWriter, code, message, details string) {
	JSON(w, http.StatusInternalServerError, &APIResponse{
		Success: false,
		Message: "Internal Server Error",
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(w http.ResponseWriter, code, message, details string) {
	JSON(w, http.StatusUnauthorized, &APIResponse{
		Success: false,
		Message: "Unauthorized",
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// Forbidden sends a 403 Forbidden response
func Forbidden(w http.ResponseWriter, code, message, details string) {
	JSON(w, http.StatusForbidden, &APIResponse{
		Success: false,
		Message: "Forbidden",
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// ValidationError sends a 422 Unprocessable Entity response for validation errors
func ValidationError(w http.ResponseWriter, code, message, details string) {
	JSON(w, http.StatusUnprocessableEntity, &APIResponse{
		Success: false,
		Message: "Validation Error",
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// HandleError analyzes an error and sends the appropriate HTTP response
func HandleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *domain.ProductNotFoundError:
		NotFound(w, "PRODUCT_NOT_FOUND", e.Error(), "Please verify the product ID and try again")
	case *domain.InvalidProductIDError:
		BadRequest(w, "INVALID_PRODUCT_ID", e.Error(), "Product ID must be a valid non-empty string")
	case *domain.ValidationError:
		ValidationError(w, "VALIDATION_ERROR", e.Error(), "Please check your input and try again")
	default:
		InternalServerError(w, "INTERNAL_ERROR", "An unexpected error occurred", "Please try again later or contact support if the problem persists")
	}
}
