package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

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