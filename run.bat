@echo off
echo =====================================
echo     MELI PRODUCTS API - RUN
echo =====================================
echo.

echo Iniciando servidor en puerto 8080...
echo.
echo URLs importantes:
echo   API Base: http://localhost:8080/api/v1
echo   Swagger:  http://localhost:8080/swagger/index.html
echo   Health:   http://localhost:8080/api/v1/health
echo.
echo Presiona Ctrl+C para detener el servidor
echo =====================================
echo.

swag init -g cmd/api/main.go -o .\docs\ --parseDependency --parseInternal
go run cmd/api/main.go
