package routes

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
)

func toSnake(s string) string {
	if s == "" {
		return s
	}

	var out []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				out = append(out, '_')
			}
			out = append(out, unicode.ToLower(r))
		} else {
			out = append(out, r)
		}
	}

	return string(out)
}

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
