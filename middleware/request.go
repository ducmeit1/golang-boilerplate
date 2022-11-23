package middleware

import (
	"context"
	"tvf-saleshub-backend/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestCtxMiddleware uses to add X-Request-ID to header of request
func RequestCtxMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID != "" {
			_, err := uuid.Parse(requestID)
			if err != nil {
				requestID = uuid.New().String()
			}
		} else {
			requestID = uuid.New().String()
		}

		c.Header("X-Request-ID", requestID)
		ctx := context.WithValue(c.Request.Context(), util.RequestIDCK, requestID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
