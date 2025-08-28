@echo off
echo =====================================
echo   VERIFICACION DE COMPILACION
echo =====================================
echo.

echo [1/3] Verificando sintaxis Go...
go fmt ./...

echo.
echo [2/3] Ejecutando go vet...
go vet ./...
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: go vet encontro problemas!
    pause
    exit /b 1
)

echo.
echo [3/3] Compilando proyecto...
go build -o bin/test.exe cmd/api/main.go
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: El proyecto no compila!
    pause
    exit /b 1
) else (
    echo EXITO: El proyecto compila correctamente!
    del bin\test.exe 2>nul
)

echo.
echo =====================================
echo      VERIFICACION COMPLETADA
echo =====================================
echo.
echo El proyecto esta listo para ejecutar!
echo Usa: setup.bat (si es la primera vez)
echo  o : run.bat (para ejecutar)
echo.
pause
