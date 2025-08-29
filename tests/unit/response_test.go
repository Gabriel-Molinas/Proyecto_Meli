package unit

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"meli-products-api/domain"
	"meli-products-api/pkg/response"
)

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	
	apiResponse := &response.APIResponse{
		Success: true,
		Message: "Test message",
		Data:    "test data",
	}
	
	response.JSON(w, http.StatusOK, apiResponse)
	
	// Verificar código de estado
	if w.Code != http.StatusOK {
		t.Errorf("JSON() status code = %v, want %v", w.Code, http.StatusOK)
	}
	
	// Verificar header Content-Type
	expectedContentType := "application/json"
	contentType := w.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("JSON() Content-Type = %v, want %v", contentType, expectedContentType)
	}
	
	// Verificar respuesta JSON
	var result response.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("JSON() failed to unmarshal response: %v", err)
		return
	}
	
	if result.Success != true {
		t.Errorf("JSON() Success = %v, want true", result.Success)
	}
	
	if result.Message != "Test message" {
		t.Errorf("JSON() Message = %v, want 'Test message'", result.Message)
	}
	
	if result.Data != "test data" {
		t.Errorf("JSON() Data = %v, want 'test data'", result.Data)
	}
}

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	testData := map[string]interface{}{
		"id":   "TEST001",
		"name": "Test Product",
	}
	
	response.Success(w, testData, "Operation successful")
	
	if w.Code != http.StatusOK {
		t.Errorf("Success() status code = %v, want %v", w.Code, http.StatusOK)
	}
	
	var result response.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("Success() failed to unmarshal response: %v", err)
		return
	}
	
	if result.Success != true {
		t.Errorf("Success() Success = %v, want true", result.Success)
	}
	
	if result.Message != "Operation successful" {
		t.Errorf("Success() Message = %v, want 'Operation successful'", result.Message)
	}
	
	if result.Error != nil {
		t.Errorf("Success() Error = %v, want nil", result.Error)
	}
}

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	
	response.BadRequest(w, "INVALID_INPUT", "Invalid input provided", "Please check your request parameters")
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("BadRequest() status code = %v, want %v", w.Code, http.StatusBadRequest)
	}
	
	var result response.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("BadRequest() failed to unmarshal response: %v", err)
		return
	}
	
	if result.Success != false {
		t.Errorf("BadRequest() Success = %v, want false", result.Success)
	}
	
	if result.Message != "Bad Request" {
		t.Errorf("BadRequest() Message = %v, want 'Bad Request'", result.Message)
	}
	
	if result.Error == nil {
		t.Error("BadRequest() Error is nil")
		return
	}
	
	if result.Error.Code != "INVALID_INPUT" {
		t.Errorf("BadRequest() Error.Code = %v, want 'INVALID_INPUT'", result.Error.Code)
	}
	
	if result.Error.Message != "Invalid input provided" {
		t.Errorf("BadRequest() Error.Message = %v, want 'Invalid input provided'", result.Error.Message)
	}
	
	if result.Error.Details != "Please check your request parameters" {
		t.Errorf("BadRequest() Error.Details = %v, want 'Please check your request parameters'", result.Error.Details)
	}
}

func TestHandleError(t *testing.T) {
	t.Run("ProductNotFoundError", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := &domain.ProductNotFoundError{ID: "NONEXISTENT"}
		
		response.HandleError(w, err)
		
		if w.Code != http.StatusNotFound {
			t.Errorf("HandleError() status code = %v, want %v", w.Code, http.StatusNotFound)
		}
		
		var result response.APIResponse
		jsonErr := json.Unmarshal(w.Body.Bytes(), &result)
		if jsonErr != nil {
			t.Errorf("HandleError() failed to unmarshal response: %v", jsonErr)
			return
		}
		
		if result.Error.Code != "PRODUCT_NOT_FOUND" {
			t.Errorf("HandleError() Error.Code = %v, want 'PRODUCT_NOT_FOUND'", result.Error.Code)
		}
	})
	
	t.Run("InvalidProductIDError", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := &domain.InvalidProductIDError{ID: ""}
		
		response.HandleError(w, err)
		
		if w.Code != http.StatusBadRequest {
			t.Errorf("HandleError() status code = %v, want %v", w.Code, http.StatusBadRequest)
		}
		
		var result response.APIResponse
		jsonErr := json.Unmarshal(w.Body.Bytes(), &result)
		if jsonErr != nil {
			t.Errorf("HandleError() failed to unmarshal response: %v", jsonErr)
			return
		}
		
		if result.Error.Code != "INVALID_PRODUCT_ID" {
			t.Errorf("HandleError() Error.Code = %v, want 'INVALID_PRODUCT_ID'", result.Error.Code)
		}
	})
	
	t.Run("ValidationError", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := &domain.ValidationError{Field: "price", Message: "must be positive"}
		
		response.HandleError(w, err)
		
		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("HandleError() status code = %v, want %v", w.Code, http.StatusUnprocessableEntity)
		}
		
		var result response.APIResponse
		jsonErr := json.Unmarshal(w.Body.Bytes(), &result)
		if jsonErr != nil {
			t.Errorf("HandleError() failed to unmarshal response: %v", jsonErr)
			return
		}
		
		if result.Error.Code != "VALIDATION_ERROR" {
			t.Errorf("HandleError() Error.Code = %v, want 'VALIDATION_ERROR'", result.Error.Code)
		}
	})
	
	t.Run("Generic error", func(t *testing.T) {
		w := httptest.NewRecorder()
		err := errors.New("unexpected error")
		
		response.HandleError(w, err)
		
		if w.Code != http.StatusInternalServerError {
			t.Errorf("HandleError() status code = %v, want %v", w.Code, http.StatusInternalServerError)
		}
		
		var result response.APIResponse
		jsonErr := json.Unmarshal(w.Body.Bytes(), &result)
		if jsonErr != nil {
			t.Errorf("HandleError() failed to unmarshal response: %v", jsonErr)
			return
		}
		
		if result.Error.Code != "INTERNAL_ERROR" {
			t.Errorf("HandleError() Error.Code = %v, want 'INTERNAL_ERROR'", result.Error.Code)
		}
	})
}