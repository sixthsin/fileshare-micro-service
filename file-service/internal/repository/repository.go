package repository

import (
	"file-service/internal/model"
	"file-service/pkg/db"
)

type FileRepository struct {
	DB *db.Db
}

func NewFileRepository(db *db.Db) *FileRepository {
	return &FileRepository{
		DB: db,
	}
}

func (r *FileRepository) SaveFile(filePath string, fileMetadata *model.FileMetadata) error {
	result := r.DB.Create(fileMetadata)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *FileRepository) GetFileByHash(hash string) (*model.FileMetadata, error) {
	var file model.FileMetadata
	result := r.DB.DB.Where("hash = ?", hash).First(&file)
	if result.Error != nil {
		return nil, result.Error
	}
	return &file, nil
}
