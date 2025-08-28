package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

// CompareProductsHandler maneja las solicitudes CompareProductsQuery
type CompareProductsHandler struct {
	repo domain.ProductRepository
}

// NewCompareProductsHandler crea un nuevo CompareProductsHandler
func NewCompareProductsHandler(repo domain.ProductRepository) *CompareProductsHandler {
	return &CompareProductsHandler{repo: repo}
}

// Handle procesa CompareProductsQuery y devuelve productos para comparación
func (h *CompareProductsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.CompareProductsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for CompareProductsHandler")
	}

	products, err := h.repo.GetByIDs(query.ProductIDs)
	if err != nil {
		return nil, fmt.Errorf("error retrieving products for comparison: %w", err)
	}

	// Devolver respuesta de comparación con metadatos adicionales
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