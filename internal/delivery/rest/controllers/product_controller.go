package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	// "meli-products-api/domain"
	"meli-products-api/internal/application/mediator"
	"meli-products-api/internal/application/queries/product"
	"meli-products-api/pkg/response"
)

// ProductController handles HTTP requests for product operations
type ProductController struct {
	mediator mediator.Mediator
}

// NewProductController creates a new ProductController
func NewProductController(mediator mediator.Mediator) *ProductController {
	return &ProductController{
		mediator: mediator,
	}
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Retrieve detailed information about a specific product by its unique identifier
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" example("PHONE001")
// @Success 200 {object} response.APIResponse{data=domain.Product} "Product retrieved successfully"
// @Failure 400 {object} response.APIResponse "Invalid product ID"
// @Failure 404 {object} response.APIResponse "Product not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /products/{id} [get]
func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.BadRequest(c.Writer, "INVALID_PRODUCT_ID", "Product ID is required", "Please provide a valid product ID in the URL path")
		return
	}

	query := &product.GetProductQuery{ID: id}
	result, err := pc.mediator.Send(c.Request.Context(), query)

	if err != nil {
		response.HandleError(c.Writer, err)
		return
	}

	response.Success(c.Writer, result, "Product retrieved successfully")
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products with optional filtering by category and price range
// @Tags products
// @Accept json
// @Produce json
// @Param category query string false "Filter by category" example("Smartphones")
// @Param min_price query number false "Minimum price filter" example(100.00)
// @Param max_price query number false "Maximum price filter" example(2000.00)
// @Success 200 {object} response.APIResponse{data=[]domain.Product} "Products retrieved successfully"
// @Failure 400 {object} response.APIResponse "Invalid query parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /products [get]
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	// Parse query parameters
	category := c.Query("category")
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")

	var minPrice, maxPrice float64
	var err error

	// Parse price filters if provided
	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil || minPrice < 0 {
			response.BadRequest(c.Writer, "INVALID_MIN_PRICE", "Invalid minimum price", "Minimum price must be a valid positive number")
			return
		}
	}

	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil || maxPrice < 0 {
			response.BadRequest(c.Writer, "INVALID_MAX_PRICE", "Invalid maximum price", "Maximum price must be a valid positive number")
			return
		}
	}

	// Validate price range
	if minPrice > 0 && maxPrice > 0 && minPrice > maxPrice {
		response.BadRequest(c.Writer, "INVALID_PRICE_RANGE", "Invalid price range", "Minimum price cannot be greater than maximum price")
		return
	}

	query := &product.GetAllProductsQuery{
		Category: category,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}

	result, err := pc.mediator.Send(c.Request.Context(), query)
	if err != nil {
		response.HandleError(c.Writer, err)
		return
	}

	response.Success(c.Writer, result, "Products retrieved successfully")
}

// CompareProducts godoc
// @Summary Compare multiple products
// @Description Retrieve and compare multiple products by their IDs for feature comparison
// @Tags products
// @Accept json
// @Produce json
// @Param ids query string true "Comma-separated product IDs" example("PHONE001,PHONE002,PHONE003")
// @Success 200 {object} response.APIResponse "Products comparison retrieved successfully"
// @Failure 400 {object} response.APIResponse "Invalid product IDs or insufficient products for comparison"
// @Failure 404 {object} response.APIResponse "One or more products not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /products/compare [get]
func (pc *ProductController) CompareProducts(c *gin.Context) {
	idsParam := c.Query("ids")

	if idsParam == "" {
		response.BadRequest(c.Writer, "MISSING_PRODUCT_IDS", "Product IDs are required", "Please provide comma-separated product IDs in the 'ids' query parameter")
		return
	}

	// Parse and validate product IDs
	ids := strings.Split(idsParam, ",")
	var cleanIDs []string

	for _, id := range ids {
		trimmedID := strings.TrimSpace(id)
		if trimmedID != "" {
			cleanIDs = append(cleanIDs, trimmedID)
		}
	}

	if len(cleanIDs) < 2 {
		response.BadRequest(c.Writer, "INSUFFICIENT_PRODUCTS", "At least 2 products required for comparison", "Please provide at least 2 valid product IDs separated by commas")
		return
	}

	if len(cleanIDs) > 10 {
		response.BadRequest(c.Writer, "TOO_MANY_PRODUCTS", "Too many products for comparison", "Please provide at most 10 products for comparison")
		return
	}

	query := &product.CompareProductsQuery{ProductIDs: cleanIDs}
	result, err := pc.mediator.Send(c.Request.Context(), query)

	if err != nil {
		response.HandleError(c.Writer, err)
		return
	}

	response.Success(c.Writer, result, "Products comparison retrieved successfully")
}

// SearchProducts godoc
// @Summary Search products
// @Description Search for products by name, description, brand, or category
// @Tags products
// @Accept json
// @Produce json
// @Param q query string true "Search query" example("Samsung Galaxy")
// @Success 200 {object} response.APIResponse "Products search completed successfully"
// @Failure 400 {object} response.APIResponse "Invalid or missing search query"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /products/search [get]
func (pc *ProductController) SearchProducts(c *gin.Context) {
	searchQuery := strings.TrimSpace(c.Query("q"))

	if searchQuery == "" {
		response.BadRequest(c.Writer, "MISSING_SEARCH_QUERY", "Search query is required", "Please provide a search query in the 'q' parameter")
		return
	}

	if len(searchQuery) < 2 {
		response.BadRequest(c.Writer, "INVALID_SEARCH_QUERY", "Search query too short", "Search query must be at least 2 characters long")
		return
	}

	query := &product.SearchProductsQuery{Query: searchQuery}
	result, err := pc.mediator.Send(c.Request.Context(), query)

	if err != nil {
		response.HandleError(c.Writer, err)
		return
	}

	response.Success(c.Writer, result, "Products search completed successfully")
}

// GetCategories godoc
// @Summary Get all available categories
// @Description Retrieve a list of all available product categories
// @Tags metadata
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse "Categories retrieved successfully"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /categories [get]
func (pc *ProductController) GetCategories(c *gin.Context) {
	query := &product.GetCategoriesQuery{}
	result, err := pc.mediator.Send(c.Request.Context(), query)

	if err != nil {
		response.HandleError(c.Writer, err)
		return
	}

	response.Success(c.Writer, result, "Categories retrieved successfully")
}

// GetBrands godoc
// @Summary Get all available brands
// @Description Retrieve a list of all available product brands
// @Tags metadata
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse "Brands retrieved successfully"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /brands [get]
func (pc *ProductController) GetBrands(c *gin.Context) {
	query := &product.GetBrandsQuery{}
	result, err := pc.mediator.Send(c.Request.Context(), query)

	if err != nil {
		response.HandleError(c.Writer, err)
		return
	}

	response.Success(c.Writer, result, "Brands retrieved successfully")
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API is running and healthy
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse "API is healthy"
// @Router /health [get]
func (pc *ProductController) HealthCheck(c *gin.Context) {
	healthData := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format("2006-01-02T15:04:05Z"),
		"service":   "meli-products-api",
		"version":   "1.0.0",
	}

	response.Success(c.Writer, healthData, "API is healthy")
}
