package controllers

import "meli-products-api/domain"

// ProductComparisonResponse represents the response for product comparison API
// @Description Response model for product comparison
type ProductComparisonResponse struct {
	// List of products being compared
	Products []*domain.Product `json:"products"`
	
	// Total number of products in the comparison
	TotalCount int `json:"total_count" example:"3"`
	
	// List of requested product IDs
	RequestedIDs []string `json:"requested_ids"`
}

// ProductSearchResponse represents the response for product search API
// @Description Response model for product search
type ProductSearchResponse struct {
	// List of products matching the search
	Products []*domain.Product `json:"products"`
	
	// Search query used
	Query string `json:"query" example:"Samsung Galaxy"`
	
	// Number of products found
	Count int `json:"count" example:"2"`
}

// CategoriesResponse represents the response for categories API
// @Description Response model for categories
// @Example ["Smartphones", "Laptops", "Audífonos"]
type CategoriesResponse []string

// BrandsResponse represents the response for brands API  
// @Description Response model for brands
// @Example ["Samsung", "Apple", "Google", "Dell", "Sony"]
type BrandsResponse []string

// HealthResponse represents the response for health check API
// @Description Response model for health check
type HealthResponse struct {
	// Service status
	Status string `json:"status" example:"healthy"`
	
	// Current timestamp
	Timestamp string `json:"timestamp" example:"2024-01-15T10:30:00Z"`
	
	// Service name
	Service string `json:"service" example:"meli-products-api"`
	
	// Service version
	Version string `json:"version" example:"1.0.0"`
}
