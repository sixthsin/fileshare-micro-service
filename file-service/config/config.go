package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Grpc GrpcConfig
}

type GrpcConfig struct {
	Host string
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading .env file")
	}
	return &Config{
		Grpc: GrpcConfig{
			Host: os.Getenv("GRPC_HOST"),
			Port: os.Getenv("GRPC_PORT"),
		},
	}
}
