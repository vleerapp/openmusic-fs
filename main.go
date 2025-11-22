package main

import (
	"log"

	"github.com/vleerapp/openmusic-fs/internal/api"
	_ "github.com/vleerapp/openmusic-fs/internal/api/routes"
	"github.com/vleerapp/openmusic-fs/internal/musicfs"
)

func main() {
	musicfs.Scan()

	watcher, err := musicfs.StartWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	api.Start()
}
