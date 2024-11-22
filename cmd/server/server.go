package main

import (
	"neuro-most/auth-service/config"
	"neuro-most/auth-service/internal/infra"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	app := infra.NewConfig(cfg)
	app.JWT().Router().Start()
}
