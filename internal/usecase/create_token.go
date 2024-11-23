package usecase

import (
	"context"
	"fmt"
	"neuro-most/auth-service/internal/utils"
)

type (
	CreateTokenUseCase interface {
		Execute(ctx context.Context, input CreateTokenUseCaseInput) (CreateTokenUseCaseOutput, error)
	}

	CreateTokenUseCaseInput struct {
		Username string
		Password string
	}

	CreateTokenUseCaseOutput struct {
		Token string
	}

	createTokenInteractor struct {
		adminUser     string
		adminPassword string
		jwt           utils.JWKSHandler
	}
)

func NewCreateTokenInteractor(jwt utils.JWKSHandler, adminUser, adminPassword string) CreateTokenUseCase {
	return &createTokenInteractor{
		jwt:           jwt,
		adminUser:     adminUser,
		adminPassword: adminPassword,
	}
}

func (i *createTokenInteractor) Execute(ctx context.Context, input CreateTokenUseCaseInput) (CreateTokenUseCaseOutput, error) {
	if input.Username != i.adminUser || input.Password != i.adminPassword {
		return CreateTokenUseCaseOutput{}, fmt.Errorf("invalid email or password")
	}
	token, err := i.jwt.Generate(input.Username)
	if err != nil {
		return CreateTokenUseCaseOutput{}, err
	}
	return CreateTokenUseCaseOutput{
		Token: token,
	}, nil
}
