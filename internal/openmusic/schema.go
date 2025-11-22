package openmusic

type Song struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ArtistName  string `json:"artistName,omitempty"`
	AlbumName   string `json:"albumName,omitempty"`
	Genre       string `json:"genre,omitempty"`
	ReleaseYear int    `json:"releaseYear,omitempty"`
	Track       int    `json:"track,omitempty"`
	TrackTotal  int    `json:"trackTotal,omitempty"`
	Disc        int    `json:"disc,omitempty"`
	DiscTotal   int    `json:"discTotal,omitempty"`
	CoverURL    string `json:"coverURL,omitempty"`
	Lyrics      string `json:"lyrics,omitempty"`
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
