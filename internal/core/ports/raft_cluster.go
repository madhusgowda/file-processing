package ports

type RaftCluster interface {
	//SendMessage(message string) error
	UpdateFileMapping(fileName string, fileSize int64) error
}
