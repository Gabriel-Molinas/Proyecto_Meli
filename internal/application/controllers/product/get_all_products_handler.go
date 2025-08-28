package product

import (
	"context"
	"fmt"

	"meli-products-api/domain"
	"meli-products-api/internal/application/queries/product"
)

// GetAllProductsHandler maneja las solicitudes GetAllProductsQuery
type GetAllProductsHandler struct {
	repo domain.ProductRepository
}

// NewGetAllProductsHandler crea un nuevo GetAllProductsHandler
func NewGetAllProductsHandler(repo domain.ProductRepository) *GetAllProductsHandler {
	return &GetAllProductsHandler{repo: repo}
}

// Handle procesa GetAllProductsQuery y devuelve productos filtrados
func (h *GetAllProductsHandler) Handle(ctx context.Context, request interface{}) (interface{}, error) {
	query, ok := request.(*product.GetAllProductsQuery)
	if !ok {
		return nil, fmt.Errorf("invalid request type for GetAllProductsHandler")
	}

	return h.repo.GetAll(query.Category, query.MinPrice, query.MaxPrice)
}