package services

import (
	"fileProcessing/internal/core/ports"
	"fileProcessing/internal/repositories/redis"
)

type FileService struct {
	repo ports.RedisRepository
}

func NewFileService(repo *redis.RedisRepository) *FileService {
	return &FileService{
		repo: repo,
	}
}

func (fs *FileService) GetFileSize(filename string) (int64, error) {
	size, err := fs.repo.GetFileSize(filename)
	if err != nil {
		return 0, err
	}
	return size, nil
}
