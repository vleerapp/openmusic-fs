package routes

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func stream(c *gin.Context) {
	path, ok := musicfs.GetPathByID(c.Param("id"))

	if !ok {
		c.JSON(http.StatusNotFound, helpers.CreateError("Song file not found", nil))
		return
	}

	f, err := os.Open(path)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("Could not read file: %v", err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Printf("Could not get file stats: %v", err)
		return
	}

	ext := filepath.Ext(path)
	mimeType := "audio/mpeg"
	switch ext {
	case ".flac":
		mimeType = "audio/flac"
	case ".ogg", ".oga":
		mimeType = "audio/ogg"
	case ".wav":
		mimeType = "audio/wav"
	}

	c.Header("Content-Type", mimeType)
	http.ServeContent(c.Writer, c.Request, fi.Name(), fi.ModTime(), f)
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("/song/:id/stream", stream)
		g.HEAD("/song/:id/stream", stream)
	})
}
