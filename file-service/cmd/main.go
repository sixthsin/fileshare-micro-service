package main

import (
	"context"
	"file-processor-service/config"
	grpcauth "file-processor-service/internal/grpc/auth"
	"fmt"
	"log"
)

func main() {
	config := config.LoadConfig()

	// gin.SetMode(gin.DebugMode)
	// router := gin.Default()

	authClient := grpcauth.NewAuthClient(config.Grpc.Host + ":" + config.Grpc.Port)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlRlc3RFbWFpbDJAdGVzdC5jb20iLCJleHAiOjE3NDQ2NTQ5NTEsInVzZXJuYW1lIjoiVGVzdFVzZXJuYW1lMiJ9.AySvRnpdNCawlYtouD9cDJ5mSPrFrMH6YS4T3oudV5I"

	response, err := authClient.ValidateToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(response)
}
