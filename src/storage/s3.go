package storage

type S3 struct {
	AbstractFile
}

func (s *S3) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *S3) Read(p []byte) (n int, err error) {
	return 0, nil
}
