package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	NetworkConfig *NetworkConfig `yaml:"network"`
	LoggerConfig  *LoggerConfig  `yaml:"logger"`
	EngineConfig  *EngineConfig  `yaml:"engine"`
}

type NetworkConfig struct {
	Address        string        `yaml:"address"`
	MaxConnections int           `yaml:"max_connections"`
	MaxMessageSize string        `yaml:"max_message_size"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
}

type LoggerConfig struct {
	FilePath string `yaml:"file_path"`
}

type EngineConfig struct {
	Type string `yaml:"type"`
}

func Load(filename string) (*Config, error) {
	var conf Config

	if filename == "" {
		return &Config{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if err = yaml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &conf, nil
}
