package config

import (
	logger "github.com/igordevopslabs/bjjgame/pkg"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}
)

func NewConfig() (*Config, error) {
	var err error

	cfg := &Config{}
	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		logger.LogError("Erro ao ler as variaveis de ambiente", err)
		return nil, err
	}

	return cfg, nil
}
