package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

	out, err := os.Create("/tmp/file")
	defer out.Close()

	if err != nil {
		fmt.Fprintf(w, "Failed to open the file for writing")
		return
	}

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

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
