/*
Package json implementa el repositorio de productos utilizando archivos JSON como 
fuente de datos.

Este repositorio proporciona una implementación completa de domain.ProductRepository 
que carga y mantiene los productos en memoria para un acceso rápido. Es ideal para 
desarrollo, demos y aplicaciones que no requieren persistencia compleja.

Características:
- Carga de datos desde archivos JSON al inicializar
- Operaciones de búsqueda y filtrado en memoria
- Manejo de errores específicos del dominio
- Soporte para búsqueda por texto en múltiples campos
- Extracción de metadatos (categorías y marcas únicas)
*/
package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"meli-products-api/domain"
)

// ProductRepository implementa domain.ProductRepository utilizando archivos JSON
type ProductRepository struct {
	filePath string
	products []*domain.Product
}

// NewProductRepository crea un nuevo repositorio de productos basado en JSON
func NewProductRepository(filePath string) (*ProductRepository, error) {
	repo := &ProductRepository{
		filePath: filePath,
	}

	// Cargar productos desde archivo JSON durante la inicialización
	if err := repo.loadProducts(); err != nil {
		return nil, fmt.Errorf("failed to load products: %w", err)
	}

	return repo, nil
}

// loadProducts carga los productos desde el archivo JSON a memoria
func (r *ProductRepository) loadProducts() error {
	file, err := os.Open(r.filePath)
	if err != nil {
		return fmt.Errorf("failed to open products file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read products file: %w", err)
	}

	if err := json.Unmarshal(bytes, &r.products); err != nil {
		return fmt.Errorf("failed to parse products JSON: %w", err)
	}

	return nil
}

// GetByID obtiene un producto por su ID
func (r *ProductRepository) GetByID(id string) (*domain.Product, error) {
	if id == "" {
		return nil, &domain.InvalidProductIDError{ID: id}
	}

	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, &domain.ProductNotFoundError{ID: id}
}

// GetAll obtiene todos los productos con filtrado opcional
func (r *ProductRepository) GetAll(category string, minPrice, maxPrice float64) ([]*domain.Product, error) {
	var filteredProducts []*domain.Product

	for _, product := range r.products {
		// Filtrar por categoría si está especificada
		if category != "" && !strings.EqualFold(product.Category, category) {
			continue
		}

		// Filtrar por rango de precio si está especificado
		if minPrice > 0 && product.Price < minPrice {
			continue
		}
		if maxPrice > 0 && product.Price > maxPrice {
			continue
		}

		filteredProducts = append(filteredProducts, product)
	}

	return filteredProducts, nil
}

// GetByIDs obtiene múltiples productos por sus IDs para comparación
func (r *ProductRepository) GetByIDs(ids []string) ([]*domain.Product, error) {
	if len(ids) == 0 {
		return []*domain.Product{}, nil
	}

	var products []*domain.Product
	var notFoundIDs []string

	for _, id := range ids {
		product, err := r.GetByID(id)
		if err != nil {
			if _, ok := err.(*domain.ProductNotFoundError); ok {
				notFoundIDs = append(notFoundIDs, id)
				continue
			}
			return nil, err
		}
		products = append(products, product)
	}

	// Si algunos productos no fueron encontrados, devolver error con detalles
	if len(notFoundIDs) > 0 {
		return products, fmt.Errorf("products not found: %v", notFoundIDs)
	}

	return products, nil
}

// Search busca productos por nombre o descripción
func (r *ProductRepository) Search(query string) ([]*domain.Product, error) {
	if query == "" {
		return r.GetAll("", 0, 0)
	}

	var matchingProducts []*domain.Product
	queryLower := strings.ToLower(query)

	for _, product := range r.products {
		// Buscar en nombre, descripción, marca y categoría
		if strings.Contains(strings.ToLower(product.Name), queryLower) ||
			strings.Contains(strings.ToLower(product.Description), queryLower) ||
			strings.Contains(strings.ToLower(product.Brand), queryLower) ||
			strings.Contains(strings.ToLower(product.Category), queryLower) {
			matchingProducts = append(matchingProducts, product)
		}
	}

	return matchingProducts, nil
}

// GetProductCount devuelve el número total de productos
func (r *ProductRepository) GetProductCount() int {
	return len(r.products)
}

// GetCategories devuelve todas las categorías únicas
func (r *ProductRepository) GetCategories() []string {
	categoryMap := make(map[string]bool)
	var categories []string

	for _, product := range r.products {
		if !categoryMap[product.Category] {
			categoryMap[product.Category] = true
			categories = append(categories, product.Category)
		}
	}

	return categories
}

// GetBrands devuelve todas las marcas únicas
func (r *ProductRepository) GetBrands() []string {
	brandMap := make(map[string]bool)
	var brands []string

	for _, product := range r.products {
		if !brandMap[product.Brand] {
			brandMap[product.Brand] = true
			brands = append(brands, product.Brand)
		}
	}

	return brands
}
