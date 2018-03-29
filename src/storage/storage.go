package storage

import (
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
)

// AbstractFile is a simple file with associated metadata that can be handled
// differently depending on the storage system it's being written to.
type AbstractFile struct {
	Name     string
	Tags     []string
	Path     []string
	Handler  multipart.File
	Metadata map[string]string
}

// NewAbstractFileFromRequest from http.Request
func NewAbstractFileFromRequest(r *http.Request) (AbstractFile, error) {
	vars := mux.Vars(r)
	headers := r.Header
	file, _, err := r.FormFile("file")

	if err != nil {
		return AbstractFile{}, err
	}

	return AbstractFile{
		Name:    vars["name"],
		Tags:    headers["tags"],
		Path:    headers["path"],
		Handler: file,
	}, nil
}

// Backend contains the interface needed for adding in a new storage layer
type Backend interface {
	LoadConfig()
	WriteFile(AbstractFile) error
}
