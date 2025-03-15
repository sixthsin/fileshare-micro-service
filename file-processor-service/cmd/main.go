package main

import (
	"context"
	grpcauth "file-processor-service/internal/grpc/auth"
	"fmt"
	"log"
)

func main() {
	// main.go или другой файл
	authClient := grpcauth.NewAuthClient("localhost:50051")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlRlc3RFbWFpbDEyMzQ1Njc4QHRlc3QuY29tIiwiZXhwIjoxNzQyMTM2ODc2LCJ1c2VyX2lkIjoyMX0.uhFMt5APfBNRKfi4QYfxlGFaR1Wmhppe3GsFqszspio"

	response, err := authClient.ValidateToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(response)
}
