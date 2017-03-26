package git

import (
	"github.com/src-d/go-git-fixtures"
	"github.com/karfield/go-git/plumbing/transport"
	"github.com/karfield/go-git/plumbing/transport/test"

	. "github.com/karfield/go-check"
)

type UploadPackSuite struct {
	test.UploadPackSuite
	fixtures.Suite
}

var _ = Suite(&UploadPackSuite{})

func (s *UploadPackSuite) SetUpSuite(c *C) {
	s.Suite.SetUpSuite(c)

	s.UploadPackSuite.Client = DefaultClient

	ep, err := transport.NewEndpoint("git://github.com/git-fixtures/basic.git")
	c.Assert(err, IsNil)
	s.UploadPackSuite.Endpoint = ep

	ep, err = transport.NewEndpoint("git://github.com/git-fixtures/empty.git")
	c.Assert(err, IsNil)
	s.UploadPackSuite.EmptyEndpoint = ep

	ep, err = transport.NewEndpoint("git://github.com/git-fixtures/non-existent.git")
	c.Assert(err, IsNil)
	s.UploadPackSuite.NonExistentEndpoint = ep

}
