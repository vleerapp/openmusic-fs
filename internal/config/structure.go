package config

type Links struct {
	Homepage         string `validate:"required,url"`
	PrivacyStatement string `toml:"privacy_statement" validate:"url"`
	Donate           string `validate:"url"`
}

type Branding struct {
	Name  string `validate:"required"`
	Short string `validate:"required,min=3,max=5,uppercase"`
	Logo  string `validate:"url"`
	Theme string `validate:"hexcolor"`
	Links Links  `validate:"required"`
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

type Contact struct {
	Name  string
	Email string
}

type Server struct {
	Port   string `default:"8080"`
	Secret *string
}

type Config struct {
	Branding Branding `validate:"required"`
	Details  Details  `validate:"required"`
	Contact  Contact
	Server   Server
}
