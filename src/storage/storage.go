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
	Created  time.Time
	Updated  time.Time
}

// Fetch will return to you an abstract file that matches the name and path
// you send it. This is used for finding if a file already exists and you should
// try updating the file rather than creating a new one.
func (a AbstractFile) Fetch() (AbstractFile, error) {

	af := AbstractFile{}

	if err := session.Query("SELECT name, path, tags, metadata, created, updated FROM files WHERE name = ? AND path = ? LIMIT 1",
		a.Name, a.Path).Scan(&af.Name, &af.Path, &af.Tags, &af.Metadata, &af.Created, &af.Updated); err != nil {
		return af, err
	}

	return af, nil
}

// Update updates only tags metadata and updated timestamp. Everything else is
// unable to be altered and will result in a failure.
func (a AbstractFile) Update() {
	if err := session.Query("UPDATE files SET tags = ?, metadata = ?, updated = ? WHERE name = ? AND path = ?",
		a.Tags, a.Metadata, time.Now(), a.Name, a.Path).Exec(); err != nil {
		log.Fatal(err)
	}
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

// NewAbstractFileFromDatabase will look for an already existing entry in the database
// and create an AbstractFile for use.
func NewAbstractFileFromDatabase(r *http.Request) (AbstractFile, error) {
	vars := mux.Vars(r)

	fileFromDB, err := AbstractFile{
		Name: vars["name"],
		Path: []string{vars["bucket"]},
	}.Fetch()

	return fileFromDB, err
}

// Backend contains the interface needed for adding in a new storage layer
type Backend interface {
	LoadConfig()
	WriteFile(AbstractFile) error
	ReadFile(AbstractFile, http.ResponseWriter) error
}
