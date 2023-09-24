package services

import "errors"

type FileProcessor interface {
	CalculateFileSize(data []byte) (int64, error)
}

type SampleFileProcessor struct{}

func NewFileProcessor() FileProcessor { return &SampleFileProcessor{} }

func (fp *SampleFileProcessor) CalculateFileSize(data []byte) (int64, error) {
	if len(data) == 0 {
		return 0, errors.New("empty file")
	}
	return int64(len(data)), nil
}
