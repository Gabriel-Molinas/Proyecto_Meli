package product

import (
	"context"
	"fmt"

	"meli-products-api/internal/application/queries/product"
)

// GetBrandsHandler maneja las solicitudes GetBrandsQuery
type GetBrandsHandler struct {
	repo interface {
		GetBrands() []string
	}
}

// NewGetBrandsHandler crea un nuevo GetBrandsHandler
func NewGetBrandsHandler(repo interface {
	GetBrands() []string
}) *GetBrandsHandler {
	return &GetBrandsHandler{repo: repo}
}

// Handle procesa GetBrandsQuery y devuelve las marcas disponibles
func (h *GetBrandsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	_, ok := request.(*product.GetBrandsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetBrandsHandler")
	}

	return h.repo.GetBrands(), nil
}