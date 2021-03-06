package object

import (
	"bytes"

	"github.com/karfield/go-git/utils/merkletrie"
	"github.com/karfield/go-git/utils/merkletrie/noder"
)

// DiffTree compares the content and mode of the blobs found via two
// tree objects.
func DiffTree(a, b *Tree) (Changes, error) {
	from := newTreeNoder(a)
	to := newTreeNoder(b)

	hashEqual := func(a, b noder.Hasher) bool {
		return bytes.Equal(a.Hash(), b.Hash())
	}

	merkletrieChanges, err := merkletrie.DiffTree(from, to, hashEqual)
	if err != nil {
		return nil, err
	}

	return newChanges(merkletrieChanges)
}
