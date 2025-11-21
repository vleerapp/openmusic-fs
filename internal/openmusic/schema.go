package openmusic

import "github.com/dhowden/tag"

type Song struct {
	ID          string
	Title       string
	ArtistName  string       // Optional
	AlbumName   string       // Optional
	Genre       string       // Optional
	ReleaseYear int          // Optional
	Track       int          // Optional
	TrackTotal  int          // Optional
	Disc        int          // Optional
	DiscTotal   int          // Optional
	CoverURL    *tag.Picture // Optional
	Lyrics      string       // Optional
}

type Album struct {
	ID          string
	Title       string
	ArtistName  string // Optional
	ArtistID    int    // Optional
	Genre       string // Optional
	ReleaseYear int    // Optional
	TrackCount  int    // Optional
	CoverURL    string // Optional
}

type Artist struct {
	ID       string
	Name     string
	Genre    string // Optional
	ImageURL string // Optional
}

type Playlist struct {
	ID          string
	Name        string
	Description string // Optional
	ImageURL    string // Optional
}
