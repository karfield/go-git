// Package filesystem is a storage backend base on filesystems
package filesystem

import (
	"github.com/karfield/go-git/storage/filesystem/internal/dotgit"

	"github.com/karfield/go-billy"
)

// Storage is an implementation of git.Storer that stores data on disk in the
// standard git format (this is, the .git directory). Zero values of this type
// are not safe to use, see the NewStorage function below.
type Storage struct {
	fs billy.Filesystem

	ObjectStorage
	ReferenceStorage
	IndexStorage
	ShallowStorage
	ConfigStorage
	ModuleStorage
}

// NewStorage returns a new Storage backed by a given `fs.Filesystem`
func NewStorage(fs billy.Filesystem) (*Storage, error) {
	dir := dotgit.New(fs)
	o, err := newObjectStorage(dir)
	if err != nil {
		return nil, err
	}

	return &Storage{
		fs: fs,

		ObjectStorage:    o,
		ReferenceStorage: ReferenceStorage{dir: dir},
		IndexStorage:     IndexStorage{dir: dir},
		ShallowStorage:   ShallowStorage{dir: dir},
		ConfigStorage:    ConfigStorage{dir: dir},
		ModuleStorage:    ModuleStorage{dir: dir},
	}, nil
}

// Filesystem returns the underlying filesystem
func (s *Storage) Filesystem() billy.Filesystem {
	return s.fs
}
