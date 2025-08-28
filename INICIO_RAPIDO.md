# MELI Products API - Instrucciones Rápidas

## 🚀 INICIO RÁPIDO

### 1️⃣ Primera vez (configurar todo):
```
setup.bat
```

### 2️⃣ Ejecutar aplicación:
```
run.bat
```

### 3️⃣ Verificar código (opcional):
```
verify.bat
```

## 🌐 URLs importantes:
- **API**: http://localhost:8080/api/v1
- **Swagger**: http://localhost:8080/swagger/index.html  
- **Health**: http://localhost:8080/api/v1/health

## 📊 Endpoints principales:
- `GET /api/v1/products` - Todos los productos
- `GET /api/v1/products/{id}` - Producto específico
- `GET /api/v1/products/compare?ids=PHONE001,PHONE002` - Comparar
- `GET /api/v1/products/search?q=Samsung` - Buscar

## 🧪 Pruebas rápidas:
```bash
curl "http://localhost:8080/api/v1/products/PHONE001"
curl "http://localhost:8080/api/v1/products/compare?ids=PHONE001,PHONE002"
curl "http://localhost:8080/api/v1/products/search?q=Samsung"
```

## 🛠️ Solución de problemas:
Si hay errores de compilación, ejecuta:
```
verify.bat
```

## 📂 Scripts disponibles:
- **setup.bat** - Configuración completa inicial
- **run.bat** - Ejecutar aplicación 
- **verify.bat** - Verificar código
- **setup.ps1** - Setup para PowerShell

---
✅ **¡Todo listo! Ejecuta setup.bat para comenzar**
