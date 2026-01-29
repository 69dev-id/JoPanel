package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RequireSecret() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Agent-Secret")
		expected := os.Getenv("AGENT_SECRET")
		if expected == "" {
			expected = "fallback_secret_change_me" // Default for dev
		}

		if token != expected {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized Agent Access"})
			c.Abort()
			return
		}
		c.Next()
	}
}
