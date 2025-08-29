package unit

import (
	"os"
	"path/filepath"
	"testing"

	"meli-products-api/domain"
	jsonRepo "meli-products-api/internal/repository/json"
)

// createTestFile crea un archivo JSON temporal para las pruebas
func createTestFile(t *testing.T, content string) string {
	t.Helper()
	
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test_products.json")
	
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	return filePath
}

func TestNewProductRepository(t *testing.T) {
	testData := `[
		{
			"id": "TEST001",
			"name": "Test Product",
			"image_url": "https://example.com/test.jpg",
			"description": "Test description",
			"price": 299.99,
			"rating": 4.5,
			"category": "Electronics",
			"brand": "TestBrand",
			"available": true,
			"specifications": []
		}
	]`

	t.Run("Crear repositorio exitosamente", func(t *testing.T) {
		filePath := createTestFile(t, testData)
		
		repo, err := jsonRepo.NewProductRepository(filePath)
		if err != nil {
			t.Errorf("NewProductRepository() error = %v, wantErr nil", err)
			return
		}
		
		if repo == nil {
			t.Error("NewProductRepository() returned nil repository")
		}
	})

	t.Run("Error con archivo inexistente", func(t *testing.T) {
		_, err := jsonRepo.NewProductRepository("nonexistent_file.json")
		if err == nil {
			t.Error("NewProductRepository() expected error for nonexistent file, got nil")
		}
	})

	t.Run("Error con JSON inválido", func(t *testing.T) {
		invalidJSON := `{"invalid": json}`
		filePath := createTestFile(t, invalidJSON)
		
		_, err := jsonRepo.NewProductRepository(filePath)
		if err == nil {
			t.Error("NewProductRepository() expected error for invalid JSON, got nil")
		}
	})
}

func TestRepositoryGetByID(t *testing.T) {
	testData := `[
		{
			"id": "TEST001",
			"name": "Test Product 1",
			"image_url": "https://example.com/test1.jpg",
			"description": "Test description 1",
			"price": 299.99,
			"rating": 4.5,
			"category": "Electronics",
			"brand": "TestBrand",
			"available": true,
			"specifications": []
		},
		{
			"id": "TEST002",
			"name": "Test Product 2",
			"image_url": "https://example.com/test2.jpg",
			"description": "Test description 2",
			"price": 499.99,
			"rating": 4.0,
			"category": "Electronics",
			"brand": "TestBrand",
			"available": false,
			"specifications": []
		}
	]`

	filePath := createTestFile(t, testData)
	repo, err := jsonRepo.NewProductRepository(filePath)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	t.Run("Obtener producto existente", func(t *testing.T) {
		product, err := repo.GetByID("TEST001")
		if err != nil {
			t.Errorf("GetByID() error = %v, wantErr nil", err)
			return
		}
		
		if product.ID != "TEST001" {
			t.Errorf("GetByID() ID = %v, want TEST001", product.ID)
		}
		if product.Name != "Test Product 1" {
			t.Errorf("GetByID() Name = %v, want Test Product 1", product.Name)
		}
	})

	t.Run("Producto no encontrado", func(t *testing.T) {
		_, err := repo.GetByID("NONEXISTENT")
		if err == nil {
			t.Error("GetByID() expected error for nonexistent product, got nil")
		}
		
		if _, ok := err.(*domain.ProductNotFoundError); !ok {
			t.Errorf("GetByID() error type = %T, want *domain.ProductNotFoundError", err)
		}
	})

	t.Run("ID vacío", func(t *testing.T) {
		_, err := repo.GetByID("")
		if err == nil {
			t.Error("GetByID() expected error for empty ID, got nil")
		}
		
		if _, ok := err.(*domain.InvalidProductIDError); !ok {
			t.Errorf("GetByID() error type = %T, want *domain.InvalidProductIDError", err)
		}
	})
}

func TestRepositorySearch(t *testing.T) {
	testData := `[
		{
			"id": "TEST001",
			"name": "Samsung Galaxy",
			"image_url": "https://example.com/test1.jpg",
			"description": "Smartphone with advanced camera",
			"price": 299.99,
			"rating": 4.5,
			"category": "Smartphones",
			"brand": "Samsung",
			"available": true,
			"specifications": []
		},
		{
			"id": "TEST002",
			"name": "iPhone Pro",
			"image_url": "https://example.com/test2.jpg",
			"description": "Apple smartphone with excellent performance",
			"price": 999.99,
			"rating": 4.8,
			"category": "Smartphones",
			"brand": "Apple",
			"available": true,
			"specifications": []
		}
	]`

	filePath := createTestFile(t, testData)
	repo, err := jsonRepo.NewProductRepository(filePath)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	t.Run("Buscar por nombre", func(t *testing.T) {
		products, err := repo.Search("Samsung")
		if err != nil {
			t.Errorf("Search() error = %v, wantErr nil", err)
			return
		}
		
		if len(products) != 1 {
			t.Errorf("Search() count = %v, want 1", len(products))
		}
		
		if products[0].Brand != "Samsung" {
			t.Errorf("Search() result brand = %v, want Samsung", products[0].Brand)
		}
	})

	t.Run("Buscar por descripción", func(t *testing.T) {
		products, err := repo.Search("camera")
		if err != nil {
			t.Errorf("Search() error = %v, wantErr nil", err)
			return
		}
		
		if len(products) != 1 {
			t.Errorf("Search() by description count = %v, want 1", len(products))
		}
	})

	t.Run("Búsqueda sin resultados", func(t *testing.T) {
		products, err := repo.Search("NonExistent")
		if err != nil {
			t.Errorf("Search() error = %v, wantErr nil", err)
			return
		}
		
		if len(products) != 0 {
			t.Errorf("Search() no results count = %v, want 0", len(products))
		}
	})

	t.Run("Búsqueda vacía", func(t *testing.T) {
		products, err := repo.Search("")
		if err != nil {
			t.Errorf("Search() error = %v, wantErr nil", err)
			return
		}
		
		// Búsqueda vacía debe devolver todos los productos
		if len(products) != 2 {
			t.Errorf("Search() empty query count = %v, want 2", len(products))
		}
	})
}