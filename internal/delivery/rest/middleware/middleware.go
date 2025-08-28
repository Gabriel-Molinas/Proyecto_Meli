/*
Package middleware contiene la implementación de todos los middlewares HTTP 
utilizados en la aplicación REST.

Los middlewares proporcionan funcionalidades transversales que se aplican a 
todas las rutas HTTP, incluyendo configuración de CORS, logging de requests, 
manejo de errores, headers de seguridad y generación de IDs únicos de request.

Middlewares implementados:
- CORS: Configuración de Cross-Origin Resource Sharing
- Logger: Registro detallado de requests HTTP
- RequestID: Generación de IDs únicos para trazabilidad
- Recovery: Manejo y recuperación de panics
- SecurityHeaders: Headers de seguridad estándar
*/
package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware configura los headers de Cross-Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Permite requests desde cualquier origen durante desarrollo
		// En producción, especificar orígenes permitidos explícitamente
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		// Manejar requests preflight
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

// LoggerMiddleware registra requests HTTP con información detallada
func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			// Formato de log personalizado con más detalles
			return fmt.Sprintf("[%s] %s %s %d %s \"%s\" %s \"%s\" %s\n",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ClientIP,
				param.ErrorMessage,
				param.Keys,
			)
		},
		Output: gin.DefaultWriter,
	})
}

// RequestIDMiddleware agrega un ID único de request a cada solicitud
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)
		c.Next()
	}
}

// RecoveryMiddleware se recupera de panics y devuelve una respuesta de error apropiada
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Printf("Panic recovered: %s", err)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "An unexpected error occurred",
				"details": "Please try again later or contact support",
			},
		})
	})
}

// SecurityHeadersMiddleware agrega headers HTTP relacionados con seguridad
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Next()
	}
}

// generateRequestID genera un ID de request simple (en producción, usar UUID)
func generateRequestID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}
