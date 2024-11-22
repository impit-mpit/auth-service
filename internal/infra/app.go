package infra

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"neuro-most/auth-service/config"
	"neuro-most/auth-service/internal/infra/router"
	"neuro-most/auth-service/internal/utils"
)

type app struct {
	conf config.Config
	grpc router.RouterGrpc
	http router.RouterHTTP
	jwt  utils.JWKSHandler
}

func NewConfig(cfg config.Config) *app {
	return &app{
		conf: cfg,
	}
}

func (a *app) JWT() *app {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	a.jwt = utils.NewJWKSHandler(privateKey)
	return a
}

func (a *app) Router() *app {
	a.grpc = router.NewRouterGrpc()
	a.http = router.NewRouterHTTP(a.jwt)
	return a
}

func (a *app) Start() {
	go a.grpc.Listen()
	a.http.Listen()
}
