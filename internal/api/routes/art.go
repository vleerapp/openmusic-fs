package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func art(c *gin.Context) {
	data, mimeType, ok, err := musicfs.GetCoverByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.CreateError("Internal Server Error", nil))
		log.Printf("Could not read cover: %v", err)
		return
	}

	if !ok || data == nil {
		c.JSON(http.StatusNotFound, helpers.CreateError("Song cover not found", nil))
	}

	c.Data(http.StatusOK, mimeType, data)
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/song/:id/art", art)
	})
}
