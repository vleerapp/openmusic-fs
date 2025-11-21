package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vleerapp/openmusic-fs/internal/api/middlewares"
	"github.com/vleerapp/openmusic-fs/internal/config"
)

type RouteFunc func(*gin.RouterGroup)

var routeList []RouteFunc

func Register(r RouteFunc) {
	routeList = append(routeList, r)
}

func Start() {
	cfg := config.Load()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middlewares.ConfigMiddleware(cfg))
	r.Use(middlewares.AuthMiddleware())

	api := r.Group("/api")

	for _, route := range routeList {
		route(api)
	}

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
