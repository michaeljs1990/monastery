package config

import (
	"flag"
	"os"
)

var (
	S3Bucket          string
	S3AccessKeyID     string
	S3SecretAccessKey string
)

func init() {
	flag.StringVar(&S3Bucket, "S3Bucket", os.Getenv("S3_BUCKET"), "Set bucket to upload files to in S3")
	flag.StringVar(&S3AccessKeyID, "S3AccessKeyID", os.Getenv("S3_ACCESS_KEY_ID"), "Set access key for S3")
	flag.StringVar(&S3SecretAccessKey, "S3SecretAccessKey", os.Getenv("S3_SECRET_ACCESS_KEY"), "Set secret key for S3")
}
