# Products Comparison API

## 📋 Descripción

API REST diseñada para la comparación de artículos/productos, implementada siguiendo los principios de Clean Architecture con patrón Mediator y CQRS. Proporciona endpoints eficientes para recuperar información detallada de productos para funcionalidades de comparación.

## 🏗️ Arquitectura

El proyecto implementa Clean Architecture con las siguientes capas:

```
meli-products-api/
├── cmd/api/                    # Punto de entrada de la aplicación
├── domain/                     # Entidades y reglas de negocio
├── internal/
│   ├── delivery/rest/          # Capa de presentación (HTTP)
│   │   ├── controllers/        # Controladores HTTP
│   │   └── middleware/         # Middleware HTTP
│   ├── application/            # Capa de aplicación
│   │   ├── handlers/           # Handlers de comandos/queries
│   │   ├── queries/            # Definiciones de queries (CQRS)
│   │   └── mediator/           # Patrón Mediator
│   └── repository/             # Capa de acceso a datos
│       └── json/               # Implementación JSON
├── pkg/response/               # Paquetes compartidos
├── data/                       # Datos de ejemplo (JSON)
└── docs/                       # Documentación Swagger
```

## 🚀 Características

- ✅ **Clean Architecture** - Separación clara de responsabilidades
- ✅ **Patrón Mediator** - Desacoplamiento entre controladores y lógica de negocio
- ✅ **CQRS** - Separación de comandos y queries
- ✅ **Swagger/OpenAPI** - Documentación automática de API
- ✅ **Manejo de errores** - Respuestas consistentes y detalladas
- ✅ **Middleware** - CORS, logging, recovery, security headers
- ✅ **Datos JSON** - Sin dependencias de bases de datos externas
- ✅ **Validación** - Validación de entrada robusta
- ✅ **Health Check** - Endpoint de verificación de salud

## 📊 Endpoints Disponibles

### Productos
- `GET /api/v1/products` - Obtener todos los productos (con filtros opcionales)
- `GET /api/v1/products/{id}` - Obtener producto por ID
- `GET /api/v1/products/search?q={query}` - Buscar productos
- `GET /api/v1/products/compare?ids={id1,id2,id3}` - Comparar productos

### Metadatos
- `GET /api/v1/categories` - Obtener todas las categorías
- `GET /api/v1/brands` - Obtener todas las marcas

### Sistema
- `GET /api/v1/health` - Health check
- `GET /swagger/index.html` - Documentación Swagger

## 🛠️ Instalación y Ejecución

### Requisitos
- Go 1.21 o superior
- Swag CLI (para generar documentación Swagger)

### Opción 1: Setup Automático (Recomendado) 🚀

**Windows (CMD):**
```cmd
cd C:\repoDisco\Proyecto_Meli
setup.bat
```

**Windows (PowerShell):**
```powershell
cd C:\repoDisco\Proyecto_Meli
.\setup.ps1
```

### Opción 2: Setup Manual 🔧

```bash
cd C:\repoDisco\Proyecto_Meli

# 1. Descargar dependencias
go mod tidy
go mod download

# 2. Instalar Swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# 3. Generar documentación Swagger
swag init -g cmd/api/main.go -o docs/ --parseDependency --parseInternal

# 4. Ejecutar la aplicación
go run cmd/api/main.go
```

### Solución de Problemas 🔧

Si obtienes errores como "missing go.sum entry":
```bash
go clean -modcache
go mod tidy
go mod download
```

La API estará disponible en: `http://localhost:8080`
Documentación Swagger: `http://localhost:8080/swagger/index.html`

## 🔧 Makefile

Puedes usar estos comandos para facilitar el desarrollo:

```bash
# Construir la aplicación
make build

# Ejecutar la aplicación
make run

# Generar documentación Swagger
make swagger

# Ejecutar tests
make test

# Limpiar archivos generados
make clean

# Ver todos los comandos disponibles
make help
```

## 📝 Ejemplos de Uso

### 1. Obtener un producto específico
```bash
curl -X GET "http://localhost:8080/api/v1/products/PHONE001"
```

### 2. Buscar productos
```bash
curl -X GET "http://localhost:8080/api/v1/products/search?q=Samsung"
```

### 3. Comparar productos
```bash
curl -X GET "http://localhost:8080/api/v1/products/compare?ids=PHONE001,PHONE002,PHONE003"
```

### 4. Filtrar productos por categoría y precio
```bash
curl -X GET "http://localhost:8080/api/v1/products?category=Smartphones&min_price=1000&max_price=1500"
```

## 🏷️ Estructura de Respuesta

Todas las respuestas siguen este formato estándar:

```json
{
    "success": true,
    "message": "Request completed successfully",
    "data": { ... },
    "error": null,
    "meta": {
        "timestamp": "2024-01-15T10:30:00Z",
        "request_id": "req-12345-abcde",
        "version": "v1"
    }
}
```

## 📊 Modelo de Datos

### Producto
```json
{
    "id": "PHONE001",
    "name": "Samsung Galaxy S24 Ultra",
    "image_url": "https://images.samsung.com/galaxy-s24-ultra.jpg",
    "description": "El smartphone más avanzado de Samsung...",
    "price": 1299.99,
    "rating": 4.6,
    "category": "Smartphones",
    "brand": "Samsung",
    "available": true,
    "specifications": [
        {
            "name": "Pantalla",
            "value": "6.8",
            "unit": "pulgadas"
        }
    ]
}
```

## 🧪 Testing

Para ejecutar los tests:
```bash
make test
# o
go test ./...
```

## 📚 Documentación

La documentación completa de la API está disponible en Swagger UI cuando la aplicación está ejecutándose:
`http://localhost:8080/swagger/index.html`

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia Apache 2.0 - ver el archivo [LICENSE](LICENSE) para detalles.

## 📞 Soporte

Para soporte y preguntas:
- Email: support@example.com
- Issues: [GitHub Issues](https://github.com/your-repo/issues)

---

**Desarrollado con ❤️ usando Go y Clean Architecture**
