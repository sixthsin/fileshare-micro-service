package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Rest RestConfig
	Grpc GrpcConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type RestConfig struct {
	Port string
}

type GrpcConfig struct {
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading .env file")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("AUTH_SERVICE_DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Rest: RestConfig{
			Port: os.Getenv("AUTH_SERVICE_REST_PORT"),
		},
		Grpc: GrpcConfig{
			Port: os.Getenv("GRPC_PORT"),
		},
	}
}
