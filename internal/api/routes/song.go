package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func song(c *gin.Context) {
	song, ok := musicfs.GetSongByID(c.Param("id"))

	if !ok {
		c.JSON(http.StatusNotFound, helpers.CreateError("Song not found", nil))
	}

	c.JSON(http.StatusOK, song)
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/song/:id", song)
	})
}
