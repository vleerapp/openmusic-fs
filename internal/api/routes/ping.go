package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
)

func ping(c *gin.Context) {
	v, _ := c.Get("auth")
	auth, _ := v.(bool)

	c.JSON(http.StatusOK, gin.H{"message": "pong", "auth": auth})
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/ping", ping)
	})
}
