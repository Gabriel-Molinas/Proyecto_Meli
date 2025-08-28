package domain

import "fmt"

// Product representa una entidad de producto para comparación
// @Description Product model for comparison
type Product struct {
	// Identificador único del producto
	ID string `json:"id" example:"PHONE001"`
	
	// Nombre del producto
	Name string `json:"name" example:"Samsung Galaxy S24 Ultra" validate:"required"`
	
	// URL de la imagen del producto
	ImageURL string `json:"image_url" example:"https://images.example.com/samsung-s24.jpg" validate:"required,url"`
	
	// Descripción detallada del producto
	Description string `json:"description" example:"Latest Samsung flagship smartphone with advanced camera technology" validate:"required"`
	
	// Precio del producto en formato decimal
	Price float64 `json:"price" example:"1299.99" validate:"required,gt=0"`
	
	// Calificación del producto (escala 1-5)
	Rating float32 `json:"rating" example:"4.5" validate:"required,gte=0,lte=5"`
	
	// Especificaciones técnicas del producto
	Specifications []Specification `json:"specifications"`
	
	// Categoría del producto
	Category string `json:"category" example:"Smartphones"`
	
	// Marca del producto
	Brand string `json:"brand" example:"Samsung"`
	
	// Estado de disponibilidad
	Available bool `json:"available" example:"true"`
}

// Specification representa una especificación técnica de un producto
// @Description Technical specification model
type Specification struct {
	// Nombre de la especificación
	Name string `json:"name" example:"Display Size" validate:"required"`
	
	// Valor de la especificación
	Value string `json:"value" example:"6.8 inches" validate:"required"`
	
	// Unidad de medida si aplica
	Unit string `json:"unit,omitempty" example:"inches"`
}

// ProductRepository define la interfaz para acceso a datos de productos
type ProductRepository interface {
	// GetByID obtiene un producto por su ID
	GetByID(id string) (*Product, error)
	
	// GetAll obtiene todos los productos con filtrado opcional
	GetAll(category string, minPrice, maxPrice float64) ([]*Product, error)
	
	// GetByIDs obtiene múltiples productos por sus IDs para comparación
	GetByIDs(ids []string) ([]*Product, error)
	
	// Search busca productos por nombre o descripción
	Search(query string) ([]*Product, error)
}

// ProductNotFoundError representa un error cuando no se encuentra un producto
type ProductNotFoundError struct {
	ID string
}

func (e *ProductNotFoundError) Error() string {
	return fmt.Sprintf("product with ID '%s' not found", e.ID)
}

// InvalidProductIDError representa un error cuando el ID del producto es inválido
type InvalidProductIDError struct {
	ID string
}

func (e *InvalidProductIDError) Error() string {
	return fmt.Sprintf("invalid product ID: '%s'", e.ID)
}

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}
