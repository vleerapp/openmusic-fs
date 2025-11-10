package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/config"
)

func GetConfig(c *gin.Context) *config.Config {
	return c.MustGet("config").(*config.Config)
}
