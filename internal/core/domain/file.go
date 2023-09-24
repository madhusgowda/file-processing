package domains

type File struct {
	FileName string
	Size     int64
}

func NewFile(fileName string, size int64) *File {
	return &File{
		FileName: fileName,
		Size:     size,
	}
}
