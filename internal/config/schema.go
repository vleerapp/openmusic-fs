package config

type Links struct {
	Homepage         string  `json:"homepage" validate:"required,url"`
	PrivacyStatement *string `json:"privacyStatement" toml:"privacy_statement" validate:"omitempty,url"`
	Donate           *string `json:"donate" validate:"omitempty,url"`
}

type Branding struct {
	Name  string  `json:"name" validate:"required"`
	Email string  `json:"email" validate:"required,email"`
	Short string  `json:"short" validate:"required,min=3,max=5,uppercase"`
	Logo  *string `json:"logo" validate:"omitempty,url"`
	Theme *string `json:"theme" validate:"omitempty,hexcolor"`
	Links Links   `json:"links" validate:"required"`
}

type Capabilities struct {
	// Query
	Search            bool `json:"search" default:"true"`
	Songs             bool `json:"songs" default:"true"`
	Streaming         bool `json:"streaming" default:"true"`
	Filters           bool `json:"filters"`
	Artists           bool `json:"artists"`
	Albums            bool `json:"albums"`
	Playlists         bool `json:"playlists"`
	Lyrics            bool `json:"lyrics"`
	Popular           bool `json:"popular"`
	Similar           bool `json:"similar"`
	OpenMusicMetadata bool `json:"open_music_metadata"`
	Radio             bool `json:"radio"`
	Random            bool `json:"random"`
	List              bool `json:"list"`

	// Mutation
	Upload   bool `json:"upload"`
	Favorite bool `json:"favorite"`
	Rate     bool `json:"rate"`
	Edit     bool `json:"edit"`
}

type Details struct {
	Version      string       `json:"version" validate:"required,semver"`
	Capabilities Capabilities `json:"capabilities"`
}

type Server struct {
	Port      string `default:"8080"`
	Secret    *string
	MusicPath string `toml:"music_path"`
}

type Config struct {
	Branding Branding `validate:"required"`
	Details  Details  `validate:"required"`
	Server   Server
}
