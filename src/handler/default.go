package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaeljs1990/monastery/src/storage"
)

// DefaultUpload is the general purpose file upload function that will feature lots
// of extra logic so users don't have to think about things like file size. Other
// upload methods will be available later in the case where the user knows more about
// the data he is going to be uploading in advanced.
func DefaultUpload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	file, header, err := r.FormFile("file")
	defer file.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s3uploader := &storage.S3{}
	s3uploader.LoadConfig()
	abf := storage.AbstractFile{
		Name:    vars["name"],
		Handler: file,
	}
	s3uploader.WriteFile(abf)

	// the header contains useful info, like the original file name
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded successfully.", header.Filename)
	fmt.Fprintf(w, "Bucket: %v\n", vars["bucket"])
}

func DefaultDownload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bucket: %v\n", vars["bucket"])
}
