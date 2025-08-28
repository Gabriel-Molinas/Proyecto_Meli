package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

// SearchProductsHandler maneja las solicitudes SearchProductsQuery
type SearchProductsHandler struct {
	repo domain.ProductRepository
}

// NewSearchProductsHandler crea un nuevo SearchProductsHandler
func NewSearchProductsHandler(repo domain.ProductRepository) *SearchProductsHandler {
	return &SearchProductsHandler{repo: repo}
}

// Handle procesa SearchProductsQuery y devuelve productos coincidentes
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