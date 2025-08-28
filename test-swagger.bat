@echo off
echo =====================================
echo   PRUEBA DE GENERACION SWAGGER
echo =====================================
echo.

echo [1/4] Verificando sintaxis Go...
go fmt ./...

echo.
echo [2/4] Ejecutando go vet...
go vet ./...
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: go vet encontro problemas!
    pause
    exit /b 1
)

echo.
echo [3/4] Verificando que swag esta instalado...
swag --version >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo Swag no esta instalado. Instalando...
    go install github.com/swaggo/swag/cmd/swag@latest
    if %ERRORLEVEL% NEQ 0 (
        echo ERROR: No se pudo instalar swag!
        pause
        exit /b 1
    )
)

echo.
echo [4/4] Generando documentacion Swagger...
swag init -g cmd/api/main.go -o docs/ --parseDependency --parseInternal
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: No se pudo generar la documentacion Swagger!
    echo Verifica que los comentarios de Swagger esten correctos.
    pause
    exit /b 1
)

echo.
echo =====================================
echo    SWAGGER GENERADO EXITOSAMENTE!
echo =====================================
echo.
echo Archivos generados:
echo   docs/docs.go
echo   docs/swagger.json
echo   docs/swagger.yaml
echo.
echo Ahora puedes ejecutar: run.bat
echo.
pause
