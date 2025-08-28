package product

import (
	"context"
	"fmt"

	"meli-products-api/internal/application/queries/product"
)

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