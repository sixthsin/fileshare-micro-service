package main

import (
	"context"
	grpcauth "file-processor-service/internal/grpc/auth"
	"fmt"
	"log"
)

func main() {
	authClient := grpcauth.NewAuthClient("localhost:50051")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlRlc3RFbWFpbDg4QHRlc3QuY29tIiwiZXhwIjoxNzQzNjc3MjY0LCJ1c2VyX2lkIjoyMn0.k0XfaSQLPpZ-YZ_XZ933vS9re8_N9XUfYqx9FxKPYDs"

	response, err := authClient.ValidateToken(context.Background(), token)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(response)
}
