package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	API APIConfig `toml:"api"`
}

type APIConfig struct {
	Endpoint string `toml:"endpoint"`
	Model    string `toml:"model"`
}

func DefaultConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "aid", "config.toml")
}

func LoadConfig() (*Config, error) {
	path := DefaultConfigPath()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %s", path)
	}

	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	if c.API.Endpoint == "" {
		return fmt.Errorf("api.endpoint is required in config")
	}
	if c.API.Model == "" {
		return fmt.Errorf("api.model is required in config")
	}
	return nil
}

const configTemplate = `[api]
endpoint = "http://localhost:11434/v1/chat/completions"
model = "mistral:7b"`

func InitConfig() error {
	path := DefaultConfigPath()

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("config file already exists at %s", path)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(path, []byte(configTemplate), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
