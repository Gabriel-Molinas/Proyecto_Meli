package product

// GetProductQuery representa una consulta para obtener un producto por ID
type GetProductQuery struct {
	ID string `json:"id" validate:"required" example:"PHONE001"`
}

// GetAllProductsQuery representa una consulta para obtener todos los productos con filtros opcionales
type GetAllProductsQuery struct {
	Category string  `json:"category,omitempty" example:"Smartphones"`
	MinPrice float64 `json:"min_price,omitempty" example:"0"`
	MaxPrice float64 `json:"max_price,omitempty" example:"2000"`
}

// CompareProductsQuery representa una consulta para comparar múltiples productos
type CompareProductsQuery struct {
	ProductIDs []string `json:"product_ids" validate:"required,min=2" example:"PHONE001,PHONE002"`
}

// SearchProductsQuery representa una consulta para buscar productos
type SearchProductsQuery struct {
	Query string `json:"query" validate:"required" example:"Samsung Galaxy"`
}

// GetCategoriesQuery representa una consulta para obtener todas las categorías disponibles
type GetCategoriesQuery struct{}

// GetBrandsQuery representa una consulta para obtener todas las marcas disponibles
type GetBrandsQuery struct{}
