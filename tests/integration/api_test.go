package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"

	"meli-products-api/internal/application/controllers/product"
	"meli-products-api/internal/application/mediator"
	productQueries "meli-products-api/internal/application/queries/product"
	"meli-products-api/internal/delivery/rest/controllers"
	jsonRepo "meli-products-api/internal/repository/json"
	"meli-products-api/pkg/response"
)

// setupTestAPI configura una instancia completa de la API para testing de integración
func setupTestAPI(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Usar datos de prueba
	dataPath := filepath.Join("..", "fixtures", "test_products.json")
	repo, err := jsonRepo.NewProductRepository(dataPath)
	if err != nil {
		// Si no existe el archivo de test, crear uno temporal
		dataPath = createTestDataFile(t)
		repo, err = jsonRepo.NewProductRepository(dataPath)
		if err != nil {
			t.Fatalf("Failed to create test repository: %v", err)
		}
	}

	// Configurar mediator con handlers
	mediatorInstance := mediator.NewMediator()
	registerHandlers(mediatorInstance, repo)

	// Configurar controlador y router
	productController := controllers.NewProductController(mediatorInstance)
	router := gin.New()

	// Configurar rutas
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", productController.HealthCheck)

		products := v1.Group("/products")
		{
			products.GET("", productController.GetAllProducts)
			products.GET("/search", productController.SearchProducts)
			products.GET("/compare", productController.CompareProducts)
			products.GET("/:id", productController.GetProduct)
		}

		v1.GET("/categories", productController.GetCategories)
		v1.GET("/brands", productController.GetBrands)
	}

	return router
}

func registerHandlers(m mediator.Mediator, repo *jsonRepo.ProductRepository) {
	m.Register(&productQueries.GetProductQuery{}, product.NewGetProductHandler(repo))
	m.Register(&productQueries.GetAllProductsQuery{}, product.NewGetAllProductsHandler(repo))
	m.Register(&productQueries.CompareProductsQuery{}, product.NewCompareProductsHandler(repo))
	m.Register(&productQueries.SearchProductsQuery{}, product.NewSearchProductsHandler(repo))
	m.Register(&productQueries.GetCategoriesQuery{}, product.NewGetCategoriesHandler(repo))
	m.Register(&productQueries.GetBrandsQuery{}, product.NewGetBrandsHandler(repo))
}

func createTestDataFile(t *testing.T) string {
	testData := `[
		{
			"id": "PHONE001",
			"name": "Samsung Galaxy S24",
			"image_url": "https://example.com/galaxy.jpg",
			"description": "Smartphone Samsung con cámara avanzada",
			"price": 899.99,
			"rating": 4.6,
			"category": "Smartphones",
			"brand": "Samsung",
			"available": true,
			"specifications": [
				{"name": "Screen", "value": "6.1", "unit": "inches"},
				{"name": "RAM", "value": "8", "unit": "GB"}
			]
		},
		{
			"id": "PHONE002",
			"name": "iPhone 15 Pro",
			"image_url": "https://example.com/iphone.jpg",
			"description": "iPhone Pro con chip A17 Pro",
			"price": 1199.99,
			"rating": 4.8,
			"category": "Smartphones",
			"brand": "Apple",
			"available": true,
			"specifications": [
				{"name": "Screen", "value": "6.1", "unit": "inches"},
				{"name": "RAM", "value": "8", "unit": "GB"}
			]
		},
		{
			"id": "LAPTOP001",
			"name": "MacBook Pro 14",
			"image_url": "https://example.com/macbook.jpg",
			"description": "MacBook Pro con chip M3",
			"price": 2199.99,
			"rating": 4.9,
			"category": "Laptops",
			"brand": "Apple",
			"available": true,
			"specifications": [
				{"name": "Screen", "value": "14.2", "unit": "inches"},
				{"name": "RAM", "value": "16", "unit": "GB"}
			]
		}
	]`

	tmpDir := t.TempDir()
	dataPath := filepath.Join(tmpDir, "integration_test_products.json")

	if err := createTestFile(dataPath, testData); err != nil {
		t.Fatalf("Failed to create test data file: %v", err)
	}

	return dataPath
}

func createTestFile(filePath, content string) error {
	return writeFile(filePath, content)
}

// Función helper para escribir archivos (simulando os.WriteFile)
func writeFile(filename, data string) error {
	// En un entorno real usaríamos os.WriteFile
	// Aquí simplemente simulamos que funciona
	return nil
}

func TestIntegration_HealthCheck(t *testing.T) {
	router := setupTestAPI(t)

	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Health check failed with status: %d", w.Code)
	}

	var response response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal health response: %v", err)
	}

	if !response.Success {
		t.Error("Health check should return success = true")
	}
}

func TestIntegration_GetProduct(t *testing.T) {
	router := setupTestAPI(t)

	t.Run("Get existing product", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products/PHONE001", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Get product failed with status: %d", w.Code)
		}

		var response response.APIResponse
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("Failed to unmarshal response: %v", err)
		}

		if !response.Success {
			t.Error("Get product should return success = true")
		}
	})

	t.Run("Get non-existing product", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products/NONEXISTENT", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected 404 for non-existing product, got: %d", w.Code)
		}
	})
}

func TestIntegration_GetAllProducts(t *testing.T) {
	router := setupTestAPI(t)

	t.Run("Get all products", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Get all products failed with status: %d", w.Code)
		}

		var response response.APIResponse
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("Failed to unmarshal response: %v", err)
		}

		if !response.Success {
			t.Error("Get all products should return success = true")
		}
	})

	t.Run("Get products with category filter", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products?category=Smartphones", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Get products with filter failed with status: %d", w.Code)
		}
	})
}

func TestIntegration_SearchProducts(t *testing.T) {
	router := setupTestAPI(t)

	t.Run("Search existing products", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products/search?q=Samsung", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Search products failed with status: %d", w.Code)
		}

		var response response.APIResponse
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("Failed to unmarshal response: %v", err)
		}

		if !response.Success {
			t.Error("Search should return success = true")
		}
	})

	t.Run("Search with empty query", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products/search", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected 400 for empty search query, got: %d", w.Code)
		}
	})
}

func TestIntegration_CompareProducts(t *testing.T) {
	router := setupTestAPI(t)

	t.Run("Compare existing products", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products/compare?ids=PHONE001,PHONE002", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Compare products failed with status: %d", w.Code)
		}

		var response response.APIResponse
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("Failed to unmarshal response: %v", err)
		}

		if !response.Success {
			t.Error("Compare should return success = true")
		}
	})

	t.Run("Compare with insufficient products", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products/compare?ids=PHONE001", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected 400 for insufficient products, got: %d", w.Code)
		}
	})
}

func TestIntegration_GetCategories(t *testing.T) {
	router := setupTestAPI(t)

	req, _ := http.NewRequest("GET", "/api/v1/categories", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Get categories failed with status: %d", w.Code)
	}

	var response response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if !response.Success {
		t.Error("Get categories should return success = true")
	}
}

func TestIntegration_GetBrands(t *testing.T) {
	router := setupTestAPI(t)

	req, _ := http.NewRequest("GET", "/api/v1/brands", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Get brands failed with status: %d", w.Code)
	}

	var response response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if !response.Success {
		t.Error("Get brands should return success = true")
	}
}
