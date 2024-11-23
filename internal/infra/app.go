package infra

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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

func parsePrivateKey(pemString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	var ok bool
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return privateKey, nil
}

func (a *app) JWT() *app {
	privateKey, err := parsePrivateKey(a.conf.JWTPrivateKey)
	if err != nil {
		panic(err)
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
