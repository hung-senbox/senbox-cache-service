package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Active string `yaml:"active"`

		MySQL struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Name     string `yaml:"name"`
		} `yaml:"mysql"`

		MongoDB struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
			Name string `yaml:"name"`
		} `yaml:"mongodb"`

		RedisCache struct {
			TTLSeconds int    `yaml:"ttlSeconds"`
			Host       string `yaml:"host"`
			Port       string `yaml:"port"`
			Password   string `yaml:"password"`
			DB         int    `yaml:"db"`
		} `yaml:"redis_cache"`
	} `yaml:"database"`

	Consul struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"consul"`

	Registry struct {
		Host string `yaml:"host"`
	} `yaml:"registry"`
}

// LoadConfig đọc YAML từ file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
