package ports

type RedisRepository interface {
	SaveFileMapping(filename string, size int64) error
	GetFileSize(filename string) (int64, error)
	Set(fileName string, fileSize int64) error
}
