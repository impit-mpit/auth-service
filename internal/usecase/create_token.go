package usecase

import (
	"context"
	"fmt"
	"neuro-most/auth-service/internal/utils"
)

type (
	CreateTokenUseCase interface {
		Execute(ctx context.Context, input CreateTokenUseCaseInput) CreateTokenUseCaseOutput
	}

	CreateTokenUseCaseInput struct {
		UserID int64
		Email  string
		Name   string
	}

	CreateTokenUseCaseOutput struct {
		Token string
	}

	createTokenInteractor struct {
		jwt utils.JWKSHandler
	}
)

func NewCreateTokenInteractor(jwt utils.JWKSHandler) CreateTokenUseCase {
	return &createTokenInteractor{
		jwt: jwt,
	}
}

func (i *createTokenInteractor) Execute(ctx context.Context, input CreateTokenUseCaseInput) CreateTokenUseCaseOutput {
	token, _ := i.jwt.Generate(fmt.Sprintf("%d", input.UserID), input.Email, input.Name)
	return CreateTokenUseCaseOutput{
		Token: token,
	}
}
