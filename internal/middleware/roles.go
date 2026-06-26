package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// /abc -> only admin can access -> 2 level check -> auth ? -> 2nd RequireAdmin
// /xyz -> only admin can access -> 1 level check -> auth ?
// /ccc -> anyone can access -> no check needed

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := GetRole(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		if !strings.EqualFold(role, "admin") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "This route can only be acesssed by admin",
			})

			return
		}

		c.Next()
	}
}
