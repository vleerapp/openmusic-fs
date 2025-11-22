package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func list(c *gin.Context) {
	songs := musicfs.ListSongs()

	c.JSON(http.StatusOK, songs)
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/songs/list", list)
	})
}
