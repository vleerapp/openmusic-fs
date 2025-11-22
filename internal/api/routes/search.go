package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func search(c *gin.Context) {
	songs := musicfs.SearchSongs(c.Query("query"))

	if songs == nil {
		songs = []musicfs.Item{}
	}

	c.JSON(http.StatusOK, songs)
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/songs/search", search)
	})
}
