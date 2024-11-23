package action

import (
	"context"
	authv1 "neuro-most/auth-service/gen/go/auth/v1"
	"neuro-most/auth-service/internal/usecase"
)

type CreateTokenAction struct {
	uc usecase.CreateTokenUseCase
}

func NewCreateTokenAction(uc usecase.CreateTokenUseCase) CreateTokenAction {
	return CreateTokenAction{uc: uc}
}

func (a *CreateTokenAction) Execute(ctx context.Context, input *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	token, err := a.uc.Execute(ctx, usecase.CreateTokenUseCaseInput{Username: input.Username, Password: input.Password})
	if err != nil {
		return nil, err
	}

	return &authv1.LoginResponse{Token: token.Token}, nil

}
