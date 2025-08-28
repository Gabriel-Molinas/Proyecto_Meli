package response

import (
	"encoding/json"
	"net/http"

	"meli-products-api/domain"
)

// APIResponse representa la estructura estándar de respuesta de la API
type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Request completed successfully"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// ErrorInfo representa la información de error en la respuesta
type ErrorInfo struct {
	Code    string `json:"code" example:"PRODUCT_NOT_FOUND"`
	Message string `json:"message" example:"Product with ID 'INVALID_ID' not found"`
	Details string `json:"details,omitempty" example:"Please check the product ID and try again"`
}

// Meta representa la información de metadatos en la respuesta
type Meta struct {
	Timestamp  string `json:"timestamp" example:"2024-01-15T10:30:00Z"`
	RequestID  string `json:"request_id,omitempty" example:"req-12345-abcde"`
	Version    string `json:"version" example:"v1"`
	TotalCount int    `json:"total_count,omitempty" example:"150"`
	Page       int    `json:"page,omitempty" example:"1"`
	PageSize   int    `json:"page_size,omitempty" example:"20"`
	TotalPages int    `json:"total_pages,omitempty" example:"8"`
}

// JSON envía una respuesta JSON con el código de estado dado
func JSON(w http.ResponseWriter, statusCode int, response *APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Success envía una respuesta JSON exitosa
func Success(w http.ResponseWriter, data interface{}, message string) {
	JSON(w, http.StatusOK, &APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessWithMeta envía una respuesta JSON exitosa con metadatos
func SuccessWithMeta(w http.ResponseWriter, data interface{}, message string, meta *Meta) {
	JSON(w, http.StatusOK, &APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Created envía una respuesta 201 Created
func Created(w http.ResponseWriter, data interface{}, message string) {
	JSON(w, http.StatusCreated, &APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// BadRequest envía una respuesta 400 Bad Request
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

// NotFound envía una respuesta 404 Not Found
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

// InternalServerError envía una respuesta 500 Internal Server Error
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

// Unauthorized envía una respuesta 401 Unauthorized
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

// Forbidden envía una respuesta 403 Forbidden
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

// ValidationError envía una respuesta 422 Unprocessable Entity para errores de validación
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

// HandleError analiza un error y envía la respuesta HTTP apropiada
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
