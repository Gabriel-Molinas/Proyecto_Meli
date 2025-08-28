package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

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