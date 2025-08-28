/*
Aplicación principal de la API de Comparación de Productos.

Esta aplicación implementa una API REST para comparación de productos utilizando 
los patrones Clean Architecture, Mediator y CQRS. Proporciona endpoints para 
consultar, buscar y comparar productos de manera eficiente.

Arquitectura:
- Clean Architecture con capas bien definidas
- Patrón Mediator para desacoplar controladores de handlers
- CQRS para separar operaciones de lectura
- Repository pattern para abstracción de datos
- Middleware completo para logging, CORS, seguridad

La aplicación utiliza Gin como framework web y Swagger para documentación automática.
*/
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
	// Inicializar repositorio con datos JSON
	dataPath := filepath.Join("data", "products.json")
	repo, err := jsonRepo.NewProductRepository(dataPath)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Inicializar mediator
	mediatorInstance := mediator.NewMediator()

	// Registrar handlers con el mediator
	registerHandlers(mediatorInstance, repo)

	// Inicializar controlador
	productController := controllers.NewProductController(mediatorInstance)

	// Configurar router de Gin
	router := setupRouter(productController)

	// Iniciar servidor
	log.Println("Starting Products Comparison API on port 8080...")
	log.Println("Swagger documentation available at: http://localhost:8080/swagger/index.html")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// registerHandlers registra todos los handlers de queries con el mediator
func registerHandlers(m mediator.Mediator, repo *jsonRepo.ProductRepository) {
	// Registrar handlers de productos
	m.Register(&productQueries.GetProductQuery{}, product.NewGetProductHandler(repo))
	m.Register(&productQueries.GetAllProductsQuery{}, product.NewGetAllProductsHandler(repo))
	m.Register(&productQueries.CompareProductsQuery{}, product.NewCompareProductsHandler(repo))
	m.Register(&productQueries.SearchProductsQuery{}, product.NewSearchProductsHandler(repo))

	// Registrar handlers de metadatos
	m.Register(&productQueries.GetCategoriesQuery{}, product.NewGetCategoriesHandler(repo))
	m.Register(&productQueries.GetBrandsQuery{}, product.NewGetBrandsHandler(repo))
}

// setupRouter configura y devuelve el router de Gin con todas las rutas y middleware
func setupRouter(productController *controllers.ProductController) *gin.Engine {
	// Establecer Gin en modo release para producción (comentar para desarrollo)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Agregar middleware
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.SecurityHeadersMiddleware())

	// Ruta de documentación Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grupo de rutas de API
	v1 := router.Group("/api/v1")
	{
		// Rutas del sistema
		v1.GET("/health", productController.HealthCheck)

		// Rutas de productos
		products := v1.Group("/products")
		{
			products.GET("", productController.GetAllProducts)
			products.GET("/search", productController.SearchProducts)
			products.GET("/compare", productController.CompareProducts)
			products.GET("/:id", productController.GetProduct)
		}

		// Rutas de metadatos
		v1.GET("/categories", productController.GetCategories)
		v1.GET("/brands", productController.GetBrands)
	}

	// Redirección de raíz a swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	return router
}
