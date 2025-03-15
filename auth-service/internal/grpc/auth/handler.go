package grpcauth

import (
	"auth-service/generated/auth"
	"auth-service/internal/service"

	"context"
	"errors"
)

type GrpcHandler struct {
	auth.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewGrpcHandler(service *service.AuthService) *GrpcHandler {
	return &GrpcHandler{service: service}
}

func (handler *GrpcHandler) ValidateToken(ctx context.Context, req *auth.TokenRequest) (*auth.UserResponse, error) {
	user, err := handler.service.ValidateToken(req.Token)
	if err != nil || !user.IsValid {
		return &auth.UserResponse{IsValid: false}, errors.New("invalid token")
	}
	return &auth.UserResponse{
		UserId:  user.ID,
		Email:   user.Email,
		IsValid: true,
	}, nil
}
