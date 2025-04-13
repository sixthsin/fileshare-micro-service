package grpcauth

import (
	"context"
	"file-processor-service/pkg/grpc/auth"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client auth.AuthServiceClient
}

func NewAuthClient(address string) *AuthClient {
	// creds, err := credentials.NewClientTLSFromFile("path/to/cert.pem", "")
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	return &AuthClient{
		client: auth.NewAuthServiceClient(conn),
	}
}

func (c *AuthClient) ValidateToken(ctx context.Context, token string) (*auth.UserResponse, error) {
	resp, err := c.client.ValidateToken(ctx, &auth.TokenRequest{Token: token})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
