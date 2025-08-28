package domain

import "fmt"

// Product represents a product entity for comparison
// @Description Product model for comparison
type Product struct {
	// Unique identifier for the product
	ID string `json:"id" example:"PHONE001"`
	
	// Name of the product
	Name string `json:"name" example:"Samsung Galaxy S24 Ultra" validate:"required"`
	
	// URL of the product image
	ImageURL string `json:"image_url" example:"https://images.example.com/samsung-s24.jpg" validate:"required,url"`
	
	// Detailed description of the product
	Description string `json:"description" example:"Latest Samsung flagship smartphone with advanced camera technology" validate:"required"`
	
	// Price of the product in decimal format
	Price float64 `json:"price" example:"1299.99" validate:"required,gt=0"`
	
	// Rating of the product (1-5 scale)
	Rating float32 `json:"rating" example:"4.5" validate:"required,gte=0,lte=5"`
	
	// Technical specifications of the product
	Specifications []Specification `json:"specifications"`
	
	// Category of the product
	Category string `json:"category" example:"Smartphones"`
	
	// Brand of the product
	Brand string `json:"brand" example:"Samsung"`
	
	// Availability status
	Available bool `json:"available" example:"true"`
}

// Specification represents a technical specification of a product
// @Description Technical specification model
type Specification struct {
	// Name of the specification
	Name string `json:"name" example:"Display Size" validate:"required"`
	
	// Value of the specification
	Value string `json:"value" example:"6.8 inches" validate:"required"`
	
	// Unit of measurement if applicable
	Unit string `json:"unit,omitempty" example:"inches"`
}

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	// GetByID retrieves a product by its ID
	GetByID(id string) (*Product, error)
	
	// GetAll retrieves all products with optional filtering
	GetAll(category string, minPrice, maxPrice float64) ([]*Product, error)
	
	// GetByIDs retrieves multiple products by their IDs for comparison
	GetByIDs(ids []string) ([]*Product, error)
	
	// Search products by name or description
	Search(query string) ([]*Product, error)
}

// ProductNotFoundError represents an error when a product is not found
type ProductNotFoundError struct {
	ID string
}

func (e *ProductNotFoundError) Error() string {
	return fmt.Sprintf("product with ID '%s' not found", e.ID)
}

// InvalidProductIDError represents an error when product ID is invalid
type InvalidProductIDError struct {
	ID string
}

func (e *InvalidProductIDError) Error() string {
	return fmt.Sprintf("invalid product ID: '%s'", e.ID)
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}
