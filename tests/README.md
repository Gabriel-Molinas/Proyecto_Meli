# Tests del Proyecto API Comparación de Productos

Esta carpeta contiene todos los tests organizados por tipo y propósito.

## Estructura de Tests

```
tests/
├── unit/                    # Tests unitarios
│   ├── domain_test.go      # Tests de entidades de dominio
│   ├── repository_test.go  # Tests del repositorio
│   ├── mediator_test.go    # Tests del patrón Mediator
│   └── response_test.go    # Tests de utilidades HTTP
├── integration/            # Tests de integración
│   └── api_test.go        # Tests de API completa
├── e2e/                   # Tests end-to-end (futuro)
├── fixtures/              # Datos de prueba
│   └── test_products.json # Productos para testing
├── mocks/                 # Mocks reutilizables (futuro)
└── README.md             # Esta documentación
```

## Tipos de Tests

### 1. Tests Unitarios (`unit/`)

Prueban componentes individuales de forma aislada:

- **`domain_test.go`**: Entidades, errores custom, validaciones
- **`repository_test.go`**: Operaciones del repositorio JSON
- **`mediator_test.go`**: Patrón Mediator y handlers
- **`response_test.go`**: Utilidades de respuesta HTTP

### 2. Tests de Integración (`integration/`)

Prueban la interacción entre múltiples componentes:

- **`api_test.go`**: Tests completos de endpoints REST
- Configuración completa de la aplicación
- Uso de datos reales de prueba

### 3. Fixtures (`fixtures/`)

Datos de prueba compartidos:

- **`test_products.json`**: Conjunto de productos para testing
- Datos estructurados para diferentes escenarios
- Productos en múltiples categorías (Smartphones, Laptops, etc.)

## Ejecutar Tests

### Todos los tests en la nueva estructura:
```bash
go test ./tests/... -v
```

### Solo tests unitarios:
```bash
go test ./tests/unit/... -v
```

### Solo tests de integración:
```bash
go test ./tests/integration/... -v
```

### Con cobertura:
```bash
go test ./tests/... -cover -coverprofile=coverage.out
```

### Tests específicos:
```bash
# Test específico
go test ./tests/unit -run TestProductNotFoundError -v

# Test específico de integración
go test ./tests/integration -run TestIntegration_GetProduct -v
```

## Comandos Make

Usar el Makefile especializado para testing:

```bash
# Tests unitarios separados
make -f Makefile.test test-unit

# Tests de integración separados  
make -f Makefile.test test-integration

# Todos los tests con cobertura
make -f Makefile.test test-coverage
```

## Convenciones

### Naming
- Tests unitarios: `TestFunctionName`
- Tests de integración: `TestIntegration_FeatureName`
- Helper functions: `setupTestX`, `createTestY`

### Estructura de Test
```go
func TestSomething(t *testing.T) {
    t.Run("caso exitoso", func(t *testing.T) {
        // Arrange
        // Act  
        // Assert
    })
    
    t.Run("caso de error", func(t *testing.T) {
        // Arrange
        // Act
        // Assert
    })
}
```

### Mocks y Fixtures
- Usar mocks para dependencias externas
- Datos de prueba en `fixtures/`
- Helpers para setup/teardown compartidos

## Migración desde Tests Originales

Los tests originales (junto al código) pueden mantenerse o eliminarse:

```bash
# Ver tests originales
find . -name "*_test.go" -not -path "./tests/*"

# Para eliminar tests originales después de validar
# rm domain/product_test.go
# rm internal/repository/json/product_repository_test.go
# etc.
```

## Beneficios de esta Estructura

1. **Separación clara** entre tipos de test
2. **Datos centralizados** en fixtures
3. **Ejecución selectiva** por tipo
4. **Mejor organización** para proyectos grandes
5. **Tests de integración** más realistas
6. **Mocks reutilizables** (futuro)

## Próximos Pasos

- [ ] Agregar tests E2E en `e2e/`
- [ ] Crear mocks reutilizables en `mocks/`
- [ ] Benchmarks de performance
- [ ] Tests de carga/stress
- [ ] Coverage analysis detallado