package domain

type UploadFilesInfo struct {
	FilePaths []string
	FileNames []string
	FileSizes []int64
}

type SecToDHM struct {
	Day  int
	Hour int
	Min  int
	Sec  int
}
