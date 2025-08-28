package product

// GetProductQuery represents a query to get a single product by ID
type GetProductQuery struct {
	ID string `json:"id" validate:"required" example:"PHONE001"`
}

// GetAllProductsQuery represents a query to get all products with optional filters
type GetAllProductsQuery struct {
	Category string  `json:"category,omitempty" example:"Smartphones"`
	MinPrice float64 `json:"min_price,omitempty" example:"0"`
	MaxPrice float64 `json:"max_price,omitempty" example:"2000"`
}

// CompareProductsQuery represents a query to compare multiple products
type CompareProductsQuery struct {
	ProductIDs []string `json:"product_ids" validate:"required,min=2" example:"PHONE001,PHONE002"`
}

// SearchProductsQuery represents a query to search products
type SearchProductsQuery struct {
	Query string `json:"query" validate:"required" example:"Samsung Galaxy"`
}

// GetCategoriesQuery represents a query to get all available categories
type GetCategoriesQuery struct{}

// GetBrandsQuery represents a query to get all available brands
type GetBrandsQuery struct{}
