package main

import (
	"auth-service/config"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	db := db.NewDb(config)
	router := gin.Default()

	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	authService := service.NewAuthService(userRepository)

	// Handlers
	handler.NewAuthHandler(router, handler.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	router.Run(":" + config.Port.Port)
}
