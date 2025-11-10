package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/config"
)

func ConfigMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	}
}
