package controllers

import "meli-products-api/domain"

// ProductComparisonResponse representa la respuesta para la API de comparación de productos
// @Description Response model for product comparison
type ProductComparisonResponse struct {
	// Lista de productos que se están comparando
	Products []*domain.Product `json:"products"`
	
	// Número total de productos en la comparación
	TotalCount int `json:"total_count" example:"3"`
	
	// Lista de IDs de productos solicitados
	RequestedIDs []string `json:"requested_ids"`
}

// ProductSearchResponse representa la respuesta para la API de búsqueda de productos
// @Description Response model for product search
type ProductSearchResponse struct {
	// Lista de productos que coinciden con la búsqueda
	Products []*domain.Product `json:"products"`
	
	// Consulta de búsqueda utilizada
	Query string `json:"query" example:"Samsung Galaxy"`
	
	// Número de productos encontrados
	Count int `json:"count" example:"2"`
}

// CategoriesResponse representa la respuesta para la API de categorías
// @Description Response model for categories
// @Example ["Smartphones", "Laptops", "Audífonos"]
type CategoriesResponse []string

// BrandsResponse representa la respuesta para la API de marcas  
// @Description Response model for brands
// @Example ["Samsung", "Apple", "Google", "Dell", "Sony"]
type BrandsResponse []string

// HealthResponse representa la respuesta para la API de health check
// @Description Response model for health check
type HealthResponse struct {
	// Estado del servicio
	Status string `json:"status" example:"healthy"`
	
	// Timestamp actual
	Timestamp string `json:"timestamp" example:"2024-01-15T10:30:00Z"`
	
	// Nombre del servicio
	Service string `json:"service" example:"meli-products-api"`
	
	// Versión del servicio
	Version string `json:"version" example:"1.0.0"`
}
