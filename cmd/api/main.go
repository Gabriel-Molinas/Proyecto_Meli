package main

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"meli-products-api/internal/application/controllers/product"
	"meli-products-api/internal/application/mediator"
	productQueries "meli-products-api/internal/application/queries/product"
	"meli-products-api/internal/delivery/rest/controllers"
	"meli-products-api/internal/delivery/rest/middleware"
	jsonRepo "meli-products-api/internal/repository/json"

	// Import docs for swagger generation
	_ "meli-products-api/docs"
)

// @title           Products Comparison API
// @version         1.0
// @description     API for product comparison with detailed product information
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Initialize repository with JSON data
	dataPath := filepath.Join("data", "products.json")
	repo, err := jsonRepo.NewProductRepository(dataPath)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize mediator
	mediatorInstance := mediator.NewMediator()

	// Register handlers with mediator
	registerHandlers(mediatorInstance, repo)

	// Initialize controller
	productController := controllers.NewProductController(mediatorInstance)

	// Setup Gin router
	router := setupRouter(productController)

	// Start server
	log.Println("Starting Products Comparison API on port 8080...")
	log.Println("Swagger documentation available at: http://localhost:8080/swagger/index.html")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// registerHandlers registers all query handlers with the mediator
func registerHandlers(m mediator.Mediator, repo *jsonRepo.ProductRepository) {
	// Register product handlers
	m.Register(&productQueries.GetProductQuery{}, product.NewGetProductHandler(repo))
	m.Register(&productQueries.GetAllProductsQuery{}, product.NewGetAllProductsHandler(repo))
	m.Register(&productQueries.CompareProductsQuery{}, product.NewCompareProductsHandler(repo))
	m.Register(&productQueries.SearchProductsQuery{}, product.NewSearchProductsHandler(repo))

	// Register metadata handlers
	m.Register(&productQueries.GetCategoriesQuery{}, product.NewGetCategoriesHandler(repo))
	m.Register(&productQueries.GetBrandsQuery{}, product.NewGetBrandsHandler(repo))
}

// setupRouter configures and returns the Gin router with all routes and middleware
func setupRouter(productController *controllers.ProductController) *gin.Engine {
	// Set Gin to release mode for production (comment out for development)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Add middleware
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.SecurityHeadersMiddleware())

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes group
	v1 := router.Group("/api/v1")
	{
		// System routes
		v1.GET("/health", productController.HealthCheck)

		// Product routes
		products := v1.Group("/products")
		{
			products.GET("", productController.GetAllProducts)
			products.GET("/search", productController.SearchProducts)
			products.GET("/compare", productController.CompareProducts)
			products.GET("/:id", productController.GetProduct)
		}

		// Metadata routes
		v1.GET("/categories", productController.GetCategories)
		v1.GET("/brands", productController.GetBrands)
	}

	// Root redirect to swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	return router
}
