package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Grpc    GrpcConfig
	Db      DbConfig
	Storage StorageConfig
	Rest    RestConfig
}

type GrpcConfig struct {
	Host string
	Port string
}

type DbConfig struct {
	Dsn string
}

type StorageConfig struct {
	UploadDir string
}

type RestConfig struct {
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
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Storage: StorageConfig{
			UploadDir: os.Getenv("UPLOAD_DIR"),
		},
		Rest: RestConfig{
			Port: os.Getenv("REST_PORT"),
		},
	}
}
