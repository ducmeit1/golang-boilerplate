package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// CorsMiddleware is a middleware for CORS requests.
func CorsMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type", "Accept", "Origin", "X-Request-With", "X-Request-ID", "Authorization"},
			ExposedHeaders: []string{"Content-Type", "Content-Length", "X-Request-ID"},
			MaxAge:         0,
			Debug:          false,
		}).HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}
