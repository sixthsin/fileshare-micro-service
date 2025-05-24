package grpcauth

import (
	"context"
	"errors"
	"file-service/pkg/grpc/auth"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client auth.AuthServiceClient
}

func NewAuthClient(address string) (*AuthClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	return &AuthClient{client: auth.NewAuthServiceClient(conn)}, nil
}

func (c *AuthClient) ValidateToken(ctx context.Context, token string) (*auth.UserResponse, error) {
	resp, err := c.client.ValidateToken(ctx, &auth.TokenRequest{Token: token})
	if err != nil {
		return nil, err
	}
	if !resp.IsValid {
		return nil, errors.New("invalid token")
	}
	return resp, nil
}
