package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AuthService string `env:"AUTH_SERVICE" env-default:"localhost:50051"`
}

func NewLoadConfig() (Config, error) {
	var cfg Config
	cleanenv.ReadConfig(".env", &cfg)
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
