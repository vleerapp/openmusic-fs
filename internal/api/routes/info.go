package routes

import (
	"net/http"
	"unicode"

	"github.com/fatih/structs"
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

	links := gin.H{
		"homepage":         cfg.Branding.Links.Homepage,
		"privacyStatement": cfg.Branding.Links.PrivacyStatement,
		"donate":           cfg.Branding.Links.Donate,
	}

	branding := gin.H{
		"name":  cfg.Branding.Name,
		"email": cfg.Branding.Email,
		"short": cfg.Branding.Short,
		"logo":  cfg.Branding.Logo,
		"theme": cfg.Branding.Theme,
		"links": links,
	}

	caps := gin.H{}
	for k, v := range structs.Map(cfg.Details.Capabilities) {
		caps[toSnake(k)] = v
	}

	details := gin.H{
		"version":      cfg.Details.Version,
		"capabilities": caps,
	}

	c.JSON(http.StatusOK, gin.H{
		"branding": branding,
		"details":  details,
	})
}

func init() {
	api.Register(func(g *gin.RouterGroup) {
		g.GET("", info)
	})
}
