package main

import (
	"os"

	"github.com/vleerapp/openmusic-fs/internal/api"
	_ "github.com/vleerapp/openmusic-fs/internal/api/routes"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func main() {
	musicfs.Scan()
	os.Exit(0)
	api.Start()
}
