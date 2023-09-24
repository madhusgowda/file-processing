package services

import (
	domains "fileProcessing/internal/core/domain"
	"fileProcessing/internal/core/ports"
	"fileProcessing/internal/repositories/redis"
	"fmt"
)

type FileService struct {
	repo ports.RedisRepository
}

func NewFileService(repo *redis.RedisRepository) *FileService {
	return &FileService{
		repo: repo,
	}
}

func (fs *FileService) ProcessFile(file domains.File) error {
	fmt.Printf("Processing file: %s, Size: %d\n", file.FileName, file.Size)

	err := fs.repo.SaveFileMapping(file.FileName, file.Size)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileService) GetFileSize(filename string) (int64, error) {
	size, err := fs.repo.GetFileSize(filename)
	if err != nil {
		return 0, err
	}
	return size, nil
}
