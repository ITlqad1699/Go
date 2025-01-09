package middleware

import (
	"github.com/anonydev/e-commerce-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Todo: Implement the middleware
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorInvalidToken, "invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
