package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" yaml:"port"`
	}
)

func NewConfig() (*Config, error) {
	var err error

	cfg := &Config{}

	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
