package server

import (
	"github.com/karfield/go-git/plumbing/storer"
	"github.com/karfield/go-git/plumbing/transport"
	"github.com/karfield/go-git/storage/filesystem"

	"github.com/karfield/go-billy"
	"github.com/karfield/go-billy/osfs"
)

// DefaultLoader is a filesystem loader ignoring host and resolving paths to /.
var DefaultLoader = NewFilesystemLoader(osfs.New("/"))

// Loader loads repository's storer.Storer based on an optional host and a path.
type Loader interface {
	// Load loads a storer.Storer given a transport.Endpoint.
	// Returns transport.ErrRepositoryNotFound if the repository does not
	// exist.
	Load(ep transport.Endpoint) (storer.Storer, error)
}

type fsLoader struct {
	base billy.Filesystem
}

// NewFilesystemLoader creates a Loader that ignores host and resolves paths
// with a given base filesystem.
func NewFilesystemLoader(base billy.Filesystem) Loader {
	return &fsLoader{base}
}

// Load looks up the endpoint's path in the base file system and returns a
// storer for it. Returns transport.ErrRepositoryNotFound if a repository does
// not exist in the given path.
func (l *fsLoader) Load(ep transport.Endpoint) (storer.Storer, error) {
	fs := l.base.Dir(ep.Path)
	if _, err := fs.Stat("config"); err != nil {
		return nil, transport.ErrRepositoryNotFound
	}

	return filesystem.NewStorage(fs)
}

// MapLoader is a Loader that uses a lookup map of storer.Storer by
// transport.Endpoint.
type MapLoader map[transport.Endpoint]storer.Storer

// Load returns a storer.Storer for given a transport.Endpoint by looking it up
// in the map. Returns transport.ErrRepositoryNotFound if the endpoint does not
// exist.
func (l MapLoader) Load(ep transport.Endpoint) (storer.Storer, error) {
	s, ok := l[ep]
	if !ok {
		return nil, transport.ErrRepositoryNotFound
	}

	return s, nil
}
