package service

import (
	"file-service/config"
	"file-service/internal/model"
	"file-service/internal/repository"
	"io"
	"math/rand"
	"os"
	"path/filepath"
)

type FileService struct {
	*config.Config
	*repository.FileRepository
}

type FileServiceDeps struct {
	*config.Config
	*repository.FileRepository
}

func NewFileService(deps FileServiceDeps) *FileService {
	return &FileService{
		Config:         deps.Config,
		FileRepository: deps.FileRepository,
	}
}

func generateFileHash() string {
	return RandStringRunes(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *FileService) SaveFile(file io.Reader, filename string, fileSize int64, contentType string, author string) (string, error) {
	savePath := filepath.Join(s.Config.Storage.UploadDir, filename)
	outFile, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		return "", err
	}

	fileHash := generateFileHash()
	metadata := &model.FileMetadata{
		Filename:    filename,
		Author:      author,
		ContentType: contentType,
		FileSize:    fileSize,
		Hash:        fileHash,
	}

	if err := s.FileRepository.SaveFile(savePath, metadata); err != nil {
		return "", err
	}

	return fileHash, nil
}

func (s *FileService) GetFileByHash(hash string) (*model.FileMetadata, error) {
	file, err := s.FileRepository.GetFileByHash(hash)
	if err != nil {
		return nil, err
	}
	return file, nil
}
