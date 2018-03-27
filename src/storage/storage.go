package storage

// AbstractFile is a simple file with associated metadata that can be handled
// differently depending on the storage system it's being written to.
type AbstractFile struct {
	Name     string
	Tags     []string
	Path     []string
	Metadata map[string]string
}
