package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"meli-products-api/domain"
)

// ProductRepository implements domain.ProductRepository using JSON file
type ProductRepository struct {
	filePath string
	products []*domain.Product
}

// NewProductRepository creates a new JSON-based product repository
func NewProductRepository(filePath string) (*ProductRepository, error) {
	repo := &ProductRepository{
		filePath: filePath,
	}

	// Load products from JSON file on initialization
	if err := repo.loadProducts(); err != nil {
		return nil, fmt.Errorf("failed to load products: %w", err)
	}

	return repo, nil
}

// loadProducts loads products from JSON file into memory
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

// GetByID retrieves a product by its ID
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

// GetAll retrieves all products with optional filtering
func (r *ProductRepository) GetAll(category string, minPrice, maxPrice float64) ([]*domain.Product, error) {
	var filteredProducts []*domain.Product

	for _, product := range r.products {
		// Filter by category if specified
		if category != "" && !strings.EqualFold(product.Category, category) {
			continue
		}

		// Filter by price range if specified
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

// GetByIDs retrieves multiple products by their IDs for comparison
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

	// If some products were not found, return an error with details
	if len(notFoundIDs) > 0 {
		return products, fmt.Errorf("products not found: %v", notFoundIDs)
	}

	return products, nil
}

// Search products by name or description
func (r *ProductRepository) Search(query string) ([]*domain.Product, error) {
	if query == "" {
		return r.GetAll("", 0, 0)
	}

	var matchingProducts []*domain.Product
	queryLower := strings.ToLower(query)

	for _, product := range r.products {
		// Search in name, description, brand, and category
		if strings.Contains(strings.ToLower(product.Name), queryLower) ||
			strings.Contains(strings.ToLower(product.Description), queryLower) ||
			strings.Contains(strings.ToLower(product.Brand), queryLower) ||
			strings.Contains(strings.ToLower(product.Category), queryLower) {
			matchingProducts = append(matchingProducts, product)
		}
	}

	return matchingProducts, nil
}

// GetProductCount returns the total number of products
func (r *ProductRepository) GetProductCount() int {
	return len(r.products)
}

// GetCategories returns all unique categories
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

// GetBrands returns all unique brands
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
