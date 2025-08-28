# MELI Products API - Setup Script
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host "     MELI PRODUCTS API - SETUP" -ForegroundColor Yellow
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host

Write-Host "[1/5] Limpiando cache de modulos..." -ForegroundColor Green
go clean -modcache

Write-Host
Write-Host "[2/5] Descargando dependencias..." -ForegroundColor Green
go mod tidy

Write-Host
Write-Host "[3/5] Verificando dependencias..." -ForegroundColor Green
go mod download

Write-Host
Write-Host "[4/5] Instalando herramienta Swagger..." -ForegroundColor Green
go install github.com/swaggo/swag/cmd/swag@latest

Write-Host
Write-Host "[5/5] Generando documentacion Swagger..." -ForegroundColor Green
swag init -g cmd/api/main.go -o docs/

Write-Host
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host "        SETUP COMPLETADO!" -ForegroundColor Yellow
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host

Write-Host "Para ejecutar la aplicacion:" -ForegroundColor White
Write-Host "  go run cmd/api/main.go" -ForegroundColor Magenta
Write-Host

Write-Host "URLs importantes:" -ForegroundColor White
Write-Host "  API: " -NoNewline -ForegroundColor White
Write-Host "http://localhost:8080/api/v1" -ForegroundColor Cyan
Write-Host "  Swagger: " -NoNewline -ForegroundColor White
Write-Host "http://localhost:8080/swagger/index.html" -ForegroundColor Cyan
Write-Host

Read-Host "Presiona Enter para continuar"
