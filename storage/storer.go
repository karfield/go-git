package storage

import (
	"github.com/karfield/go-git/config"
	"github.com/karfield/go-git/plumbing/storer"
)

// Storer is a generic storage of objects, references and any information
// related to a particular repository. The package github.com/karfield/go-git/storage
// contains two implementation a filesystem base implementation (such as `.git`)
// and a memory implementations being ephemeral
type Storer interface {
	storer.EncodedObjectStorer
	storer.ReferenceStorer
	storer.ShallowStorer
	storer.IndexStorer
	config.ConfigStorer
	ModuleStorer
}

// ModuleStorer allows interact with the modules' Storers
type ModuleStorer interface {
	// Module returns a Storer reprensting a submodule, if not exists returns a
	// new empty Storer is returned
	Module(name string) (Storer, error)
}
