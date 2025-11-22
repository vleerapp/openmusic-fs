package musicfs

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/dhowden/tag"
	"github.com/fsnotify/fsnotify"
	"github.com/vleerapp/openmusic-fs/internal/config"
	"github.com/vleerapp/openmusic-fs/internal/openmusic"
)

type Item struct {
	openmusic.Song
	path string
}

var (
	Items   []Item
	itemsMu sync.RWMutex
)

func ListSongs() []Item {
	itemsMu.RLock()
	defer itemsMu.RUnlock()

	out := make([]Item, len(Items))
	copy(out, Items)
	return out
}

func GetSongByID(id string) (*openmusic.Song, bool) {
	itemsMu.RLock()
	defer itemsMu.RUnlock()

	for i := range Items {
		if Items[i].Song.ID == id {
			song := Items[i].Song
			return &song, true
		}
	}

	return nil, false
}

func SearchSongs(query string) []Item {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil
	}
	q = strings.ToLower(q)

	itemsMu.RLock()
	defer itemsMu.RUnlock()

	var out []Item
	for _, it := range Items {
		title := strings.ToLower(it.Title)
		artist := strings.ToLower(it.ArtistName)
		album := strings.ToLower(it.AlbumName)

		if strings.Contains(title, q) ||
			strings.Contains(artist, q) ||
			strings.Contains(album, q) {

			out = append(out, it)
		}
	}

	return out
}

func indexFile(path string) {
	if !IsAudio(path) {
		log.Printf("Not audio (skip): %s", path)
		return
	}

	id, err := FileSHA256(path)
	if err != nil {
		log.Printf("failed to hash file %s: %v", path, err)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %s: %v", path, err)
		return
	}
	info, err := tag.ReadFrom(file)
	file.Close()
	if err != nil {
		log.Printf("failed to read tags for %s: %v", path, err)
		return
	}

	track, trackTotal := info.Track()
	disc, discTotal := info.Disc()

	title := info.Title()
	if title == "" {
		title = FallbackTitleFromPath(path)
	}
	artist := info.Artist()
	if artist == "" {
		artist = "Unknown Artist"
	}
	album := info.Album()
	if album == "" {
		album = "Unknown Album"
	}

	var coverURL string
	if info.Picture() != nil {
		coverURL = "/art/" + id
	}

	item := Item{
		Song: openmusic.Song{
			ID:          id,
			Title:       title,
			ArtistName:  artist,
			AlbumName:   album,
			Genre:       info.Genre(),
			ReleaseYear: info.Year(),
			Track:       track,
			TrackTotal:  trackTotal,
			Disc:        disc,
			DiscTotal:   discTotal,
			Lyrics:      info.Lyrics(),
			CoverURL:    coverURL,
		},
		path: path,
	}

	itemsMu.Lock()
	defer itemsMu.Unlock()

	for i := range Items {
		if Items[i].Song.ID == id {
			Items[i] = item
			log.Printf("Updated %+v", item)
			return
		}
	}

	Items = append(Items, item)
	log.Printf("Added %+v", item)
}

func removeByPath(path string) {
	itemsMu.Lock()
	defer itemsMu.Unlock()

	if len(Items) == 0 {
		return
	}

	filtered := Items[:0]
	for _, it := range Items {
		if it.path != path {
			filtered = append(filtered, it)
		}
	}
	Items = filtered
	log.Printf("Removed items for path %s", path)
}

func Scan() {
	conf := config.Load()

	files, err := WalkFiles(conf.Server.MusicPath)
	if err != nil {
		log.Fatal("Could not walk files.")
	}

	itemsMu.Lock()
	Items = nil
	itemsMu.Unlock()

	for _, path := range files {
		indexFile(path)
	}
}

func StartWatcher() (*fsnotify.Watcher, error) {
	conf := config.Load()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("Could not create directory watcher: %v", err)
		return nil, err
	}

	go func() {
		defer watcher.Close()

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Create) {
					log.Println("created:", event.Name)
					indexFile(event.Name)
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified:", event.Name)
					indexFile(event.Name)
				}
				if event.Has(fsnotify.Remove) {
					log.Println("removed:", event.Name)
					removeByPath(event.Name)
				}
				if event.Has(fsnotify.Rename) {
					log.Println("renamed:", event.Name)
					removeByPath(event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watch error:", err)
			}
		}
	}()

	if err := watcher.Add(conf.Server.MusicPath); err != nil {
		log.Printf("Could not watch directory: %v", err)
		watcher.Close()
		return nil, err
	}

	return watcher, nil
}
