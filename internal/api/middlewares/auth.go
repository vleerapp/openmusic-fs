package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := helpers.GetConfig(c)

		if cfg.Server.Secret == nil {
			c.Next()
			return
		}

		c.Set("auth", true)

		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.Set("auth", false)
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		if token != *cfg.Server.Secret {
			c.Set("auth", false)
		}

		c.Next()
	}
}
