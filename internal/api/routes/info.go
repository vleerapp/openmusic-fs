package routes

import (
	"net/http"
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api"
	"github.com/vleerapp/openmusic-fs/internal/api/helpers"
)

func toCamel(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func info(c *gin.Context) {
	cfg := helpers.GetConfig(c)

	links := gin.H{
		"homepage":          cfg.Branding.Links.Homepage,
		"privacy_statement": cfg.Branding.Links.PrivacyStatement,
		"donate":            cfg.Branding.Links.Donate,
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
		caps[toCamel(k)] = v
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
