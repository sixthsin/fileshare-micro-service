package storage

import (
	"file-service/config"
	"log"
	"os"
)

func InitStorage(config *config.Config) {
	if _, err := os.Stat(config.Storage.UploadDir); os.IsNotExist(err) {
		err := os.Mkdir(config.Storage.UploadDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		log.Println("Storage initialized successfully")
	}
}
