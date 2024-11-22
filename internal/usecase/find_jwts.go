package usecase

import (
	"context"
	"neuro-most/auth-service/internal/utils"
)

type (
	FindJwtsUseCase interface {
		Execute(ctx context.Context) FindJwtsUseCaseOutput
	}

	FindJwtsUseCaseOutput map[string]interface{}

	FindJwtsInteractor struct {
		jwt utils.JWKSHandler
	}
)

func NewFindJwtsInteractor(jwt utils.JWKSHandler) FindJwtsUseCase {
	return &FindJwtsInteractor{
		jwt: jwt,
	}
}

func (uc FindJwtsInteractor) Execute(ctx context.Context) FindJwtsUseCaseOutput {
	jwts := uc.jwt.Validate()
	return FindJwtsUseCaseOutput(jwts)
}
