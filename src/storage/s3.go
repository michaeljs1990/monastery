package storage

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws/defaults"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"

	"github.com/michaeljs1990/monastery/src/config"
)

// S3 struct for interacting with AWS
type S3 struct {
	config aws.Config
}

// LoadConfig handles the needed credentials for connecting to the desired
// aws account and bucket for file uploads.
func (s *S3) LoadConfig() {
	s.config = defaults.Config()

	s.config.Region = config.S3Region

	s.config.Credentials = &aws.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     config.S3AccessKeyID,
			SecretAccessKey: config.S3SecretAccessKey,
			Source:          "Monastery S3 Config",
		},
	}
}

// WriteFile takes a file and ensure that it is written up into S3
func (s *S3) WriteFile(f AbstractFile) (err error) {

	reader := io.Reader(f.Handler)

	upload := s3manager.NewUploader(s.config)
	_, err = upload.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(f.Name),
		Body:   reader,
	})

	return err
}

type Tests struct {
	http.ResponseWriter
}

// TODO: I am a bad person and this does bad things.
func (w Tests) WriteAt(b []byte, off int64) (n int, err error) {
	fmt.Println("bytes: " + strconv.Itoa(len(b)))
	fmt.Println("offset: " + strconv.FormatInt(off, 10))
	return w.Write(b)
}

// ReadFile from S3 and write to the client
func (s *S3) ReadFile(f AbstractFile, w http.ResponseWriter) error {
	// Since the client isn't able to write files with offset
	// we have to download all the bytes in order without any concurrency
	// which in testing did little to speed up the download anyway.
	downloader := s3manager.NewDownloader(s.config)
	downloader.Concurrency = 1

	t := Tests{w}
	_, err := downloader.Download(t, &s3.GetObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(f.Name),
	})

	return err
}
