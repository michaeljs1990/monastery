package storage

import (
	"net/http"

	"github.com/michaeljs1990/monastery/src/config"
)

// Download handles what backend is picked for uploading a file
func Download(a AbstractFile, w http.ResponseWriter) error {

	if config.S3Enabled {
		triggerDownloader(a, &S3{}, w)
	}

	a.Update()

	return nil
}

func triggerDownloader(a AbstractFile, b Backend, w http.ResponseWriter) error {
	b.LoadConfig()
	return b.ReadFile(a, w)
}
