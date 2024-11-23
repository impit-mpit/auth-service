package router

import (
	"context"
	"log"
	"net"
	"neuro-most/auth-service/config"
	authv1 "neuro-most/auth-service/gen/go/auth/v1"
	"neuro-most/auth-service/internal/adapters/api/action"
	"neuro-most/auth-service/internal/usecase"
	"neuro-most/auth-service/internal/utils"

	"google.golang.org/grpc"
)

type RouterGrpc struct {
	cfg config.Config
	jwt utils.JWKSHandler
	authv1.UnimplementedAuthServiceServer
}

func NewRouterGrpc(jwt utils.JWKSHandler, cfg config.Config) RouterGrpc {
	return RouterGrpc{
		jwt: jwt,
		cfg: cfg,
	}
}

func (r *RouterGrpc) Listen() {
	port := ":3001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	authv1.RegisterAuthServiceServer(srv, r)

	log.Printf("Starting gRPC server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (r *RouterGrpc) Login(ctx context.Context, input *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	var (
		uc  = usecase.NewCreateTokenInteractor(r.jwt, r.cfg.AdminUser, r.cfg.AdminPassword)
		act = action.NewCreateTokenAction(uc)
	)
	return act.Execute(ctx, input)
}
