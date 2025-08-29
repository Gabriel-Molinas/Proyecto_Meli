package unit

import (
	"testing"
	
	"meli-products-api/domain"
)

func TestProductNotFoundError(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "Error con ID válido",
			id:       "PHONE001",
			expected: "product with ID 'PHONE001' not found",
		},
		{
			name:     "Error con ID vacío",
			id:       "",
			expected: "product with ID '' not found",
		},
		{
			name:     "Error con ID especial",
			id:       "TEST@123",
			expected: "product with ID 'TEST@123' not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &domain.ProductNotFoundError{ID: tt.id}
			if err.Error() != tt.expected {
				t.Errorf("ProductNotFoundError.Error() = %v, want %v", err.Error(), tt.expected)
			}
		})
	}
}

func TestInvalidProductIDError(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "Error con ID inválido",
			id:       "INVALID",
			expected: "invalid product ID: 'INVALID'",
		},
		{
			name:     "Error con ID vacío",
			id:       "",
			expected: "invalid product ID: ''",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &domain.InvalidProductIDError{ID: tt.id}
			if err.Error() != tt.expected {
				t.Errorf("InvalidProductIDError.Error() = %v, want %v", err.Error(), tt.expected)
			}
		})
	}
}

func TestValidationError(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		message  string
		expected string
	}{
		{
			name:     "Error de validación simple",
			field:    "price",
			message:  "must be greater than 0",
			expected: "validation error on field 'price': must be greater than 0",
		},
		{
			name:     "Error con campo vacío",
			field:    "",
			message:  "field is required",
			expected: "validation error on field '': field is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &domain.ValidationError{Field: tt.field, Message: tt.message}
			if err.Error() != tt.expected {
				t.Errorf("ValidationError.Error() = %v, want %v", err.Error(), tt.expected)
			}
		})
	}
}

func TestProductStruct(t *testing.T) {
	// Test para verificar que la estructura Product se puede crear correctamente
	product := domain.Product{
		ID:          "TEST001",
		Name:        "Test Product",
		ImageURL:    "https://example.com/test.jpg",
		Description: "Test description",
		Price:       299.99,
		Rating:      4.5,
		Category:    "Electronics",
		Brand:       "TestBrand",
		Available:   true,
		Specifications: []domain.Specification{
			{Name: "Color", Value: "Black", Unit: ""},
		},
	}

	// Verificar que todos los campos se asignen correctamente
	if product.ID != "TEST001" {
		t.Errorf("Expected ID 'TEST001', got %v", product.ID)
	}
	if product.Price != 299.99 {
		t.Errorf("Expected Price 299.99, got %v", product.Price)
	}
	if product.Rating != 4.5 {
		t.Errorf("Expected Rating 4.5, got %v", product.Rating)
	}
	if !product.Available {
		t.Error("Expected Available to be true")
	}
	if len(product.Specifications) != 1 {
		t.Errorf("Expected 1 specification, got %v", len(product.Specifications))
	}
}

func TestSpecificationStruct(t *testing.T) {
	// Test para verificar que la estructura Specification se puede crear correctamente
	spec := domain.Specification{
		Name:  "RAM",
		Value: "8",
		Unit:  "GB",
	}

	if spec.Name != "RAM" {
		t.Errorf("Expected Name 'RAM', got %v", spec.Name)
	}
	if spec.Value != "8" {
		t.Errorf("Expected Value '8', got %v", spec.Value)
	}
	if spec.Unit != "GB" {
		t.Errorf("Expected Unit 'GB', got %v", spec.Unit)
	}
}

func TestSpecificationWithoutUnit(t *testing.T) {
	// Test para especificación sin unidad
	spec := domain.Specification{
		Name:  "Processor",
		Value: "Snapdragon 8 Gen 3",
		Unit:  "",
	}

	if spec.Unit != "" {
		t.Errorf("Expected empty Unit, got %v", spec.Unit)
	}
}