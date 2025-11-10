package main

import (
	"github.com/vleerapp/openmusic-fs/internal/api"
	_ "github.com/vleerapp/openmusic-fs/internal/api/routes"
)

func main() {
	api.Start()
}
