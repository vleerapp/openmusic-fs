package musicfs

import (
	"log"
	"os"

	"github.com/dhowden/tag"
	"github.com/vleerapp/openmusic-fs/internal/config"
	"github.com/vleerapp/openmusic-fs/internal/openmusic"
)

type Item struct {
	openmusic.Song
	path string
}

var items []Item

func Scan() {
	conf := config.Load()

	files, err := WalkFiles(conf.Server.MusicPath)
	if err != nil {
		log.Fatal("Could not walk files.")
	}

	for _, path := range files {
		file, err := os.Open(path)

		if err != nil {
			log.Printf("failed to open file %s: %v", path, err)
			continue
		}

		defer file.Close()

		if isAudio := IsAudio(path); isAudio == false {
			log.Printf("File is not an audio file %s", path)
			continue
		}

		id, err := FileSHA256(path)

		if err != nil {
			log.Printf("failed to open file %s: %v", path, err)
			continue
		}

		info, err := tag.ReadFrom(file)

		if err != nil {
			log.Printf("failed to open file %s: %v", path, err)
			continue
		}

		track, trackTotal := info.Track()
		disc, discTotal := info.Disc()

		item := Item{
			Song: openmusic.Song{
				ID:          id,
				Title:       info.Title(),
				ArtistName:  info.Artist(),
				AlbumName:   info.Album(),
				Genre:       info.Genre(),
				ReleaseYear: info.Year(),
				Track:       track,
				TrackTotal:  trackTotal,
				Disc:        disc,
				DiscTotal:   discTotal,
				Lyrics:      info.Lyrics(),
				CoverURL:    info.Picture(),
			},
			path: file.Name(),
		}

		items = append(items, item)

		log.Printf("Loaded %+v", item)
	}
}
