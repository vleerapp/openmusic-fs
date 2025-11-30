package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
)

func info(c *gin.Context) {
	cfg := helpers.GetConfig(c)

	c.JSON(http.StatusOK, gin.H{
		"branding": cfg.Branding,
		"details":  cfg.Details,
	})
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("", info)
	})
}
