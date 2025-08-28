package product

import (
	"context"
	"fmt"

	"meli-products-api/internal/application/queries/product"
)

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