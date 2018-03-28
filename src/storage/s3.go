package storage

import (
	"io"

	"github.com/aws/aws-sdk-go-v2/aws/defaults"

	"github.com/aws/aws-sdk-go-v2/aws"
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

	// cfg, err := external.LoadDefaultAWSConfig()

	upload := s3manager.NewUploader(s.config)
	_, err = upload.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(f.Name),
		Body:   reader,
	})

	return err
}

func (s *S3) ReadFile(p []byte) (n int, err error) {
	return 0, nil
}
