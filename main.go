package main

import (
	"log"

	"github.com/vleerapp/openmusic-fs/internal/config"
)

func main() {
	conf, err := config.Load()

	if err != nil {
		log.Println("error:", err)
		return
	}

	log.Printf("config: %+v\n", conf.Details.Capabilities)
}
