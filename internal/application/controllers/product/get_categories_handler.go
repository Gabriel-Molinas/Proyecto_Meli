package product

import (
	"context"
	"fmt"

	"meli-products-api/internal/application/queries/product"
)

// GetCategoriesHandler maneja las solicitudes GetCategoriesQuery
type GetCategoriesHandler struct {
	repo interface {
		GetCategories() []string
	}
}

// NewGetCategoriesHandler crea un nuevo GetCategoriesHandler
func NewGetCategoriesHandler(repo interface {
	GetCategories() []string
}) *GetCategoriesHandler {
	return &GetCategoriesHandler{repo: repo}
}

// Handle procesa GetCategoriesQuery y devuelve las categorías disponibles
func (h *GetCategoriesHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	_, ok := request.(*product.GetCategoriesQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetCategoriesHandler")
	}

	return h.repo.GetCategories(), nil
}