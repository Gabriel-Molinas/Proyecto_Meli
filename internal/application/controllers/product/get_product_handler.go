package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

// GetProductHandler maneja las solicitudes GetProductQuery
type GetProductHandler struct {
	repo domain.ProductRepository
}

// NewGetProductHandler crea un nuevo GetProductHandler
func NewGetProductHandler(repo domain.ProductRepository) *GetProductHandler {
	return &GetProductHandler{repo: repo}
}

// Handle procesa GetProductQuery y devuelve un producto individual
func (h *GetProductHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.GetProductQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetProductHandler")
	}

	return h.repo.GetByID(query.ID)
}