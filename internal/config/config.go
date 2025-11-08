package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
)

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

type Capability string

const (
	CapFilter  Capability = "filter"
	CapArtists Capability = "artists"
)

type Details struct {
	Version      string       `validate:"required,semver"`
	Capabilities []Capability `validate:"required,dive,oneof=filter artists"`
}

type Config struct {
	Branding Branding `validate:"required"`
	Details  Details  `validate:"required"`
}

var validate = validator.New()

func Load() (*Config, error) {

	path := "om.conf.toml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("om.conf.toml not found")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &cfg, nil
}
