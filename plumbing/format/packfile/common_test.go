package packfile

import (
	"testing"

	"github.com/karfield/go-git/plumbing"

	. "github.com/karfield/go-check"
)

func Test(t *testing.T) { TestingT(t) }

func newObject(t plumbing.ObjectType, cont []byte) plumbing.EncodedObject {
	o := plumbing.MemoryObject{}
	o.SetType(t)
	o.SetSize(int64(len(cont)))
	o.Write(cont)

	return &o
}

type piece struct {
	val   string
	times int
}

func genBytes(elements []piece) []byte {
	var result []byte
	for _, e := range elements {
		for i := 0; i < e.times; i++ {
			result = append(result, e.val...)
		}
	}

	return result
}
