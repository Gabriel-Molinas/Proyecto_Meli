@echo off
echo =====================================
echo     MELI PRODUCTS API - SETUP
echo =====================================
echo.

echo [1/6] Limpiando cache de modulos...
go clean -modcache

echo.
echo [2/6] Descargando dependencias...
go mod tidy

echo.
echo [3/6] Verificando dependencias...
go mod download

echo.
echo [4/6] Verificando sintaxis del codigo...
go fmt ./...
go vet ./...
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: Hay problemas en el codigo!
    pause
    exit /b 1
)

echo.
echo [5/6] Instalando herramienta Swagger...
go install github.com/swaggo/swag/cmd/swag@latest

echo.
echo [6/6] Generando documentacion Swagger...
swag init -g cmd/api/main.go -o docs/ --parseDependency --parseInternal

echo.
echo =====================================
echo        SETUP COMPLETADO!
echo =====================================
echo.
echo Para ejecutar la aplicacion:
echo   run.bat  (o)  go run cmd/api/main.go
echo.
echo URLs importantes:
echo   API: http://localhost:8080/api/v1
echo   Swagger: http://localhost:8080/swagger/index.html
echo.
pause
