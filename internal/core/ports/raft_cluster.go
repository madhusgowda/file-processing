package ports

type RaftCluster interface {
	UpdateFileMapping(fileName string, fileSize int64) error
}
