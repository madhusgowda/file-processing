package ports

type FileProcessor interface {
	CalculateFileSize(filePath string) (int64, error)
}
