package storage

import "github.com/michaeljs1990/monastery/src/config"

// Upload handles what backend is picked for uploading a file
func Upload(a AbstractFile) error {

	if config.S3Enabled {
		return triggerUploader(a, &S3{})
	}

	return nil
}

func triggerUploader(a AbstractFile, b Backend) error {
	b.LoadConfig()
	return b.WriteFile(a)
}
