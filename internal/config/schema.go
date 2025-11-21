package config

type Links struct {
	Homepage         string  `validate:"required,url"`
	PrivacyStatement *string `toml:"privacy_statement" validate:"omitempty,url"`
	Donate           *string `validate:"omitempty,url"`
}

type Branding struct {
	Name  string  `validate:"required"`
	Email string  `validate:"required,email"`
	Short string  `validate:"required,min=3,max=5,uppercase"`
	Logo  *string `validate:"omitempty,url"`
	Theme *string `validate:"omitempty,hexcolor"`
	Links Links   `validate:"required"`
}

type Capabilities struct {
	// Query
	Search            bool `default:"true"`
	Songs             bool `default:"true"`
	Streaming         bool `default:"true"`
	Filters           bool
	Artists           bool
	Albums            bool
	Playlists         bool
	Lyrics            bool
	Popular           bool
	Similar           bool
	OpenMusicMetadata bool
	Radio             bool
	Random            bool
	List              bool

	// Mutation
	Upload   bool
	Favorite bool
	Rate     bool
	Edit     bool
}

type Details struct {
	Version      string `validate:"required,semver"`
	Capabilities Capabilities
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
