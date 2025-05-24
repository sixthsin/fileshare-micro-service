package main

import (
	"file-service/config"
	grpcauth "file-service/internal/grpc/auth"
	"file-service/internal/handler"
	"file-service/internal/repository"
	"file-service/internal/service"
	"file-service/pkg/db"
	"file-service/pkg/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	db := db.NewDb(config)
	storage.InitStorage(config)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// gRPC authorization client
	authClient, err := grpcauth.NewAuthClient(config.Grpc.Host + ":" + config.Grpc.Port)
	if err != nil {
		log.Fatalf("Failed to create auth client: %v", err)
	}

	// Repository
	fileRepository := repository.NewFileRepository(db)

	// Service
	fileService := service.NewFileService(service.FileServiceDeps{
		Config:         config,
		FileRepository: fileRepository,
	})

	// Handler
	handler.NewFileShareHandler(router, handler.FileShareHandlerDeps{
		AuthClient:  authClient,
		Config:      config,
		FileService: fileService,
	})

	if err := router.Run(":" + config.Rest.Port); err != nil {
		log.Fatalf("Failed to start REST API server: %v", err)
	}

	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlRlc3RFbWFpbDJAdGVzdC5jb20iLCJleHAiOjE3NDQ2NTQ5NTEsInVzZXJuYW1lIjoiVGVzdFVzZXJuYW1lMiJ9.AySvRnpdNCawlYtouD9cDJ5mSPrFrMH6YS4T3oudV5I"

	// response, err := authClient.ValidateToken(context.Background(), token)
	// if err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }

	// fmt.Println(response)
}
