package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string   `json:"username"`
	Issuer   string   `json:"iss"`
	Roles    []string `json:"roles"`
	Subject  string   `json:"sub"`
	Audience []string `json:"aud"`
}

type JWKSHandler struct {
	privateKey *rsa.PrivateKey
	issuer     string
}

func NewJWKSHandler(privateKey *rsa.PrivateKey) JWKSHandler {
	return JWKSHandler{privateKey: privateKey, issuer: "auth-service"}
}

func (m *JWKSHandler) Generate(username string) (string, error) {
	now := time.Now()
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   username,
			Issuer:    m.issuer,
		},
		Username: username,
		Roles:    []string{"admin"},
		Audience: []string{"grpc-gateway"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "key1"

	return token.SignedString(m.privateKey)
}

func (j *JWKSHandler) Validate() map[string]interface{} {
	publicKey := j.privateKey.Public().(*rsa.PublicKey)

	jwks := map[string]interface{}{
		"keys": []map[string]interface{}{
			{
				"kty": "RSA",
				"kid": "key1",
				"n":   base64.RawURLEncoding.EncodeToString(publicKey.N.Bytes()),
				"e":   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(publicKey.E)).Bytes()),
				"alg": "RS256",
				"use": "sig",
			},
		},
	}
	return jwks
}
