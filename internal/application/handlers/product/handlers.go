package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

// GetProductHandler handles GetProductQuery requests
type GetProductHandler struct {
	repo domain.ProductRepository
}

// NewGetProductHandler creates a new GetProductHandler
func NewGetProductHandler(repo domain.ProductRepository) *GetProductHandler {
	return &GetProductHandler{repo: repo}
}

// Handle processes GetProductQuery and returns a single product
func (h *GetProductHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.GetProductQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetProductHandler")
	}

	return h.repo.GetByID(query.ID)
}

// GetAllProductsHandler handles GetAllProductsQuery requests
type GetAllProductsHandler struct {
	repo domain.ProductRepository
}

// NewGetAllProductsHandler creates a new GetAllProductsHandler
func NewGetAllProductsHandler(repo domain.ProductRepository) *GetAllProductsHandler {
	return &GetAllProductsHandler{repo: repo}
}

// Handle processes GetAllProductsQuery and returns filtered products
func (h *GetAllProductsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.GetAllProductsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetAllProductsHandler")
	}

	return h.repo.GetAll(query.Category, query.MinPrice, query.MaxPrice)
}

// CompareProductsHandler handles CompareProductsQuery requests
type CompareProductsHandler struct {
	repo domain.ProductRepository
}

// NewCompareProductsHandler creates a new CompareProductsHandler
func NewCompareProductsHandler(repo domain.ProductRepository) *CompareProductsHandler {
	return &CompareProductsHandler{repo: repo}
}

// Handle processes CompareProductsQuery and returns products for comparison
func (h *CompareProductsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.CompareProductsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for CompareProductsHandler")
	}

	products, err := h.repo.GetByIDs(query.ProductIDs)
	if err != nil {
		return nil, fmt.Errorf("error retrieving products for comparison: %w", err)
	}

	// Return comparison response with additional metadata
	result := struct {
		Products     []*domain.Product `json:"products"`
		TotalCount   int               `json:"total_count"`
		RequestedIDs []string          `json:"requested_ids"`
	}{
		Products:     products,
		TotalCount:   len(products),
		RequestedIDs: query.ProductIDs,
	}

	return result, nil
}

// SearchProductsHandler handles SearchProductsQuery requests
type SearchProductsHandler struct {
	repo domain.ProductRepository
}

// NewSearchProductsHandler creates a new SearchProductsHandler
func NewSearchProductsHandler(repo domain.ProductRepository) *SearchProductsHandler {
	return &SearchProductsHandler{repo: repo}
}

// Handle processes SearchProductsQuery and returns matching products
func (h *SearchProductsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.SearchProductsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for SearchProductsHandler")
	}

	products, err := h.repo.Search(query.Query)
	if err != nil {
		return nil, err
	}

	result := struct {
		Products []*domain.Product `json:"products"`
		Query    string            `json:"query"`
		Count    int               `json:"count"`
	}{
		Products: products,
		Query:    query.Query,
		Count:    len(products),
	}

	return result, nil
}

// GetCategoriesHandler handles GetCategoriesQuery requests
type GetCategoriesHandler struct {
	repo interface {
		GetCategories() []string
	}
}

// NewGetCategoriesHandler creates a new GetCategoriesHandler
func NewGetCategoriesHandler(repo interface {
	GetCategories() []string
}) *GetCategoriesHandler {
	return &GetCategoriesHandler{repo: repo}
}

// Handle processes GetCategoriesQuery and returns available categories
func (h *GetCategoriesHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	_, ok := request.(*product.GetCategoriesQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetCategoriesHandler")
	}

	return h.repo.GetCategories(), nil
}

// GetBrandsHandler handles GetBrandsQuery requests
type GetBrandsHandler struct {
	repo interface {
		GetBrands() []string
	}
}

// NewGetBrandsHandler creates a new GetBrandsHandler
func NewGetBrandsHandler(repo interface {
	GetBrands() []string
}) *GetBrandsHandler {
	return &GetBrandsHandler{repo: repo}
}

// Handle processes GetBrandsQuery and returns available brands
func (h *GetBrandsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	_, ok := request.(*product.GetBrandsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetBrandsHandler")
	}

	return h.repo.GetBrands(), nil
}
