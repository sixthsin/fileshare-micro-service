package main

import (
	"auth-service/config"
	"auth-service/generated/auth"
	"auth-service/internal/authgrpc"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/db"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	config := config.LoadConfig()
	db := db.NewDb(config)

	// REST API
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	authService := service.NewAuthService(service.AuthServiceDeps{
		Config:         config,
		UserRepository: userRepository,
	})

	// REST Handlers
	handler.NewAuthHandler(router, handler.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	go func(port string) {
		fmt.Println(port)
		if err := router.Run(":" + port); err != nil {
			log.Fatalf("Failed to start REST API server: %v", err)
		}
	}(config.Rest.Port)

	// gRPC
	listenner, err := net.Listen("tcp", ":"+config.Grpc.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC Handlers
	grpcHandler := authgrpc.NewGrpcHandler(authService)

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, grpcHandler)
	log.Printf("Starting gRPC server on : %s\n", config.Grpc.Port)

	if err := grpcServer.Serve(listenner); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
