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
	file, err := storage.NewAbstractFileFromRequest(r)
	defer file.Handler.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = storage.Upload(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// the header contains useful info, like the original file name
	w.WriteHeader(http.StatusOK)
}

func DefaultDownload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bucket: %v\n", vars["bucket"])
}
