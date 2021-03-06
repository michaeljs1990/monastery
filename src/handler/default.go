package handler

import (
	"net/http"

	"github.com/michaeljs1990/monastery/src/storage"
)

// DefaultUpload is the general purpose file upload function that will feature lots
// of extra logic so users don't have to think about things like file size. Other
// upload methods will be available later in the case where the user knows more about
// the data he is going to be uploading in advanced.
func DefaultUpload(w http.ResponseWriter, r *http.Request) {
	file, err := storage.NewAbstractFileFromRequest(r)
	defer file.Handler.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = storage.Upload(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DefaultDownload handles a simple fetch of files backend data sources.
func DefaultDownload(w http.ResponseWriter, r *http.Request) {
	file, err := storage.NewAbstractFileFromDatabase(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = storage.Download(file, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
