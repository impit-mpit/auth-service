package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	JWTPrivateKey string `env:"JWT_PRIVATE_KEY" env-required:"true"` // RSA Private Key in PEM format
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
