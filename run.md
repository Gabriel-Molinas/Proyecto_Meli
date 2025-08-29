### Requisitos Previos
- **Go 1.21** o superior
- **Swag CLI** (para generación de documentación Swagger)

### Opción 1: Setup Automático (Recomendado)

**Windows (CMD)**:
```cmd
cd C:\repoDisco\Proyecto_Meli
setup.bat
```

**Windows (PowerShell)**:
```powershell
cd C:\repoDisco\Proyecto_Meli
.\setup.ps1
```

### Opción 2: Setup Manual

```bash
# 1. Navegar al directorio del proyecto
cd C:\[su_directorio]\Proyecto_Meli

# 2. Descargar e instalar dependencias
go mod tidy
go mod download

# 3. Instalar Swag CLI para documentación
go install github.com/swaggo/swag/cmd/swag@latest

# 4. Generar documentación Swagger
swag init -g cmd/api/main.go -o docs/ --parseDependency --parseInternal

# 5. Ejecutar la aplicación
go run cmd/api/main.go
```

### Solución de Problemas

Si encuentras errores de dependencias:
```bash
go clean -modcache
go mod tidy
go mod download
```

**URLs de acceso**:
- API: `http://localhost:8080`
- Documentación Swagger: `http://localhost:8080/swagger/index.html`

### Comandos Make Disponibles

El proyecto incluye un Makefile con comandos útiles:

```bash
make build      # Compilar la aplicación
make run        # Ejecutar la aplicación
make swagger    # Generar documentación Swagger
make test       # Ejecutar suite de tests
make clean      # Limpiar archivos generados
make help       # Mostrar ayuda de comandos
```