package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func random(c *gin.Context) {
	song, ok := musicfs.GetRandomSong()

	if !ok || song == nil {
		c.JSON(http.StatusNotFound, helpers.CreateError("No songs found", nil))
	}

	c.JSON(http.StatusOK, song)
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/songs/random", random)
	})
}
