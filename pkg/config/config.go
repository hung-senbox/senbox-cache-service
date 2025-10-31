package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type RedisConfig struct {
	TTLSeconds int    `yaml:"ttlSeconds"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Password   string `yaml:"password"`
	DB         int    `yaml:"db"`
}

type Config struct {
	Redis RedisConfig `yaml:"redis_cache"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot parse yaml: %w", err)
	}

	return &cfg, nil
}
