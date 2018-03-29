package storage

import (
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/michaeljs1990/monastery/src/cassandra"
)

var session = cassandra.CreateSession()

// AbstractFile is a simple file with associated metadata that can be handled
// differently depending on the storage system it's being written to.
type AbstractFile struct {
	Name     string
	Path     []string
	Tags     []string
	Handler  multipart.File
	Metadata map[string]string
}

// Create persists an Abstract file to the database for future lookup
func (a AbstractFile) Create() {
	if err := session.Query("INSERT INTO files (name, path, tags, created, updated) VALUES (?, ?, ?, ?, ?)",
		a.Name, a.Path, a.Tags, time.Now(), time.Now()).Exec(); err != nil {
		log.Fatal(err)
	}
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
		Path:    []string{vars["bucket"]},
		Tags:    headers["tags"],
		Handler: file,
	}, nil
}

// Backend contains the interface needed for adding in a new storage layer
type Backend interface {
	LoadConfig()
	WriteFile(AbstractFile) error
}
