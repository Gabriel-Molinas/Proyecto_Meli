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