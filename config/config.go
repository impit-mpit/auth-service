package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	JWTPrivateKey string `env:"JWT_PRIVATE_KEY" env-required:"true"` // RSA Private Key in PEM format
	AdminUser     string `env:"ADMIN_USER" env-required:"true"`      // Admin username
	AdminPassword string `env:"ADMIN_PASSWORD" env-required:"true"`  // Admin password
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
