package config

import (
	"fmt"
	"log"
	"os"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
)

var validate = validator.New()

func LoadErr() (*Config, error) {
	path := "om.conf.toml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("om.conf.toml not found")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := defaults.Set(&cfg); err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &cfg, nil
}

func Load() *Config {
	cfg, err := LoadErr()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
