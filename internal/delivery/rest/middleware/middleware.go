package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware configures Cross-Origin Resource Sharing headers
func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Allow requests from any origin during development
		// In production, specify allowed origins explicitly
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

// LoggerMiddleware logs HTTP requests with detailed information
func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			// Custom log format with more details
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

// RequestIDMiddleware adds a unique request ID to each request
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

// RecoveryMiddleware recovers from panics and returns a proper error response
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

// SecurityHeadersMiddleware adds security-related HTTP headers
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Next()
	}
}

// generateRequestID generates a simple request ID (in production, use UUID)
func generateRequestID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}
