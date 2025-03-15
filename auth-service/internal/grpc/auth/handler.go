package grpcauth

import (
	"auth-service/generated/shared-proto/auth"
	"auth-service/internal/service"
	"fmt"

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
		fmt.Println(err)
		return &auth.UserResponse{IsValid: false}, errors.New("invalid token")
	}
	return &auth.UserResponse{
		UserId:  user.ID,
		Email:   user.Email,
		IsValid: true,
	}, nil
}
